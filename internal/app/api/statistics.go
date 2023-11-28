package api

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/service"
	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/utils"
	database "github.com/AbdulRafayZia/Gorilla-mux/internal/infrastructure/Database"
)

func Statistics(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var request utils.StatsRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Unable to get data", http.StatusBadRequest)
		return
	}
	tokenString, err := service.GetToken(w, r)
	if tokenString == "" || err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "could not provide autherization bearer", http.StatusUnauthorized)
		return
	}

	Claims, err := service.VerifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, " Could not get Claims ", http.StatusUnauthorized)

		return
	}
	validrole := service.CheckStaffRole(Claims.Role)
	if !validrole {
		http.Error(w, "Not a Staff Member", http.StatusUnauthorized)
		return

	}

	average_execution_time, err := database.GetAvergeExeTime(request)
	if err != nil {
		http.Error(w, "Could not fetch average execution time from database", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(average_execution_time)
}
