package model

import (
	"github.com/jinzhu/gorm"
)

type Workspace struct {
	gorm.Model
	RoomID   int
	Name     string
	MaxSeats int
	Desks    []Desk `gorm:"foreignkey:RoomID"`
}

type Desk struct {
	gorm.Model
	DeskID                int
	RoomID                int
	Name                  string
	DockingStationPresent bool
	AdjustableDeskPresent bool
	NumberOfMonitors      int
	Bookings              []Booking `gorm:"foreignkey:DeskID"`
}

type Booking struct {
	gorm.Model
	BookingID int
	DeskID    int
	PersonID  string
	Date      string
}
