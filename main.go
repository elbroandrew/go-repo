package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", imageLoadHandler)
	if err := router.Run(); err != nil {
		panic(err)
	}
}

/*
	func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"title": "Main website",
			})
		}
*/
func imageLoadHandler(c *gin.Context) {
	//content := "Download file"
	//fileName := "hello.txt"
	//c.Header("Content-Disposition", "attachment; filename="+fileName)
	//c.Header("Content-Type", "application/text/plain")
	//c.Header("Accept-Length", fmt.Sprintf("%d", len(content)))
	//_, err := c.Writer.Write([]byte(content))
	//if err != nil {
	//	return
	//}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website",
	}) //невозможно отрендерирть хтмл темплейт и отправить файл, надло или то, или то сделать.
	fileName := "hello.txt"
	content := "Download file"
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/text/plain")
	c.Header("Accept-Length", fmt.Sprintf("%d", len(content)))
	c.Writer.Write([]byte(content))
}
