package pkg

// GenerateLedgerFromScoreboard creates ledger consist of point difference for each player
// Calling this function with nil scoreboard is a no-op

func GenerateLedgerFromScoreboard(scoreboard map[string]int) [][]any {
	players := make([]string, 0)
	for player, _ := range scoreboard {
		players = append(players, player)
	}

	ledger := make([][]any, 0)
	if scoreboard != nil {
		for i := 0; i < len(players); i++ {
			for j := i + 1; j < len(players); j++ {
				score1, player1exist := scoreboard[players[i]]
				score2, player2exist := scoreboard[players[j]]
				if player2exist && player1exist && score1 != score2 {
					transaction := []any{players[i], players[j], score1 - score2}
					ledger = append(ledger, transaction)
				}
			}
		}
	}

	return ledger
}
