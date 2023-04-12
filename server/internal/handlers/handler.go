package handlers

import (
	"encoding/json"
	"net/http"
	"server/internal/dto"
)

type Handler struct {
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(dto.NewCommonResponse(data)); err != nil {
		panic(err)
	}
}

func (h Handler) GetData(w http.ResponseWriter, r *http.Request) {
	crimeData := dto.CrimeData{
		Image:       "https://webdamn.com/wp-content/uploads/2020/06/golang-min.png",
		Title:       "Gopher",
		Description: "Being a Gopher",
	}
	response := []dto.CrimeData{
		crimeData, crimeData, crimeData, crimeData,
	}
	writeResponse(w, 200, response)
}
