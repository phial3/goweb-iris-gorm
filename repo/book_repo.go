package repo

import (
	"github.com/spf13/cast"
)

import (
	"goweb-iris-gorm/datasource"
	"goweb-iris-gorm/model"
)

type BookRepository interface {
	GetBookList(m map[string]interface{}) (total int, books []model.Book)
	SaveBook(book model.Book) (err error)
	GetBook(id uint) (book model.Book, err error)
	DelBook(id uint) (err error)
}

func NewBookRepository() BookRepository {
	return &bookRepository{}
}

var db = datasource.GetDB()

type bookRepository struct{}

func (n bookRepository) GetBookList(m map[string]interface{}) (total int, books []model.Book) {
	db.Table("book").Count(&total)
	err := db.Limit(cast.ToInt(m["size"])).Offset((cast.ToInt(m["page"]) - 1) * cast.ToInt(m["size"])).Find(&books).Error
	if err != nil {
		panic("select Error")
	}
	return
}

func (n bookRepository) SaveBook(book model.Book) (err error) {
	if book.ID != 0 {
		err := db.Save(&book).Error
		return err
	} else {
		err := db.Create(&book).Error
		return err
	}
}

func (n bookRepository) GetBook(id uint) (book model.Book, err error) {
	err = db.First(&book, id).Error
	return
}

func (n bookRepository) DelBook(id uint) (err error) {
	var book model.Book
	book.ID = id
	err = db.Unscoped().Delete(&book).Error //如果直接Delete是软删除
	return
}
