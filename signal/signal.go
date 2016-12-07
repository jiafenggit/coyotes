package signal

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mylxsw/task-runner/config"
)

// 初始化信号接受处理程序
func InitSignalReceiver(runtime *config.Runtime) {
	signalChan := make(chan os.Signal)
	signal.Notify(
		signalChan,
		syscall.SIGHUP,
		syscall.SIGUSR2,
		syscall.SIGINT,
		syscall.SIGKILL,
	)
	go func() {
		for {
			sig := <-signalChan
			switch sig {
			case syscall.SIGUSR2, syscall.SIGHUP, syscall.SIGINT, syscall.SIGKILL:
				runtime.StopRunning = true
				//close(command)
				for i := 0; i < runtime.Concurrent; i++ {
					runtime.StopRunningChan <- struct{}{}
				}
				log.Print("Received exit signal.")
			}
		}
	}()

}
