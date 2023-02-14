package dbcore

import "strings"

const (
	ASC  = "ASC"
	DESC = "DESC"
)

type order string

var Order = new(order)

func (order) Ascending() order {
	return ASC
}

func (o order) IsAscending() bool {
	return o == ASC
}

func (order) Descending() order {
	return DESC
}

func (o order) IsDescending() bool {
	return o == DESC
}

type SortItem struct {
	Field string
	Order order
}

func ParseSort(sort string) []SortItem {
	result := make([]SortItem, 0)
	if sort == "" {
		return result
	}
	sorts := strings.Split(sort, ",")
	for _, sort := range sorts {
		order := Order.Ascending()
		if strings.HasPrefix(sort, "-") {
			order = Order.Descending()
			sort = sort[1:]
		}
		result = append(result, SortItem{
			Field: sort,
			Order: order,
		})
	}
	return result
}
