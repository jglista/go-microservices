FROM golang:1.12 as builder

WORKDIR /go/src/github.com/jglista/go-microservices/user-service

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o user-service .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/jglista/go-microservices/user-service/user-service .

CMD ["./user-service"]