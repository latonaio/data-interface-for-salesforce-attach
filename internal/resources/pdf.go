package resources

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
)

type File struct {
	method   string
	metadata map[string]interface{}
}

func NewFile(metadata map[string]interface{}) (*File, error) {
	rawMethod, ok := metadata["method"]
	if !ok {
		return nil, errors.New("missing required parameters: method")
	}
	method, ok := rawMethod.(string)
	if !ok {
		return nil, errors.New("failed to convert interface{} to string")
	}
	return &File{
		method:   method,
		metadata: metadata,
	}, nil
}

func (f *File) getMetadata() (map[string]interface{}, error) {
	return nil, nil
}

func (f *File) postMetadata() (map[string]interface{}, error) {
	objectNameIF, exists := f.metadata["object_name"]
	if !exists {
		return nil, errors.New("object_name is required")
	}
	objectName, ok := objectNameIF.(string)
	if !ok {
		return nil, errors.New("failed to convert interface{} to string")
	}
	idIF, exists := f.metadata["object_id"]
	if !exists {
		return nil, errors.New("object_id is required")
	}
	id, ok := idIF.(string)
	if !ok {
		return nil, errors.New("failed to convert interface{} to string")
	}
	fileNameIF, exists := f.metadata["file_name"]
	if !exists {
		return nil, errors.New("file_name is required")
	}
	fileName, ok := fileNameIF.(string)
	if !ok {
		return nil, errors.New("failed to convert interface{} to string")
	}
	pathIF, exists := f.metadata["path"]
	if !exists {
		return nil, errors.New("path is required")
	}
	path, ok := pathIF.(string)
	if !ok {
		return nil, errors.New("failed to convert interface{} to string")
	}
	fileExtensionIF, exists := f.metadata["file_extension"]
	if !exists {
		return nil, errors.New("file_extension is required")
	}
	fileExtension, ok := fileExtensionIF.(string)
	if !ok {
		return nil, errors.New("failed to convert interface{} to string")
	}

	fName := fileName + "." + fileExtension
	filePath := path + "/" + fName
	body, err := ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}
	return buildBase64Metadata(
		"attach",
		"post",
		objectName,
		id,
		map[string]string{fileExtension + "Name": fName},
		body,
	), nil
}

func (f *File) updateMetadata() (map[string]interface{}, error) {
	objectNameIF, exists := f.metadata["object_name"]
	if !exists {
		return nil, errors.New("object_name is required")
	}
	objectName, ok := objectNameIF.(string)
	if !ok {
		return nil, errors.New("failed to convert interface{} to string")
	}
	idIF, exists := f.metadata["object_id"]
	if !exists {
		return nil, errors.New("object_id is required")
	}
	id, ok := idIF.(string)
	if !ok {
		return nil, errors.New("failed to convert interface{} to string")
	}
	fileNameIF, exists := f.metadata["file_name"]
	if !exists {
		return nil, errors.New("file_name is required")
	}
	fileName, ok := fileNameIF.(string)
	if !ok {
		return nil, errors.New("failed to convert interface{} to string")
	}
	pathIF, exists := f.metadata["path"]
	if !exists {
		return nil, errors.New("path is required")
	}
	path, ok := pathIF.(string)
	if !ok {
		return nil, errors.New("failed to convert interface{} to string")
	}
	fileExtensionIF, exists := f.metadata["file_extension"]
	if !exists {
		return nil, errors.New("file_extension is required")
	}
	fileExtension, ok := fileExtensionIF.(string)
	if !ok {
		return nil, errors.New("failed to convert interface{} to string")
	}

	fName := fileName + "." + fileExtension
	filePath := path + "/" + fName
	body, err := ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}
	return buildBase64Metadata(
		"attach",
		"put",
		objectName,
		id,
		map[string]string{fileExtension + "Name": fName},
		body,
	), nil
}

// BuildMetadata
func (f *File) BuildMetadata() (map[string]interface{}, error) {
	switch f.method {
	case "get":
		return f.getMetadata()
	case "post":
		return f.postMetadata()
	case "put":
		return f.updateMetadata()
	}
	return nil, errors.New("invalid method")
}

func buildBase64Metadata(connectionKey, method, object, pathParam string, queryParams map[string]string, body []byte) map[string]interface{} {
	metadata := map[string]interface{}{
		"method":         method,
		"object":         object,
		"connection_key": connectionKey,
		"is_body_base64": true,
	}
	if len(pathParam) > 0 {
		metadata["path_param"] = pathParam
	}
	if queryParams != nil {
		metadata["query_params"] = queryParams
	}
	if len(body) > 0 {
		metadata["body"] = base64.StdEncoding.EncodeToString(body)
	}
	return metadata
}

var ReadFile = func(path string) ([]byte, error) {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: file path: %s: %v", path, err)
	}
	return body, nil
}
