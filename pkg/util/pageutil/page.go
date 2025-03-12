package pageutil

import (
	"errors"
	"reflect"
	"strings"

	"github.com/gunawanpras/be-product-service/pkg/util/constant"
)

func ValidateSortDirection(availSort []string, reqSort, reqDirection string) error {
	if !ItemExists(availSort, strings.ToLower(reqSort)) {
		return errors.New(constant.ErrInvalidSort)
	}

	if !IsValidDirection(reqDirection) {
		return errors.New(constant.ErrInvalidSortDirection)
	}

	return nil
}

func ItemExists(arrayType, item any) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Slice && arr.Kind() != reflect.Array {
		panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func IsValidDirection(direction string) bool {
	if strings.TrimSpace(direction) == "" {
		return true
	}

	return ItemExists(constant.ValidSortDirection, strings.ToLower(direction))
}
