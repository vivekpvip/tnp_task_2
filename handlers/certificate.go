package handlers

import (
	"certificate-app/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetCertificateByID handles fetching a certificate by its ID.
func GetCertificateByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var certificate models.Certificate
	if err := models.DB.First(&certificate, id).Error; err != nil {
		http.Error(w, "Certificate not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(certificate)
}

// CreateCertificate handles creating a new certificate.
func CreateCertificate(w http.ResponseWriter, r *http.Request) {
	var newCertificate models.Certificate
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newCertificate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := models.DB.Create(&newCertificate).Error; err != nil {
		http.Error(w, "Error creating certificate", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCertificate)
}
