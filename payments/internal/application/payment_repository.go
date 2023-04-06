package application

import (
	"context"
	"eda-in-go/payments/internal/models"
)

type PaymentRepository interface {
	Save(Ctx context.Context, payment *models.Payment) error
	Find(ctx context.Context, paymentID string) (*models.Payment, error)
}
