package express

import (
	"fmt"
	"go.temporal.io/sdk/workflow"
	"time"
)

func OrderProcessingWorkflow(ctx workflow.Context, orderID string) (string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	})
	if err := workflow.ExecuteActivity(ctx, SendOrderConfirmationEmail,
		orderID).Get(ctx, nil); err != nil {
		return "", err
	}

	// 更新订单状态为支付成功
	if err := workflow.ExecuteActivity(ctx, UpdateOrderStatus,
		orderID, "paid").Get(ctx, nil); err != nil {
		return "", err
	}

	// 检查库存并通知物流
	if err := workflow.ExecuteActivity(ctx, CheckInventoryAndNotifyShipping,
		orderID).Get(ctx, nil); err != nil {
		return "", err
	}

	// 订单处理完成
	return fmt.Sprintf("Order processing completed for orderID: %s", orderID), nil
}
