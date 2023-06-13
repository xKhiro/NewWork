package handlers

import (
	"NewWork/pkg/model"
	"NewWork/pkg/storge/mysql"
	"context"
	"log"
)

type Handler struct {
	storage *mysql.GormStorage
}

func NewHandler() *Handler {

	storage := mysql.NewGormStorage()

	return &Handler{storage: storage}
}

func (h *Handler) GetWorkspaces(params model.WorkspaceFilter) []model.WorkspaceDTO {
	log.Print("GetWorkspaces with params")

	workspaces, err := h.storage.GetAllWorkspaces(context.TODO(), &params)
	if err != nil {
		return nil
	}

	log.Println("Workspaces: ", workspaces)

	return nil
}

func (h *Handler) GetBookings(personId string) []model.BookingDTO {
	// Hier können Sie die Logik hinzufügen, um die Buchungen zu filtern, basierend auf der personId
	return nil
}

func (h *Handler) CreateBooking(personId string, newBooking model.BookingDTO) model.BookingDTO {
	// Hier können Sie die Logik hinzufügen, um eine neue Buchung zu erstellen
	//h.bookings = append(h.bookings, newBooking)
	return newBooking
}

func (h *Handler) CancelBooking(personId string, bookingId int) error {
	// Hier können Sie die Logik hinzufügen, um eine Buchung zu stornieren
	return nil
}
