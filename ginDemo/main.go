package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var (
	DB *gorm.DB
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMySQL() (err error) {
	dsn := "root:12345abc@tcp(127.0.0.1:3306)/temp_db?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func main() {
	//创建数据库连接

	err := initMySQL()
	if err != nil {
		panic(err)
	}
	defer DB.Close() //程序退出，关闭连接

	//模型绑定
	//会自动在数据库中创建一个表，叫todos，不用我们自己手动创建
	DB.AutoMigrate(&Todo{})

	r := gin.Default() //返回默认的路由引擎
	//告诉gin去哪里找静态文件
	r.Static("/static", "static")
	//告诉gin去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	//v1
	v1Group := r.Group("/v1")
	{
		//添加
		v1Group.POST("/todo", func(ctx *gin.Context) {
			//
			var todo Todo
			ctx.BindJSON(&todo)
			//fmt.Printf("插入的数据：%#v", todo)
			//存入数据库
			err2 := DB.Create(&todo).Error
			if err2 != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err2.Error()})
			} else {
				ctx.JSON(http.StatusOK, todo) //返回todo
			}
			//返回响应
		})

		//查看所有待办
		v1Group.GET("/todo", func(ctx *gin.Context) {
			//查询表里所有事项，用切片
			var todoList []Todo
			err2 := DB.Find(&todoList).Error
			if err2 != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err2.Error})
			} else {
				ctx.JSON(http.StatusOK, todoList)
			}
		})

		//查看某一个//无
		v1Group.GET("/todo/:id", func(ctx *gin.Context) {

		})

		//删掉某一个事项
		v1Group.DELETE("todo/:id", func(ctx *gin.Context) {
			id, ok := ctx.Params.Get("id")
			if len(id) == 0 || !ok {
				ctx.JSON(http.StatusOK, gin.H{"error": "无效的id"})
				return
			}
			err2 := DB.Where("id=?", id).Delete(Todo{}).Error
			if err2 != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err2.Error})
			} else {
				ctx.JSON(http.StatusOK, gin.H{id: "已删除"})
			}
		})

		//修改（完成）某一个代办事项
		v1Group.PUT("/todo/:id", func(ctx *gin.Context) {
			id, ok := ctx.Params.Get("id")
			if !ok || len(id) == 0 {
				ctx.JSON(http.StatusOK, gin.H{"error": "无效的id"})
				return
			}
			var todo Todo
			err2 := DB.Where("id = ?", id).First(&todo).Error //此时todoo是从db查出来的id=？的记录
			if err2 != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err2.Error()})
				return
			}
			ctx.BindJSON(&todo) //这里是解析请求中的数据，绑定到todoo中，此时todoo中的记录是前端传过来的那份了
			//DB.Save() 操作的行为是“存在即更新，不存在即创建”
			err2 = DB.Save(&todo).Error
			if err2 != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err2.Error()})
			} else {
				ctx.JSON(http.StatusOK, todo)
			}

		})

	}
	r.Run(":9090")

}
