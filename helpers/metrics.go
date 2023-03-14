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
