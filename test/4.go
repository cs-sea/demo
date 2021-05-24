package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func readMemStats() {

	var ms runtime.MemStats

	runtime.ReadMemStats(&ms)

	log.Printf(" ===> Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
}

func main() {

	//创建trace文件
	//f, err := os.Create("trace.out")
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer f.Close()
	//
	////启动trace goroutine
	//err = trace.Start(f)
	//if err != nil {
	//	panic(err)
	//}
	//defer trace.Stop()

	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:10000", nil))
	}()

	fmt.Println("执行开始")
	readMemStats()
	test()
	fmt.Println("执行完成")
	readMemStats()
	time.Sleep(time.Second * 10)
	runtime.GC()
	fmt.Println("gc 完成")
	readMemStats()
	time.Sleep(time.Minute * 3)
}

func test() {
	res := make([]int, 0)

	//main
	var i = 0
	for i < 100000000 {
		res = append(res, 5)
		i++
	}
}
