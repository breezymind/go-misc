package misc

import (
	"encoding/json"
	"fmt"
	"hash/crc32"
	"hash/crc64"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/breezymind/gq"
)

// RequireJSONFile 는 filepath로 정의한 주석을 제거한 json 파일을 로드 하여 gq.SyncMap 형태로 리턴한다.
func RequireJSONFile(filepath string) (*gq.SyncMap, error) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("Read File Error : %s", err)
	}
	re := regexp.MustCompile("(?s)[^https?:]//.*?\n|/\\*.*?\\*/")
	validfile := re.ReplaceAll(file, nil)
	wrap := gq.NewSyncMapByJSONByte(validfile)
	if wrap == nil {
		return nil, fmt.Errorf("Read File Error : %s", err)
	}
	return wrap, nil
}

// GoroutineID 는 고루틴 고유값을 리턴 한다
func GoroutineID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func SignalListener(cb func()) chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(
		sig,
		syscall.SIGSEGV,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	go func() {
		for {
			s := <-sig
			switch s {
			case syscall.SIGSEGV:
				log.Println("[SIGSEGV]")
			case syscall.SIGHUP:
				log.Println("[SIGHUP]")
			case syscall.SIGINT:
				log.Println("[SIGINT]")
			case syscall.SIGTERM:
				log.Println("[SIGTERM]")
			case syscall.SIGQUIT:
				log.Println("[SIGQUIT]")
			}
			cb()
		}
	}()
	return sig
}

// SetTimeout 는 ms 만큼 후에 cb 함수를 실행 한다
func SetTimeout(cb func(), ms time.Duration) {
	time.Sleep(time.Millisecond * ms)
	cb()
}

// SetInterval 는 ms 만큼 반복하며 cb 함수에 true 리턴하면 종료 한다
func SetInterval(cb func() bool, ms time.Duration) {
	t := time.NewTicker(time.Millisecond * ms).C
	for {
		select {
		case <-t:
			if cb() == true {
				return
			}
		}
	}
}

/*
func LoadFiles(dir string, ext string) map[string]interface{} {
	cts := make(map[string]interface{}, 0)
	nodes, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, node := range nodes {
		nodename := node.Name()
		nodepath := dir + "/" + nodename
		if node.IsDir() {
			for k, v := range LoadFiles(nodepath, ext) {
				cts[k] = v
			}
		} else {
			res, _ := ioutil.ReadFile(nodepath)
			cts[strings.Replace(nodename, ext, "", -1)] = res
		}
	}
	return cts
}
*/

// IsJSON 는 string 이 json 포멧인지 검증한다
func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// InArray 는 pool 배열에서 needle 값이 있는 지 확인한다
func InArray(needle interface{}, pool interface{}) (bool, int) {
	val := reflect.Indirect(reflect.ValueOf(pool))
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if ok := (needle == val.Index(i).Interface()); ok {
				return ok, i
			}
		}
	}
	return false, -1
}

// StringSplitApply 는 문자열에서 특정 문자를 기준으로 슬라이스를 만들고 각 요소에 특정함수를 적용한다
func StringSplitApply(s, sep, join string, cb func(part string) string) string {
	tmp := make([]string, 0)
	for _, part := range strings.Split(s, sep) {
		part = cb(part)
		if len(part) > 0 {
			tmp = append(tmp, part)
		}
	}
	return strings.Join(tmp, join)
}

func CRC32(raw string) string {
	return fmt.Sprintf("%08x", crc32.ChecksumIEEE(
		[]byte(raw),
	))
}

func CRC64(raw string) string {
	return fmt.Sprintf("%x", crc64.Checksum(
		[]byte(raw),
		crc64.MakeTable(crc64.ECMA),
	))
}
