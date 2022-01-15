package dto

import (
	"time"
)

type AuthorResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name,omitempty"`
	Bio       string    `json:"bio,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

//func AuthorResponseFromGrpc(product *readerService.Product) *ProductResponse {
//	return &ProductResponse{
//		ProductID:   product.GetProductID(),
//		Name:        product.GetName(),
//		Description: product.GetDescription(),
//		Price:       product.GetPrice(),
//		CreatedAt:   product.GetCreatedAt().AsTime(),
//		UpdatedAt:   product.GetUpdatedAt().AsTime(),
//	}
//}
