package handlers

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/latonaio/aion-core/pkg/log"
	"github.com/latonaio/data-interface-for-salesforce-attach/pkg/db/models"
)

func HandleContractPDF(metadata map[string]interface{}) error {
	jsonIF, exists := metadata["content"]
	if !exists {
		return errors.New("content is required")
	}
	jsonStr, ok := jsonIF.(string)
	if !ok {
		return errors.New("failed to convert interface{} to string")
	}
	var resp Resp
	if err := json.Unmarshal([]byte(jsonStr), &resp); err != nil {
		return fmt.Errorf("failed to json unmarshal: %v", err)
	}
	if *resp.Id == "" || *resp.FileName == "" {
		return errors.New("invalid contract id or file name")
	}
	cPdf := models.ContractPDF{
		SfContractID: resp.Id,
		FileName:     resp.FileName,
	}
	if err := cPdf.Register(); err != nil {
		return fmt.Errorf("failed to register contract PDF: %v", err)
	}
	log.Print("registered pdf.")
	return nil
}

type Resp struct {
	Id       *string `json:"id"`
	FileName *string `json:"file_name"`
}
