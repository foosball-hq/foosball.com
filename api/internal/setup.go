package internal

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/kms/kmsiface"
	"github.com/caarlos0/env"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DatabaseConnectionOptions contains the set of fields used to connect to the foosball.com database
type DatabaseConnectionOptions struct {
	// Environment is the environment where the database is located
	Environment string `env:ENV`

	// Host is the host of the database
	Host string `env:"POSTGRES_HOST"`

	// Port is the port number of the database
	Port string `env:"POSTGRES_PORT"`

	// DatabaseName is the name of the database
	DatabaseName string `env:"POSTGRES_DBNAME"`

	// Username is the name of the user to connect to the database as
	Username string `env:"POSTGRES_USERNAME"`

	// Password is the password to use to connect to the database
	Password string `env:"POSTGRES_PASSWORD"`
}

// CreateDatabaseConnection creates a new connection to the Postgres database. All connection information must be
// provided via a DatabaseConnectionOptions instance. If the connections fields aren't set, then they will instead
// be read in from the environment
func CreateDatabaseConnection(svc kmsiface.KMSAPI, options DatabaseConnectionOptions) (*gorm.DB, error) {
	err := env.Parse(&options)
	if err != nil {
		return nil, fmt.Errorf("unable to parse environment variables: %v", err)
	}

	// We assume that if the options.Environment == "local", then the password is already set in
	// the options and doesn't need to be decrypted though KMS. This wll allow us to to test
	// methods/functions dependent on this one in a local environment.
	var password string
	if options.Environment == "local" {
		password, err = decrypt(svc, options.Password)
		if err != nil {
			return nil, err
		}
	}

	return gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", options.Host, options.Port, options.Username, options.DatabaseName, password))
}

func decrypt(svc kmsiface.KMSAPI, encrypted string) (string, error) {
	blob := []byte(encrypted)
	result, err := svc.Decrypt(&kms.DecryptInput{CiphertextBlob: blob})

	if err != nil {
		return "", fmt.Errorf("unable to decrypt string: %v", err)
	}

	return string(result.Plaintext), nil
}
