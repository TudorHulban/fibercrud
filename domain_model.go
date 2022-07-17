package main

import (
	"errors"
	"fmt"
)

type Company struct {
	*CompanyData

	repo *RepoCompany
}

func NewCompany(data *CompanyData, repo *RepoCompany) (*Company, error) {
	if data == nil {
		return nil, errors.New("passed data is nil")
	}

	if repo == nil {
		return nil, errors.New("passed repo is nil")
	}

	return &Company{
		CompanyData: data,
		repo:        repo,
	}, nil
}

func NewCompanyEmpty(repo *RepoCompany) (*Company, error) {
	if repo == nil {
		return nil, errors.New("passed repo is nil")
	}

	return &Company{
		repo: repo,
	}, nil
}

func (c *Company) RepoNewCompany() (int, error) {
	errInsert := c.repo.DBConn.Create(c.CompanyData).Error

	return int(c.CompanyData.ID), errInsert
}

func (c *Company) RepoGetCompanyByID(id uint) (*CompanyData, error) {
	var res CompanyData

	rows := c.repo.DBConn.First(&res, id).RowsAffected

	if rows == 1 {
		return &res, nil
	}

	return nil, errRecordNotFound
}

func (c *Company) RepoDeleteCompany(id uint) error {
	var res CompanyData

	rows := c.repo.DBConn.Delete(&res, id).RowsAffected

	if rows == 1 {
		return nil
	}

	return errRecordNotFound
}

func (c *Company) RepoGetCompanies() ([]CompanyData, error) {
	var res []CompanyData

	errSelect := c.repo.DBConn.Find(&res).Error

	return res, errSelect
}

// RepoUpdateCompany should update the entity.
func (c *Company) RepoUpdateCompany() error {
	fmt.Println("")

	fmt.Printf("RepoUpdateCompany:\n%#v\n", *c.CompanyData)

	rows := c.repo.DBConn.Model(c.CompanyData).Updates(c.CompanyData).RowsAffected

	if rows == 1 {
		return nil
	}

	return errRecordNotFound
}
