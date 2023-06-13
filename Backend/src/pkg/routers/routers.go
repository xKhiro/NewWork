package routers

import (
	"NewWork/pkg/handlers"
	"NewWork/pkg/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Router struct {
	workspaceHandler *handlers.WorkspaceHandler
	bookingHandler   *handlers.BookingHandler
}

func NewRouter() *Router {
	workspaceHandler := handlers.NewWorkspaceHandler()
	bookingHandler := handlers.NewBookingHandler()

	return &Router{
		workspaceHandler: workspaceHandler,
		bookingHandler:   bookingHandler,
	}
}

func (r *Router) GetWorkspaces(w http.ResponseWriter, req *http.Request) {
	params := model.GetWorkspacesParams{
		Date:              req.URL.Query().Get("date"),
		Booked:            req.URL.Query().Get("booked"),
		HasDockingStation: req.URL.Query().Get("hasDockingStation"),
		HasAdjustableDesk: req.URL.Query().Get("hasAdjustableDesk"),
		HasTwoScreens:     req.URL.Query().Get("hasTwoScreens"),
		RoomId:            req.URL.Query().Get("roomId"),
		WorkspaceId:       req.URL.Query().Get("workspaceId"),
	}
	workspaces := r.workspaceHandler.GetWorkspaces(params)
	json.NewEncoder(w).Encode(workspaces)
}

func (r *Router) GetBookings(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	personId := vars["personId"]

	bookings := r.bookingHandler.GetBookings(personId)

	json.NewEncoder(w).Encode(bookings)
}

func (r *Router) CreateBooking(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	personId := vars["personId"]

	var booking model.Booking
	_ = json.NewDecoder(req.Body).Decode(&booking)

	newBooking := r.bookingHandler.CreateBooking(personId, booking)

	json.NewEncoder(w).Encode(newBooking)
}

func (r *Router) CancelBooking(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	personId := vars["personId"]

	bookingIdString := vars["bookingId"]

	bokkingId, err := strconv.Atoi(bookingIdString)
	if err != nil {
		return
	}

	err = r.bookingHandler.CancelBooking(personId, bokkingId)
	if err != nil {
		return
	}
}
