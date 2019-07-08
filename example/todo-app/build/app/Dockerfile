FROM golang:1.12 AS base
WORKDIR /go/src/github.com/prksu/todoapp
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod download

FROM base AS build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /todo-app cmd/main.go

FROM alpine:latest
COPY --from=build /todo-app /todo-app
EXPOSE 9000
ENTRYPOINT [ "/todo-app" ]