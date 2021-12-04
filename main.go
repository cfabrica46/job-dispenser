package main

import (
	"fmt"
	"log"

	"github.com/cfabrica46/job-dispenser/cmd"
	"github.com/cfabrica46/job-dispenser/job"
)

/* Para aprovechar la concurrencia utilizare 4 tipos de jobs
1. agrupadas: las cuales seran un conjunto de tareas que tendran un orden para ejecutarse
2. independientes: estas se ejecutaran de manera independiente de las demas tareas
3. primordiales: mientras halla tareas primordiales pendientes las demas tareas no podran ejecutarse hasta que se complete la ultima de las tareas primordiales
4. de cierre: tareas que se deberan ejecutar una vez no halla tareas de ningun otro tipo

*/
func main() {
	var idJob int

	c := make(chan bool)

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

	check := job.VerifyIsExistByID(idJob)
	if !check {
		log.Fatal("Job ID invalid")
	}

	err = job.ExecuteJobByID(idJob, c)
	if err != nil {
		log.Fatal(err)
	}

	<-c

}
