package models

import (
	"os"

	er "github.com/pkg/errors"

	"github.com/joho/godotenv"
)

type ResponseKeys struct {
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

type RequestDTO struct {
	PublicKey string `json:"public_key"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type ResponseAddress struct {
	Address string `json:"address"`
}

type UserConfig struct {
	AccessKeyID     string
	SecretAccessKey string
	MyRegion        string
}

func (u *UserConfig) NewUserConfig() error {
	err := godotenv.Load()
	if err != nil {
		// Handle the error, e.g., log it or exit the program.
		panic("Error loading .env file")
	}

	u.AccessKeyID = GetEnvWithKey("AWS_ACCESS_KEY_ID")
	u.SecretAccessKey = GetEnvWithKey("AWS_SECRET_ACCESS_KEY")
	u.MyRegion = GetEnvWithKey("AWS_REGION")

	if u.AccessKeyID == "" || u.SecretAccessKey == "" || u.MyRegion == "" {
		return er.New("error_getting_env_vars")
	}
	return nil
}

// get env value
func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}
