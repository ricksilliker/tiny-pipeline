package api

type Asset struct {
	Name string
	Status *Status
	Project *Project
	Tags []string
}