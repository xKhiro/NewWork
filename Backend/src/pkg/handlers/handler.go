package handlers

import (
	"NewWork/pkg/model"
	"NewWork/pkg/storge/mysql"
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"time"
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

	log.Println("GetBookings with personId")

	bookings, err := h.storage.GetAllBookings(context.Background(), personId)

	if err != nil {
		return nil
	}

	log.Println("Bookings: ", bookings)

	bookingDTO := mapBookingsToBookingDTO(bookings)

	return bookingDTO
}

func mapBookingsToBookingDTO(bookings []model.Booking) []model.BookingDTO {

	var bookingsDTO []model.BookingDTO

	for _, booking := range bookings {
		bookingDTO := model.BookingDTO{
			BookingId:   booking.BookingId,
			PersonId:    booking.PersonId,
			WorkspaceId: booking.WorkspaceId,
			RoomName:    "Romm1",
			Date:        booking.Date.String(),
		}

		bookingsDTO = append(bookingsDTO, bookingDTO)
	}

	return bookingsDTO

}

func (h *Handler) CreateBooking(personId string, newBooking model.BookingDTO) error {

	log.Print("CreateBooking for personId: ", personId)

	booking, err := MapToBooking(newBooking)
	if err != nil {
		log.Panicln("Error: ", err)
		return err
	}

	log.Println("New Booking: ", booking)
	err = h.storage.CreateBooking(context.TODO(), booking)
	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) CancelBooking(personId string, bookingId int) error {
	// Hier können Sie die Logik hinzufügen, um eine Buchung zu stornieren
	return nil
}

func MapToBooking(dto model.BookingDTO) (model.Booking, error) {
	layout := "2006-01-02" // This is the standard Go layout for date. Adjust this to the format of your input date.
	t, err := time.Parse(layout, dto.Date)
	if err != nil {
		fmt.Println(err)
		return model.Booking{}, err
	}

	return model.Booking{
		Model:       gorm.Model{ID: uint(dto.BookingId)}, // gorm.Model.ID is of type uint
		BookingId:   dto.BookingId,
		WorkspaceId: dto.WorkspaceId,
		PersonId:    dto.PersonId,
		Date:        t,
	}, nil
}
