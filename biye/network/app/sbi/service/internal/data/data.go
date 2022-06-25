package data

import (
	"network/app/sbi/service/internal/conf"

	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewRedis, NewSbiRepo, NewKafkaProducer)

type Data struct {
	db  *gorm.DB
	rdb *redis.Client
	kp  sarama.AsyncProducer
	log *log.Helper
}

func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "sbi-service/data/gorm"))
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
	if err != nil {
		log.Fatal("failed opening connection to mysql: %v", err)
	}
	if err := db.AutoMigrate(&Sbi{}, &Orders{}); err != nil {
		log.Fatal(err)
	}

	return db
}

func NewRedis(conf *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

// NewData .
func NewData(db *gorm.DB, rdb *redis.Client, kp sarama.AsyncProducer, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "sbi-service/data"))

	d := &Data{
		db:  db,
		rdb: rdb,
		kp:  kp,
		log: log,
	}
	return d, func() {

	}, nil
}

func NewKafkaProducer(conf *conf.Data) sarama.AsyncProducer {
	c := sarama.NewConfig()
	p, err := sarama.NewAsyncProducer(conf.Kafka.Addr, c)
	if err != nil {
		panic(err)
	}
	return p
}
