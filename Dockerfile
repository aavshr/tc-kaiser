FROM golang:1.13.0 AS base 
WORKDIR /app/proxy
COPY *.go go.mod go.sum ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

FROM alpine:latest as certs
RUN apk --no-cache add ca-certificates

FROM nginx:stable-alpine
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
WORKDIR /root/
COPY --from=base /app/proxy/nginx/default.conf /etc/nginx/conf.d
COPY --from=base /app/proxy/server .
COPY --from=base /app/proxy/nginx/start.sh .
ENV S3_BUCKET tc-deta-bucket
ENV AWS_REGION eu-central-1
CMD ["./start.sh"]