package pkg

import (
	"fmt"
	"main/pkg/utils"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func SquareNumbersWithShutdown() {
	var num = 1
	var signalChannel = make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT)
	exit := false

	// запускается перед всем началом цикла, и начинает ожидать значения
	// не упало вроде из-за того что сделал буферезированный канал
	go utils.WaitExit(signalChannel, &exit)

	for {
		if exit {
			break
		}
		SquareNumber(num)
		num++

	}

}

func SquareNumber(num int) {
	time.Sleep(1 * time.Second)
	fmt.Println(num * num)
}
