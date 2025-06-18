package sql

import (
	"strconv"
	"strings"
)

// InSliceInt64 拼接SQL in
func InSliceInt64(opt, field string, s []int64) string {
	var b strings.Builder
	for i, v := range s {
		if i > 0 {
			b.WriteString(",")
		}
		b.WriteString(strconv.FormatInt(v, 10))
	}
	sql := b.String()
	if len(sql) == 0 {
		return ""
	}
	return " " + opt + " `" + field + "` IN (" + sql + ")"
}

// InSliceString 拼接SQL in
func InSliceString(opt, field string, s []string) string {
	ss := strings.Join(s, "','")
	if len(ss) > 0 {
		return " " + opt + " `" + field + "` IN ('" + ss + "')"
	}
	return ""
}
