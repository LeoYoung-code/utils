package sql

import (
	"strconv"
	"strings"

	"utils"
)

// InSliceInt64 拼接SQL in
func InSliceInt64(opt, field string, s []int64) string {
	sql := utils.Reduce(s, "", func(sql string, i int64) string {
		sql = sql + "," + strconv.FormatInt(i, 10)
		return sql
	})
	return " " + opt + " " + field + " IN (" + sql[1:] + ")"
}

// InSliceString 拼接SQL in
func InSliceString(opt, field string, s []string) string {
	ss := strings.Join(s, "','")
	if len(ss) > 0 {
		return " " + opt + " " + field + " IN ('" + ss + "')"
	}
	return ""
}
