package controllers_test

import (
	"testing"

	"github.com/an-0305/study-go-api/controllers"
	"github.com/an-0305/study-go-api/controllers/testdata"
	_ "github.com/go-sql-driver/mysql"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)

	m.Run()
}
