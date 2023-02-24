package goroutine

import (
	redis2 "demon/components/redis"
	"fmt"
)

const (
	ListDetail          = "list:detail"
	HashDetail          = "hash:rad_online:%s"
	HashProxy           = "hash:proxy:%s"
	KeyOnlineCheckout   = "key:online_checkout:%s"
	HashOnlineCheckout  = "hash:online_checkout:%s"
	HashRadOnlineUpdate = "hash:rad_online:update:%s"
)

var hashChan = make(chan map[string]string, 1000)

func init() {
	quit := make(chan bool) // 退出开关
	redis := redis2.Rdb
	temp, err := redis.LRange(redis.Context(), ListDetail, 0, -1).Result()
	if err != nil {
		checkError(err, "读取队列发生错误")
		return
	}

	if len(temp) == 0 {
		fmt.Println("队列长度为0，程序退出")
	}

	for _, id := range temp {
		go func(hashId string) {
			getHashDetail(hashId)
		}(id)
	}

	for {
		select {
		case detail := <-hashChan:
			fmt.Println("detail = ", detail)
		case <-quit:
			return
		}
	}

}

/**
获取在线详情
*/
func getHashDetail(hashId string) {
	val, err := redis2.Rdb.HGetAll(redis2.Rdb.Context(), fmt.Sprintf(HashDetail, hashId)).Result()
	if err != nil {
		checkError(err, "获取在线信息时发生错误")
		return
	}
	if val == nil {
		checkError(err, "未获取到在线信息，去获取代理hash")
		go func() {
			getHashProxy(hashId)
		}()
		go func() {
			handleDrop(hashId)
		}()
	}
	fmt.Println("入channel", hashId)
	hashChan <- val
}

/**
获取代理信息
*/
func getHashProxy(hashId string) {
	val, err := redis2.Rdb.HGetAll(redis2.Rdb.Context(), fmt.Sprintf(HashProxy, hashId)).Result()
	if err != nil {
		checkError(err, "获取代理信息时发生错误")
		return
	}
	if val == nil {
		checkError(err, "未获取到代理信息,删除用户在线key")
		redis2.Rdb.Del(redis2.Rdb.Context(), fmt.Sprintf(KeyOnlineCheckout, hashId))
		redis2.Rdb.Del(redis2.Rdb.Context(), fmt.Sprintf(HashDetail, hashId))
		redis2.Rdb.Del(redis2.Rdb.Context(), fmt.Sprintf(HashProxy, hashId))
		return
	}
	hashChan <- val
}

func handleDrop(hashId string) {
	val, err := redis2.Rdb.HGetAll(redis2.Rdb.Context(), fmt.Sprintf(HashOnlineCheckout, hashId)).Result()
	if err != nil {
		checkError(err, "查询在线结算信息时发生错误")
		return
	}
	if val != nil {
		redis2.Rdb.Del(redis2.Rdb.Context(), fmt.Sprintf(HashOnlineCheckout, hashId))
		redis2.Rdb.HSet(redis2.Rdb.Context(), fmt.Sprintf(HashDetail, hashId), "drop_reason", "0") // 恢复下线原因
	} else {
		redis2.Rdb.Del(redis2.Rdb.Context(), fmt.Sprintf(HashDetail, hashId))          //删除hash
		redis2.Rdb.Del(redis2.Rdb.Context(), fmt.Sprintf(HashProxy, hashId))           // 删除代理
		redis2.Rdb.Del(redis2.Rdb.Context(), fmt.Sprintf(HashRadOnlineUpdate, hashId)) // 删除UPDATE流量缓存
	}
}
func checkError(err error, msg string) {
	fmt.Printf("%s:%v", msg, err.Error())
}
