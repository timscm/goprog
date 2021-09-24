package map_test

import "testing"

func TestMapInit(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))

	// make(map[key-type]value-type, cap)
	m3 := make(map[int]int, 10)
	// cap 的参数不能是map
	// t.Logf("cap m3=%d", cap(m3))
	t.Logf("len m3=%d", len(m3))
}

func TestMapAccessKeyNotExist(t *testing.T) {
	m1 := map[int]int{}
	// 不存在时，居然返回了值0
	t.Log(m1[1])

	m1[2] = 0
	t.Log(m1[2])

	// 判断key是否存在的正确方法
	m1[3] = 0
	if v, ok := m1[3]; ok {
		t.Logf("key 3's value is %d", v)
	} else {
		t.Log("key 3 is not existing.")
	}
}

func TestMapTravel(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
