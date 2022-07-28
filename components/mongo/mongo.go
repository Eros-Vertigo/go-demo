package mongo

import (
	"context"
	"demon/common"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var (
	Mdb *mongo.Client
	Ctx context.Context
)

func init() {
	var err error
	dns := fmt.Sprintf("mongodb://%s:%d", common.Config.Mongo.Host, common.Config.Mongo.Port)
	// 设置连接超时
	Ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2))
	defer cancel()

	o := options.Client().ApplyURI(dns)
	o.SetMaxPoolSize(uint64(50))
	// 发起连接
	Mdb, err = mongo.Connect(Ctx, o)
	if err != nil {
		log.Fatalf("mongo.init err : [%v]", err)
	}
	if err = Mdb.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatalf("mongo.init err : [%v]", err)
	}
	log.Infof("Mongo [%s:%d] 初始化完成", common.Config.Mongo.Host, common.Config.Mongo.Port)
}
