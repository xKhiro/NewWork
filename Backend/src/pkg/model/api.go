package model

type BookingDTO struct {
	BookingId   int    `json:"bookingId"`
	PersonId    string `json:"personId"`
	WorkspaceId int    `json:"workspaceId"`
	RoomName    string `json:"roomName"`
	Date        string `json:"date"`
}

type WorkspaceDTO struct {
	WorkspaceId       int          `json:"workspaceId"`
	Name              string       `json:"name"`
	HasDockingStation bool         `json:"hasDockingStation"`
	HasAdjustableDesk bool         `json:"hasAdjustableDesk"`
	HasTwoScreens     bool         `json:"hasTwoScreens"`
	RoomName          string       `json:"roomName"`
	Bookings          []BookingDTO `json:"bookings"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type WorkspaceFilter struct {
	Date              string
	Booked            string
	HasDockingStation string
	HasAdjustableDesk string
	HasTwoScreens     string
	RoomId            string
	WorkspaceId       string
}
