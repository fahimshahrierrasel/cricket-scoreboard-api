package responsemodels

import (
	"cricket-scoreboard-api/src/models"
)

// Player godoc
// @Summary Define Player model
type Player struct {
	ID         string
	Name       string
	PlayerType models.PlayerType
	TeamID     string
}
