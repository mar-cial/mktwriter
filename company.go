package main

import (
	"os"
)

type Company struct {
	About    string    `json:"about"`
	Products []Product `json:"products"`
}

func readAbout() (string, error) {
	aboutBts, err := os.ReadFile("company/about.md")
	if err != nil {
		return "", err
	}

	return string(aboutBts), nil
}

func loadCompany() (*Company, error) {
	about, err := readAbout()
	if err != nil {
		return nil, err

	}

	products, err := readProducts()
	if err != nil {
		return nil, err
	}

	company := &Company{
		About:    about,
		Products: products,
	}

	return company, nil
}
