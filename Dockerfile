FROM golang:1.13.0 AS base 
WORKDIR /app/proxy
COPY *.go go.mod go.sum ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

FROM alpine:latest as certs
RUN apk --no-cache add ca-certificates

FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
WORKDIR /root/
COPY --from=base /app/proxy/server .
CMD ["./server"]