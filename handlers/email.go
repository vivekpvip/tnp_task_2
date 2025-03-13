package handlers

import (
	"certificate-app/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gomail/v2"
	"github.com/gorilla/mux"
)

// SendCertificate handles sending the certificate via email.
func SendCertificate(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var request struct {
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid email", http.StatusBadRequest)
		return
	}

	// Fetch certificate from DB
	var certificate models.Certificate
	if err := models.DB.First(&certificate, id).Error; err != nil {
		http.Error(w, "Certificate not found", http.StatusNotFound)
		return
	}

	// Send email
	if err := sendEmail(request.Email, "Here is your certificate!"); err != nil {
		http.Error(w, "Error sending email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Certificate sent successfully!")
}

func sendEmail(to string, content string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "your_email@gmail.com")
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", "Certificate")
	mailer.SetBody("text/plain", content)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, "your_email@gmail.com", "your_email_password")
	return dialer.DialAndSend(mailer)
}
