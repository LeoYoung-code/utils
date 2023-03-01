package utils

import (
	"strconv"
	"strings"
)

// SqlInSliceInt64 拼接SQL in
func SqlInSliceInt64(opt, field string, s []int64) string {
	sql := Reduce(s, "", func(sql string, i int64) string {
		sql = sql + "," + strconv.FormatInt(i, 10)
		return sql
	})
	return " " + opt + " " + field + " IN (" + sql[1:] + ")"
}

// SqlInSliceString 拼接SQL in
func SqlInSliceString(opt, field string, s []string) string {
	ss := strings.Join(s, "','")
	if len(ss) > 0 {
		return " " + opt + " " + field + " IN ('" + ss + "')"
	}
	return ""
}
