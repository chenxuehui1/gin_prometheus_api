package main
import (
	"github.com/chenxuehui1/gin_prometheus_api/collector"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

/*func init() {
	rand.Seed(time.Now().Unix())
}*/

func main() {
	r := gin.Default()
	r.Use(collector.PromMiddleware(nil))

	r.GET("/metrics", collector.PromHandler(promhttp.Handler()))

	r.GET("/", func(c *gin.Context) {
		collector.TimeSleep()
		c.JSON(http.StatusOK, gin.H{
			"message": "home",
		})
	})

	r.GET("/index", func(c *gin.Context) {
		collector.TimeSleep()
		c.JSON(http.StatusOK, gin.H{
			"message": "index",
		})
	})

	r.GET("/hello", func(c *gin.Context) {
		collector.TimeSleep()
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})


	r.GET("/forbidden", func(c *gin.Context) {
		collector.TimeSleep()
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "forbidden",
		})
	})

	r.GET("/badreq", func(c *gin.Context) {
		collector.TimeSleep()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "badreq",
		})
	})

	if err := r.Run(":4433"); err != nil {
		log.Fatalf("run server error: %v", err)
	}
}
