package data

import (
	"context"
	"kratos-realworld/internal/conf"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewMysql, NewMongoDB, NewRedis, NewUserRepo, NewProfileRepo)

// Data .
type Data struct {
	mysql   *gorm.DB
	mongodb *mongo.Client
	redis   *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, mysql *gorm.DB, mongodb *mongo.Client, redis *redis.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{mysql: mysql, mongodb: mongodb, redis: redis}, cleanup, nil
}

func NewMongoDB(c *conf.Data) *mongo.Client {
	var uri = "mongodb://root:a123456@10.0.10.13:23301,10.0.10.13:23302," +
		"10.0.10.13:23303/?maxPoolSize=100&w=majority&authSource=admin"
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic("failed to connect mongodb database")
	}
	return client
}

func NewRedis(c *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.0.10.13:23200",
		Password: "a123456", // no password set
		DB:       0,         // use default DB
	})
	return rdb
}

func NewMysql(c *conf.Data) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.Database.GetDsn()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to connect mysql database")
	}
	InitDB(db)
	return db
}

func InitDB(db *gorm.DB) {
	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}
}
