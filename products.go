package main

import (
	"encoding/json"
	"os"
)

type Product struct {
	ID             string            `json:"id"`
	Name           string            `json:"name"`
	Airflow        string            `json:"airflow"`
	Description    string            `json:"description"`
	TechnicalSpecs map[string]string `json:"technical_specs"`
}

func readProducts() ([]Product, error) {
	productsBts, err := os.ReadFile("company/products.json")
	if err != nil {
		return nil, err
	}

	var products []Product
	if err := json.Unmarshal(productsBts, &products); err != nil {
		return nil, err
	}

	return products, nil
}
