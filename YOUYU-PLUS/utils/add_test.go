package utils

import "testing"

func TestAdd(t *testing.T) {
	
	got := Add(1, 2)
	if got != 3 {
		t.Errorf("测试失败!期望得到3，实际得到的是:%d", got)
	}
}