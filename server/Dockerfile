FROM golang:1.21-alpine3.18
WORKDIR /app
COPY . ./server
RUN mv ./server/mod/* .
RUN go mod tidy
RUN go build -o main ./server/cmd/main.go

EXPOSE 5300 8081
CMD [ "/app/main" ]