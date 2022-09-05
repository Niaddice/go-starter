package utils

import (
	"database/sql/driver"
	"regexp"
	"strings"
	"time"
)

type LocalTime time.Time

func (t *LocalTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")

	reg, _ := regexp.Match(`^[1-2][0-9][0-9][0-9]-[0-1]{0,1}[0-9]-[0-3]{0,1}[0-9]$`, []byte(timeStr))
	if reg {
		t1, _ := time.Parse("2006-01-02", timeStr)
		*t = LocalTime(t1)
	} else {
		t1, _ := time.Parse("2006-01-02 15:04:05", timeStr)
		*t = LocalTime(t1)
	}
	return err
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len("2006-01-02 15:04:05")+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, "2006-01-02 15:04:05")
	b = append(b, '"')
	return b, nil
}

func (t LocalTime) Value() (driver.Value, error) {
	if t.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return []byte(time.Time(t).Format("2006-01-02 15:04:05")), nil
}

func (t *LocalTime) Scan(v interface{}) error {
	tTime, _ := time.Parse("2006-01-02 15:04:05 +0800 CST", v.(time.Time).String())
	*t = LocalTime(tTime)
	return nil
}

func (t LocalTime) String() string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}
