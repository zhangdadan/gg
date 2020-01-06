package unique

import (
	"crypto/md5"
	"encoding/hex"
	"sync"
	"time"
)

var ai int64
var aiMutex sync.Mutex // 定义互斥锁变量 mutex

/*
得到自增
*/
func GetAI() int64 {
	aiMutex.Lock()
	defer aiMutex.Unlock()
	if ai > 100000000000 {
		ai = 0
	}
	ai++
	return ai
}

/*
得到MD5
*/
func GetMD5() string {
	m := md5.Sum([]byte(string(time.Now().UnixNano()) + string(GetAI())))
	return hex.EncodeToString(m[:])
}
