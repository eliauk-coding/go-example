FROM golang:1.16-alpine
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR $GOPATH/src/GoUpUp/docker_go

COPY .. .

RUN go build .

EXPOSE 8080
ENTRYPOINT ["./docker_go"]