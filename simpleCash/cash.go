package simpleCash

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
)

type iCash interface {
	GetCashStruct(num int) (interface{}, error)
	SetCash(num int, data interface{}) error
}

type Cash struct {
	currentName string
}

func (c *Cash) GetCashStruct(num int, v interface{}) (exist bool) {
	c.setName()
	fullname := fmt.Sprintf("%s%s_%d.json", os.TempDir(), c.currentName, num)
	if !fileExists(fullname) {
		return false
	}
	fileBytes, err := ioutil.ReadFile(fullname)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(fileBytes, v); err != nil {
		panic(err)
	}
	return true
}

func (c *Cash) SetCash(num int, data interface{}) {
	c.setName()
	fullname := fmt.Sprintf("%s%s_%d.json", os.TempDir(), c.currentName, num)
	marshal, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(fullname, marshal, 777); err != nil {
		panic(err)
	}
	return
}

func (c *Cash) setName() {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	c.currentName = f.Name()
	return
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
