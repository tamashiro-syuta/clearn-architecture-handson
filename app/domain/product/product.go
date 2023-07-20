package product

import (
	"unicode/utf8"

	"github.com/code-kakitai/go-pkg/errors"
	"github.com/code-kakitai/go-pkg/ulid"
)

type Product struct {
	id          string
	ownerID     string
	name        string
	description string
	price       int64
	inventory   int
}

func NewProduct(
	ownerID string,
	name string,
	description string,
	price int64,
	inventory int,
) (*Product, error) {
	// ownerIDのバリデーション
	if !ulid.IsValid(ownerID) {
		return nil, errors.NewError("オーナーIDの値が不正です。")
	}
	// 名前のバリデーション
	if utf8.RuneCountInString(name) < nameLengthMin || utf8.RuneCountInString(name) > nameLengthMax {
		return nil, errors.NewError("商品名の値が不正です。")
	}

	// 商品説明のバリデーション
	if utf8.RuneCountInString(description) < descriptionLengthMin || utf8.RuneCountInString(description) > descriptionLengthMax {
		return nil, errors.NewError("商品説明の値が不正です。")
	}

	// 価格のバリデーション
	if price < 1 {
		return nil, errors.NewError("価格の値が不正です。")
	}

	// 在庫数のバリデーション
	// 在庫はないけど、商品は登録したい等あるあるため、0は許容する
	if inventory < 0 {
		return nil, errors.NewError("在庫数の値が不正です。")
	}
	return &Product{
		id:          ulid.NewULID(),
		ownerID:     ownerID,
		name:        name,
		description: description,
		price:       price,
		inventory:   inventory,
	}, nil
}

func (p *Product) OwnerID() string {
	return p.ownerID
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) Description() string {
	return p.description
}

func (p *Product) Price() int64 {
	return p.price
}

const (
	// 名前の最大値/最小値
	nameLengthMin = 1
	nameLengthMax = 100

	// 説明の最大値/最小値
	descriptionLengthMin = 1
	descriptionLengthMax = 1000
)