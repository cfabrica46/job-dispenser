package job

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func myPrint() (err error) {
	fmt.Println("What do you want to print?")

	reader := bufio.NewReader(os.Stdin)
	messageBytes, err := reader.ReadString('\n')
	if err != nil {
		return
	}

	message := strings.Split(string(messageBytes), "")

	fmt.Println("Job Sended")
	go func() {
		time.Sleep(time.Duration(len(message)) * time.Second)
		fmt.Println(string(messageBytes))
		fmt.Println("Job Completed")
	}()
	return
}
