package order

import (
	"context"
	"testing"
	"time"

	"github.com/code-kakitai/go-pkg/ulid"
	"go.uber.org/mock/gomock"

	cartDomain "github/code-kakitai/code-kakitai/domain/cart"
	orderDomain "github/code-kakitai/code-kakitai/domain/order"
)

func TestOrderUseCase_Run(t *testing.T) {
	// usecase準備
	ctrl := gomock.NewController(t)
	mockOrderDomainService := orderDomain.NewMockOrderDomainService(ctrl)
	mockCartRepo := cartDomain.NewMockCartRepository(ctrl)
	uc := NewOrderUseCase(mockOrderDomainService, mockCartRepo)

	// 各種テストデータ準備
	now := time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)
	userID := ulid.NewULID()
	dtos := []OrderUseCaseDto{
		{
			ProductID: ulid.NewULID(),
			Count:     1,
		},
		{
			ProductID: ulid.NewULID(),
			Count:     3,
		},
	}
	cart, _ := cartDomain.NewCart(userID)
	for _, dto := range dtos {
		cart.AddProduct(dto.ProductID, dto.Count)
	}

	tests := []struct {
		name     string
		dtos     []OrderUseCaseDto
		mockFunc func()
		wantErr  bool
	}{
		{
			name: "work",
			dtos: dtos,
			mockFunc: func() {
				gomock.InOrder(
					mockCartRepo.EXPECT().FindByUserID(gomock.Any(), userID).Return(cart, nil),
					mockOrderDomainService.EXPECT().OrderProducts(gomock.Any(), cart, now).Return("", nil),
				)
			},
			wantErr: false,
		},
		{
			name: "cartの中身とdtosの中身が一致しない",
			dtos: []OrderUseCaseDto{
				{
					ProductID: ulid.NewULID(),
					Count:     1,
				},
			},
			mockFunc: func() {
				gomock.InOrder(
					mockCartRepo.EXPECT().FindByUserID(gomock.Any(), userID).Return(cart, nil),
				)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			_, err := uc.Run(context.Background(), userID, tt.dtos, now)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderUseCase.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
