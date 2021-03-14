package main

import (
	"fmt"
	"strconv"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"

	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	count := 0
	mvc.Configure(app.Party("/root"), func(mvcApp *mvc.Application) {
		mvcApp.Router.Use(func(context context.Context) {
			if context.Path() == "/root/test" {
				count++
				fmt.Println("/root/test 请求次数:" + strconv.Itoa(count))
			}
			// 注意我们在中间件函数的最后需要加上context.Next()，该函数会告诉 iris 继续执行下面的 controller，否则代码不会执行到 controller。
			context.Next()

		})
		mvcApp.Handle(new(MyController))

	})
	_ = app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

type MyController struct {
	Ctx iris.Context
}

func (m *MyController) Get() string {
	return "Hey"
}

// 约定大于配置
func (m *MyController) GetBy(id int64) interface{} {
	return map[string]interface{}{"id": id}
}

func (m *MyController) GetHelloWorld() interface{} {
	return map[string]string{"message": "Hello world!"}
}

func (m *MyController) GetHello() interface{} {
	return map[string]string{"message": "Hello Iris!"}
}

func (m *MyController) AnyTest() {
	_, _ = m.Ctx.HTML("<h1>test</h1>")
}
