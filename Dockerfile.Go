FROM golang

ADD Makefile /
WORKDIR /

ADD . /go/src/lean-golang-docker-container/
WORKDIR /go/src/lean-golang-docker-container/
RUN make setup
WORKDIR /

RUN make buildgo
CMD ["/bin/bash"]