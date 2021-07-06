FROM golang:1.15.1-alpine AS builder
RUN apk update && apk --no-cache add ca-certificates tzdata \
  bash git openssh gcc g++ pkgconfig build-base curl \
  zlib-dev librdkafka-dev pkgconf && rm -rf /var/cache/apk/*
# Copy app and run go mod.
WORKDIR /go/src/github.com/ElvertMora/quasar-fire-app
COPY ./go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -tags musl -a -installsuffix cgo -o app .
# Copy and run the app.
FROM alpine:3.6
WORKDIR /usr/
COPY --from=builder /go/src/github.com/ElvertMora/quasar-fire-app/app .
CMD /usr/app