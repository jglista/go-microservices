FROM golang:1.12.0 as builder

WORKDIR /go/src/github.com/jglista/go-microservices/user-cli

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o user-cli .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/jglista/go-microservices/user-cli/user-cli .

ENTRYPOINT ["./user-cli"]

CMD ["./user-cli"]