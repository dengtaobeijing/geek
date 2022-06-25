package data

import (
	"network/app/sbi/task/internal/conf"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewTaskRepo)

type Data struct {
	db  *gorm.DB
	log *log.Helper
}

func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "sbi-job/data/gorm"))
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
	if err != nil {
		log.Fatal("failed opening connection to mysql: %v", err)
	}
	if err := db.AutoMigrate(&IntegratingCount{}); err != nil {
		log.Fatal(err)
	}

	return db
}

// NewData .
func NewData(db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "sbi-task/data"))

	d := &Data{
		db:  db,
		log: log,
	}
	return d, func() {

	}, nil
}
