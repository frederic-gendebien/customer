package serialization

import "github.com/frederic-gendebien/customer/domain"

type Serializer interface {
	Marshal(customer *domain.Customer) ([]byte, error)
	Unmarshal([]byte) (*domain.Customer, error)
}

type Format int

const (
	Protobuf 	Format = 0
	Json 		Format = 1
	Yaml 		Format = 2
)

var serializers = []Serializer{
	NewProtobufSerializer(),
	NewJsonSerializer(),
	NewYamlSerializer(),
}

func init() {
}

func (format Format) Serializer() Serializer {
	return serializers[format]
}
