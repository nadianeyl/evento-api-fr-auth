package request

import "github.com/nadianeyl/evento-api-fr-auth/internal/domain"

type TicketTypeRequest struct {
	Name  domain.TicketTypeName `json:"name"`
	Price float64               `json:"price"`
}
