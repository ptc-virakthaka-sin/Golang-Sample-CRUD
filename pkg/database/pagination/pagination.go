package pagination

import (
	"math"
	"strings"
)

func GetPageResponse(page, limit, totalData int64) PageResponse {
	return PageResponse{
		Page:       page,
		Limit:      limit,
		Total:      totalData,
		TotalPages: GetTotalPages(totalData, limit),
		HasNext:    GetHasNext(page, totalData, limit),
		HasPrev:    GetHasPrevious(page),
		NextPage:   GetNextPage(page, totalData, limit),
		PrevPage:   GetPrevPage(page),
		From:       GetStartNoOfPage(page, limit),
		To:         GetEndNoOfPage(page, totalData, limit),
	}
}

func GetTotalPages(total, limit int64) int64 {
	d := safeDivideFloat(float64(total), float64(limit))
	return int64(math.Ceil(d))
}

func GetHasNext(currentPage, total, limit int64) bool {
	totalPages := GetTotalPages(total, limit)
	return currentPage < totalPages
}

func GetHasPrevious(currentPage int64) bool {
	return currentPage > defaultPage
}

func GetNextPage(currentPage, total, limit int64) int64 {
	hasNext := GetHasNext(currentPage, total, limit)
	if hasNext {
		return currentPage + pageStep
	}
	return currentPage
}

func GetPrevPage(currentPage int64) int64 {
	hasPrev := GetHasPrevious(currentPage)
	if hasPrev {
		return currentPage - pageStep
	}
	return currentPage
}

func GetStartNoOfPage(currentPage, limit int64) int64 {
	if currentPage > defaultPage {
		// page 2 -> ((2-1) * 20) + 1 -> 21
		return ((currentPage - 1) * limit) + 1
	}
	// page <= 1 -> 1
	return 1
}

func GetEndNoOfPage(currentPage, total, limit int64) int64 {
	endNo := currentPage * limit
	if endNo > total {
		return total
	}
	return endNo
}

func GetMongoSortOrder(s string) int {
	if s != "" && (strings.EqualFold(s, "desc") || strings.EqualFold(s, "-1")) {
		return -1
	}
	return 1
}

func safeDivideFloat(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}

// --------------------------------------

type Paginate struct {
	Page    int
	Offset  int
	Limit   int
	OrderBy string
}

func ToPaginate(params PageRequest) *Paginate {
	page := params.GetPage()
	if page == 0 {
		page = 1
	}

	limit := params.GetLimit()
	switch {
	case limit > maxLimitPerPage:
		limit = maxLimitPerPage
	case limit <= 0:
		limit = defaultLimitPerPage
	}

	orderBy := ""

	if params.GetSort() != "" {
		// ?sort=field1,field2:asc,field3:desc
		sorts := strings.Split(params.GetSort(), ",")
		orders := []string{}

		for _, s := range sorts {
			fields := strings.Split(s, ":")
			field := fields[0]

			if len(fields) > 1 {
				direction := GetOrderDirection(fields[1])
				orders = append(orders, field+" "+direction)
			} else {
				orders = append(orders, field)
			}
		}
		// ORDER BY age desc, name;
		orderBy = strings.Join(orders, ", ")
	}

	offset := (page - 1) * limit

	return &Paginate{
		Page:    page,
		Offset:  offset,
		Limit:   limit,
		OrderBy: orderBy,
	}
}

func GetOrderDirection(s string) string {
	if s != "" && strings.EqualFold(s, "desc") {
		return "desc"
	}
	return "asc"
}

func (p *PageRequest) GetPage() int {
	if p.Page < defaultPage {
		return defaultPage
	}
	return int(p.Page)
}

/*
In vuetify data table items per page, Setting this prop to -1 will display all items on the page
https://vuetifyjs.com/en/api/v-data-table/#props-items-per-page

In mongodb zero value i.e. limit(0) is equivalent to setting no limit.
link https://www.mongodb.com/docs/v4.4/reference/method/cursor.limit/#zero-value
*/
func (p *PageRequest) GetLimit() int {
	if p.Limit == 0 || p.Limit == -1 {
		return 0
	}
	if p.Limit < -1 {
		return defaultLimitPerPage
	} else if p.Limit > maxLimitPerPage {
		return maxLimitPerPage
	}
	return int(p.Limit)
}

func (p *PageRequest) GetOffset() int {
	if p.Page > 1 {
		return int((p.Page - 1) * p.Limit)
	}
	return 0
}

func (p *PageRequest) GetSort() string {
	return p.Sort
}

func (p *PageRequest) GetSQLSort() interface{} {
	panic("no implementation")
}

func (p *PageRequest) SetPage(page int) {
	if page < defaultPage {
		p.Page = defaultPage
		return
	}
	p.Page = int64(page)
}

func (p *PageRequest) SetLimit(limit int) {
	if limit < 1 {
		p.Limit = defaultLimitPerPage
		return
	} else if limit > maxLimitPerPage {
		p.Limit = maxLimitPerPage
		return
	}
	p.Limit = int64(limit)
}

func (p *PageRequest) SetSort(sort string) {
	p.Sort = sort
}
