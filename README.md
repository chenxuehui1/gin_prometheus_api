<h1 align="center">📡 ginprom</h1>
<p align="center">
    <em>Prometheus metrics exporter for Gin.Inspired by <a href="https://github.com/chenxuehui1/gin_prometheus_api">Depado/ginprom.</a></em>
</p>

### 🔰 Installation

```shell
$ go get -u https://github.com/chenxuehui1/gin_prometheus_api
```

### 📝 Usage

It's easy to get started with ginprom, only a few lines of code needed.

```golang
import (
	"github.com/chenxuehui1/gin_prometheus_api/collector"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    r := gin.Default()
	// use prometheus metrics exporter middleware.
	//
	// collector.PromMiddleware() expects a collector.PromOpts{} poniter.
	// It was used for filtering labels with regex. `nil` will pass every requests.
	//
	// collector promethues-labels: 
	//   `status`, `endpoint`, `method`
	//
	// for example:
	// 1). I want not to record the 404 status request. That's easy for it.
	// collector.PromMiddleware(&collector.PromOpts{ExcludeRegexStatus: "404"})
	//
	// 2). And I wish ignore endpoint start with `/prefix`.
	// collector.PromMiddleware(&collector.PromOpts{ExcludeRegexEndpoint: "^/prefix"})
	r.Use(collector.PromMiddleware(nil))

    // register the `/metrics` route.
	r.GET("/metrics", collector.PromHandler(promhttp.Handler()))

    // your working routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "home"})
    })
}
```

### 🎉 Metrics

Details about exposed Prometheus metrics.

| Name | Type | Exposed Information |
| ---- | ---- | ---------------------|
| service_uptime						| Counter	| HTTP service uptime. |
| http_request_count_total		        | Counter	| Total number of HTTP requests made. |
| http_request_duration_seconds         | Histogram | HTTP request latencies in seconds. |
| http_request_size_bytes 		        | Summary	| HTTP request sizes in bytes. |
| http_response_size_bytes 		        | Summary	|HTTP request sizes in bytes. |


### 📊 Grafana

Although Promethues offers a simple dashboard, Grafana is clearly a better choice. [Grafana configuration](./ginprom-service.json).

![](https://user-images.githubusercontent.com/19553554/65812184-19a5a000-e1f6-11e9-8881-e0c260196bc9.png)


### 📃 LICENSE

MIT [©chenxuehui1](https://github.com/chenxuehui1)
