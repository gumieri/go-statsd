package statsd

import (
	"testing"
)

func TestSimpleCount(t *testing.T) {
	result := (&Metric{
		Name:  "test",
		Value: 1,
		Type:  Counter,
	}).String()

	if expected := "test:1|c"; expected != result {
		t.Errorf("metric string is different.\nExpected: \"%s\".\n     Got: \"%s\"", expected, result)
	}
}

func TestCountWithSample(t *testing.T) {
	result := (&Metric{
		Name:       "test",
		Value:      1.11111,
		Type:       Counter,
		SampleRate: 0.33,
	}).String()

	if expected := "test:1.11111|c|@0.33"; expected != result {
		t.Errorf("metric string is different.\nExpected: \"%s\".\n     Got: \"%s\"", expected, result)
	}
}

func TestCountWithTag(t *testing.T) {
	result := (&Metric{
		Name:       "test",
		Value:      1.11111,
		Type:       Counter,
		SampleRate: 0.33,
		Tags: []Tag{
			{
				Key:   "oneKey",
				Value: "oneValue",
			},
		},
	}).String()

	if expected := "test:1.11111|c|@0.33|#oneKey:oneValue"; expected != result {
		t.Errorf("metric string is different.\nExpected: \"%s\".\n     Got: \"%s\"", expected, result)
	}
}

func TestCountWithMultipleTags(t *testing.T) {
	result := (&Metric{
		Name:  "test",
		Value: 1,
		Type:  Counter,
		Tags: []Tag{
			{
				Key:   "oneKey",
				Value: "oneValue",
			},
			{
				Key:   "twoKey",
				Value: "twoValue",
			},
		},
	}).String()

	if expected := "test:1|c|#oneKey:oneValue,twoKey:twoValue"; expected != result {
		t.Errorf("metric string is different.\nExpected: \"%s\".\n     Got: \"%s\"", expected, result)
	}
}
