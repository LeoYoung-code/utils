package utils

import (
	"hash/fnv"
	"strconv"
)

func Num2Version(innerVersion string) (outVersion string) {
	innr, err := strconv.Atoi(innerVersion)
	if err != nil {
		return innerVersion
	}
	if len(innerVersion) >= 7 {
		innr = innr / 100
	}
	big := innr / 10000
	innr = innr % 10000
	mid := innr / 100
	innr = innr % 100
	min := ""
	if innr > 0 {
		min = "." + strconv.Itoa(innr)
	}
	outVersion = strconv.Itoa(big) + "." + strconv.Itoa(mid) + min
	return
}

func GetSum32(data string) uint32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(data))
	return h.Sum32()
}
