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
FROM btwiuse/xuper
WORKDIR /home/xchain
#RUN apt update&& apt install -y build-essential
COPY --from=builder /home/xchain/output /usr/bin
RUN ln -s $PWD/bin/* /usr/bin
EXPOSE 8085 37101 47101
CMD bash control.sh start; jackie
