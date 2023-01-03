package main

import (
	_ "gin/common"
	"gin/config"
	"github.com/gin-gonic/gin"
	"html/template"

	"net/http"
	"strconv"
)

func main() {
	//person := new(provider.Person)
	//
	//person.CreatePerson()
	//fmt.Printf("person:%v \n", &person)
	//	1.创建路由
	r := gin.Default()

	r.StaticFS("./stc", http.Dir("./static"))
	//r.Static("/controller", "controller")
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.LoadHTMLGlob("views/*")
	r.GET("/", func(c *gin.Context) {
		//c.String(http.StatusOK, `Hello Gin World!`)
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	//routes.InitUserRoutes(r)
	// 3.监听端口， 默认是8080
	// Run(里面不指定端口号就是8080)
	err := r.Run(config.GlobalConfig.ServerHost + ":" + strconv.Itoa(config.GlobalConfig.ServerPort))
	if err != nil {

		return
	}

}
