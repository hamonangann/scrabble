package pkg

import "errors"

// Error Codes:
// 	001: scoreboard hasn't been created
//  002: players name not existed

var scoreboard map[string]int

// NewScoreboard create new scoreboard with provided players
// It returns valid to false if any duplicated player name
func NewScoreboard(players []string) bool {
	scoreboard = make(map[string]int, 0)

	for _, player := range players {
		scoreboard[player] = 0
	}

	return validateNewScoreboard(players, scoreboard)
}

// validateNewScoreboard validate all player have a unique name (no duplicates)
func validateNewScoreboard(players []string, scoreboard map[string]int) bool {
	// if length of players slice are more/less than scoreboard map
	// it means that there are duplicated player name or unsuccessful scoreboard creation
	return len(players) == len(scoreboard)
}

// AddScore adds score to current player
// The score could be negative
func AddScore(player string, score int) error {
	if err := validateNewScore(player); err != nil {
		return err
	}
	scoreboard[player] += score
	return nil
}

// validateNewScore validate new score input
// Scoreboard must be created and player name must be existed on scoreboard
func validateNewScore(player string) error {
	if scoreboard == nil {
		return errors.New("001")
	}
	_, exist := scoreboard[player]
	if !exist {
		return errors.New("002")
	}
	return nil
}

// GetCurrentScore returns current scoreboard
func GetCurrentScore() (map[string]int, bool) {
	// if scoreboard has not been initialized, return false
	if scoreboard == nil {
		return nil, false
	}
	return scoreboard, true
}
