package dtos

import (
	"github.com/google/uuid"
)

type OrderItemDTO struct {
	ProductID uuid.UUID
	Quantity  uint64
}
