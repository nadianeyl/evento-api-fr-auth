package response

import "github.com/nadianeyl/evento-api-fr-auth/internal/domain"

type CustomerResponse struct {
	ID       int64           `json:"id"`
	Username string          `json:"username"`
	Balance  float64         `json:"balance"`
	Orders   []*domain.Order `json:"orders"`
}
