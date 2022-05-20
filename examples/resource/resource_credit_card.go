package resource

import (
	"github.com/manyminds/api2go"
	"github.com/manyminds/api2go/examples/model"
	"github.com/manyminds/api2go/examples/storage"
	"net/http"
)

type CreditCardResource struct {
	UserStorage *storage.UserStorage
}

func (c CreditCardResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var creditCards []model.CreditCard
	creditCards = append(creditCards,model.CreditCard{ID: "1"})
	return &Response{
		Res:  creditCards,
		Code: http.StatusOK,
	}, nil
}
func (c CreditCardResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	return &Response{
		Res: &model.CreditCard{ID:"foo"},
		Code: http.StatusOK,
	}, nil
}

