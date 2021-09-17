package store

import (
	"fmt"
	"parking_lot/errors"
	"parking_lot/schema"
	"parking_lot/utils"
)

type leaveStore struct {
	*store
}

// NewLeaveStore returns new store object
func NewLeaveStore(st *store) *leaveStore {
	pl := &leaveStore{st}
	return pl
}

func (pl *leaveStore) IsHelp(arg string) (string, bool) {
	if arg == string(schema.CMDHelp) {
		return schema.CMDLeaveHint, true
	}
	return "", false
}

// Execute - `leave` Command will takes registration number and colour as Arguments
// the system checks for a first availabe slot to park, if slot available
// slot will allocated to the vehicle.
// This will checks if the vehicle registration number is duplicate or not.
func (pl *leaveStore) Execute(cmd *schema.Command) (string, error) {
	var slot *schema.Slot
	var slotNum int
	var err error

	if res, isHelp := pl.IsHelp(cmd.Arguments[0]); isHelp {
		return res, nil
	}
	if ParkingLot == nil {
		return "", errors.ErrNoParkingLot
	}

	// validate if slot is > 0
	if slotNum, err = validateSlotNumber(cmd.Arguments); err != nil {
		return "", err
	}

	// TODO get a slot by slot number
	if slot, err = ParkingLot.GetAnUnoccupiedSlot(slotNum); err != nil {
		return "", err
	}

	// vehicle leave parking lot
	slot.ExitPark()

	return fmt.Sprintf(SlotIsFreeInfo, slot.GetID()), nil
}

func validateSlotNumber(args []string) (int, error) {
	val, bool := utils.IsNaturalNumber(args[0])
	if bool == false {
		return val, errors.ErrInvalidSlotID
	}

	return val, nil
}
