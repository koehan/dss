package v1

import "github.com/gin-gonic/gin"

func Register(v1 *gin.RouterGroup) {
	// this rule defines the scanning host and port
	rule := v1.Group("/rule")
	{
		rule.POST("", Rule.Post)     // add rule
		rule.GET("", Rule.Get)       // list rule
		rule.PUT("", Rule.Put)       // modify rule
		rule.DELETE("", Rule.Delete) // delete rule
	}
	// task scheduler
	task := v1.Group("/task")
	{
		task.POST("", Task.Post) // push task to redis
	}
	//port scan result
	port := v1.Group("/port")
	{
		port.GET("", Port.Get)           // list port scan result
		port.GET("trend", Port.Trend)    // last 7 days scan result trend
		port.GET("Stats", Port.Stats)    // last 7 days scan result stats
		port.GET("remind", Port.Remind)  // send notification if new port open  (schedule)
		port.DELETE("clear", Port.Clear) // save last 7 days result (schedule)
	}
}
