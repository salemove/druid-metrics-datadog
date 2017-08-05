package main

import (
	"github.com/DataDog/datadog-go/statsd"
	"github.com/antonholmquist/jason"
	"log"
	"strings"
)

const SAMPLE_RATE = 1.0

var excludedTags = [...]string{"timestamp", "service", "metric", "value", "context", "id", "interval", "taskId"}

type DatadogPublisher struct {
	client      *statsd.Client
	definitions map[string]Definition
}

func CreateDatadogPublisher(addr string) (*DatadogPublisher, error) {
	datadogClient, err := statsd.New(addr)
	if err != nil {
		return nil, err
	}
	publisher := DatadogPublisher{client: datadogClient, definitions: MetricDefinitions()}
	return &publisher, nil
}

func (p *DatadogPublisher) Publish(metric *jason.Object) {
	metricName, err := metric.GetString("metric")
	if err != nil {
		log.Println("Found a metric without a name", err, metric)
		return
	}

	serviceName, err := metric.GetString("service")
	if err != nil {
		log.Println("Found a metric without a service name", err, metric)
		return
	}

	if serviceName == "druid/peon" {
		// Skipping all metrics from peon. These are task specific and need some
		// thinking before sending to DataDog.
		return
	}

	definition, ok := p.definitions[metricName]
	if !ok {
		log.Println("Found a metric without a definition", metric)
		return
	}

	metricKey := PrepareMetricKey(serviceName, metricName)

	metricValue, err := metric.GetFloat64("value")
	if err != nil {
		log.Println("Unable to prepare metric value", err, metric)
	}

	if definition.ConvertRange {
		metricValue = metricValue * 100
	}

	tags := PrepareTags(metric)

	switch metricType := definition.Type; metricType {
	case Count:
		p.client.Count(metricKey, int64(metricValue), tags, SAMPLE_RATE)
	case Gauge:
		p.client.Gauge(metricKey, metricValue, tags, SAMPLE_RATE)
	case Timer:
		p.client.TimeInMilliseconds(metricKey, metricValue, tags, SAMPLE_RATE)
	default:
		log.Println("Unknown metric type", metricType)
	}
}

func PrepareMetricKey(serviceName string, metricName string) string {
	return strings.Replace(serviceName+"."+metricName, "/", ".", -1)
}

func PrepareTags(metric *jason.Object) []string {
	var tags []string
	for fieldName, fieldValue := range metric.Map() {
		if !IsExcludedTag(fieldName) {
			if stringValue, err := fieldValue.String(); err == nil {
				tags = append(tags, fieldName+":"+stringValue)
			}
		}
	}

	return tags
}

func IsExcludedTag(tag string) bool {
	for _, excludedTag := range excludedTags {
		if tag == excludedTag {
			return true
		}
	}
	return false
}
