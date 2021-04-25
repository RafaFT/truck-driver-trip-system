package presenter

import (
	"encoding/json"

	"github.com/rafaft/truck-driver-trip-system/adapter/controller/rest"
)

type deleteDriverOutputError struct {
	Error string `json:"error"`
}

// output port implementation - Presenter
type deleteDriver struct {
}

func NewDeleteDriver() rest.DeleteDriverPresenter {
	return deleteDriver{}
}

func (p deleteDriver) OutputError(err error) []byte {
	output := deleteDriverOutputError{
		Error: err.Error(),
	}

	b, _ := json.Marshal(&output)

	return b
}
