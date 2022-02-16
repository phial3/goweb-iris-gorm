package controller

import (
	"github.com/kataras/iris/v12"
	"log"
)

import (
	"goweb-iris-gorm/model"
	"goweb-iris-gorm/service"
)

type UserController struct {
	Ctx     iris.Context
	Service service.UserService
}

func NewUserController() *UserController {
	return &UserController{Service: service.NewUserServices()}
}

func (g *UserController) PostLogin() model.Result {
	var m map[string]string
	err := g.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}
	result := g.Service.Login(m)
	return result
}

func (g *UserController) PostSave() (result model.Result) {
	var user model.User
	if err := g.Ctx.ReadJSON(&user); err != nil {
		log.Println(err)
		result.Msg = "数据错误"
		return
	}
	return g.Service.Save(user)
}
