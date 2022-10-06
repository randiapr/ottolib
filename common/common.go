// put your code here for common function, library, etc
package common

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	jsontime "github.com/liamylian/jsontime/v2/v2"
	"github.com/randiapr/ottolib/constant"
	"github.com/randiapr/ottolib/model"
)

// setup custom date time format
var (
	SharedJson = jsontime.ConfigWithCustomTimeFormat
)

func init() {
	jktTimezone, _ := time.LoadLocation("Asia/Jakarta")
	jsontime.AddTimeFormatAlias("datetime", constant.DATE_TIME_FORMAT)
	jsontime.AddTimeFormatAlias("date", constant.DATE_FORMAT_1)
	jsontime.AddLocaleAlias("jkt", jktTimezone)
}

// Validate struct with validator
func Validate(interfacez interface{}) error {
	validate := validator.New()
	err := validate.Struct(interfacez)
	if err != nil {
		errMsg := ""
		for _, errs := range err.(validator.ValidationErrors) {
			errMsg = errs.Field() + " must valid " + errs.Tag()
			break
		}
		return errors.New(errMsg)
	}
	return nil
}

// DateStringToMilliSeconds ...
func DateStringToMilliSeconds(date string) int64 {
	dates, _ := time.Parse("2006-01-02 15:04:05", date)
	return DateTimeToMilliSeconds(dates)
}

// DateTimeToMilliSeconds ...
func DateTimeToMilliSeconds(date time.Time) int64 {
	return date.UnixNano() / int64(time.Millisecond)
}

func OffsetRecord(page, size int64) int64 {
	return (page - 1) * size
}

// TotalPages ..
func totalPages(count int64, limit int64) int64 {
	return int64(math.Ceil(float64(count) / float64(limit)))
}

// NextPage ..
func nextPage(page, total, limit int64) int64 {
	if page == totalPages(total, limit) || total == 0 {
		return page
	}
	return page + 1
}

// PrevPage ..
func prevPage(page int64) int64 {
	if page > 1 {
		return page - 1
	}
	return page
}

// MetaPageBuilder total, limit, offset, page
func MetaPageBuilder(total, limit, page int64) model.PaginationMeta {
	if limit == 0 {
		limit = 10
	}
	if page == 0 {
		page = 1
	}
	return model.PaginationMeta{
		TotalRecords: total,
		TotalPages:   totalPages(total, limit),
		Offset:       OffsetRecord(page, limit),
		Limit:        limit,
		Page:         page,
		PrevPage:     prevPage(page),
		NextPage:     nextPage(page, total, limit),
	}
}

// SortedBy query param []string
func SortedBy(sort []string) []string {
	var sorted []string
	for _, v := range sort {
		split := strings.Split(v, ",")
		sorted = append(sorted, fmt.Sprintf("%s %s", split[0], split[1]))
	}
	return sorted
}
