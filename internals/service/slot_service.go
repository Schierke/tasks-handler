package service

import (
	"context"

	"github.com/Schierke/tasks-handler/internals/models"
)

type SlotRepository interface {
	GetSlotsByShiftID(ctx context.Context, id string) ([]models.Slot, error)
}

type slotserviceImpl struct {
	repo SlotRepository
}

func NewSlotService(repo SlotRepository) *slotserviceImpl {
	return &slotserviceImpl{
		repo: repo,
	}
}
