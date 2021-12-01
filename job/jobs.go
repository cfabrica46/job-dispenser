package job

import (
	"fmt"
	"time"
)

func myPrint(args []string) (err error) {
	fmt.Println("Job Sended")
	go func() {
		time.Sleep(time.Duration(len(args)) * time.Second)
		fmt.Println(args)
	}()
	return
}
