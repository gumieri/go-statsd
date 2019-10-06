# Go Statsd

A extremely simple Statsd client for Go.

## Usage

```go
import "github.com/gumieri/statsd"

client := statsd.New("localhost", 8125)

err := client.Send(&statsd.Metric{
  Name:  "test", // string
  Value: 1, // float64
  Type:  statsd.Counter, // statsd.Gauge, statsd.Millisecond, statsd.Histogram or statsd.Set
  SampleRate: 0.33, // float64
  Tags: []statsd.Tag{
    {Key: "oneKey", Value: "oneValue"},
    {Key: "twoKey", Value: "twoValue"},
  },
})
```

There is also a `TrySend` which does not return any error:

```go
client.TrySend(&statsd.Metric{
  Name:  "test",
  Value: 1,
  Type:  statsd.Counter,
})
```

## Code Status

[![Go Report Card](https://goreportcard.com/badge/github.com/gumieri/go-statsd)](https://goreportcard.com/report/github.com/gumieri/go-statsd)
[![Build Status](https://travis-ci.org/gumieri/go-statsd.svg?branch=master)](https://travis-ci.org/gumieri/go-statsd)

## License

Note is released under the [MIT License](http://www.opensource.org/licenses/MIT).

