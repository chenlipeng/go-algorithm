package demo

import (
	"log"
	"sync"
)

/*
 *
 * D.1 数组/切片索引越界
 * D.2 空指针调用
 * D.3 过早关闭 HTTP 响应体
 * D.4 除以零
 * D.5 向已关闭的通道发消息
 * D.6 重复关闭通道
 * D.7 关闭未初始化的通道
 * D.8 给未初始化map赋值
 * D.9 跨协程的恐慌处理
 * D.10 sync 计数为负值
 *
 */
type Person struct {
	Name string
}

func (p *Person) Say() {
	log.Println("Hello. My name is", p.Name)
}

func PanicMain() {
	DivideZero()
	NilPtrCall()
	log.Println("ok")
	return

	wg := sync.WaitGroup{}
	wg.Add(2)

	go exceedIndex(&wg)

	//go deadLock(&wg)
	go concurrentMap(&wg)

	//调用不存在的函数——编译期间就报错
	//unknownFunc()
	wg.Wait()
}

//除数为0——可捕获
func DivideZero() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed:", err)
		}
	}()

	x, y := 1, 0
	log.Println(x / y)
}

//空指针调用 —— 可捕获
func NilPtrCall() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed:", err)
		}
	}()

	p := &Person{"chenlipeng"}
	p = nil
	p.Say()
}

//不加锁并发读写map——不可捕捉
func concurrentMap(wg1 *sync.WaitGroup) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed:", err)
			wg1.Done()
		}
	}()

	mp := map[int]int{}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		for idx := 0; idx < 1000; idx++ {
			mp[idx] = idx
		}
		wg.Done()
	}()

	go func() {
		for idx := 0; idx < 1000; idx++ {
			mp[idx] = idx
		}
		wg.Done()
	}()
	wg.Wait()
}

//死锁——不可捕捉
func deadLock(wg1 *sync.WaitGroup) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed:", err)
			wg1.Done()
		}
	}()

	a, b := make(chan bool, 1), make(chan bool, 1)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		select {
		case <-a:
			log.Println("a")
			wg.Done()
		}
	}()

	go func() {
		select {
		case <-b:
			log.Println("b")
			//a <- true
			wg.Done()
		}
	}()

	//a <- true
	b <- true
	wg.Wait()
}

func exceedIndex(wg *sync.WaitGroup) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed:", err)
			wg.Done()
		}
	}()

	log.Println("xxxxxxxxxx")
	//越界——可捕捉
	arr := []int{}
	log.Println(arr[1990])
}
