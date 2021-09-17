package schema

import (
	"fmt"
	"parking_lot/errors"
	"time"
)

// ParkingLot struct holds all the parking lot information and parking history
type ParkingLot struct {
	Name        string         `json:"name"`
	Floor       string         `json:"floor"`
	TotalBlocks int            `json:"total_blocks"`
	BlockHeight int            `json:"block_height"`
	TotalSlots  int            `json:"total_slots"`
	Address     string         `json:"address"`
	Pincode     string         `json:"pincode"`
	Slots       []*Slot        `json:"slots"`
	ParkHistory []*ParkHistory `json:"park_history"`
}

// ParkHistory holds the parking information
type ParkHistory struct {
	SlotID             uint
	RegistrationNumber string
	Colour             string
	CreatedAt          time.Time
}

// FirstAvailableSlot returns the first available slot to park Vehicle
func (pl *ParkingLot) FirstAvailableSlot() (*Slot, error) {
	for _, slot := range pl.Slots {
		if slot.IsSlotAvailable() {
			return slot, nil
		}
	}

	return nil, errors.ErrParkingSlotsFull
}

// IsDuplicateRegNo checks if car already registered in the parking lot
func (pl *ParkingLot) IsDuplicateRegNo(regno string) error {
	for _, slot := range pl.Slots {
		if slot.IsSlotOccupied() && slot.Vehicle.IsVehicleRegNoMatched(regno) {
			return errors.ErrDuplicateVehicle(regno)
		}
	}

	return nil
}

// GetAnUnoccupiedSlot return a slot if occupied else nil
func (pl *ParkingLot) GetAnUnoccupiedSlot(slotNum int) (*Slot, error) {

	if slotNum > len(pl.Slots) {
		return nil, errors.ErrSlotAlreadyAvailable
	}
	return pl.Slots[slotNum], nil
}

// GetAllRegistrationsByCarColor return list of registration numbers based on car color
func (pl *ParkingLot) GetAllRegistrationsByCarColor(carColor string) (string, error) {
	res := ""
	for _, slot := range pl.Slots {
		if slot.IsSlotOccupied() && slot.Vehicle.IsVehicleColurMatched(carColor) {
			res += fmt.Sprintf("%s, ", slot.Vehicle.GetRegNumber())
		}
	}

	if len(res) > 0 {
		res = res[0 : len(res)-2]
	}

	return res, nil
}

// GetAllSlotsByCarColor return list of registration numbers based on car color
func (pl *ParkingLot) GetAllSlotsByCarColor(carColor string) (string, error) {
	res := ""
	for _, slot := range pl.Slots {
		if slot.IsSlotOccupied() && slot.Vehicle.IsVehicleColurMatched(carColor) {
			res += fmt.Sprintf("%d, ", slot.GetID())
		}
	}

	if len(res) > 0 {
		res = res[0 : len(res)-2]
	}

	return res, nil
}

// GetSlotByRegNo return a slot by car registration number
func (pl *ParkingLot) GetSlotByRegNo(regNo string) (string, error) {
	res := ""
	for _, slot := range pl.Slots {
		if slot.IsSlotOccupied() && slot.Vehicle.IsVehicleRegNoMatched(regNo) {
			return fmt.Sprintf("%d", slot.GetID()), nil
		}
	}

	return res, errors.ErrCarNotFound
}
