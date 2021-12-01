package main

import (
	"fmt"
	"log"

	"github.com/cfabrica46/job-dispenser/cmd"
	"github.com/cfabrica46/job-dispenser/job"
)

func main() {
	var idJob int

	err := cmd.ClearTerminal()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Welcome to job-dispenser")
	fmt.Println("What work you wish to dispense?")
	fmt.Print(job.PrintAllJobs())
	fmt.Print("> ")
	_, err = fmt.Scan(&idJob)
	if err != nil {
		log.Fatal(err)
	}

	if idJob > job.GetLastID() || idJob <= 0 {
		log.Fatal("ID job not valid")
	}

}
