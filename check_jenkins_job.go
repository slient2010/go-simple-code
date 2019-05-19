package main

import (
	"fmt"
	"github.com/bndr/gojenkins"
)

func main() {
	jenkins := gojenkins.CreateJenkins(nil, "http://loalhost:8080/", "admin", "admin")
	// Provide CA certificate if server is using self-signed certificate
	// caCert, _ := ioutil.ReadFile("/tmp/ca.crt")
	// jenkins.Requester.CACert = caCert
	_, err := jenkins.Init()

	if err != nil {
		panic("Something Went Wrong")
	}

	nodes, _ := jenkins.GetAllNodes()

	for _, node := range nodes {

		// Fetch Node Data
		node.Poll()
		result, _ := node.IsOnline()
		if result {
			fmt.Println("Node is Online")
		}
	}

	if err != nil {
		panic("Last SuccessBuild does not exist")
	}

	jobs, _ := jenkins.GetAllJobs()
	for _, i := range jobs {
		jobIsRunning, _ := i.IsRunning()
		fmt.Println(i.Raw.DisplayName, jobIsRunning)
	}
}
