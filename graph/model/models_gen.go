// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"ktp-fix/internal/models"
)

//  Pagination is default input pagination
type Pagination struct {
	First  int      `json:"first"`
	Offset int      `json:"offset"`
	After  *string  `json:"after"`
	Query  string   `json:"query"`
	Sort   []string `json:"sort"`
}

//  Object that is being paginated
type PaginationEdgeKtp struct {
	Node   *models.Ktp `json:"node"`
	Cursor string      `json:"cursor"`
}

//  Information about pagination
type PaginationInfoKtp struct {
	EndCursor   string `json:"endCursor"`
	HasNextPAge bool   `json:"hasNextPAge"`
}

//  Result while querying list using graphql
type PaginationResultKtp struct {
	Totalcount int                  `json:"totalcount"`
	Edges      []*PaginationEdgeKtp `json:"edges"`
	PageInfo   *PaginationInfoKtp   `json:"pageInfo"`
}
