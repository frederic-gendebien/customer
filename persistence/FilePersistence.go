package persistence

import (
	"github.com/frederic-gendebien/customer/domain"
	"github.com/frederic-gendebien/customer/serialization"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const GOPATH = "GOPATH"

var gopath = ""
var folder = ""
var serializer = serialization.Protobuf.Serializer()

func init() {
	verifyEnvironment()
	createFolders()
}

func verifyEnvironment() string {
	gopath = os.Getenv(GOPATH)
	if gopath == "" {
		log.Fatalf("Undefined %s environment variable", GOPATH)
	}

	return gopath
}

func createFolders() {
	folder = filepath.Join(gopath, "var", "customer")
	if folder == "" {
		log.Fatal("could not create folders")
	}

	err := os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		log.Fatal("could not create folders")
	}
}

func filePath(alias string) string {
	filePath := filepath.Join(folder, alias)

	return filePath
}

func Save(alias string, customer *domain.Customer) error {

	bytes, err := serializer.Marshal(customer)
	if err != nil {
		return err
	}

	filePath := filePath(alias)
	err = ioutil.WriteFile(filePath, bytes, os.ModePerm)

	return err
}

func Delete(alias string) error {
	err := os.Remove(filePath(alias))

	return err
}

func List() ([]string, error) {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	aliases := make([]string, len(files))
	for index, file := range files {
		aliases[index] = file.Name()
	}

	return aliases, nil
}

func Get(alias string) (*domain.Customer, error) {
	bytes, err := ioutil.ReadFile(filePath(alias))
	if err != nil {
		return nil, err
	}

	return serializer.Unmarshal(bytes)
}
