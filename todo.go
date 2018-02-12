package main

type CV struct {
	Name       string            `json:"name"`
	Profession string            `json:"profession"`
	City       string            `json:"city"`
	Contacts   map[string]string `json:"contacts"`
}
