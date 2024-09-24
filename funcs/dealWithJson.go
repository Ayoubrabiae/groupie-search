package funcs

import (
	"encoding/json"
	"errors"
	"net/http"
)

func GetAndParse(url string, v any) error {
	response, err := http.Get(url)
	if response == nil {
		return errors.New("internet connection")
	}

	if err != nil {
		return err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(v)
	if err != nil {
		return err
	}
	return nil
}
