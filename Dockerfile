# Build stage
FROM golang:alpine AS build-env

RUN apk add --no-cache git

ADD . ./src/druid-metrics-datadog
RUN cd ./src/druid-metrics-datadog && \
  go get && \
  CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags '-s -w' -o /druid-metrics-datadog

# Final stage
FROM scratch
COPY --from=build-env /druid-metrics-datadog /
ENTRYPOINT ["/druid-metrics-datadog"]
