package helpers

import (
	"encoding/json"
	"fmt"

	"os"
)

func CreateModel(modelName string, v interface{}) error {
	bytes, err := json.MarshalIndent(v, "", "   ")
	_ = checkReturnError(err)
	file, err := os.OpenFile(fmt.Sprintf("./datastore/%s.json", modelName), os.O_CREATE|os.O_RDWR, 0600)
	_ = checkReturnError(err)
	defer file.Close()

	//chunk of code for json write build structure
	fstat, err := file.Stat()
	_ = checkReturnError(err)
	fSize := fstat.Size()

	if fSize < 1 {
		_, err = file.WriteAt([]byte(`[`), 0)
		_ = checkReturnError(err)
		bytes = append(bytes, ']')
		_, err = file.WriteAt(bytes, fSize+1)
		_ = checkReturnError(err)
	}
	if fSize > 1 {
		bytes = append(bytes, ']')
		_, err = file.WriteAt([]byte(`, \n`), fSize-1)
		_, err = file.WriteAt(bytes, fSize+1)
		_ = checkReturnError(err)
	}
	return err
}

func checkReturnError(err error) error {
	if err != nil {
		err.Error()
	}
	return err
}
