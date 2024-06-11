package utils

import (
	"fmt"
	"reflect"
	"time"
)

var Debug = new(dbg)

type dbg struct {
}

func (d dbg) Print(i interface{}) {
	var kv = make(map[string]interface{})
	vValue := reflect.ValueOf(i)
	vType := reflect.TypeOf(i)
	for i := 0; i < vValue.NumField(); i++ {
		kv[vType.Field(i).Name] = vValue.Field(i)
	}
	fmt.Println("------ debug print ------")
	for k, v := range kv {
		fmt.Println(k, " : ", v)
	}
}

func (d dbg) PrintAndWaiting(data map[string]interface{}) {
	for k, v := range data {
		fmt.Printf("========%s=========: %v", k, v)
		fmt.Println("")
	}
	time.Sleep(time.Second * 1000000)
}

func (d dbg) Waiting(mark ...interface{}) {
	if len(mark) > 0 {
		for i, v := range mark {
			fmt.Printf("Waiting at mark[%d] value[%v]\n", i, v)
		}
	}
	time.Sleep(time.Second * 1000000)
}
