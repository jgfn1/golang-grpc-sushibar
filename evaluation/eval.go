// Team members
// Gerson Fialho | jgfn1
// Arthur Frade | afa4
// César Silva | accs2

package main

import (
	"fmt"
	"os/exec"
	"sync"
	"time"
)

var clientId = 0
var rttsArray = []float64{}
var rttsSum = 0

func main() {
	numberOfGoRoutinesArray := []int{1, 5, 10, 20, 40, 80}
	var wg sync.WaitGroup

	for _, numberOfGoRoutines := range numberOfGoRoutinesArray {
		goRoutinesLimitChannel := make(chan int, numberOfGoRoutines)
		for i := 0; i < 10000; i++  {
			goRoutinesLimitChannel <- 1
			wg.Add(1)
			go func() {
				defer wg.Done()
				createClientProcess()
				<- goRoutinesLimitChannel
			}()
		}
		wg.Wait()
		fmt.Printf("RTT mean for %d processe(s): %d\n", numberOfGoRoutines, rttsSum/len(rttsArray))
	}
}

func createClientProcess() {
	start := time.Now().UnixMilli()
	clientId++

	bash := exec.Command("bash", "-c", "go run client/client.go")
	_, err := bash.Output()
	//stdout, err := bash.Output()
	end := time.Now().UnixMilli()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	rtt := end - start
	rttsArray = append(rttsArray, float64(rtt)) 
	rttsSum += int(rtt)
}