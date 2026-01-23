package domain

type ShippingItem struct {
	ProductCode string
	Quantity    int32
}

type Shipping struct {
	ID           int64
	OrderID      int64
	Items        []ShippingItem
	DeliveryDays int32
}

// CalculateDeliveryDays calcula os dias de entrega conforme a regra:
// - Prazo mínimo: 1 dia
// - A cada 5 unidades: +1 dia adicional
func (s *Shipping) CalculateDeliveryDays() {
	totalQuantity := int32(0)
	for _, item := range s.Items {
		totalQuantity += item.Quantity
	}

	// Mínimo 1 dia + 1 dia a cada 5 unidades
	s.DeliveryDays = 1 + (totalQuantity / 5)
}
