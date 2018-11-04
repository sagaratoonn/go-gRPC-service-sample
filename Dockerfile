FROM golang:latest
COPY server.go ./
COPY pb ./pb/
RUN go get -u google.golang.org/grpc
RUN go build ./server.go
ENTRYPOINT ["./server"]
