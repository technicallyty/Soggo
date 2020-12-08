package game

import "github.com/influxdata/influxdb/client"

// State - methods for handling game state
type State struct {
	Clients []*client.Client
}
