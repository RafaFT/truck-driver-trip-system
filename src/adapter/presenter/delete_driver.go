package presenter

import (
	"encoding/json"

	"github.com/rafaft/truck-driver-trip-system/adapter/controller/rest"
)

type deleteDriverOutputError struct {
	Error string `json:"error"`
}

// output port (presenter) implementation
type deleteDriverPresenter struct {
}

func NewDeleteDriverPresenter() rest.DeleteDriverPresenter {
	return deleteDriverPresenter{}
}

func (p deleteDriverPresenter) OutputError(err error) []byte {
	output := deleteDriverOutputError{
		Error: err.Error(),
	}

	b, _ := json.Marshal(&output)

	return b
}
