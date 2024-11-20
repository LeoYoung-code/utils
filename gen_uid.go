package utils

import (
	"os"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/spaolacci/murmur3"
	"github.com/yitter/idgenerator-go/idgen"
)

const (
	workerIdBitLength = 9
)

func init() {
	podUID := os.Getenv("POD_UID")
	if podUID == "" {
		log.Warn("POD_UID is not set")
	} else {
		log.Infof("pod uid: %s", podUID)
	}
	workerId := podUIDToUint16(podUID)
	log.Infof("worker id: %d", workerId)
	options := idgen.NewIdGeneratorOptions(workerId)
	options.WorkerIdBitLength = workerIdBitLength // 默认值6，限定 WorkerId 最大值为2^6-1，即默认最多支持64个节点。
	options.SeqBitLength = 6                      // 默认值6，限制每毫秒生成的ID个数。若生成速度超过5万个/秒，建议加大 SeqBitLength 到 10。
	idgen.SetIdGenerator(options)
}

func podUIDToUint16(uid string) uint16 {
	hash := murmur3.Sum32([]byte(uid))
	return uint16(hash % (1 << workerIdBitLength))
}

func GenUid() int64 {
	return idgen.NextId()
}
