package model

import (
	"database/sql/driver"
	"fmt"
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
	WorkspaceId int
	PersonId    string
	Date        time.Time
}

type Date struct {
	time.Time
}

func (t *Date) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		newTime, err := time.Parse("2006-01-02", string(v))
		if err != nil {
			return err
		}
		*t = Date{newTime}
	case string:
		newTime, err := time.Parse("2006-01-02", v)
		if err != nil {
			return err
		}
		*t = Date{newTime}
	default:
		return fmt.Errorf("invalid type for Date")
	}
	return nil
}

func (t Date) Value() (driver.Value, error) {
	return t.Format("2006-01-02"), nil
}
