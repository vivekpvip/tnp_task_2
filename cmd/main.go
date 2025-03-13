package main

import (
	"certificate-app/handlers"
	"certificate-app/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database
	utils.InitDB()

	r := mux.NewRouter()

	// Define routes
	r.Handle("/certificates/{id}", handlers.AuthMiddleware(http.HandlerFunc(handlers.GetCertificateByID))).Methods("GET")
	r.Handle("/certificates", handlers.AuthMiddleware(http.HandlerFunc(handlers.CreateCertificate))).Methods("POST")
	r.Handle("/certificates", handlers.AuthMiddleware(http.HandlerFunc(handlers.GetAllCertificates))).Methods("GET")
	r.Handle("/certificates/{id}", handlers.AuthMiddleware(http.HandlerFunc(handlers.UpdateCertificate))).Methods("PUT")
	r.Handle("/send/{id}", handlers.AuthMiddleware(http.HandlerFunc(handlers.SendCertificate))).Methods("POST")
	r.Handle("/send_bulk", handlers.AuthMiddleware(http.HandlerFunc(handlers.SendBulkEmail))).Methods("POST")

	// Start the server
	log.Fatal(http.ListenAndServe(":8000", r))
}
