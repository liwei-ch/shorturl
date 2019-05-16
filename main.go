package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"shorturl/url"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	engine.LoadHTMLGlob("resources/*")

	// 直接使用/generate/https://www.baidu.com会报错，还是使用参数吧
	engine.GET("/generate", func(c *gin.Context) {
		param_url := c.Query("url")
		surl := url.GenerateUrl(param_url)
		c.JSON(http.StatusOK, gin.H{
			"url":  param_url,
			"surl": surl,
		})
	})
	// 这个可以这么写，不过要和generate里面的链接对应
	engine.GET("/to/:surl", func(c *gin.Context) {
		surl := c.Param("surl")
		url, ok := url.ParseSurl(surl)
		if ok {
			c.Redirect(http.StatusMovedPermanently, url)
		} else {
			c.HTML(http.StatusNotFound, "404.tmpl", nil)
		}
	})

	//go func() {
	//	ticker := time.NewTicker(time.Second * 5)
	//	for {
	//		select {
	//		case <-ticker.C:
	//			//按时输出goroutine数目
	//			fmt.Println(runtime.NumGoroutine())
	//		}
	//	}
	//}()

	engine.Run(":8080")
}
