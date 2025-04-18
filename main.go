// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2025 Datadog, Inc.

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	corev1 "k8s.io/api/core/v1"

	chaosv1beta1 "github.com/DataDog/chaos-controller/api/v1beta1"
	"github.com/DataDog/chaos-controller/cloudservice"
	"github.com/DataDog/chaos-controller/config"
	"github.com/DataDog/chaos-controller/controllers"
	"github.com/DataDog/chaos-controller/eventbroadcaster"
	"github.com/DataDog/chaos-controller/eventnotifier"
	"github.com/DataDog/chaos-controller/log"
	"github.com/DataDog/chaos-controller/o11y/metrics"
	metricstypes "github.com/DataDog/chaos-controller/o11y/metrics/types"
	"github.com/DataDog/chaos-controller/o11y/profiler"
	profilertypes "github.com/DataDog/chaos-controller/o11y/profiler/types"
	"github.com/DataDog/chaos-controller/o11y/tracer"
	tracertypes "github.com/DataDog/chaos-controller/o11y/tracer/types"
	"github.com/DataDog/chaos-controller/services"
	"github.com/DataDog/chaos-controller/targetselector"
	"github.com/DataDog/chaos-controller/utils"
	"github.com/DataDog/chaos-controller/watchers"
	chaoswebhook "github.com/DataDog/chaos-controller/webhook"

	"github.com/go-logr/zapr"
	"k8s.io/apimachinery/pkg/runtime"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/cache"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
	"sigs.k8s.io/controller-runtime/pkg/webhook"

	// +kubebuilder:scaffold:imports

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

//go:generate mockery  --config .local.mockery.yaml
//go:generate mockery  --config .vendor.mockery.yaml

var scheme = runtime.NewScheme()

func init() {
	// +kubebuilder:scaffold:scheme
	_ = clientgoscheme.AddToScheme(scheme)
	_ = chaosv1beta1.AddToScheme(scheme)
}

