package misc

import (
	"testing"
)

func Test_RequireJSON(t *testing.T) {
	parse, err := RequireJSON("test.json")
	if err != nil {
		t.Errorf("Test_RequireJSON\n %s", err)
	}
	t.Logf("Test_RequireJSON\n %s", parse.GetJsonPretty())
}
func Test_GoroutineID(t *testing.T) {
	t.Logf("Test_GoroutineID %d", GoroutineID())
}
func Test_SetTimeout(t *testing.T) {
	// t.Logf("Test_SetTimeout %d", GoroutineID())
	SetTimeout(func() {
		t.Log("after 2 seconds")
	}, 2000)
}
func Test_SetInterval(t *testing.T) {
	val := 3
	SetInterval(func() bool {
		t.Logf("loop %d sec\n", val)
		val--
		if val < 1 {
			t.Log("end")
			return true
		}
		return false
	}, 1000)
}

// func Test_LoadFiles(t *testing.T) {
// 	LoadFiles("/Users/breezymind/.gvm/pkgsets/go1.9.2/global/src/github.com/breezymind/go-misc", "")
// }
func Test_IsJSON(t *testing.T) {
	parse, err := RequireJSON("test.json")
	if err != nil {
		t.Errorf("Test_RequireJSON\n %s", err)
	}
	rawjson := parse.GetJsonString()
	t.Log(rawjson)
	t.Logf("Test_IsJSON %t", IsJSON(rawjson))
}

func Test_InArray(t *testing.T) {
	strs := []string{"apple", "banana", "orange"}
	t.Log(InArray("apple", strs))

	ints := []int{1, 2, 3}
	t.Log(InArray(2, ints))

	infs := []interface{}{222, "breezy"}
	t.Log(InArray("breezy", infs))
}
