FROM golang:1.9.2
WORKDIR /go/src/github.com/tomaszheflik/RMQ_go/
COPY builder.go converters.go error_handle.go read_mq.go send_mq.go types.go ./
RUN go get -v 
RUN CGO_ENABLE=0 GOOS=linux go build -a -o builder .

FROM debian:9.4-slim
RUN apt-get update && apt-get install ca-certificates -y
WORKDIR /root/
COPY --from=0 /go/src/github.com/tomaszheflik/RMQ_go/builder .
ENTRYPOINT ["./builder"]
