package utility

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"net/http"
	"reflect"
)

func ValidationStruct(err error, c *gin.Context) {

	if e, ok := err.(*json.UnmarshalTypeError); ok {
		msg := fmt.Sprintf("" + e.Field + " must be a " + kindOfData(e.Field).String())
		fmt.Println(msg)
		c.JSON(http.StatusBadRequest, gin.H{"error": "" + e.Field + " must be a " + kindOfData(e.Field).String()})
		return
	} else if errors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range errors {
			if fieldErr.Tag() == "required" {
				errMsg := fmt.Sprintf("Field %s is required", fieldErr.Field())
				c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
				return
			}
		}
	}
	fmt.Println(reflect.TypeOf(err))
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
}

func kindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
