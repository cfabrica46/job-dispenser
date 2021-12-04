package job

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func myPrint(c chan<- bool) (err error) {
	fmt.Println("What do you want to print?")

	reader := bufio.NewReader(os.Stdin)
	messageBytes, err := reader.ReadString('\n')
	if err != nil {
		return
	}

	message := strings.Split(string(messageBytes), " ")
	delay := time.Duration(len(message)) * time.Second

	go func() {
		fmt.Println()
		fmt.Println("Job Sended")
		time.Sleep(delay)
		fmt.Println(string(messageBytes))
		fmt.Println("Job Completed")
		fmt.Println()
		c <- true
	}()
	return
}
