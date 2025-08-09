package app

import (
	"database/sql"
	"gobackend/core/configuration"
	"gobackend/infra/cache"
	"gobackend/infra/db"
	"gobackend/infra/msgq"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Dependencies struct {
	DB        *gorm.DB
	Redis     *redis.Client
	RabbitMQ  *amqp.Connection
	SqlNative *sql.DB
}

func InitDependencies(cfg configuration.Config) *Dependencies {
	postgres, err := db.NewPostgres(cfg)
	if err != nil {
		log.Fatalf("Failed to init PostgreSQL: %v", err)
	}

	rds := cache.NewRedis(cfg)

	rmq, err := msgq.NewRabbitMQ(cfg)
	if err != nil {
		log.Fatalf("Failed to init RabbitMQ: %v", err)
	}

	return &Dependencies{
		DB:       postgres,
		Redis:    rds,
		RabbitMQ: rmq,
	}
}
