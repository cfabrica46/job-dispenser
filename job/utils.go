package job

import "fmt"

func addJobIntoMyJobs(id int, name string, myFunc MyFunc) {
	myJobs[id] = myJob{id: id, name: name, myFunc: myFunc}
}

func ExecuteJobByID(id int) (err error) {
	err = myJobs[id].myFunc()
	if err != nil {
		return
	}
	return
}

func PrintAllJobs() (result string) {
	for i := range myJobs {
		result += fmt.Sprintf("%d. %s\n", myJobs[i].id, myJobs[i].name)
	}
	return
}

func VerifyIsExistByID(id int) bool {
	return myJobs[id].myFunc != nil
}
