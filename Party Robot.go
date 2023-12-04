package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
    var chTable string
    if (table < 10) {
    	chTable = fmt.Sprintf("00%d", table)
    } else if (table < 100) {
    	chTable = fmt.Sprintf("0%d", table)
    } else {
    	chTable = fmt.Sprintf("%d", table)  
    }
	return fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table %s. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.", name, chTable, direction, distance, neighbor)
}
