package utils

import (
	"testing"
	"time"
)

func TestSetGoCache(t *testing.T) {
	type name struct {
		Name string
		Val  string
	}

	ctime := time.Second * 5

	v1 := 123
	v2 := "r234r"
	v3 := name{"111", "222"}

	SetGoCache("K1", v1, ctime)
	SetGoCache("K2", v2, ctime)
	SetGoCache("K3", v3, ctime)
	for i := 0; i < 8; i++ {
		var rs1 int
		val1, ok1 := GetGoCache("K1", rs1)

		var rs2 string
		val2, ok2 := GetGoCache("K2", rs2)

		var rs3 name
		val3, ok3 := GetGoCache("K3", rs3)

		var rs4 name
		val4, ok4 := GetGoCache("K4", rs4)
		t.Logf("1 - val(%T): %v, ok: %+v", val1, val1, ok1)
		t.Logf("2 - val(%T): %v, ok: %+v", val2, val2, ok2)
		t.Logf("3 - val(%T): %v, ok: %+v", val3, val3, ok3)
		t.Logf("4 - val(%T): %v, ok: %+v", val4, val4, ok4)

		DelGoCache("K2")

		time.Sleep(time.Second * 1)
	}
}
