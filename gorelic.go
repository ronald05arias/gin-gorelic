package gorelic

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yvasiyarov/go-metrics"
	"github.com/yvasiyarov/gorelic"
)

var agent *gorelic.Agent

//Metric defines an interface for monitoring metrics
type Metric interface {
	GetName() string
	GetUnits() string
	GetValue() (float64, error)
}

//addCustomMetric allows the addition of a metric for monitoring
func addCustomMetric(metric Metric) {
	if agent != nil {
		agent.AddCustomMetric(metric)
	}
}

//addCustomMetrics allows the addition of new metrics for monitoring
func addCustomMetrics(metrics []Metric) {
	for _, metric := range metrics {
		addCustomMetric(metric)
	}
}

// Handler for the agent
func Handler(c *gin.Context) {
	startTime := time.Now()
	c.Next()
	if agent != nil {
		agent.HTTPTimer.UpdateSince(startTime)
	}
}

// InitNewrelicAgent creates the new relic agent
func InitNewrelicAgent(license string, appname string, verbose bool, customMetrics []Metric) (gin.HandlerFunc, error) {
	if license == "" {
		return nil, fmt.Errorf("Please specify a NewRelic license")
	}

	agent = gorelic.NewAgent()
	agent.NewrelicLicense = license
	agent.HTTPTimer = metrics.NewTimer()
	agent.CollectHTTPStat = true
	agent.Verbose = verbose
	agent.NewrelicName = appname

	if customMetrics != nil {
		addCustomMetrics(customMetrics)
	}

	return Handler, agent.Run()
}
