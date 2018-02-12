package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func PrintCV(w http.ResponseWriter, r *http.Request) {
	cv := CV{
		Name:       "Alessandro Baffa",
		Profession: "Senior Software Engineer",
		City:       "Tokyo, Japan",
		Contacts: map[string]string{
			"phone":    "(+81)080-2556-1845)",
			"email":    "contact@alebaffa.com",
			"blog":     "alebaffa.com",
			"twitter":  "@alebaffa",
			"linkedIn": "in/alessandrobaffa"},
		Experiences: []Experience{NewExperience(
			"2012, 2018",
			"Amadeus IT Group",
			"Developer Advocate")}}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cv); err != nil {
		panic(err)
	}
}
