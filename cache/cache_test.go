package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

		cacheDriver.Delete("K2")

		time.Sleep(time.Second * 1)
	}
}

func TestGetGoCache(t *testing.T) {
	type args struct {
		key        string
		defaultVal any
	}

	tests := []struct {
		name     string
		setup    func()
		args     args
		wantVal  any
		wantBool bool
	}{
		{
			name: "Normal Case",
			setup: func() {
				SetGoCache[string]("test_key", "test_value", defaultExpire)
			},
			args: args{
				key:        "test_key",
				defaultVal: "default",
			},
			wantVal:  "test_value",
			wantBool: true,
		},
		{
			name: "Type Mismatch",
			setup: func() {
				SetGoCache[int]("test_key", 123, defaultExpire)
			},
			args: args{
				key:        "test_key",
				defaultVal: "default",
			},
			wantVal:  "default",
			wantBool: false,
		},
		{
			name: "Nil Value",
			setup: func() {
				// Go的泛型不支持直接使用nil作为类型参数，改用空结构体代替
				type emptyStruct struct{}
				SetGoCache[emptyStruct]("nil_key", emptyStruct{}, defaultExpire)
			},
			args: args{
				key:        "nil_key",
				defaultVal: "default",
			},
			wantVal:  "default",
			wantBool: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			gotVal, gotBool := GetGoCache(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.wantBool, gotBool)
			assert.Equal(t, tt.wantVal, gotVal)
		})
	}
}

func TestDelGoCache(t *testing.T) {
	// 设置测试键值
	SetGoCache[string]("delete_key", "value_to_delete", defaultExpire)

	// 删除缓存
	cacheDriver.Delete("delete_key")

	// 验证删除是否成功
	_, exists := GetGoCache[string]("delete_key", "")
	assert.False(t, exists, "缓存删除失败")
}

func TestDeepCopy(t *testing.T) {
	type sourceStruct struct {
		Name string
		Age  int
	}

	type targetStruct struct {
		Name string
		Age  int
	}

	tests := []struct {
		name       string
		source     any
		target     any
		wantErr    bool
		beforeFunc func()
	}{
		{
			name:   "Valid Copy",
			source: &sourceStruct{Name: "Alice", Age: 30},
			target: &targetStruct{},
			beforeFunc: func() {
				// 初始化目标对象
			},
			wantErr: false,
		},
		{
			name:   "Invalid Type",
			source: "not a struct",
			target: &targetStruct{},
			beforeFunc: func() {
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeFunc()
			err := DeepCopy(tt.target, tt.source)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
