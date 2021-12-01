package job

import "fmt"

func PrintAllJobs() (result string) {
	for i := range MyJobs {
		result += fmt.Sprintf("%d. %s\n", i+1, MyJobs[i].Name)
	}
	return
}

func GetLastID() int {
	return MyJobs[len(MyJobs)-1].ID
}
