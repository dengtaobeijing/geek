package data

import (
	"network/app/sbi/job/internal/conf"

	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewJobRepo, NewKafkaConsumer)

type Data struct {
	db  *gorm.DB
	kc  sarama.Consumer
	log *log.Helper
}

func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "sbi-job/data/gorm"))
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
	if err != nil {
		log.Fatal("failed opening connection to mysql: %v", err)
	}
	if err := db.AutoMigrate(&Integrating{}); err != nil {
		log.Fatal(err)
	}

	return db
}

// NewData .
func NewData(db *gorm.DB, kc sarama.Consumer, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "sbi-job/data"))

	d := &Data{
		db:  db,
		kc:  kc,
		log: log,
	}
	return d, func() {

	}, nil
}

func NewKafkaConsumer(conf *conf.Data) sarama.Consumer {
	c := sarama.NewConfig()
	p, err := sarama.NewConsumer(conf.Kafka.Addr, c)
	if err != nil {
		panic(err)
	}
	return p
}
