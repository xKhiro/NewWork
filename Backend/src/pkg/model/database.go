package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Room struct {
	gorm.Model
	RoomId     int
	Name       string
	MaxSeats   int
	Workspaces []Workspace `gorm:"foreignkey:RoomId"`
}

type Workspace struct {
	gorm.Model
	WorkspaceId           int
	RoomId                int
	Name                  string
	DockingStationPresent bool
	AdjustableDeskPresent bool
	NumberOfMonitors      int
	Bookings              []Booking `gorm:"foreignkey:WorkspaceId"`
}

type Booking struct {
	gorm.Model
	BookingId   int
	WorkspaceId int
	PersonId    string
	Date        time.Time
}
