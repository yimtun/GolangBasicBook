package main

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"
)

var (
	jobRunning bool                  = false
	mu         sync.RWMutex          //读写互斥锁 适合读多写少
	jobDone    = make(chan struct{}) // 无缓冲的通道channel
)

func main() {
	go watchJobEnd()                       // 后台常驻  用于从通道接收 job运行完成  而后修改全局变量  jobRunning为false
	http.HandleFunc("/run", runJobHandler) // 运行job
	http.HandleFunc("/cat", catJobHandler) // 查看job 运行状态
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}

func Test_job(t *testing.T) {
	go watchJobEnd()                       // 后台常驻  用于从通道接收 job运行完成  而后修改全局变量  jobRunning为false
	http.HandleFunc("/", runJobHandler)    // 运行job
	http.HandleFunc("/cat", catJobHandler) // 查看job 运行状态
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}

func runJobHandler(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()
	if jobRunning {
		w.Write([]byte("runJobHandler:job 已经运行 无法再次启动"))
	}
	go runJobV1()
	w.Write([]byte("Job created\n"))
}

// 作业开始后立即释放了锁，以便其他 Goroutine 可以同时访问共享资源。
func runJobV1() {
	mu.Lock()
	defer mu.Unlock()
	if jobRunning {
		fmt.Println("runJobV1:  job 已经运行 无法再次启动")
		return
	}
	jobRunning = true // 修改状态后  而后goroutine 并不用等待 job执行完 就可以释放锁

	go func() {
		fmt.Println("Job 启动")
		time.Sleep(10 * time.Second) // 模拟job运行   执行job的代码不维护锁 效率高
		jobDone <- struct{}{}
	}()
}

// 作业完成后才释放了锁，这意味着持有锁的时间更长，可能会影响其他 Goroutine 对共享资源的访问。
func runJobV2() {
	go func() {
		mu.Lock()
		defer mu.Unlock()
		if jobRunning {
			fmt.Println("runJob check   Job is already in progress")
			return
		}
		jobRunning = true
		fmt.Println("Job started")
		time.Sleep(10 * time.Second) // 模拟job运行
		jobDone <- struct{}{}        // 等待job 运行完 才会释放锁
	}()
}

func catJobHandler(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()
	if jobRunning {
		w.Write([]byte("jobRunning\n"))
		return
	}
	w.Write([]byte("job not Running\n"))
}

func watchJobEnd() {
	for {
		select {
		case <-jobDone:
			mu.Lock()
			jobRunning = false
			mu.Unlock()
			fmt.Println("Job 完成")
		}
	}
}
