package service

import (
	"fmt"
	"log"
)

func New(attr InitAttribute) *ProductService {
	if err := attr.validate(); err != nil {
		log.Panic(err)
	}

	return &ProductService{
		cache:  attr.Cache,
		repo:   attr.Repo,
		config: attr.Config,
	}
}

func (attr InitAttribute) validate() error {
	if !attr.Repo.validate() {
		return fmt.Errorf("missing product repo : %+v", attr.Repo.ProductRepo)
	}

	return nil
}

func (repo RepoAttribute) validate() bool {
	return repo.ProductRepo != nil
}
