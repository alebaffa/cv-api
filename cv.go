package main

type CV struct {
	Name        string            `json:"name"`
	Profession  string            `json:"profession"`
	City        string            `json:"city"`
	Contacts    map[string]string `json:"contacts"`
	Experiences []Experience      `json:"experiences"`
}

type Experience map[string]string

func NewExperience(Dates string, Company string, Position string) Experience {
	return map[string]string{
		"Dates":    Dates,
		"Company":  Company,
		"Position": Position}
}