func main() {
	logger, err := log.NewZapLogger()
	if err != nil {
		ctrl.Log.WithName("setup").Error(err, "error creating controller logger")
		os.Exit(1)
	}

	desugaredLogger := zapr.NewLogger(logger.Desugar())

	// Set any singleton loggers of underlying libraries to use our zap logger
	ctrl.SetLogger(desugaredLogger)
	klog.SetLogger(desugaredLogger)

	// get controller node name
	controllerNodeName, exists := os.LookupEnv("CONTROLLER_NODE_NAME")
	if !exists {
		logger.Fatal("missing required CONTROLLER_NODE_NAME environment variable")
	}

	restConfig, err := rest.InClusterConfig()
	if err != nil {
		logger.Fatalw("error creating in-cluster rest config", "error", err)
	}

	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		logger.Fatalw("error creating kubernetes clientset", "error", err)
	}

	configMapClient := clientset.CoreV1().ConfigMaps(os.Getenv("POD_NAMESPACE"))

	cfg, err := config.New(configMapClient, logger, os.Args[1:])
	if err != nil {
		logger.Fatalw("unable to create a valid configuration", "error", err)
	}

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 scheme,
		HealthProbeBindAddress: cfg.Controller.HealthProbeBindAddr,
		Metrics: metricsserver.Options{
			BindAddress: cfg.Controller.MetricsBindAddr,
		},
		LeaderElection:   cfg.Controller.LeaderElection,
		LeaderElectionID: "75ec2fa4.datadoghq.com",
		Logger:           desugaredLogger,
		WebhookServer: webhook.NewServer(webhook.Options{
			Host:    cfg.Controller.Webhook.Host,
			Port:    cfg.Controller.Webhook.Port,
			CertDir: cfg.Controller.Webhook.CertDir,
		}),
	})

	if err != nil {
		logger.Fatalw("unable to start manager", "error", err)
	}

	broadcaster := eventbroadcaster.EventBroadcaster()

	// event notifiers
	notifiers, err := eventnotifier.CreateNotifiers(cfg.Controller.Notifiers, logger)
	if err != nil {
		logger.Errorw("error(s) while creating notifiers", "error", err)
	}

	eventbroadcaster.RegisterNotifierSinks(mgr, broadcaster, notifiers, logger)

	metricsSink := initMetricsSink(cfg.Controller.MetricsSink, logger, metricstypes.SinkAppController)

	defer closeMetricsSink(logger, metricsSink)

	profilerSink, err := profiler.GetSink(logger, profilertypes.SinkDriver(cfg.Controller.ProfilerSink))
	if err != nil {
		logger.Errorw("error while creating profiler sink, switching to noop", "error", err)

		profilerSink, _ = profiler.GetSink(logger, profilertypes.SinkDriverNoop)
	}
	// handle profiler sink close on exit
	defer func() {
		logger.Infow("closing profiler sink client before exiting", "sink", profilerSink.GetSinkName())
		profilerSink.Stop()
	}()

	tracerSink, err := tracer.GetSink(logger, tracertypes.SinkDriver(cfg.Controller.TracerSink))
	if err != nil {
		logger.Errorw("error while creating profiler sink, switching to noop", "error", err)

		tracerSink, _ = tracer.GetSink(logger, tracertypes.SinkDriverNoop)
	}
	// handle tracer sink close on exit
	defer func() {
		logger.Infow("closing tracer sink client before exiting", "sink", tracerSink.GetSinkName())

		if err := tracerSink.Stop(); err != nil {
			logger.Errorw("error closing tracer sink client", "sink", metricsSink.GetSinkName(), "error", err)
		}
	}()

	// initiate Open Telemetry, set it up with the sink Provider, use TraceContext for propagation through the CRD
	otel.SetTracerProvider(tracerSink.GetProvider())
	otel.SetTextMapPropagator(propagation.TraceContext{})

	if err = metricsSink.MetricRestart(); err != nil {
		logger.Errorw("error sending MetricRestart", "sink", metricsSink.GetSinkName(), "error", err)
	}

	// target selector
	targetSelector := targetselector.NewRunningTargetSelector(cfg.Controller.EnableSafeguards, controllerNodeName)

	var gcPtr *time.Duration
	if cfg.Controller.ExpiredDisruptionGCDelay >= 0 {
		gcPtr = &cfg.Controller.ExpiredDisruptionGCDelay
	}

	// initialize the cloud provider manager which will handle ip ranges files updates
	cloudProviderManager, err := cloudservice.New(logger, cfg.Controller.CloudProviders, nil)
	if err != nil {
		logger.Fatalw("error initializing CloudProviderManager", "error", err)
	}

	cloudProviderManager.StartPeriodicPull()

	chaosPodService, err := services.NewChaosPodService(services.ChaosPodServiceConfig{
		Client:         mgr.GetClient(),
		Log:            logger,
		ChaosNamespace: cfg.Injector.ChaosNamespace,
		TargetSelector: targetSelector,
		Injector: services.ChaosPodServiceInjectorConfig{
			ServiceAccount:                cfg.Injector.ServiceAccount,
			Image:                         cfg.Injector.Image,
			Annotations:                   cfg.Injector.Annotations,
			Labels:                        cfg.Injector.Labels,
			Tolerations:                   cfg.Injector.Tolerations,
			NetworkDisruptionAllowedHosts: cfg.Injector.NetworkDisruption.AllowedHosts,
			DNSDisruptionDNSServer:        cfg.Injector.DNSDisruption.DNSServer,
			DNSDisruptionKubeDNS:          cfg.Injector.DNSDisruption.KubeDNS,
			ImagePullSecrets:              cfg.Injector.ImagePullSecrets,
			LogLevel:                      cfg.Injector.LogLevel,
		},
		ImagePullSecrets: cfg.Injector.ImagePullSecrets,
		MetricsSink:      metricsSink,
	})

	if err != nil {
		logger.Fatalw("error initializing ChaosPodService", "error", err)
	}

	// create disruption reconciler
	disruptionReconciler := &controllers.DisruptionReconciler{
		Client:                     mgr.GetClient(),
		BaseLog:                    logger,
		Scheme:                     mgr.GetScheme(),
		Recorder:                   broadcaster.NewRecorder(mgr.GetScheme(), corev1.EventSource{Component: chaosv1beta1.SourceDisruptionComponent}),
		MetricsSink:                metricsSink,
		TracerSink:                 tracerSink,
		TargetSelector:             targetSelector,
		ExpiredDisruptionGCDelay:   gcPtr,
		ChaosPodService:            chaosPodService,
		CloudService:               cloudProviderManager,
		DisruptionsDeletionTimeout: cfg.Controller.DisruptionDeletionTimeout,
		DeleteOnly:                 cfg.Controller.DeleteOnly,
		FinalizerDeletionDelay:     cfg.Controller.FinalizerDeletionDelay,
	}

	informerClient := kubernetes.NewForConfigOrDie(ctrl.GetConfigOrDie())
	kubeInformerFactory := kubeinformers.NewSharedInformerFactoryWithOptions(informerClient, time.Hour*24, kubeinformers.WithNamespace(cfg.Injector.ChaosNamespace))

	cont, err := disruptionReconciler.SetupWithManager(mgr, kubeInformerFactory)
	if err != nil {
		logger.Fatalw("unable to create controller", "controller", chaosv1beta1.DisruptionKind, "error", err)
	}

	watchersFactoryConfig := watchers.FactoryConfig{
		Log:            logger,
		MetricSink:     metricsSink,
		Reader:         mgr.GetAPIReader(),
		Recorder:       disruptionReconciler.Recorder,
		ChaosNamespace: cfg.Injector.ChaosNamespace,
	}
	watcherFactory := watchers.NewWatcherFactory(watchersFactoryConfig)
	disruptionReconciler.DisruptionsWatchersManager = watchers.NewDisruptionsWatchersManager(cont, watcherFactory, mgr.GetAPIReader(), logger)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		ticker := time.NewTicker(time.Minute * 5)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				logger.Debugw("Check if we need to remove any expired watchers...")
				disruptionReconciler.DisruptionsWatchersManager.RemoveAllExpiredWatchers()

			case <-ctx.Done():
				// Context canceled, terminate the goroutine
				return
			}
		}
	}()

	defer cancel()

	stopCh := make(chan struct{})
	kubeInformerFactory.Start(stopCh)

	go disruptionReconciler.ReportMetrics(ctx)

	if cfg.Controller.DisruptionRolloutEnabled {
		// create deployment and statefulset informers
		globalInformerFactory := kubeinformers.NewSharedInformerFactory(informerClient, time.Hour*24)
		deploymentInformer := globalInformerFactory.Apps().V1().Deployments().Informer()
		statefulsetInformer := globalInformerFactory.Apps().V1().StatefulSets().Informer()

		deploymentHandler := watchers.NewDeploymentHandler(mgr.GetClient(), logger)
		statefulsetHandler := watchers.NewStatefulSetHandler(mgr.GetClient(), logger)

		_, err = deploymentInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
			AddFunc:    deploymentHandler.OnAdd,
			UpdateFunc: deploymentHandler.OnUpdate,
			DeleteFunc: deploymentHandler.OnDelete,
		})
		if err != nil {
			logger.Fatalw("unable to add event handler for Deployments", "error", err)
		}

		_, err = statefulsetInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
			AddFunc:    statefulsetHandler.OnAdd,
			UpdateFunc: statefulsetHandler.OnUpdate,
			DeleteFunc: statefulsetHandler.OnDelete,
		})
		if err != nil {
			logger.Fatalw("unable to add event handler for StatefulSets", "error", err)
		}

		// wait for the deployment and statefulset informer caches to be synced
		synced := globalInformerFactory.WaitForCacheSync(ctx.Done())
		for informerType, ok := range synced {
			if !ok {
				logger.Errorw("failed to wait for informer cache to sync", "informer", informerType)
				return
			}
		}

		// start the deployment and statefulset informers
		globalInformerFactory.Start(stopCh)

		// create disruption rollout reconciler
		disruptionRolloutReconciler := &controllers.DisruptionRolloutReconciler{
			Client:  mgr.GetClient(),
			BaseLog: logger,
			Scheme:  mgr.GetScheme(),
			// new metrics sink for rollout controller
			MetricsSink:                    initMetricsSink(cfg.Controller.MetricsSink, logger, metricstypes.SinkAppRolloutController),
			TargetResourceMissingThreshold: cfg.Controller.TargetResourceMissingThreshold,
		}

		defer closeMetricsSink(logger, disruptionRolloutReconciler.MetricsSink)

		if err := disruptionRolloutReconciler.SetupWithManager(mgr); err != nil {
			logger.Errorw("unable to create controller", "controller", "DisruptionRollout", "error", err)
			os.Exit(1) //nolint:gocritic
		}

		// add the indexer on target resource for disruption rollouts
		err = mgr.GetCache().IndexField(context.Background(), &chaosv1beta1.DisruptionRollout{}, "targetResource", func(obj client.Object) []string {
			dr, ok := obj.(*chaosv1beta1.DisruptionRollout)
			if !ok {
				return []string{""}
			}

			targetResource := fmt.Sprintf("%s-%s-%s", dr.Spec.TargetResource.Kind, dr.GetNamespace(), dr.Spec.TargetResource.Name)

			return []string{targetResource}
		})
		if err != nil {
			logger.Fatalw("unable to add index", "controller", "DisruptionRollout", "error", err)
		}
	}

	// register disruption validating webhook
	setupWebhookConfig := utils.SetupWebhookWithManagerConfig{
		Manager:                       mgr,
		Logger:                        logger,
		MetricsSink:                   metricsSink,
		TracerSink:                    tracerSink,
		Recorder:                      disruptionReconciler.Recorder,
		NamespaceThresholdFlag:        cfg.Controller.SafeMode.NamespaceThreshold,
		ClusterThresholdFlag:          cfg.Controller.SafeMode.ClusterThreshold,
		EnableSafemodeFlag:            cfg.Controller.SafeMode.Enable,
		AllowNodeFailure:              cfg.Controller.SafeMode.AllowNodeFailure,
		AllowNodeLevel:                cfg.Controller.SafeMode.AllowNodeLevel,
		DisabledDisruptions:           cfg.Controller.DisabledDisruptions,
		DeleteOnlyFlag:                cfg.Controller.DeleteOnly,
		HandlerEnabledFlag:            cfg.Handler.Enabled,
		DefaultDurationFlag:           cfg.Controller.DefaultDuration,
		MaxDurationFlag:               cfg.Controller.MaxDuration,
		ChaosNamespace:                cfg.Injector.ChaosNamespace,
		CloudServicesProvidersManager: cloudProviderManager,
		Environment:                   cfg.Controller.SafeMode.Environment,
		PermittedUserGroups:           cfg.Controller.SafeMode.PermittedUserGroups,
	}

	logger.Debug("setup webhook for disruption")

	if err = (&chaosv1beta1.Disruption{}).SetupWebhookWithManager(setupWebhookConfig); err != nil {
		logger.Fatalw("unable to create webhook for disruption", "webhook", chaosv1beta1.DisruptionKind, "error", err)
	}

	if cfg.Controller.DisruptionCronEnabled {
		disruptionCronRecorder := broadcaster.NewRecorder(mgr.GetScheme(), corev1.EventSource{Component: chaosv1beta1.SourceDisruptionCronComponent})
		disruptionCronMetricsSink := initMetricsSink(cfg.Controller.MetricsSink, logger, metricstypes.SinkAppCronController)

		defer closeMetricsSink(logger, disruptionCronMetricsSink)

		// create disruption cron reconciler
		disruptionCronReconciler := &controllers.DisruptionCronReconciler{
			Client:  mgr.GetClient(),
			BaseLog: logger,
			Scheme:  mgr.GetScheme(),
			// separate metrics sink and recorder for cron controller
			Recorder:                       disruptionCronRecorder,
			MetricsSink:                    disruptionCronMetricsSink,
			FinalizerDeletionDelay:         cfg.Controller.FinalizerDeletionDelay,
			TargetResourceMissingThreshold: cfg.Controller.TargetResourceMissingThreshold,
		}

		defer closeMetricsSink(logger, disruptionCronReconciler.MetricsSink)

		if err := disruptionCronReconciler.SetupWithManager(mgr); err != nil {
			logger.Errorw("unable to create controller", "controller", "DisruptionCron", "error", err)
			os.Exit(1) //nolint:gocritic
		}

		disruptionCronSetupWebhookConfig := utils.SetupWebhookWithManagerConfig{
			Manager:                          mgr,
			Logger:                           logger,
			Recorder:                         disruptionCronRecorder,
			DeleteOnlyFlag:                   cfg.Controller.DeleteOnly,
			PermittedUserGroups:              cfg.Controller.SafeMode.PermittedUserGroups,
			DefaultCronDelayedStartTolerance: cfg.Controller.DefaultCronDelayedStartTolerance,
			MinimumCronFrequency:             cfg.Controller.MinimumCronFrequency,
			DefaultDurationFlag:              cfg.Controller.DefaultDuration,
			MetricsSink:                      disruptionCronMetricsSink,
		}

		if err = (&chaosv1beta1.DisruptionCron{}).SetupWebhookWithManager(disruptionCronSetupWebhookConfig); err != nil {
			logger.Fatalw("unable to create webhook for disruption cron", "webhook", chaosv1beta1.DisruptionCronKind, "error", err)
		}
	}

	webhookDecoder := admission.NewDecoder(scheme)

	if cfg.Handler.Enabled {
		// register chaos handler init container mutating webhook
		mgr.GetWebhookServer().Register("/mutate-v1-pod-chaos-handler-init-container", &webhook.Admission{
			Handler: &chaoswebhook.ChaosHandlerMutator{
				Client:     mgr.GetClient(),
				Log:        logger,
				Image:      cfg.Handler.Image,
				Timeout:    cfg.Handler.Timeout,
				MaxTimeout: cfg.Handler.MaxTimeout,
				Decoder:    webhookDecoder,
				ResourceList: &corev1.ResourceList{
					corev1.ResourceCPU:    resource.MustParse(cfg.Handler.CPU),
					corev1.ResourceMemory: resource.MustParse(cfg.Handler.Memory),
				},
			},
		})
	}

	if cfg.Controller.UserInfoHook {
		// register user info mutating webhook
		mgr.GetWebhookServer().Register("/mutate-chaos-datadoghq-com-v1beta1-user-info", &webhook.Admission{
			Handler: &chaoswebhook.UserInfoMutator{
				Client:  mgr.GetClient(),
				Log:     logger,
				Decoder: webhookDecoder,
			},
		})
	}

	mgr.GetWebhookServer().Register("/mutate-chaos-datadoghq-com-v1beta1-disruption-span-context", &webhook.Admission{
		Handler: &chaoswebhook.SpanContextMutator{
			Client:  mgr.GetClient(),
			Log:     logger,
			Decoder: webhookDecoder,
		},
	})

	// for safety purposes: as long as no event is emitted and mgr.Start(ctx.Context) isn't
	// called, the broadcaster isn't actually initiated
	defer broadcaster.Shutdown()

	// +kubebuilder:scaffold:builder

	logger.Infow("starting chaos-controller")

	if err := mgr.AddHealthzCheck("healthz", func(req *http.Request) error {
		return mgr.GetWebhookServer().StartedChecker()(req)
	}); err != nil {
		logger.Fatalw("Unable to set up health check", "error", err)
	}

	if err := mgr.AddReadyzCheck("readyz", func(req *http.Request) error {
		return mgr.GetWebhookServer().StartedChecker()(req)
	}); err != nil {
		logger.Fatalw("Unable to set up ready check", "error", err)
	}

	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		stopCh <- struct{}{} // stop the informer

		logger.Fatalw("problem running manager", "error", err)
	}
}

// initialize metrics sink
func initMetricsSink(sink string, logger *zap.SugaredLogger, app metricstypes.SinkApp) metrics.Sink {
	metricsSink, err := metrics.GetSink(logger, metricstypes.SinkDriver(sink), app)
	if err != nil {
		logger.Errorw("error while creating metric sink, switching to noop", "error", err)

		metricsSink, _ = metrics.GetSink(logger, metricstypes.SinkDriverNoop, app)
	}

	return metricsSink
}

// handle metrics sink client close on exit
func closeMetricsSink(logger *zap.SugaredLogger, metricsSink metrics.Sink) {
	logger.Infow("closing metrics sink client before exiting", "sink", metricsSink.GetSinkName())

	if err := metricsSink.Close(); err != nil {
		logger.Errorw("error closing metrics sink client", "sink", metricsSink.GetSinkName(), "error", err)
	}
}
