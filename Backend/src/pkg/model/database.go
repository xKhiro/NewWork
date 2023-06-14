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
	RoomId                int `gorm:"column:room_Id"`
	Name                  string
	DockingStationPresent bool
	AdjustableDeskPresent bool
	HasTwoScreens         bool
	Bookings              []Booking `gorm:"foreignkey:WorkspaceId"`
}

type Booking struct {
	gorm.Model
	WorkspaceId int       `gorm:"column:workspace_Id"`
	PersonId    string    `gorm:"column:person_Id"`
	Date        time.Time `gorm:"column:date"`
}
