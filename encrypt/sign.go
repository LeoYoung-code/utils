package encrypt

import (
	"crypto/md5"
	"encoding/hex"
	"sort"
	"strings"

	"github.com/LeoYoung-code/cast"
)

const SecretKeyOfParams = "p9LAIPj01RcBt7id" //  参数待提供

// GenerateSign 生成签名
func GenerateSign(params sort.StringSlice) string {
	sort.Sort(&params)
	h := md5.New()
	_, _ = h.Write([]byte(strings.Join(params, "") + SecretKeyOfParams))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// ParamToSlice 将参数转换为切片
func ParamToSlice(param map[string]any) sort.StringSlice {
	params := make(sort.StringSlice, 0, len(param))
	for k, v := range param {
		if k == "sign" {
			continue
		}
		valueStr := cast.ToString(v)
		params = append(params, k+"="+valueStr)
	}
	return params
}

func sign() {
	param := map[string]any{
		"request_id":      "82",
		"third_party_uid": "94",
		"sign":            "in",
		"coin_cnt":        17,
	}
	sign := GenerateSign(ParamToSlice(param))
	println(sign)
}
