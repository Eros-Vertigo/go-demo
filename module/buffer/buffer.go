package buffer

import (
	"bytes"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.Info("module buffer")
	//compare([]byte{016}, []byte{16})
	//trim([]byte{'%', 'y', 't'})
	split([]byte{'l', 'y', 't', 'l', 't', 'y'})
}

// 字节比较
// 相等返回 0;
// a > b 返回 1;
// a < b 返回 -1;
func compare(a, b []byte) {
	res := bytes.Compare(a, b)
	if res == 0 {
		log.Infof("a (v = %v t = %T) == b (v = %v t = %T)", a, a, b, b)
	} else if res > 0 {
		log.Infof("a (v = %v t = %T) > b (v = %v t = %T)", a, a, b, b)
	} else {
		log.Infof("a (v = %v t = %T) < b (v = %v t = %T)", a, a, b, b)
	}
}

// ✂️切片
func trim(b []byte) {
	res := bytes.Trim(b, "%")
	log.Infof("bytes trim : %s", res)
}

// 分割字符
func split(b []byte) {
	res := bytes.Split(b, []byte{'l'})
	log.Infof("bytes split : %s", res)
}
