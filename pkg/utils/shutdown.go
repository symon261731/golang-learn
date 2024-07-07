package utils

import (
	"fmt"
	"os"
)

func WaitExit(channel chan os.Signal, isExit *bool) {
	defer close(channel)
	<-channel
	*isExit = true
	fmt.Println("Выхожу из программы")
}
