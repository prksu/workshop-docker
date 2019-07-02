FROM golang:1.12 AS base
WORKDIR /go/src/github.com/prksu/todosapp
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod download

FROM base AS build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /todos-app cmd/main.go

FROM alpine:latest
COPY --from=build /todos-app /todos-app
EXPOSE 9000
ENTRYPOINT [ "/todos-app" ]