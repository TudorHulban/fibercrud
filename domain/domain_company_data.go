package domain

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

type CompanyData struct {
	gorm.Model

	Code    string `json:"code"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Website string `json:"website"`
	Phone   string `json:"phone"` // for supporting +40 type
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

	// fmt.Printf("%#v\n", c)

	return nil
}
