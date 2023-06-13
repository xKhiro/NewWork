package handlers

import "NewWork/pkg/model"

var bookings = []model.Booking{
	{BookingId: 1, PersonId: "123", WorkspaceId: 1, RoomName: "room1", Date: "2023-01-01"},
	{BookingId: 2, PersonId: "456", WorkspaceId: 2, RoomName: "room2", Date: "2023-01-02"},
}

var workspaces = []model.Workspace{
	{WorkspaceId: 1, Name: "Workspace1", HasDockingStation: true, HasAdjustableDesk: true, HasTwoScreens: true, RoomName: "Room1", Bookings: bookings},
	{WorkspaceId: 2, Name: "Workspace2", HasDockingStation: false, HasAdjustableDesk: true, HasTwoScreens: false, RoomName: "Room2", Bookings: bookings},
}

type WorkspaceHandler struct {
	workspaces []model.Workspace
}

func NewWorkspaceHandler() *WorkspaceHandler {

	return &WorkspaceHandler{workspaces: workspaces}
}

func (h *WorkspaceHandler) GetWorkspaces(params model.GetWorkspacesParams) []model.Workspace {
	// Hier können Sie die Logik hinzufügen, um die Arbeitsplätze zu filtern, basierend auf den Parametern
	return h.workspaces
}

type BookingHandler struct {
	bookings []model.Booking
}

func NewBookingHandler() *BookingHandler {

	return &BookingHandler{bookings: bookings}
}

func (h *BookingHandler) GetBookings(personId string) []model.Booking {
	// Hier können Sie die Logik hinzufügen, um die Buchungen zu filtern, basierend auf der personId
	return h.bookings
}

func (h *BookingHandler) CreateBooking(personId string, newBooking model.Booking) model.Booking {
	// Hier können Sie die Logik hinzufügen, um eine neue Buchung zu erstellen
	h.bookings = append(h.bookings, newBooking)
	return newBooking
}

func (h *BookingHandler) CancelBooking(personId string, bookingId int) error {
	// Hier können Sie die Logik hinzufügen, um eine Buchung zu stornieren
	return nil
}
