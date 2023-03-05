package utils

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/fatih/structs"
	"github.com/go-kratos/kratos/v2/log"
	protobufStruct "github.com/golang/protobuf/ptypes/struct"
	"google.golang.org/protobuf/types/known/structpb"
)

// 将pb 的struct 转换为map
func elabValue(value *protobufStruct.Value) (interface{}, error) {
	var err error
	if value == nil {
		return nil, nil
	}
	if structValue, ok := value.GetKind().(*protobufStruct.Value_StructValue); ok {
		result := make(map[string]interface{})
		for k, v := range structValue.StructValue.Fields {
			result[k], err = elabValue(v)
			if err != nil {
				return nil, err
			}
		}
		return result, err
	}
	if listValue, ok := value.GetKind().(*protobufStruct.Value_ListValue); ok {
		result := make([]interface{}, len(listValue.ListValue.Values))
		for i, el := range listValue.ListValue.Values {
			result[i], err = elabValue(el)
			if err != nil {
				return nil, err
			}
		}
		return result, err
	}
	if _, ok := value.GetKind().(*protobufStruct.Value_NullValue); ok {
		return nil, nil
	}
	if numValue, ok := value.GetKind().(*protobufStruct.Value_NumberValue); ok {
		return numValue.NumberValue, nil
	}
	if strValue, ok := value.GetKind().(*protobufStruct.Value_StringValue); ok {
		return strValue.StringValue, nil
	}
	if boolValue, ok := value.GetKind().(*protobufStruct.Value_BoolValue); ok {
		return boolValue.BoolValue, nil
	}
	return errors.New(fmt.Sprintf("Cannot convert the value %+v", value)), nil
}

func elabEntry(entry interface{}) (*protobufStruct.Value, error) {
	var err error
	if entry == nil {
		return &protobufStruct.Value{Kind: &protobufStruct.Value_NullValue{}}, nil
	}
	rt := reflect.TypeOf(entry)
	switch rt.Kind() {
	case reflect.String:
		if realValue, ok := entry.(string); ok {
			return &protobufStruct.Value{Kind: &protobufStruct.Value_StringValue{StringValue: realValue}}, nil
		}
		return nil, errors.New("cannot convert string value")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return &protobufStruct.Value{Kind: &protobufStruct.Value_NumberValue{NumberValue: float64(reflect.ValueOf(entry).Int())}}, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return &protobufStruct.Value{Kind: &protobufStruct.Value_NumberValue{NumberValue: float64(reflect.ValueOf(entry).Uint())}}, nil
	case reflect.Float32, reflect.Float64:
		return &protobufStruct.Value{Kind: &protobufStruct.Value_NumberValue{NumberValue: reflect.ValueOf(entry).Float()}}, nil
	case reflect.Bool:
		if realValue, ok := entry.(bool); ok {
			return &protobufStruct.Value{Kind: &protobufStruct.Value_BoolValue{BoolValue: realValue}}, nil
		}
		return nil, errors.New("cannot convert boolean value")
	case reflect.Array, reflect.Slice:
		lstEntry := reflect.ValueOf(entry)

		lstValue := &protobufStruct.ListValue{Values: make([]*protobufStruct.Value, lstEntry.Len(), lstEntry.Len())}
		for i := 0; i < lstEntry.Len(); i++ {
			lstValue.Values[i], err = elabEntry(lstEntry.Index(i).Interface())
			if err != nil {
				return nil, err
			}
		}
		return &protobufStruct.Value{Kind: &protobufStruct.Value_ListValue{ListValue: lstValue}}, nil
	case reflect.Struct:
		return elabEntry(structs.Map(entry))
	case reflect.Map:
		mapEntry := make(map[string]interface{})
		entryValue := reflect.ValueOf(entry)
		for _, k := range entryValue.MapKeys() {
			mapEntry[k.String()] = entryValue.MapIndex(k).Interface()
		}
		structVal, err := Map2Struct(mapEntry)
		return &protobufStruct.Value{Kind: &protobufStruct.Value_StructValue{StructValue: structVal}}, err
	}
	return nil, errors.New(fmt.Sprintf("Cannot convert [%+v] kind:%s", entry, rt.Kind()))
}

func Map2Struct(input map[string]interface{}) (*protobufStruct.Struct, error) {
	var err error
	result := &protobufStruct.Struct{Fields: make(map[string]*protobufStruct.Value)}
	for k, v := range input {
		result.Fields[k], err = elabEntry(v)
		if err != nil {
			return nil, err
		}
	}
	return result, err
}

func Struct2Map(str *protobufStruct.Struct) (map[string]interface{}, error) {
	var err error
	result := make(map[string]interface{})
	for k, v := range str.Fields {
		result[k], err = elabValue(v)
		if err != nil {
			return nil, err
		}
	}
	return result, err
}

func PbStruct2Slice(value any) *protobufStruct.Value {
	a := structs.Values(value)
	list, err := structpb.NewList(a)
	if err != nil {
		log.Error(err.Error())
		return &protobufStruct.Value{Kind: &structpb.Value_NullValue{}}
	}
	return structpb.NewListValue(list)
}

// AnySlice2PbValue  pb结构体转切片proto   **  注意any一定是pb的struct  **
func AnySlice2PbValue[T1 any](arr []T1) (list *structpb.ListValue) {
	pp := make([]interface{}, 0, len(arr))
	for _, v := range arr {
		pp = append(pp, v)
	}
	pbArr := Map(pp, PbStruct2Slice)
	list = &structpb.ListValue{}
	for _, value := range pbArr {
		list.Values = append(list.Values, value)
	}
	return
}
