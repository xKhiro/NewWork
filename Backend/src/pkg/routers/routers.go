package routers

import (
	"NewWork/pkg/handlers"
	"NewWork/pkg/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Router struct {
	handler *handlers.Handler
}

func NewRouter() *Router {
	log.Println("New Router")

	workspaceHandler := handlers.NewHandler()

	return &Router{
		handler: workspaceHandler,
	}
}

func (r *Router) GetWorkspaces(w http.ResponseWriter, req *http.Request) {

	log.Println("Router GetWorkspaces")

	params := model.WorkspaceFilter{
		Date:              req.URL.Query().Get("date"),
		Booked:            parseBool(req.URL.Query().Get("booked")),
		HasDockingStation: parseBool(req.URL.Query().Get("hasDockingStation")),
		HasAdjustableDesk: parseBool(req.URL.Query().Get("hasAdjustableDesk")),
		HasTwoScreens:     parseBool(req.URL.Query().Get("hasTwoScreens")),
		RoomId:            req.URL.Query().Get("roomId"),
		WorkspaceId:       req.URL.Query().Get("workspaceId"),
	}

	workspaces := r.handler.GetWorkspaces(params)
	json.NewEncoder(w).Encode(workspaces)
}

func parseBool(boolStr string) bool {
	if boolStr != "" {
		booked, err := strconv.ParseBool(boolStr)
		if err != nil {
			log.Panicln("Cannot Parse: ", err)
		}
		return booked
	}
	return false

}

func (r *Router) GetBookings(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	personId := vars["personId"]

	bookings := r.handler.GetBookings(personId)

	json.NewEncoder(w).Encode(bookings)
}

func (r *Router) CreateBooking(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	personId := vars["personId"]

	var booking model.BookingDTO
	_ = json.NewDecoder(req.Body).Decode(&booking)

	error := r.handler.CreateBooking(personId, booking)

	if error.Code == 409 {
		w.WriteHeader(http.StatusConflict)
	} else if error.Code == 429 {
		w.WriteHeader(http.StatusTooManyRequests)
	} else if error.Code == 500 {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	json.NewEncoder(w).Encode(error)
}

func (r *Router) CancelBooking(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	personId := vars["personId"]

	bookingIdString := vars["bookingId"]

	bokkingId, err := strconv.Atoi(bookingIdString)
	if err != nil {
		return
	}

	err = r.handler.CancelBooking(personId, bokkingId)
	w.WriteHeader(http.StatusNoContent)

}
