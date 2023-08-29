# FROM golang:1.16-buster AS builder
# WORKDIR /app
# COPY go.* ./
# RUN go mod download
# COPY *.go ./
# RUN go build -o /hello_go_http
# EXPOSE 8080
# ENTRYPOINT ["/hello_go_http"]

FROM golang:latest
WORKDIR /pdfcpu_rest_api
COPY ./ /pdfcpu_rest_api
RUN go mod download
ENTRYPOINT go run ./main.go