//docker run -d -p 3000:3000 --name grafana --mount src=grafana-storage,dst=/var/lib/grafana -e "GF_INSTALL_PLUGINS=grafana-simple-json-datasource" grafana/grafana
package main

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/christophberger/grada"
)

func newFakeDataFunc(max int, volatility float64, responseTime int) func() float64 {
	value := rand.Float64()
	return func() float64 {
		time.Sleep(time.Duration(responseTime) * time.Millisecond) // simulate response time
		rnd := 2 * (rand.Float64() - 0.5)
		change := volatility * rnd
		change += (0.5 - value) * 0.1
		value += change
		return math.Max(0, value*float64(max))
	}
}

func main() {
	dash := grada.GetDashboard()

	CPU1metric, err := dash.CreateMetric("CPU1", 5*time.Minute, time.Second)
	if err != nil {
		log.Fatalln(err)
	}

	CPU2metric, err := dash.CreateMetricWithBufSize("CPU2", 300)
	if err != nil {
		log.Fatalln(err)
	}

	CPU1stats := newFakeDataFunc(100, 0.2, 1000)
	CPU2stats := newFakeDataFunc(100, 0.1, 1000)

	trading := func(metric *grada.Metric, dataFunc func() float64) {
		for {
			metric.Add(dataFunc())
		}
	}

	go trading(CPU1metric, CPU1stats)
	go trading(CPU2metric, CPU2stats)

	select {}
}
