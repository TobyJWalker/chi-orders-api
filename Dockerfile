FROM golang:1.19-alpine

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-library .

ENV SERVER_PORT=8082
ENV APP_ENV=production
EXPOSE 8082

CMD ["/go-library"]