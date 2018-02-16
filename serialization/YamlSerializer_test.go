package serialization

import (
	"github.com/frederic-gendebien/customer/domain"
	"reflect"
	"testing"
)

func TestNewYamlSerializer(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"success"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewYamlSerializer(); got == nil {
				t.Errorf("NewYamlSerializer() = is nil")
			}
		})
	}
}

func TestYamlSerializer_Format(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"success", "yaml"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serializer := &YamlSerializer{}
			if got := serializer.Format(); got != tt.want {
				t.Errorf("YamlSerializer.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYamlSerializer_Marshal(t *testing.T) {
	tests := []struct {
		name    string
		builder domain.CustomerBuilder
		want    string
		wantErr bool
	}{
		{
			"success",
			domain.CustomerBuilder{Name: "name", Address: "address", Vat: "vat"},
			`name: name
address: address
vat: vat
`,
			false,
		},
		{
			"without vat",
			domain.CustomerBuilder{Name: "name", Address: "address"},
			`name: name
address: address
`,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serializer := &YamlSerializer{}
			customer, err := domain.NewCustomer(tt.builder)
			if (err != nil) != tt.wantErr {
				t.Errorf("could not create customer")
				return
			}
			got, err := serializer.Marshal(customer)
			if (err != nil) != tt.wantErr {
				t.Errorf("YamlSerializer.Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			text := string(got)
			if text != tt.want {
				t.Errorf("YamlSerializer.Marshal() = %v, want %v", text, tt.want)
			}
		})
	}
}

func TestYamlSerializer_Unmarshal(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		builder domain.CustomerBuilder
		wantErr bool
	}{
		{
			"success",
			"name: name\naddress: address\nvat: vat",
			domain.CustomerBuilder{Name: "name", Address: "address", Vat: "vat"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serializer := &YamlSerializer{}
			got, err := serializer.Unmarshal([]byte(tt.value))
			if (err != nil) != tt.wantErr {
				t.Errorf("YamlSerializer.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			customer, err := domain.NewCustomer(tt.builder)
			if (err != nil) != tt.wantErr {
				t.Errorf("could not create customer")
				return
			}
			if !reflect.DeepEqual(got, customer) {
				t.Errorf("YamlSerializer.Unmarshal() = %v, want %v", got, customer)
			}
		})
	}
}
