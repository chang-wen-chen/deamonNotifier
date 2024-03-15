package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 創建一個 channel 來接收信號
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// 監聽 SIGINT (Ctrl+C) 和 SIGTERM 信號
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// 開始主要邏輯
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				fmt.Println("Daemon is running...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// 等待信號
	<-sigs
	fmt.Println("Exiting...")
	done <- true
}
