package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//StringDatabaseConnection is the string the connects to MYSQL
	StringDatabaseConnection = ""

	//Port is the port the API will be running on
	Port = 0

	//SecretKey is the key used to sign the token
	SecretKey []byte
)

// ChargeConfigs initializes environment variables
func ChargeConfigs() {
	var error error
	if error = godotenv.Load(); error != nil {
		log.Fatal(error)
	}
	Port, error = strconv.Atoi(os.Getenv("PORT_API"))
	if error != nil {
		Port = 8000
	}
	StringDatabaseConnection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("USER_DB"),
		os.Getenv("PASSWORD_DB"),
		os.Getenv("NAME_DB"))

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
