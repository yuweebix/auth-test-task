package schemas

import "github.com/Masterminds/squirrel"

type Filter interface {
	ApplyToQuery(query squirrel.SelectBuilder) squirrel.SelectBuilder
}
