package dtos

import (
	"github.com/google/uuid"
)

type OrderDTO struct {
	MemberID uuid.UUID
	Items    []OrderItemDTO
}
