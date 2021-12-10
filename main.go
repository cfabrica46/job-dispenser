package main

import (
	"fmt"
	"log"

	"github.com/cfabrica46/job-dispenser/job"
)

// owo
//holi
//jasas

/* Para aprovechar la concurrencia utilizare 4 tipos de jobs
1. agrupadas: las cuales seran un conjunto de tareas que tendran un orden para ejecutarse
2. independientes: estas se ejecutaran de manera independiente de las demas tareas
3. primordiales: mientras halla tareas primordiales pendientes las demas tareas no podran ejecutarse hasta que se complete la ultima de las tareas primordiales
4. de cierre: tareas que se deberan ejecutar una vez no halla tareas de ningun otro tipo

*/
func main() {
	var idJob int

	jobDispenser, err := job.NewJobDispenserFilled()
	if err != nil {
		log.Fatal(err)
	}

	for {
		/* err = cmd.ClearTerminal()
		if err != nil {
			log.Fatal(err)
		} */

		fmt.Println("Welcome to job-dispenser")
		fmt.Println()
		fmt.Println("What do want to do?")
		fmt.Print(job.GetStringOptions(jobDispenser))
		fmt.Print("> ")

		_, err = fmt.Scan(&idJob)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println()

		err = jobDispenser.ExecuteJobByIndex(idJob - 1)
		if err != nil {
			log.Fatal(err)
		}

		// <-c
	}

}
