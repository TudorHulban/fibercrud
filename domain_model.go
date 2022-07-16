package main

import "fmt"

type CompanyData struct {
	Code    string // PK
	Name    string
	Country string
	Website string
	Phone   string // for supporting +40 type
}

func (c CompanyData) isValid() error {
	// TODO: logic

	return nil
}

type Company struct {
	CompanyData

	repo *RepoCompany
}

func NewCompany(code, name, country, website, phone string) (*Company, error) {
	data := CompanyData{
		Code:    code,
		Name:    name,
		Country: country,
		Website: website,
		Phone:   phone,
	}

	if errValid := data.isValid(); errValid != nil {
		return nil, fmt.Errorf("NewCompany data.isValid: %w", errValid)
	}

	return &Company{
		CompanyData: data,
	}, nil
}
