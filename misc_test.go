package misc

import (
	"fmt"
	"testing"
)

// NOTE: Tests
func Test_RequireJSONFile(t *testing.T) {
	parse, err := RequireJSONFile("test.json")
	if err != nil {
		t.Errorf("Test_RequireJSONFile\n %s", err)
	}
	t.Logf("Test_RequireJSONFile\n %s", parse.GetJSONPretty())
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
	parse, err := RequireJSONFile("test.json")
	if err != nil {
		t.Errorf("Test_RequireJSONFile\n %s", err)
	}
	rawjson := parse.GetJSONString()
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

// NOTE: Examples

func ExampleRequireJSONFile() {
	parse, err := RequireJSONFile("test.json")
	if err != nil {
		fmt.Printf("Error : %s", err.Error())
	}
	fmt.Printf(parse.GetJSONPretty())
	// Output:
	//  {
	// 	"obj": {
	// 		"arr-attr": [
	// 			"gml",
	// 			"xml"
	// 		],
	// 		"int-attr": 100,
	// 		"obj-attr": {
	// 			"para": "a meta-markup language, used to create markup languages such as docbook."
	// 		},
	// 		"str-attr": "sgml"
	// 	}
	// }
}
func ExampleGoroutineID() {
	fmt.Printf("GoroutineID %d", GoroutineID())
	// Output:
	// GoroutineID 7
}

func ExampleSetTimeout() {
	SetTimeout(func() {
		fmt.Printf("after 2 seconds")
	}, 2000)
	// Output:
	// after 2 seconds
}

func ExampleSetInterval() {
	val := 3
	SetInterval(func() bool {
		fmt.Printf("loop %d sec\n", val)
		val--
		if val < 1 {
			fmt.Print("end")
			return true
		}
		return false
	}, 1000)
	// Output:
	// loop 3 sec
	// loop 2 sec
	// loop 1 sec
	// end
}

func ExampleIsJSON(t *testing.T) {
	parse, err := RequireJSONFile("test.json")
	if err != nil {
		fmt.Printf("Error : %s", err.Error())
	}
	rawjson := parse.GetJSONString()
	fmt.Print(rawjson)
	fmt.Printf("IsJSON %t", IsJSON(rawjson))
	// Output:
	// {"obj":{"arr-attr":["gml","xml"],"int-attr":100,"obj-attr":{"para":"a meta-markup language, used to create markup languages such as docbook."},"str-attr":"sgml"}}
	// IsJSON true
}

func ExampleInArray() {
	strs := []string{"apple", "banana", "orange"}
	fmt.Println(InArray("apple", strs))

	ints := []int{1, 2, 3}
	fmt.Println(InArray(2, ints))

	infs := []interface{}{222, "breezy"}
	fmt.Println(InArray("breezy", infs))

	// Output:
	// true 0
	// true 1
	// true 1
}
