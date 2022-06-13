package request

import (
	"errors"
	"net/http"
	"strconv"

	errorTypes "github.com/NicklasWallgren/go-template/adapters/driven/api/errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetParamInt(c *gin.Context, name string) (int, error) {
	val := c.Params.ByName(name)
	if val == "" {
		return 0, errors.New(name + " path parameter value is empty or not specified")
	}

	return strconv.Atoi(val)
}

func Into[T any](c *gin.Context, request T) (T, error) {
	if err := c.ShouldBindUri(&request); err != nil {
		return request, errorTypes.NewApiErrorWith(errorTypes.WithStatusAndError(http.StatusBadRequest, err))
	}

	if err := c.ShouldBindQuery(&request); err != nil {
		return request, errorTypes.NewApiErrorWith(errorTypes.WithStatusAndError(http.StatusBadRequest, err))
	}

	if c.Request.ContentLength <= 0 {
		return request, nil
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		return request, errorTypes.NewApiErrorWith(errorTypes.WithStatusAndError(http.StatusBadRequest, err))
	}

	return request, nil
}

func IntoWithDefault[T any](c *gin.Context, request T, decorator func(request *T) *T) (*T, error) {
	request, err := Into[T](c, request)
	if err != nil {
		return nil, err
	}

	return decorator(&request), nil
}

func IntoAndValidate[T any](c *gin.Context, validator binding.StructValidator, request T) (*T, error) {
	request, err := Into[T](c, request)
	if err != nil {
		return nil, err
	}

	if err := validator.ValidateStruct(&request); err != nil {
		return nil, errorTypes.NewApiErrorWith(errorTypes.WithStatusAndError(http.StatusBadRequest, err))
	}

	return &request, nil
}
