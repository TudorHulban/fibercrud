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

func (c *Company) RepoNewCompany() int {
	c.repo.DBConn.Create(c.CompanyData)

	return int(c.CompanyData.ID)
}

func (c *Company) RepoGetCompany(id uint) (*CompanyData, error) {
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

func (c *Company) RepoGetCompanies() []CompanyData {
	var res []CompanyData

	c.repo.DBConn.Find(&res)

	return res
}

// RepoUpdateCompany should update the entity.
// TODO: does not work
func (c *Company) RepoUpdateCompany() error {
	fmt.Println("")

	fmt.Printf("RepoUpdateCompany:\n%#v\n", *c.CompanyData)

	rows := c.repo.DBConn.Model(c.CompanyData).Where("id = ?", c.CompanyData.ID).Updates(c.CompanyData).RowsAffected

	if rows == 1 {
		return nil
	}

	return errRecordNotFound
}
