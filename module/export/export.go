package export

import (
	"demon/components/redis"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

func init() {
	fmt.Println("export process is start")
}

const QueueName = "list:rad_online"

func rangeList() {
	// 循环读取队列
	for {
		result, err := redis.Rdb.BRPop(redis.Rdb.Context(), time.Second, QueueName).Result()
		if err != nil {
			log.Error("Error getting item from queue:", err)
			continue
		}

		// 获取hash
		hashKey := result[1]
		hash, err := redis.Rdb.HGetAll(redis.Rdb.Context(), hashKey).Result()
		if err != nil {
			log.Error("Error getting hash:", err)
			continue
		}

		// 检查产品是否存在，如果不存在进行导出
		if _, ok := hash["field"]; !ok {
			exportData, err := json.Marshal(hash)
			if err != nil {
				log.Error("Error exporting data:", err)
				continue
			}
			fmt.Println("Exporting data:", string(exportData))
		}
	}
}
