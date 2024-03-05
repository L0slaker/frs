package main

import (
	"context"
	"fmt"
	"go.temporal.io/sdk/client"
	"hello-world-temporal/app/hello-world"
	"log"
)

func main() {
	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "greeting-workflow",
		TaskQueue: hello_world.GreetingTaskQueue,
	}

	// Start the Workflow
	we, err := c.ExecuteWorkflow(context.Background(), options, hello_world.GreetingWorkflow, "World")
	if err != nil {
		log.Fatalln("Unable to complete Workflow", err)
	}

	// Get the results
	var greeting string
	err = we.Get(context.Background(), &greeting)
	if err != nil {
		log.Fatalln("Unable to get Workflow result", err)
	}

	printResults(greeting, we.GetID(), we.GetRunID())
}

func printResults(greeting string, workflowID, runID string) {
	fmt.Printf("\nWorkflowID：%s\n RunID：%s \n", workflowID, runID)
	fmt.Printf("\n%s\n\n", greeting)
}
