package handlers

import (
	"NewWork/pkg/model"
	"NewWork/pkg/storge/mysql"
	"context"
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

	log.Print("GetWorkspaces with params: ", params)

	workspaces, err := h.storage.GetAllWorkspaces(context.TODO(), &params)
	if err != nil {
		return nil
	}

	unfiltered := MapWorkspacesToDTO(workspaces)

	for i, workspaceDto := range unfiltered {

		bookings, err := h.storage.GetBookingsWithWorkspaceId(context.Background(), workspaceDto.WorkspaceId)
		if err != nil {
			log.Panicln(err)
		}

		workspaceDto.Bookings = h.mapBookingsToBookingDTO(bookings)
		unfiltered[i] = workspaceDto
	}

	workspacesDTO := FilterWorkspaces(unfiltered, params)

	log.Println("Workspace: ", workspacesDTO)
	return workspacesDTO
}

func (h *Handler) GetBookings(personId string) []model.BookingDTO {

	log.Println("GetBookings with personId")

	bookings, err := h.storage.GetAllBookings(context.Background(), personId)

	if err != nil {
		return nil
	}

	log.Println("Bookings: ", bookings)

	bookingDTO := h.mapBookingsToBookingDTO(bookings)

	return bookingDTO
}

func (h *Handler) CreateBooking(personId string, newBooking model.BookingDTO) model.Error {

	log.Print("CreateBooking for personId: ", personId, newBooking)

	booking, err := MapToBooking(newBooking)
	if err != nil {
		log.Panicln("Error: ", err)
		return model.Error{
			Code:    500,
			Message: "Internal Error",
		}
	}

	log.Println("New Booking: ", booking)
	error := h.storage.CreateBooking(context.TODO(), booking)

	return error
}

func (h *Handler) CancelBooking(personId string, bookingId int) error {

	err := h.storage.DeleteBooking(context.Background(), bookingId)
	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) mapBookingsToBookingDTO(bookings []model.Booking) []model.BookingDTO {

	bookingsDTO := make([]model.BookingDTO, 0)
	for _, booking := range bookings {
		workspace, err := h.storage.GetWorkspace(context.Background(), booking.WorkspaceId)
		log.Println("Workspace for booking: ", workspace)
		if err != nil {
			log.Panicln("Get Workspace Error")
		}

		bookingDTO := model.BookingDTO{
			BookingId:   int(booking.ID),
			PersonId:    booking.PersonId,
			WorkspaceId: booking.WorkspaceId,
			RoomName:    checkroomID(workspace.RoomId),
			Date:        booking.Date.Format("2006-01-02"),
		}

		bookingsDTO = append(bookingsDTO, bookingDTO)
	}

	return bookingsDTO

}

func MapToBooking(dto model.BookingDTO) (model.Booking, error) {

	return model.Booking{
		//Model:       gorm.Model{ID: uint(dto.BookingId)}, // gorm.Model.ID is of type uint
		WorkspaceId: dto.WorkspaceId,
		PersonId:    dto.PersonId,
		Date:        parseTime(dto.Date),
	}, nil
}

func parseTime(date string) time.Time {

	time, err := time.Parse("2006-01-02", date)

	if err != nil {
		log.Panicln("Parse error: ", err)
	}
	return time

}

func MapWorkspacesToDTO(workspaces []model.Workspace) []model.WorkspaceDTO {
	workspacesDTO := make([]model.WorkspaceDTO, len(workspaces))

	for i, workspace := range workspaces {
		bookingsDTO := make([]model.BookingDTO, len(workspace.Bookings))
		for j, booking := range workspace.Bookings {
			bookingsDTO[j] = model.BookingDTO{
				BookingId:   int(booking.ID),
				PersonId:    booking.PersonId,
				WorkspaceId: booking.WorkspaceId,
				RoomName:    checkroomID(workspace.RoomId),
				Date:        booking.Date.Format("2006-01-02"), // Format the date as "YYYY-MM-DD"
			}
		}

		workspacesDTO[i] = model.WorkspaceDTO{
			WorkspaceId:       int(workspace.ID),
			Name:              workspace.Name,
			HasDockingStation: workspace.DockingStationPresent,
			HasAdjustableDesk: workspace.AdjustableDeskPresent,
			HasTwoScreens:     workspace.HasTwoScreens,
			RoomName:          checkroomID(workspace.RoomId),
			Bookings:          bookingsDTO,
		}
	}

	return workspacesDTO
}

func checkroomID(id int) string {

	switch id {
	case 1:
		return "Raum 01"
	case 2:
		return "Raum 02"
	case 3:
		return "Raum 03"
	case 4:
		return "Raum 04"
	case 5:
		return "Raum 05"
	default:
		return "Raum"
	}
}

func FilterWorkspaces(workspaces []model.WorkspaceDTO, filter model.WorkspaceFilter) []model.WorkspaceDTO {
	filtered := make([]model.WorkspaceDTO, 0)

	for _, workspace := range workspaces {
		if filter.Booked {
			if filter.Date != "" {

				for _, booking := range workspace.Bookings {
					if booking.Date == filter.Date {
						filtered = append(filtered, workspace)
						break
					}
				}
			} else {

				if len(workspace.Bookings) > 0 {
					filtered = append(filtered, workspace)
				}
			}
		} else if filter.Date != "" {

			hasBooking := false
			for _, booking := range workspace.Bookings {
				if booking.Date == filter.Date {
					hasBooking = true
					break
				}
			}
			if !hasBooking {
				filtered = append(filtered, workspace)
			}
		}
	}

	return filtered
}
