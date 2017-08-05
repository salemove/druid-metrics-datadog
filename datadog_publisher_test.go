package main

import (
	"github.com/antonholmquist/jason"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrepareMetricWhenServiceAndMetricPresent(t *testing.T) {
	metricBody := "{\"service\":\"druid/broker\",\"metric\":\"some/metric\"}"
	metric, _ := jason.NewObjectFromBytes([]byte(metricBody))

	key, err := PrepareMetricKey(metric)
	assert.Nil(t, err)
	assert.Equal(t, "druid.broker.some.metric", key)
}

func TestPrepareMetricWhenServiceIsMissing(t *testing.T) {
	metricBody := "{\"metric\":\"some/metric\"}"
	metric, _ := jason.NewObjectFromBytes([]byte(metricBody))

	_, err := PrepareMetricKey(metric)
	assert.NotNil(t, err)
}

func TestPrepareMetricWhenMetricIsMissing(t *testing.T) {
	metricBody := "{\"service\":\"druid/broker\"}"
	metric, _ := jason.NewObjectFromBytes([]byte(metricBody))

	_, err := PrepareMetricKey(metric)
	assert.NotNil(t, err)
}

func TestPrepareTagsToReturnTagsWithoutExcludedTags(t *testing.T) {
	metricBody := "{\"timestamp\":\"xxx\",\"name\":\"john\"}"
	metric, _ := jason.NewObjectFromBytes([]byte(metricBody))

	tags := PrepareTags(metric)
	assert.Equal(t, 1, len(tags))
	assert.Equal(t, "name:john", tags[0])
}

func TestIsExcludedTag(t *testing.T) {
	assert.Equal(t, true, IsExcludedTag("timestamp"))
	assert.Equal(t, true, IsExcludedTag("service"))
	assert.Equal(t, true, IsExcludedTag("metric"))
	assert.Equal(t, true, IsExcludedTag("value"))
	assert.Equal(t, true, IsExcludedTag("context"))
	assert.Equal(t, true, IsExcludedTag("id"))
	assert.Equal(t, true, IsExcludedTag("interval"))
	assert.Equal(t, true, IsExcludedTag("taskId"))

	assert.Equal(t, false, IsExcludedTag("feed"))
	assert.Equal(t, false, IsExcludedTag("host"))
	assert.Equal(t, false, IsExcludedTag("version"))
}
