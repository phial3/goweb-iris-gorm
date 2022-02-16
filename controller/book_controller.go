package controller

import (
	"github.com/kataras/iris/v12"
	"log"
)

import (
	"github.com/spf13/cast"
)

import (
	"goweb-iris-gorm/model"
	"goweb-iris-gorm/service"
)

type BookController struct {
	Ctx     iris.Context
	Service service.BookService
}

func NewBookController() *BookController {
	return &BookController{Service: service.NewBookService()}
}

func (g *BookController) PostList() (result model.Result) {
	var m map[string]interface{}
	err := g.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}
	if m["page"] == "" || m["page"] == nil {
		result.Code = -1
		result.Msg = "参数缺失 page"
		return
	}
	if cast.ToUint(m["page"]) == 0 {
		result.Code = -1
		result.Msg = "参数错误 page"
		return
	}
	if m["size"] == "" || m["size"] == nil {
		result.Code = -1
		result.Msg = "参数缺失 size"
		return
	}
	if cast.ToUint(m["size"]) == 0 {
		result.Code = -1
		result.Msg = "参数错误 size"
		return
	}
	return g.Service.GetBookList(m)
}

func (g *BookController) PostSave() (result model.Result) {
	var book model.Book
	if err := g.Ctx.ReadJSON(&book); err != nil {
		log.Println(err)
		result.Msg = "数据错误"
		return
	}
	return g.Service.SaveBook(book)
}

func (g *BookController) PostGet() (result model.Result) {
	var m map[string]interface{}
	err := g.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}
	if m["id"] == "" || m["id"] == nil {
		result.Code = -1
		result.Msg = "参数缺失 id"
		return
	}
	if cast.ToUint(m["id"]) == 0 {
		result.Code = -1
		result.Msg = "参数错误 id"
		return
	}
	return g.Service.GetBook(cast.ToUint(m["id"]))
}

func (g *BookController) PostDel() (result model.Result) {
	var m map[string]interface{}
	err := g.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}
	if m["id"] == "" || m["id"] == nil {
		result.Code = -1
		result.Msg = "参数缺失 id"
		return
	}
	if cast.ToUint(m["id"]) == 0 {
		result.Code = -1
		result.Msg = "参数错误 id"
		return
	}
	return g.Service.DelBook(cast.ToUint(m["id"]))
}
