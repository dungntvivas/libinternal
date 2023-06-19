package util

import "github.com/google/uuid"

func Generate_uuid() string {
	id, _ := uuid.NewUUID()
	return id.String()
}

func Calc_total_page(total int64, limit int64) int64 {
	_t := total / limit
	if (total % limit) != 0 {
		_t += 1
	}
	return _t
}
