package constants

import "time"

const (
	DATE_FMT               = "2006-01-02"
	DATE_FMT_TIGHT         = "20060102"
	DATE_TIME_FMT          = "2006-01-02 15:04:05"
	DATE_TIME_FMT_3        = "2006-01-02 15:04:05.000"
	DATE_TIME_FMT_6        = "2006-01-02 15:04:05.000000"
	DATE_HOUR_FMT          = "2006-01-02 15"
	DATE_MONTH_FMT         = "2006-01"
	DATE_HOUR_FMT_TIGHT    = "2006010215"
	DATE_SECOND_FMT_TIGHT  = "20060102150405"
	DATE_HOUR_MINUTE_FMT   = "2006-01-02 15:04"
	DATE_MONTH_DAY_FMT     = "01-02"
	TIME_FMT               = "15:04"
	HOUR_MINUTE_SECOND_FMT = "15:04:05"
	HOUR_FMT               = "15"
	MINUTE_FMT             = "04"
	SECOND_FMT             = "05"
	GMT_HOURS              = 8
	DAY                    = time.Hour * 24
	YEAR                   = time.Hour * 24 * 365

	UnixInitialTime = "1970-01-01 08:00:00"
)

var LOCALE, _ = time.LoadLocation("PRC")
