package go_pool

import (
	"testing"
	"log"
	"time"
)

func TestNewGoPool(t *testing.T) {
	goPoolObject := NewGoPool(10, CallFunc)
	go func() {
		for i := 0; i <= 10; i++ {
			goPoolObject.Push(i)
		}
		goPoolObject.Close()
	}()
	go func() {
		// 3少后自动结束
		time.Sleep(time.Second * 3)
		goPoolObject.StopChan <- 1
	}()
	goPoolObject.Run()
}

func CallFunc(val interface{}) {
	log.Println(val, "======")
}
