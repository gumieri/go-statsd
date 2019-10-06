package statsd

import "strconv"

// Type of the metric
type Type string

// Counter type of a metric
const Counter Type = "c"

// Gauge type of a metric
const Gauge Type = "g"

// Millisecond time type of a metric
const Millisecond Type = "ms"

// Histogram type of a metric
const Histogram Type = "h"

// Set type of a metric
const Set Type = "s"

// Tag property of a metric
type Tag struct {
	Key   string
	Value string
}

// Metric is a single registry to be sent to statsd
type Metric struct {
	Name       string
	Value      float64
	Type       Type
	SampleRate float64
	Tags       []Tag
}

func (m *Metric) String() (data string) {
	data = m.Name
	data = data + ":" + strconv.FormatFloat(m.Value, 'f', -1, 64)
	data = data + "|" + string(m.Type)

	if m.SampleRate > 0 && m.SampleRate < 1 {
		data = data + "|@" + strconv.FormatFloat(m.SampleRate, 'f', -1, 64)
	}

	for i, tag := range m.Tags {
		if i == 0 {
			data = data + "|#"
		} else {
			data = data + ","
		}

		data = data + tag.Key + ":" + tag.Value
	}

	return
}
