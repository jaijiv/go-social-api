package environment

import (
	"fmt"
)

/*
DBConnection contains data necessary to establish a connection
to a Mongo server.
*/
type DBConnection struct {
	Address  string
	Database string
	UserName string
	Password string
}

func (this *DBConnection) ToString() string {
	return fmt.Sprintf("%s:%s@(%s:%d)/%s", this.UserName, this.Password, this.Address, this.Database)
}
