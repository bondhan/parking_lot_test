package schema

type (
	//CMDType holds the type of commands
	CMDType = string
)

const (
	// CMDCreateParkingLot command input for create parking lot
	CMDCreateParkingLot CMDType = "create_parking_lot"
	// CMDPark command input for park a car
	CMDPark CMDType = "park"
	// CMDStatus command input for get current status of all parking lots
	CMDStatus CMDType = "status"
	// CMDHelp command input for get help hint for all the commands
	CMDHelp CMDType = "help"
	// CMDExit command input to exit from the interactive shell
	CMDExit CMDType = "exit"
	// CMDShellHistory command input to get all the interactive shell histories
	CMDShellHistory CMDType = "shell_history"
	// CMDParkingHistory command input to get all the interactive shell histories
	CMDParkingHistory CMDType = "park_history"
	// CMDLeave command input leave a parking slot
	CMDLeave CMDType = "leave"
	// CMDRegistrationNumbersWithColor command input to get all registration numbers by color
	CMDRegistrationNumbersWithColor CMDType = "registration_numbers_for_cars_with_colour"
	// CMDSlotNumbersWithColor command input get all slot by color
	CMDSlotNumbersWithColor CMDType = "slot_numbers_for_cars_with_colour"
	// CMDSlotWithRegistration command input get a slot by registration number
	CMDSlotWithRegistration CMDType = "slot_number_for_registration_number"
)

// ValidCommandsByName holds the valid commands map
var ValidCommandsByName = map[string]bool{
	string(CMDCreateParkingLot):             true,
	string(CMDPark):                         true,
	string(CMDStatus):                       true,
	string(CMDHelp):                         true,
	string(CMDExit):                         true,
	string(CMDShellHistory):                 true,
	string(CMDParkingHistory):               true,
	string(CMDLeave):                        true,
	string(CMDRegistrationNumbersWithColor): true,
	string(CMDSlotNumbersWithColor):         true,
	string(CMDSlotWithRegistration):         true,
}

// CMDArgumentLength holds the exact arguments length to read for commands
var CMDArgumentLength = map[string]int{
	string(CMDCreateParkingLot):             1,
	string(CMDPark):                         2,
	string(CMDStatus):                       0,
	string(CMDHelp):                         0,
	string(CMDExit):                         0,
	string(CMDShellHistory):                 0,
	string(CMDParkingHistory):               0,
	string(CMDLeave):                        1,
	string(CMDRegistrationNumbersWithColor): 1,
	string(CMDSlotNumbersWithColor):         1,
	string(CMDSlotWithRegistration):         1,
}
