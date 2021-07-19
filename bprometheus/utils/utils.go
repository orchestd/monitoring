package utils

import (
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/transport"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MetricsHandler(router transport.IRouter){
	if router == nil {
		return
	}
	router.GET("/metrics",PrometheusHandler())
}
func PrometheusHandler()  gin.HandlerFunc {
	h:= promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}