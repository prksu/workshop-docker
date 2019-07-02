FROM golang:1.12
WORKDIR /go/src/github.com/fossildev/hello-go
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/hello-go main.go
EXPOSE 9000
ENTRYPOINT [ "/app/hello-go" ]