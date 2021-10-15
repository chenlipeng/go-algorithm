package demo

import (
	"context"
	"log"
	"os"
	"sync"
	"time"
)

var logg *log.Logger

func someHandler() {
	wg := sync.WaitGroup{}
	//ctx, cancel := context.WithCancel(context.Background())
	//ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)

	wg.Add(1)
	go doStuff(ctx, &wg)

	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()
}

func doStuff(ctx context.Context, wg *sync.WaitGroup) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			logg.Printf("done")
			wg.Done()
			return
		default:
			logg.Printf("work")
		}
	}
}

func ContextMain() {
	logg = log.New(os.Stdout, "", log.Ltime)
	someHandler()
	logg.Printf("down")
}
