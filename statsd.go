package statsd

import (
	"fmt"
	"net"
	"strconv"
)

// Client type defines the relevant properties of a StatsD connection.
type Client struct {
	Host string
	Port int
	conn net.Conn
}

// New factory method to initialize udp connection
//
// Usage:
//
//     import "statsd"
//     client := statsd.New("localhost", 8125)
func New(host string, port int) *Client {
	client := Client{Host: host, Port: port}
	client.Open()
	return &client
}

// Open udp connection, called by default client factory
func (client *Client) Open() (err error) {
	conn, err := net.Dial("udp", client.Host+":"+strconv.Itoa(client.Port))

	if err != nil {
		return
	}

	client.conn = conn

	return
}

// Close method to close udp connection
func (client *Client) Close() {
	if client == nil || client.conn == nil {
		return
	}

	client.conn.Close()
}

// Send data to statsd
func (client *Client) Send(metric *Metric) (err error) {
	_, err = fmt.Fprintf(client.conn, metric.String())
	return
}

// TrySend Send the data but if it fail it will do nothing
func (client *Client) TrySend(metric *Metric) {
	if client == nil || client.conn == nil || metric == nil {
		return
	}

	client.Send(metric)
}
