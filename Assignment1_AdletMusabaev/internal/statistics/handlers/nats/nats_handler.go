package nats

import (
	"Assignment1_AdletMusabaev/internal/statistics/services"
	"Assignment1_AdletMusabaev/proto"
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

type StatisticsNATSHandler struct {
	svc *services.StatisticsService
}

func NewStatisticsNATSHandler(svc *services.StatisticsService) *StatisticsNATSHandler {
	return &StatisticsNATSHandler{svc: svc}
}

func (h *StatisticsNATSHandler) Subscribe(nc *nats.Conn) {
	nc.Subscribe("order.created", func(msg *nats.Msg) {
		var event proto.OrderEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Printf("Failed to unmarshal order event: %v", err)
			return
		}
		if err := h.svc.ProcessOrderEvent(&event); err != nil {
			log.Printf("Failed to process order event: %v", err)
		}
	})

	nc.Subscribe("inventory.created", func(msg *nats.Msg) {
		var event proto.InventoryEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Printf("Failed to unmarshal inventory event: %v", err)
			return
		}
		// Process inventory event if needed
	})
}
