# vessel-service/Dockerfile
FROM golang:1.11.1 as builder

WORKDIR /go/src/github.com/jglista/vessel-service

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN  dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -o vessel-service -a -installsuffix cgo -o main .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/jglista/vessel-service/main .

CMD ["./main"]