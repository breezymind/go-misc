package misc

import (
	"encoding/json"
	"fmt"
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

	"github.com/breezymind/syncmap"
)

func RequireJson(filepath string) (*syncmap.SyncMap, error) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("Read File Error : %s", err)
	}
	re := regexp.MustCompile("(?s)[^https?:]//.*?\n|/\\*.*?\\*/")
	validfile := re.ReplaceAll(file, nil)
	wrap := syncmap.NewSyncMapByJsonByte(validfile)
	if wrap == nil {
		return nil, fmt.Errorf("Read File Error : %s", err)
	}
	return wrap, nil
}

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

type SignalCallback func()

func SignalListener(cb SignalCallback) chan os.Signal {
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

func SetTimeout(cb func(), ms time.Duration) {
	time.Sleep(time.Millisecond * ms)
	cb()
}

func SetInterval(cb func(), ms time.Duration) {
	t := time.NewTicker(time.Millisecond * ms).C
	for {
		select {
		case <-t:
			cb()
		}
	}
}

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

func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

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
