package main

import (
	"fmt"
	"os/exec"
	"sync"
)

func main() {
	numberOfClientsArray := []int{1, 5, 10, 20, 40, 80}
	requestsPerClient := 1000
	var wg sync.WaitGroup
	for _, numberOfClients := range numberOfClientsArray {
		fmt.Printf("For %d client(s)\n", numberOfClients)
		for i := 0; i < numberOfClients; i++ {
			wg.Add(1)
			go createClientProcess(requestsPerClient, &wg)
		}
		wg.Wait()
	}
}

func createClientProcess(numberOfRequests int, wg *sync.WaitGroup) {
	defer wg.Done()
	invokeClientCommand := fmt.Sprintf("go run client/client.go %d", numberOfRequests)
	bash := exec.Command("bash", "-c", invokeClientCommand)
	output, err := bash.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	outputStr := string(output)
	fmt.Printf("mean %d\n", outputStr)
}
