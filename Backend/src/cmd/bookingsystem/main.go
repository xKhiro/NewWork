package main

import (
	"NewWork/pkg/routers"
	_ "github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	log.Println("Start server")
	router := mux.NewRouter()

	myRouter := routers.NewRouter()

	log.Println(router)

	router.HandleFunc("/workspaces", myRouter.GetWorkspaces).Methods("GET")
	router.HandleFunc("/users/{personId}/bookings", myRouter.GetBookings).Methods("GET")
	router.HandleFunc("/users/{personId}/bookings", myRouter.CreateBooking).Methods("POST")
	router.HandleFunc("/users/{personId}/bookings/{bookingId}", myRouter.CancelBooking).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
