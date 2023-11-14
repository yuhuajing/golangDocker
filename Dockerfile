FROM golang:alpine AS builder
WORKDIR /build
RUN adduser -u 10001 -D app-runner
ENV GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOARCH=amd64 \
    GOOS=linux \
    GO111MODULE=on
COPY --link ./go.mod .
COPY --link ./go.sum .
RUN go mod download
COPY --link . .
RUN go build -ldflags "-s -w" -a -o main ./main.go
FROM alpine:latest AS final
WORKDIR /app
COPY --from=builder --link /etc/passwd /etc/passwd
COPY --from=builder --link /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder --link /build/main /app/main
USER app-runner
EXPOSE 3000
ENTRYPOINT ["/app/main"]

# FROM golang:alpine AS build

# RUN apk add git

# RUN mkdir /src
# ADD . /src
# WORKDIR /src

# RUN go build -o /tmp/http-server ./main.go

# FROM alpine:edge

# COPY --from=build /tmp/http-server /sbin/http-server

# CMD /sbin/http-server