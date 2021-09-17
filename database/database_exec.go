package database

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"net/http"
)

func Create(c *gin.Context, value interface{}, valueName string, errorHandler func(c *gin.Context, message string, httpStatus int)) error {
	if dbc := DBConn.Create(value); dbc.Error != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(dbc.Error, &mysqlErr) && mysqlErr.Number == 1062 {
			msg := fmt.Sprintf("Create %s failed: already exists.", valueName)
			errorHandler(c, msg, http.StatusConflict)
			return dbc.Error
		}
		msg := fmt.Sprintf("Create %s failed: unkown error.", valueName)
		errorHandler(c, fmt.Sprintf(msg, valueName), http.StatusInternalServerError)
		return dbc.Error
	}
	return nil
}
