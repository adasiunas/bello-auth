package business

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Factory struct {
	db *gorm.DB
}

func NewFactory() *Factory {
	//dialect := os.Getenv("DB_DIALECT")
	sess, err := gorm.Open("mysql", "root:nesakysiu@/bello_auth?parseTime=true")
	if err != nil {
		panic(err)
	}

	sess.DB().SetConnMaxLifetime(1 * time.Hour)
	sess.DB().SetMaxIdleConns(5)
	sess.DB().SetMaxOpenConns(15)
	sess.SingularTable(true)

	return &Factory{sess}
}

func (f *Factory) DB() *gorm.DB {
	return f.db
}