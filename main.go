package main

import (
	"flag"
)

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

import (
	"goweb-iris-gorm/conf"
	"goweb-iris-gorm/router"
)

/**
缺少session实现
*/
func main() {
	flag.Parse()
	app := newApp()
	router.InitRouter(app)
	err := app.Run(iris.Addr(":"+conf.SysConfig.Port), iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		panic(err)
	}
}

func newApp() *iris.Application {
	app := iris.New()
	app.Configure(iris.WithOptimizations)
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})
	app.Use(crs)
	app.AllowMethods(iris.MethodOptions)
	return app
}
