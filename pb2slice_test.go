package utils

import (
	"testing"

	json "github.com/json-iterator/go"
	"pixiu-ad-backend/api/pixiu/common"
)

func TestAnySlice2PbValue(t *testing.T) {
	sss := make([]*common.HeaderUnionData, 0)
	sss = append(sss, &common.HeaderUnionData{
		AdUnitId: "111",
		Click:    "111",
	})
	sss = append(sss, &common.HeaderUnionData{
		AdUnitId: "222",
		Click:    "222",
	})

	type MM struct {
		Name string
		Age  int64
	}

	sssq := make([]*MM, 0)
	sssq = append(sssq, &MM{
		Name: "1111",
		Age:  123,
	})
	sssq = append(sssq, &MM{
		Name: "2222",
		Age:  444,
	})

	resp := &common.CreateAdminReply{
		Message: "test",
	}
	resp.HeaderUnionData = AnySlice2PbValue(sssq)
	// resp.HeaderUnionData = AnySlice2PbValue(sss)

	t.Log(resp)
	bss, _ := json.Marshal(sssq)
	t.Log(string(bss))
	bs, _ := json.Marshal(resp)
	t.Log(string(bs))
}
