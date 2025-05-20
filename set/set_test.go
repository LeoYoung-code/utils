package set

import "testing"

func TestMapSet(t *testing.T) {
	// 创建一个新的集合
	s := make(MapSet)

	// 测试初始状态
	if s.Size() != 0 {
		t.Errorf("Expected size 0, got %d", s.Size())
	}

	// 测试添加元素
	s.Add("a")
	s.Add("b")
	if s.Size() != 2 {
		t.Errorf("Expected size 2, got %d", s.Size())
	}
	if !s.Has("a") {
		t.Errorf("Expected to have key 'a'")
	}
	if !s.Has("b") {
		t.Errorf("Expected to have key 'b'")
	}

	// 测试重复添加
	s.Add("a")
	if s.Size() != 2 {
		t.Errorf("Expected size 2 after duplicate add, got %d", s.Size())
	}

	// 测试删除
	s.Delete("a")
	if s.Size() != 1 {
		t.Errorf("Expected size 1 after delete, got %d", s.Size())
	}
	if s.Has("a") {
		t.Errorf("Unexpected key 'a' after delete")
	}

	// 测试Values方法
	values := s.Values()
	if len(values) != 1 {
		t.Errorf("Expected 1 value, got %d", len(values))
	}
	if values[0] != "b" {
		t.Errorf("Expected value 'b', got '%s'", values[0])
	}

	// 测试Clear方法
	s.Clear()
	if s.Size() != 0 {
		t.Errorf("Expected size 0 after clear, got %d", s.Size())
	}
	if s.Has("b") {
		t.Errorf("Unexpected key 'b' after clear")
	}

	// 测试Each方法
	count := 0
	s.Add("x")
	s.Add("y")
	s.Each(func(key string) {
		count++
		if key != "x" && key != "y" {
			t.Errorf("Unexpected key: %s", key)
		}
	})
	if count != 2 {
		t.Errorf("Expected to iterate 2 items, got %d", count)
	}
}
