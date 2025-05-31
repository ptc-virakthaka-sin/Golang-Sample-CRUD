package util

import (
	"learn-fiber/pkg/database/pagination"
	"strings"

	"gorm.io/gorm"
)

func WherePaginateAndOrderBy(db *gorm.DB, paginate *pagination.Paginate) *gorm.DB {
	if paginate == nil {
		return db
	}
	db = db.Offset(paginate.Offset).Limit(paginate.Limit)
	if paginate.OrderBy != "" {
		db = db.Order(paginate.OrderBy)
	}
	return db
}

func EscapeLike(left, right, word string) string {
	var n int
	for i := range word {
		if c := word[i]; c == '%' || c == '_' || c == '\\' {
			n++
		}
	}

	// No characters to escape.
	if n == 0 {
		return left + word + right
	}

	var b strings.Builder
	b.Grow(len(word) + n)
	for _, c := range word {
		if c == '%' || c == '_' || c == '\\' {
			b.WriteByte('\\')
		}
		b.WriteRune(c)
	}
	return left + b.String() + right
}
