package database

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"net/http"
)

func Create(c *gin.Context, value interface{}, modelName string, errorHandler func(c *gin.Context, message string, httpStatus int)) error {
	if dbc := DBConn.Create(value); dbc.Error != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(dbc.Error, &mysqlErr) && mysqlErr.Number == 1062 {
			msg := fmt.Sprintf("Create %s failed: already exists.", modelName)
			errorHandler(c, msg, http.StatusConflict)
			return dbc.Error
		}
		msg := fmt.Sprintf("Create %s failed: unkown error.", modelName)
		errorHandler(c, fmt.Sprintf(msg, modelName), http.StatusInternalServerError)
		return dbc.Error
	}
	return nil
}

func CreateInBatches(c *gin.Context, value interface{}, batchSize int, modelName string, errorHandler func(c *gin.Context, message string, httpStatus int)) error {
	if dbc := DBConn.CreateInBatches(value, batchSize); dbc.Error != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(dbc.Error, &mysqlErr) && mysqlErr.Number == 1062 {
			msg := fmt.Sprintf("Create %s failed: already exists.", modelName)
			errorHandler(c, msg, http.StatusConflict)
			return dbc.Error
		}
		msg := fmt.Sprintf("Create %s failed: unkown error.", modelName)
		errorHandler(c, fmt.Sprintf(msg, modelName), http.StatusInternalServerError)
		return dbc.Error
	}
	return nil
}

func Delete(c *gin.Context, model interface{}, toDeleteID uint, modelName string, errorHandler func(c *gin.Context, message string, httpStatus int)) error {
	if dbc := DBConn.Delete(model, toDeleteID); dbc.Error != nil {
		msg := fmt.Sprintf("Delete %s failed: unkown error.", modelName)
		errorHandler(c, fmt.Sprintf(msg, modelName), http.StatusInternalServerError)
		return dbc.Error
	}
	return nil
}

func Updates(c *gin.Context, whereFields interface{}, saveFields interface{}, modelName string, errorHandler func(c *gin.Context, message string, httpStatus int)) error {
	if dbc := DBConn.Model(whereFields).Updates(saveFields); dbc.Error != nil {
		msg := fmt.Sprintf("Update %s failed: unkown error.", modelName)
		errorHandler(c, fmt.Sprintf(msg, modelName), http.StatusInternalServerError)
		return dbc.Error
	}
	return nil
}

func GetByField(matchValue interface{}, valueContainer interface{}, joins []string) {
	var dbc = DBConn
	for i := 0; i < len(joins); i++ {
		dbc = dbc.Joins(joins[i])
	}
	dbc = dbc.Where(matchValue)
	if dbc.Error != nil {
	} else {
		dbc.Find(valueContainer)
	}
}
