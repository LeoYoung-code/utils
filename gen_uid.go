package utils

import (
	"fmt"

	"C"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/yitter/idgenerator-go/idgen"
	"github.com/yitter/idgenerator-go/regworkerid"
)

const (
	workerIdBitLength = 9
)

func init() {
	workerId := genWorkerId()
	log.Infof("worker id: %d", workerId)
	options := idgen.NewIdGeneratorOptions(workerId)
	options.WorkerIdBitLength = workerIdBitLength // 默认值6，限定 WorkerId 最大值为2^6-1，即默认最多支持64个节点。
	options.SeqBitLength = 6                      // 默认值6，限制每毫秒生成的ID个数。若生成速度超过5万个/秒，建议加大 SeqBitLength 到 10。
	idgen.SetIdGenerator(options)
}

func genWorkerId() uint16 {
	address := "127.0.0.1:9001"
	password := ""
	masterName := ""
	addrChar := C.CString(address)
	passChar := C.CString(password)
	masterNameChar := C.CString(masterName)

	id := registerOne(addrChar, passChar, 4, masterNameChar, 10, 1024, 15)
	fmt.Println("注册的WorkerId:", id)
	return uint16(id)
}

// registerOne 注册一个 WorkerId，会先注销所有本机已注册的记录
// address: Redis连接地址，单机模式示例：127.0.0.1:6379，哨兵/集群模式示例：127.0.0.1:26380,127.0.0.1:26381,127.0.0.1:26382
// password: Redis连接密码
// db: Redis指定存储库，示例：1
// sentinelMasterName: Redis 哨兵模式下的服务名称，示例：mymaster，非哨兵模式传入空字符串即可
// minWorkerId: WorkerId 最小值，示例：30
// maxWorkerId: WorkerId 最大值，示例：63
// lifeTimeSeconds: WorkerId缓存时长（秒，3的倍数）
//
//export registerOne
func registerOne(address *C.char, password *C.char, db int, sentinelMasterName *C.char, minWorkerId int32, maxWorkerId int32, lifeTimeSeconds int32) int32 {
	return regworkerid.RegisterOne(regworkerid.RegisterConf{
		Address:         C.GoString(address),
		Password:        C.GoString(password),
		DB:              db,
		MasterName:      C.GoString(sentinelMasterName),
		MinWorkerId:     minWorkerId,
		MaxWorkerId:     maxWorkerId,
		LifeTimeSeconds: lifeTimeSeconds,
	})
}

func GenUid() int64 {
	return idgen.NextId()
}
