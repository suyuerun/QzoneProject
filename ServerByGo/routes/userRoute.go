package routes

import (
	"gin/common/database/provider"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Person = new(provider.Person)

// 定义接收数据的结构体
type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	Username string `form:"username" json:"username" uri:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

// 定义接收数据的结构体
type Register struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	Username string `form:"username" json:"username" uri:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
	Country  string `form:"country" json:"country" uri:"country" xml:"country" binding:"required"`
	Telcode  int    `form:"telcode" json:"telcode" uri:"telcode" xml:"telcode" binding:"required"`
	Email    string `form:"email" json:"email" uri:"email" xml:"email" binding:"required"`
	City     string `form:"city" json:"city" uri:"city" xml:"city" binding:"required"`
	Sex      string `form:"sex" json:"sex" uri:"sex" xml:"sex" binding:"required"`
}

func InitUserRoutes(r *gin.Engine) {

	//	2.绑定路由规则，执行的函数
	// gin.Context, 封装了request 和 response
	r.GET("/", func(c *gin.Context) {
		//c.String(http.StatusOK, `Hello Gin World!`)
		c.HTML(http.StatusOK, "user/index.html", gin.H{"title": "usrRoute"})
	})
	r.POST("/login", userLogin)
	r.POST("/register", userRegister)
	//r.PUT("/xxxput")
}

func userLogin(c *gin.Context) {
	// 声明接收的变量
	var form Login
	// Bind()默认解析并绑定form格式
	// 根据请求头中content-type自动推断
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 从数据库里查找玩家

	person := Person.GetPerson(form.Username)
	// 判断用户名密码是否正确
	if form.Username != person.Username || form.Password != person.Password {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})

}

func userRegister(c *gin.Context) {
	// 声明接收的变量
	var form Register
	// Bind()默认解析并绑定form格式
	// 根据请求头中content-type自动推断
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 从数据库里查找玩家

	person := Person.GetPerson(form.Username)
	// 判断用户名 是否已被占用
	if form.Username != person.Username || form.Email != person.Email {
		c.JSON(http.StatusBadRequest, gin.H{"status": "303"})
		return
	}
	personMap := map[string]interface{}{
		"username": form.Username,
		"password": form.Password,
		"sex":      form.Sex,
		"email":    form.Email,
		"country":  form.Country,
		"city":     form.City,
		"telcode":  form.Telcode,
	}
	insertRet := Person.CreatePerson(personMap)
	if insertRet == "success" {
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	} else {
		c.JSON(http.StatusOK, gin.H{"error": insertRet})
	}

}
