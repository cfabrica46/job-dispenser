package job

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func myPrint(restore, abort chan bool) (err error) {
	fmt.Println("What do you want to print?")
	fmt.Print("> ")

	reader := bufio.NewReader(os.Stdin)
	messageBytes, err := reader.ReadString('\n')
	if err != nil {
		return
	}

	restore <- true

	// fmt.Println(<-abort)

	message := strings.Split(string(messageBytes), " ")
	delay := time.Duration(len(message)) * time.Second

	/* select {
	case <-c:
		log.Fatal("holis")
		fmt.Println("Abort")
		return
	default: */
	fmt.Printf("\r\n")
	fmt.Printf("\rJob Sended\n")
	time.Sleep(delay)
	fmt.Printf("\rPRINT: %s\n", string(messageBytes))
	fmt.Println("Job Completed")
	fmt.Println()
	fmt.Print("> ")
	// }
	return
}

func myWait(restore, abort chan bool) (err error) {
	restore <- true
	return
}

func myAbort(restore, abort chan bool) (err error) {
	restore <- true
	// abort <- true
	fmt.Println("Abort All Jobs")
	return
}
