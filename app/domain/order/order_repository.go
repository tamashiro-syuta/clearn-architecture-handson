//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE
package order

import (
	"context"
)

// NOTE: リポジトリ層のインターフェースをdomainディレクトリに定義することで、
// NOTE: 依存性の注入(DI)を実現している
type OrderRepository interface {
	Save(ctx context.Context, order *Order) error
}
