package environment

import (
	"fmt"
	"os"
)

var ConnectionInfo DBConnection

/*
Reads environment variables to get Mongo server connection
information and returns a MongoConnection object. The following
constants represent the environment variables that are read:

   * ENV_MONGO_ADDRESS
   * ENV_MONGO_PORT
   * ENV_MONGO_DATABASE
   * ENV_MONGO_USERNAME
   * ENV_MONGO_PASSWORD
*/
func GetDBConnectionInformation() (DBConnection, error) {

	address := os.Getenv(ENV_MONGO_ADDRESS)
	if address == "" {
		return ConnectionInfo, fmt.Errorf("Invalid environment variable %s", ENV_MONGO_ADDRESS)
	}

	databaseName := os.Getenv(ENV_MONGO_DATABASE)
	if databaseName == "" {
		return ConnectionInfo, fmt.Errorf("Invalid environment variable %s", ENV_MONGO_DATABASE)
	}

	userName := os.Getenv(ENV_MONGO_USERNAME)
	if userName == "" {
		return ConnectionInfo, fmt.Errorf("Invalid environment variable %s", ENV_MONGO_USERNAME)
	}

	password := os.Getenv(ENV_MONGO_PASSWORD)
	if password == "" {
		return ConnectionInfo, fmt.Errorf("Invalid environment variable %s", ENV_MONGO_PASSWORD)
	}

	ConnectionInfo = DBConnection{
		Address:  address,
		Database: databaseName,
		UserName: userName,
		Password: password,
	}

	return ConnectionInfo, nil
}
