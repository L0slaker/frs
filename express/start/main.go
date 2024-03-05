package main

import (
	"context"
	"fmt"
	"go.temporal.io/sdk/client"
	"hello-world-temporal/app/express"
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
		ID:        "express-workflow",
		TaskQueue: express.ExpressTaskQueue,
	}

	// Start the Workflow
	we, err := c.ExecuteWorkflow(context.Background(), options, express.OrderProcessingWorkflow, "12138")
	if err != nil {
		log.Fatalln("Unable to complete Workflow", err)
	}

	// Get the results
	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable to get Workflow result", err)
	}

	printResults(result, we.GetID(), we.GetRunID())
}

func printResults(result string, workflowID, runID string) {
	fmt.Printf("\nWorkflowID：%s\n RunID：%s \n", workflowID, runID)
	fmt.Printf("\n%s\n\n", result)
}
