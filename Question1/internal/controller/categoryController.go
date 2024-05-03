package controller

import (
	"backend/internal/service"
	"encoding/json"
	"net/http"
)
func CategoryNameHandler(w http.ResponseWriter, r *http.Request) {
		cat := r.PathValue("categoryname")
		n := r.PathValue("n")
		res,err := service.FetchData(cat,n)
		if err!=nil{
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
    	json.NewEncoder(w).Encode(res.Body)
	}