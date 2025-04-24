package utils

import (
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
	conf := regworkerid.RegisterConf{
		Address:         "127.0.0.1:9001",
		Password:        "",
		DB:              4,
		MasterName:      "",
		MinWorkerId:     10,
		MaxWorkerId:     1024,
		LifeTimeSeconds: 15,
	}
	
	id := regworkerid.RegisterOne(conf)
	log.Infof("注册的WorkerId: %d", id)
	return uint16(id)
}

func GenUid() int64 {
	return idgen.NextId()
}
