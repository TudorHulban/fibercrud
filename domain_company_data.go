package main

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

type CompanyData struct {
	gorm.Model
	Code    string
	Name    string
	Country string
	Website string
	Phone   string // for supporting +40 type
}

var errRecordNotFound = errors.New("record not found")

func NewCompanyData(code, name, country, website, phone string) (*CompanyData, error) {
	res := CompanyData{
		Code:    code,
		Name:    name,
		Country: country,
		Website: website,
		Phone:   phone,
	}

	if errValid := res.IsValid(); errValid != nil {
		return nil, fmt.Errorf("NewCompany data.isValid: %w", errValid)
	}

	return &res, nil
}

func (c CompanyData) IsValid() error {
	// TODO: logic

	return nil
}
