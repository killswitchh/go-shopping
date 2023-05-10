FROM golang:latest

ENV APP_NAME order-service
WORKDIR /go/src/${APP_NAME}
COPY . .

RUN go install github.com/cosmtrek/air@latest
RUN go get ./
RUN go mod tidy