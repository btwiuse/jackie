FROM golang:1.18 AS builder
WORKDIR /home/xchain

RUN apt update && apt install -y unzip

# small trick to take advantage of  docker build cache
RUN ls
COPY go.* ./
COPY Makefile .
RUN make prepare

COPY . .
RUN make

# ---
#FROM ubuntu:18.04
FROM ethereum/solc:stable AS solc
FROM btwiuse/xuper
WORKDIR /home/xchain
RUN apt update
RUN apt install -y build-essential vim jq
COPY --from=builder /home/xchain/output /usr/local/bin
COPY --from=builder /home/xchain/scripts /usr/local/bin
COPY --from=builder /home/xchain/template template
COPY --from=solc /usr/bin/solc /usr/local/bin
RUN ln -s $PWD/bin/* /usr/local/bin
EXPOSE 8085 37101 47101
CMD bash control.sh start; jackie
