package store

import (
	"fmt"
	"parking_lot/errors"
	"parking_lot/schema"
	"parking_lot/utils"
)

type infoStore struct {
	*store
}

// NewInfoStore returns new store object
func NewInfoStore(st *store) *infoStore {
	pl := &infoStore{st}
	return pl
}

func (pl *infoStore) IsHelp(arg string) (string, bool) {
	if arg == string(schema.CMDHelp) {
		return schema.CMDLeaveHint, true
	}
	return "", false
}

// Execute - `leave` Command will takes registration number and colour as Arguments
// the system checks for a first availabe slot to park, if slot available
// slot will allocated to the vehicle.
// This will checks if the vehicle registration number is duplicate or not.
func (pl *infoStore) Execute(cmd *schema.Command) (string, error) {

	var err error

	if res, isHelp := pl.IsHelp(cmd.Arguments[0]); isHelp {
		return res, nil
	}
	if ParkingLot == nil {
		return "", errors.ErrNoParkingLot
	}

	response := ""
	switch cmd.Command {
	case schema.CMDRegistrationNumbersWithColor:

		if err = validateColor(cmd.Arguments); err != nil {
			return "", err
		}
		response, err = ParkingLot.GetAllRegistrationsByCarColor(cmd.Arguments[0])
		return fmt.Sprint(response), nil

	case schema.CMDSlotNumbersWithColor:
		if err = validateColor(cmd.Arguments); err != nil {
			return "", err
		}

		response, err = ParkingLot.GetAllSlotsByCarColor(cmd.Arguments[0])
		return fmt.Sprint(response), nil

	case schema.CMDSlotWithRegistration:
		if err = validateRegNo(cmd.Arguments); err != nil {
			return "", err
		}

		response, err = ParkingLot.GetSlotByRegNo(cmd.Arguments[0])
		if err != nil {
			return "", err
		}
		return fmt.Sprint(response), nil
	}

	return response, err
}

func validateRegNo(args []string) error {
	if !utils.IsRegNoValid(args[0]) {
		return errors.ErrInvalidRegNo
	}
	return nil
}

func validateColor(args []string) error {
	if !utils.IsValidString(args[0]) {
		return errors.ErrInvalidColour
	}
	return nil
}
