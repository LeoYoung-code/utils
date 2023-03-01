package utils

import jsoniter "github.com/json-iterator/go"

func Obj2String(obj interface{}) string {
	objStr, _ := jsoniter.MarshalToString(obj)
	return objStr
}
