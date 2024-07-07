package pkg

import (
	"fmt"
	"main/pkg/utils"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func SquareNumbersWithShutdown() {
	// как бы переделать на еще один канал
	var waitGroup sync.WaitGroup
	var num = 1
	var signalChannel = make(chan os.Signal, 1)
	exit := false

	// запускается перед всем началом цикла, и начинает ожидать значения
	// не упало вроде из-за того что сделал буферезированный канал
	go utils.WaitExit(signalChannel, &exit)

	for {
		// добавляем 1 в счеткик очереди для контроля горутин
		waitGroup.Add(1)

		signal.Notify(signalChannel, syscall.SIGINT)

		go SquareNumber(&waitGroup, num)
		// ниже идет ожидание выполнения горутины т.е отработки defer wg.done()
		waitGroup.Wait()
		num++

		if exit {
			break
		}
	}

	// тут вроде ставим чтобы последняя горутина с возведением в квадрат выполнилась когда когда мы уже отсигналили
	waitGroup.Wait()

}

func SquareNumber(wg *sync.WaitGroup, num int) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println(num * num)
}
