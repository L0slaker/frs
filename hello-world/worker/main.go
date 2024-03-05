package main

import (
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
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

	// This worker hosts both Workflow and Activity functions
	w := worker.New(c, hello_world.GreetingTaskQueue, worker.Options{})
	w.RegisterWorkflow(hello_world.GreetingWorkflow)
	w.RegisterActivity(hello_world.ComposeGreeting)

	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
