package express

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
	"testing"
)

func TestOrderProcessingWorkflow(t *testing.T) {
	// Set up the test suite and testing execution environment
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	// Mock activity implementation
	env.OnActivity(SendOrderConfirmationEmail, mock.Anything, "12138").
		Return("Sending order confirmation email for orderID: 12138", nil)
	env.OnActivity(UpdateOrderStatus, mock.Anything, "12138", "paid").
		Return("Updating order status for orderID: 12138 to paid", nil)
	env.OnActivity(CheckInventoryAndNotifyShipping, mock.Anything, "12138").
		Return("Checking inventory and notifying shipping for orderID: 12138", nil)

	env.ExecuteWorkflow(OrderProcessingWorkflow, "12138")
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	var result string
	require.NoError(t, env.GetWorkflowResult(&result))
	require.Equal(t, "Order processing completed for orderID: 12138", result)
}
