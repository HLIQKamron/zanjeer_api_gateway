package flespi

import "github.com/Projects/zanjeer_api_gateway/models/flespi"

type Flespi interface {
	GetTelementary() (flespi.AutoGenerated, error)
}
