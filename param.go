package utils

// RequestPagination 默认值
func RequestPagination(pageReq, pageSizeReq int64) (int64, int64) {
	page := int64(1)
	pageSize := int64(10)
	maxPageSize := int64(100)

	if pageReq > 0 {
		page = pageReq
	}

	if pageSizeReq > 0 && pageSizeReq <= maxPageSize {
		pageSize = pageSizeReq
	}
	return page, pageSize
}
