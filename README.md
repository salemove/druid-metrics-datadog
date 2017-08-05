# Druid Metrics DataDog

This is an HTTP server that receives metrics over Druid http emitter and forwards them to DataDog.

## Configuration

Configuration is done using environment variables:
* `DATADOG_ADDRESS` - Address where DD Agent is running (default: 127.0.0.1:8125)
* `PORT` - Port where this server listens on (default: 8424)

## Druid Configuration
```
druid.monitoring.emissionPeriod=PT10s
druid.monitoring.monitors=["com.metamx.metrics.JvmMonitor"]
druid.emitter=none
druid.emitter.http.flushMillis=10000
druid.emitter.http.recipientBaseUrl=http://ADDR_TO_THIS_SERVICE:8424
```
