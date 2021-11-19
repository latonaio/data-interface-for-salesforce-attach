package models

import (
	"errors"

	"github.com/latonaio/data-interface-for-salesforce-attach/pkg/db"
	"gorm.io/gorm"
)

type ContractPDF struct {
	SfContractID *string `gorm:"column:sf_contract_id"`
	FileName     *string `gorm:"column:file_name"`
}

func (c *ContractPDF) TableName() string {
	const tableName = "contract_pdf"
	return tableName
}

func (c *ContractPDF) Register() error {
	result := db.ConPool.Con.Create(c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// ContractPDFByContractID gets contractPDF records by contract id.
// Returns nil, nil if not found.
func ContractPDFByContractID(id string) ([]ContractPDF, error) {
	var contractPDF []ContractPDF
	result := db.ConPool.Con.Model(ContractPDF{}).Where("sf_contract_id = ?", id).Find(&contractPDF)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return contractPDF, nil
}
