package express

import (
	"context"
	"fmt"
)

func SendOrderConfirmationEmail(ctx context.Context, orderID string) (string, error) {
	return fmt.Sprintf("Sending order confirmation email for orderID: %s\n", orderID), nil
}

func UpdateOrderStatus(ctx context.Context, orderID, status string) (string, error) {
	return fmt.Sprintf("Updating order status for orderID: %s to %s\n", orderID, status), nil
}

func CheckInventoryAndNotifyShipping(ctx context.Context, orderID string) (string, error) {
	return fmt.Sprintf("Checking inventory and notifying shipping for orderID: %s\n", orderID), nil
}
