package sql

import (
	"strconv"
	"strings"

	"github.com/samber/lo"
)

// InSliceInt64 拼接SQL in
func InSliceInt64(opt, field string, s []int64) string {
	var b strings.Builder
	sql := lo.Reduce(s, func(agg string, item int64, _ int) string {
		b.WriteString(",")
		b.WriteString(strconv.FormatInt(item, 10))
		return b.String()
	}, "")

	if len(sql) == 0 {
		return ""
	}
	return " " + opt + " `" + field + "` IN (" + sql[1:] + ")"
}

// InSliceString 拼接SQL in
func InSliceString(opt, field string, s []string) string {
	ss := strings.Join(s, "','")
	if len(ss) > 0 {
		return " " + opt + " `" + field + "` IN ('" + ss + "')"
	}
	return ""
}
