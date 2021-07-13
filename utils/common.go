package utils

import (
	"reflect"
	"time"
)

const (
	GenesisTime     = "2020-08-25 06:00:00" //Beijing time
	GenesisUnixTime = 1598306400
)

// Note: Both src and dst are pointer structs, and the same fields (same type and field name) can be copied
func StructSubCopy(src interface{}, dst interface{}) {
	srcValue := reflect.ValueOf(src).Elem()
	dstValue := reflect.ValueOf(dst).Elem()
	dstType := dstValue.Type()
	for i := 0; i < dstType.NumField(); i++ {
		dstField := dstType.Field(i)
		srcV := srcValue.FieldByName(dstField.Name)
		if srcV.IsValid() && srcV.Type().AssignableTo(dstField.Type) {
			dstValue.FieldByName(dstField.Name).Set(srcV)
		}
	}
}

func PaginationHelper(page, pageSize, defaultPageSize int) (offset int, size int) {
	size = pageSize
	if size == 0 {
		size = defaultPageSize
	}
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * size
	}
	return
}

func GetEpochByTime(t time.Time) uint64 {
	return uint64((t.Unix() - GenesisUnixTime) / 30)
}

func GetGenesisTime() time.Time {
	return time.Unix(GenesisUnixTime, 0)
}
