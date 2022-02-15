package route

import (
	"goweb-iris-gorm/controller"
	"net/http"
)

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

import (
	"goweb-iris-gorm/constant"
	"goweb-iris-gorm/middleware"
)

func InitRouter(app *iris.Application) {
	app.Use(CrossAccess)

	mvc.New(app.Party(constant.BathUrl + constant.User)).Handle(controller.NewUserController())
	app.Use(middleware.GetJWT().Serve) // jwt

	mvc.New(app.Party(constant.BathUrl + constant.Book)).Handle(controller.NewBookController())
}

func CrossAccess11(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
func CrossAccess(ctx iris.Context) {
	ctx.ResponseWriter().Header().Add("Access-Control-Allow-Origin", "*")
}
