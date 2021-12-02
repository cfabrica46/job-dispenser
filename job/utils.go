package job

import (
	"errors"
	"fmt"
)

func AddJobIntoMyJobs(id int, name string, myFunc MyFunc) (err error) {
	if myJobs[id].myFunc != nil {
		err = errors.New("ID job is already ocupied ")
	}
	myJobs[id] = myJob{id: id, name: name, myFunc: myFunc}
	return
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
