// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0

package dbgen

import (
	"context"
)

type Querier interface {
	UpsertUser(ctx context.Context, arg UpsertUserParams) error
	UserFindById(ctx context.Context, id string) (User, error)
}

var _ Querier = (*Queries)(nil)