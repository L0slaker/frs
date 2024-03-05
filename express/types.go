package express

import "context"

type OrderActivities interface {
	SendOrderConfirmationEmail(ctx context.Context, orderID string) (string, error)
	UpdateOrderStatus(ctx context.Context, orderID, status string) (string, error)
	CheckInventoryAndNotifyShipping(ctx context.Context, orderID string) (string, error)
}
