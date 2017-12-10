# go-misc [![GoDoc](https://godoc.org/github.com/breezymind/go-misc?status.svg)](https://godoc.org/github.com/breezymind/go-misc)
golang 프로젝트들을 도와줄 라이브러리

## Todos

- [x] misc, test example 작성
- [x] misc, godoc comment
- [ ] check SignalListener

## Installation

```bash
go get "github.com/breezymind/misc"
```

```go
import "github.com/breezymind/misc"
```


## Usage

**func GoroutineID() int**

: GoroutineID 는 고루틴 고유값을 리턴 한다

```go
fmt.Printf("GoroutineID %d", GoroutineID())

// Output:
// GoroutineID 7
```

**func InArray(needle interface{}, pool interface{}) (bool, int)**

: InArray 는 pool 배열에서 needle 값이 있는 지 확인한다
Example

```go
strs := []string{"apple", "banana", "orange"}
fmt.Println(InArray("apple", strs))
// Output:
// true 0

ints := []int{1, 2, 3}
fmt.Println(InArray(2, ints))
// Output:
// true 1

infs := []interface{}{222, "breezy"}
fmt.Println(InArray("breezy", infs))
// Output:
// true 1
```

**func RequireJSONFile(filepath string)(\*syncmap.SyncMap, error)**

: RequireJSONFile 는 filepath로 정의한 주석을 제거한 json 파일을 로드 하여 syncmap.SyncMap 형태로 리턴한다

```go
parse, err := RequireJSONFile("test.json")
if err != nil {
    fmt.Printf("Error : %s", err.Error())
}
fmt.Printf(parse.GetJSONPretty())
// Output:
// {
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
```


**func IsJSON(str string) bool**

: IsJSON 는 string 이 json 포멧인지 검증한다

```go
parse, err := RequireJSONFile("test.json")
if err != nil {
    fmt.Printf("Error : %s", err.Error())
}
rawjson := parse.GetJSONString()
fmt.Print(rawjson)
// Output:
// {"obj":{"arr-attr":["gml","xml"],"int-attr":100,"obj-attr":{"para":"a meta-markup language, used to create markup languages such as docbook."},"str-attr":"sgml"}}
fmt.Printf("IsJSON %t", IsJSON(rawjson))
// Output:
// IsJSON true
```

**func SetInterval(cb func() bool, ms time.Duration)**

: SetInterval 는 ms 만큼 반복하며 cb 함수에 true 리턴하면 종료 한다

```go
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
```

**func SetTimeout(cb func(), ms time.Duration)**

: SetTimeout 는 ms 만큼 후에 cb 함수를 실행 한다

```go
SetTimeout(func() {
    fmt.Printf("after 2 seconds")
}, 2000)
// Output:
// after 2 seconds
```




```bash
$ go test -v
```

## License
[MIT license](https://opensource.org/licenses/MIT)
