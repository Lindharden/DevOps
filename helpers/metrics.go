package helpers

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	TweetsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "minitwit_number_of_tweets",
		Help: "The total number of tweets made",
	})
)

var (
	RequestResponseTime = promauto.NewHistogram(prometheus.HistogramOpts{
		Name: "minitwit_response_time",
		Help: "The response times measured for each request",
	})
)

var (
	ErrorCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "minitwit_error_codes",
		Help: "The count of codes in the error range",
	})
)
