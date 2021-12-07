package job

import "fmt"

func GetStringOptions(dispenser JobDispenser) (result string) {
	for i := range dispenser.myJobs {
		result += fmt.Sprintf("%d. %s\n", i+1, dispenser.myJobs[i].name)
	}
	return
}
