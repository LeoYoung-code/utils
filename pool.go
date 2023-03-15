package utils

import (
	"log"

	"github.com/panjf2000/ants/v2"
)

func InitPool(num int) *ants.Pool {
	pool, err := ants.NewPool(num)
	if err != nil {
		log.Fatalln(err)
	}
	return pool
}
