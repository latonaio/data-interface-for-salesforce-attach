package handlers

import "errors"

const(
	contentTypeContractPDF = "ContractPDF"
)

func Handle(metadata map[string]interface{}) error{
	keyIF, exists := metadata["key"]
	if !exists {
		return errors.New("key is required")
	}
	key, ok := keyIF.(string)
	if !ok {
		return errors.New("key is required")
	}
	if key == contentTypeContractPDF {
		return HandleContractPDF(metadata)
	}
	return nil
}