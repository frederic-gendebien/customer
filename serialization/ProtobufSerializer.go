package serialization

import (
	"github.com/frederic-gendebien/customer/domain"
	"github.com/golang/protobuf/proto"
)

type ProtobufSerializer struct{}


func NewProtobufSerializer() *ProtobufSerializer {
	return new(ProtobufSerializer)
}

func (serializer *ProtobufSerializer) Marshal(customer *domain.Customer) ([]byte, error) {
	return proto.Marshal(&ProtobufCustomer{
		Name:    proto.String(customer.GetName()),
		Address: proto.String(customer.GetAddress()),
		Vat:     proto.String(customer.GetVat()),
	})
}

func (serializer *ProtobufSerializer) Unmarshal(bytes []byte) (*domain.Customer, error) {
	protobufCustomer := ProtobufCustomer{}
	err := proto.Unmarshal(bytes, &protobufCustomer)
	if err != nil {
		return nil, err
	}

	return domain.NewCustomer(domain.CustomerBuilder{
		Name:    protobufCustomer.GetName(),
		Address: protobufCustomer.GetAddress(),
		Vat:     protobufCustomer.GetVat(),
	})
}
