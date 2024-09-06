FROM ubuntu:20.04

WORKDIR /app

RUN apt update
RUN apt install -y wget git make gcc
RUN wget https://go.dev/dl/go1.23.0.linux-amd64.tar.gz && \
    tar xvf go1.23.0.linux-amd64.tar.gz && \
    rm go1.23.0.linux-amd64.tar.gz && \
    ln -s /app/go/bin/go /bin/go

RUN git clone --depth 1 --branch v0.0.1-alpha8 https://github.com/leptonai/gpud.git
RUN cd gpud/ && CGO_ENABLED=1 make all

EXPOSE 15132
