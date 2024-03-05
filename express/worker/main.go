package main

import (
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"hello-world-temporal/app/express"
	"log"
)

func main() {
	// Create hte client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	// This worker hosts both Workflow and Activity functions
	w := worker.New(c, express.ExpressTaskQueue, worker.Options{})
	w.RegisterWorkflow(express.OrderProcessingWorkflow)
	w.RegisterActivity(express.SendOrderConfirmationEmail)
	w.RegisterActivity(express.UpdateOrderStatus)
	w.RegisterActivity(express.CheckInventoryAndNotifyShipping)

	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
