package serialization

import (
	"github.com/frederic-gendebien/customer/domain"
	"gopkg.in/yaml.v2"
)

type YamlSerializer struct{}

type YamlCustomer struct {
	Name    string `yaml:"name"`
	Address string `yaml:"address"`
	Vat     string `yaml:"vat,omitempty"`
}

func NewYamlSerializer() *YamlSerializer {
	return new(YamlSerializer)
}

func (serializer *YamlSerializer) Marshal(customer *domain.Customer) ([]byte, error) {
	return yaml.Marshal(YamlCustomer{
		Name:    customer.GetName(),
		Address: customer.GetAddress(),
		Vat:     customer.GetVat(),
	})
}

func (serializer *YamlSerializer) Unmarshal(bytes []byte) (*domain.Customer, error) {
	yamlCustomer := YamlCustomer{}
	err := yaml.Unmarshal(bytes, &yamlCustomer)
	if err != nil {
		return nil, err
	}

	return domain.NewCustomer(domain.CustomerBuilder{
		Name:    yamlCustomer.Name,
		Address: yamlCustomer.Address,
		Vat:     yamlCustomer.Vat,
	})
}
