package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"regexp"
	"shorturl/url"
)

var (
	_URL_PAT = regexp.MustCompile("https?://.*")
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	var server = url.NewUrlServer("<redis server addr>", "<password>",
		"<url hash key>", "<sql server>",
		"<generate return value prefix>", 0)

	engine := gin.Default()
	engine.LoadHTMLGlob("resources/*")

	// 直接使用/generate/https://www.baidu.com会报错，还是使用参数吧
	// 改为POST
	engine.POST("/generate", func(c *gin.Context) {
		data, err := ioutil.ReadAll(c.Request.Body)

		if err != nil {
			panic(err)
		}
		param_url := string(data)
		// 需要检测url是否有效
		if !_URL_PAT.MatchString(param_url) {
			c.JSON(http.StatusBadRequest, gin.H{
				"url":  param_url,
				"surl": "bad request",
			})
			return
		}
		surl := server.GenerateUrl(param_url)
		c.JSON(http.StatusOK, gin.H{
			"url":  param_url,
			"surl": surl,
		})
	})

	engine.GET("/generate", func(c *gin.Context) {
		c.String(http.StatusBadRequest, "please use POST")
	})
	// 这个可以这么写，不过要和generate里面的链接对应
	engine.GET("/to/:surl", func(c *gin.Context) {
		surl := c.Param("surl")
		redirect_url, ok := server.ParseSurl(surl)
		if ok {
			c.Redirect(http.StatusMovedPermanently, redirect_url)
		} else {
			c.HTML(http.StatusNotFound, "404.tmpl", nil)
		}
	})

	engine.Run(":8080")
}
