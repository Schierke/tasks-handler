package service

import (
	"context"

	"github.com/Schierke/tasks-handler/internals/models"
)

type ShiftRepository interface {
	GetShiftsByTaskID(ctx context.Context, id string) ([]models.Shift, error)
}

type shiftserviceImpl struct {
	shiftRepo ShiftRepository
	slotRepo  SlotRepository
}

func NewShiftService(shiftRepo ShiftRepository, slotRepo SlotRepository) *shiftserviceImpl {
	return &shiftserviceImpl{
		shiftRepo: shiftRepo,
		slotRepo:  slotRepo,
	}
}

func (s *shiftserviceImpl) GetShifts(ctx context.Context, taskID string) ([]models.ShiftResponse, error) {

	shifts, err := s.shiftRepo.GetShiftsByTaskID(ctx, taskID)

	if err != nil {
		return nil, err
	}

	response := make([]models.ShiftResponse, len(shifts))

	for i, shift := range shifts {
		slots, err := s.slotRepo.GetSlotsByShiftID(ctx, shift.ID)
		if err != nil {
			return nil, err
		}

		response[i].ID = shift.ID
		response[i].Time = shift.Time

		response[i].Available = shift.Slots
		count := 0
		for _, slot := range slots {
			if slot.Status == "filled" {
				count += 1
			}
		}

		response[i].Filled = count
	}

	return response, nil
}
