package utils

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

	var rs1 int
	val1, ok1 := GetGoCache("K1", rs1)
	assert.True(t, ok1)
	assert.Equal(t, v1, val1)

	var rs2 string
	val2, ok2 := GetGoCache("K2", rs2)
	assert.True(t, ok2)
	assert.Equal(t, v2, val2)

	var rs3 name
	val3, ok3 := GetGoCache("K3", rs3)
	assert.True(t, ok3)
	assert.Equal(t, v3, val3)

	var rs4 name
	val4, ok4 := GetGoCache("K4", rs4)
	assert.False(t, ok4)
	assert.Equal(t, name{}, val4)

	DelGoCache("K2")
	val2, ok2 = GetGoCache("K2", rs2)
	assert.False(t, ok2)
	assert.Equal(t, "", val2)

	time.Sleep(ctime + time.Second)
	val1, ok1 = GetGoCache("K1", rs1)
	assert.False(t, ok1)
}

func TestDeepCopy(t *testing.T) {
	type Nested struct {
		Field string
	}

	type TestStruct struct {
		Name   string
		Age    int
		Nested Nested
		Ptr    *string
	}

	t.Run("simple struct", func(t *testing.T) {
		src := TestStruct{
			Name: "test",
			Age:  20,
		}
		dst := TestStruct{}

		DeepCopy(&dst, &src)

		assert.Equal(t, src.Name, dst.Name)
		assert.Equal(t, src.Age, dst.Age)
	})

	t.Run("nested struct", func(t *testing.T) {
		str := "ptr"
		src := TestStruct{
			Name: "test",
			Age:  20,
			Nested: Nested{
				Field: "nested",
			},
			Ptr: &str,
		}
		dst := TestStruct{}

		DeepCopy(&dst, &src)

		assert.Equal(t, src.Name, dst.Name)
		assert.Equal(t, src.Age, dst.Age)
		assert.Equal(t, src.Nested.Field, dst.Nested.Field)
		assert.Equal(t, *src.Ptr, *dst.Ptr)
		assert.NotSame(t, src.Ptr, dst.Ptr)
	})

	t.Run("nil values", func(t *testing.T) {
		src := TestStruct{}
		dst := TestStruct{}

		DeepCopy(&dst, &src)

		assert.Equal(t, src, dst)
	})

	t.Run("error handling", func(t *testing.T) {
		src := "string"
		dst := 123

		DeepCopy(&dst, &src)

		assert.Equal(t, 123, dst)
	})
}
