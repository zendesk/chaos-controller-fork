FROM ubuntu:24.04 as client

COPY built_go_server /usr/local/bin/dogfood_server

ENTRYPOINT [ "/usr/local/bin/dogfood_server" ]
