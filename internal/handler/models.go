package handler

import (
	"encoding/gob"
	"encoding/json"
	"time"

	"github.com/brocaar/lorawan"
)

func init() {
	gob.Register(DataUpPayload{})
	gob.Register(JoinNotification{})
	gob.Register(ACKNotification{})
	gob.Register(ErrorNotification{})
	gob.Register(StatusNotification{})
}

// Location details.
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
}

// RXInfo contains the RX information.
type RXInfo struct {
	GatewayID lorawan.EUI64 `json:"gatewayID"`
	Name      string        `json:"name"`
	Time      *time.Time    `json:"time,omitempty"`
	RSSI      int           `json:"rssi"`
	LoRaSNR   float64       `json:"loRaSNR"`
	Location  *Location     `json:"location"`
}

// TXInfo contains the TX information.
type TXInfo struct {
	Frequency int `json:"frequency"`
	DR        int `json:"dr"`
}

// DataUpPayload represents a data-up payload.
type DataUpPayload struct {
	ApplicationID   int64         `json:"applicationID,string"`
	ApplicationName string        `json:"applicationName"`
	DeviceName      string        `json:"deviceName"`
	DevEUI          lorawan.EUI64 `json:"devEUI"`
	RXInfo          []RXInfo      `json:"rxInfo,omitempty"`
	TXInfo          TXInfo        `json:"txInfo"`
	ADR             bool          `json:"adr"`
	FCnt            uint32        `json:"fCnt"`
	FPort           uint8         `json:"fPort"`
	Data            []byte        `json:"data"`
	Object          interface{}   `json:"object,omitempty"`
}

// DataDownPayload represents a data-down payload.
type DataDownPayload struct {
	ApplicationID int64           `json:"applicationID,string"`
	DevEUI        lorawan.EUI64   `json:"devEUI"`
	Reference     string          `json:"reference"`
	Confirmed     bool            `json:"confirmed"`
	FPort         uint8           `json:"fPort"`
	Data          []byte          `json:"data"`
	Object        json.RawMessage `json:"object"`
}

// JoinNotification defines the payload sent to the application on
// a JoinNotificationType event.
type JoinNotification struct {
	ApplicationID   int64           `json:"applicationID,string"`
	ApplicationName string          `json:"applicationName"`
	DeviceName      string          `json:"deviceName"`
	DevEUI          lorawan.EUI64   `json:"devEUI"`
	DevAddr         lorawan.DevAddr `json:"devAddr"`
}

// ACKNotification defines the payload sent to the application
// on an ACK event.
type ACKNotification struct {
	ApplicationID   int64         `json:"applicationID,string"`
	ApplicationName string        `json:"applicationName"`
	DeviceName      string        `json:"deviceName"`
	DevEUI          lorawan.EUI64 `json:"devEUI"`
	Reference       string        `json:"reference"`
	Acknowledged    bool          `json:"acknowledged"`
	FCnt            uint32        `json:"fCnt"`
}

// ErrorNotification defines the payload sent to the application
// on an error event.
type ErrorNotification struct {
	ApplicationID   int64         `json:"applicationID,string"`
	ApplicationName string        `json:"applicationName"`
	DeviceName      string        `json:"deviceName"`
	DevEUI          lorawan.EUI64 `json:"devEUI"`
	Type            string        `json:"type"`
	Error           string        `json:"error"`
	FCnt            uint32        `json:"fCnt,omitempty"`
}

// StatusNotification defines the payload sent to the application
// on a device-status reporting.
type StatusNotification struct {
	ApplicationID   int64         `json:"applicationID,string"`
	ApplicationName string        `json:"applicationName"`
	DeviceName      string        `json:"deviceName"`
	DevEUI          lorawan.EUI64 `json:"devEUI"`
	Battery         int           `json:"battery"`
	Margin          int           `json:"margin"`
}
