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

	router.Use(corsMiddleware)

	myRouter := routers.NewRouter()

	router.HandleFunc("/workspaces", myRouter.GetWorkspaces).Methods("GET")
	router.HandleFunc("/users/{personId}/bookings", myRouter.GetBookings).Methods("GET")
	router.HandleFunc("/users/{personId}/bookings", myRouter.CreateBooking).Methods("POST", "OPTIONS")
	router.HandleFunc("/users/{personId}/bookings/{bookingId}", myRouter.CancelBooking).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Setze die erforderlichen CORS-Header
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Wenn der Preflight-Request (OPTIONS) gesendet wird, beantworte ihn sofort
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Rufe den nächsten Handler auf
		next.ServeHTTP(w, r)
	})
}
