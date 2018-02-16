package serialization

import (
	"encoding/json"
	"github.com/frederic-gendebien/customer/domain"
)

type JsonSerializer struct{}

type JsonCustomer struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Vat     string `json:"vat,omitempty"`
}

func NewJsonSerializer() *JsonSerializer {
	return new(JsonSerializer)
}

func (serializer *JsonSerializer) Marshal(customer *domain.Customer) ([]byte, error) {
	return json.Marshal(JsonCustomer{
		Name:    customer.GetName(),
		Address: customer.GetAddress(),
		Vat:     customer.GetVat(),
	})
}

func (serializer *JsonSerializer) Unmarshal(bytes []byte) (*domain.Customer, error) {
	jsonCustomer := JsonCustomer{}
	err := json.Unmarshal(bytes, &jsonCustomer)
	if err != nil {
		return nil, err
	}

	return domain.NewCustomer(domain.CustomerBuilder{
		Name:    jsonCustomer.Name,
		Address: jsonCustomer.Address,
		Vat:     jsonCustomer.Vat,
	})
}
