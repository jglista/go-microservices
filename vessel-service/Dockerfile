# vessel-service/Dockerfile
FROM golang:1.12 as builder

WORKDIR /go/src/github.com/jglista/go-microservices/vessel-service

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN  dep ensure -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o vessel-service .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/jglista/go-microservices/vessel-service/vessel-service .

CMD ["./vessel-service"]