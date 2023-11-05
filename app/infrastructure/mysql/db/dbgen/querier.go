// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package dbgen

import (
	"context"
)

type Querier interface {
	InsertOrder(ctx context.Context, arg InsertOrderParams) error
	InsertOrderProduct(ctx context.Context, arg InsertOrderProductParams) error
	OrderFindById(ctx context.Context, id string) (Order, error)
	OrderProductFindById(ctx context.Context, id string) (OrderProduct, error)
	ProductFetchWithOwner(ctx context.Context) ([]ProductFetchWithOwnerRow, error)
	ProductFindById(ctx context.Context, id string) (Product, error)
	ProductFindByIds(ctx context.Context, ids []string) ([]Product, error)
	UpsertProduct(ctx context.Context, arg UpsertProductParams) error
	UpsertUser(ctx context.Context, arg UpsertUserParams) error
	UserFindAll(ctx context.Context) ([]User, error)
	UserFindById(ctx context.Context, id string) (User, error)
}

var _ Querier = (*Queries)(nil)
