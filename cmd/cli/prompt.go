package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"scrabble/pkg"
	"strconv"
	"strings"
)

// printScoreboard prints current players with their respective score
// calling printScoreboard without any player/making new scoreboard is a no-op
func printScoreboard() {
	scoreboard, created := pkg.GetCurrentScore()
	if created {
		fmt.Println("Papan skor")

		for player, score := range scoreboard {
			fmt.Println(player, ":", score, "point")
		}

		fmt.Println("Semangat!")
		fmt.Println()
	}
}

// printDictionary prints current added words with their meaning
// calling printDictionary without making new dictionary is a no-op
func printDictionary() {
	dictionary, created := pkg.GetDictionary()
	if created {
		fmt.Println("Kamus")

		if len(dictionary) == 0 {
			fmt.Println("Tidak ada kata ditambahkan")
		} else {
			for word, meaning := range dictionary {
				fmt.Println(word, ":", meaning)
			}
		}

		fmt.Println("Semangat!")
		fmt.Println()
	}
}

// printLedger prints ledger from current scoreboard
// calling printLedger without making scoreboard/any player is a no-op
func printLedger() {
	scoreboard, valid := pkg.GetCurrentScore()
	if valid {
		ledger := pkg.GenerateLedgerFromScoreboard(scoreboard)

		for _, transaction := range ledger {
			diff := transaction[2].(int)
			if diff > 0 {
				fmt.Println(transaction[1], "membayar", transaction[0], diff, "poin")
			} else {
				fmt.Println(transaction[0], "membayar", transaction[1], -diff, "poin")
			}
		}
		fmt.Println()
	}

}

func listOfCommands() {
	fmt.Println("Daftar commands:")
	fmt.Println("HELP: melihat daftar commands")
	fmt.Println("EXIT: menyelesaikan permainan. Skor akhir akan tercetak")
	fmt.Println("ADD SCORE <player> <score>: menambahkan <score> poin ke <player>. <score> bisa negatif")
	fmt.Println("ADD WORD <word>: menambahkan <word> ke daftar kata yang pernah dimainkan")
}

func initialGame(reader *bufio.Reader) {
	fmt.Println("Berapa banyak player yang ingin main?")

	numberOfPlayerIsValid := false
	numberOfPlayers := 0
	fmt.Println("Masukkan jumlah player:")
	for !numberOfPlayerIsValid {
		input, err := reader.ReadString('\n')
		if err != nil { // should never happen
			log.Fatal("Terjadi kesalahan. Silakan ulangi permainan.")
		}
		numberOfPlayers, err = strconv.Atoi(strings.TrimSuffix(input, "\n"))
		if err != nil {
			fmt.Println(err)
			fmt.Println("Input tidak valid, masukkan jumlah player (0-9):")
		} else {
			numberOfPlayerIsValid = true
		}
	}
	fmt.Println()

	fmt.Println("Siapa saja yang mau main?")

	players := make([]string, 0)
	playersIsValid := false
	for !playersIsValid {
		for i := 0; i < numberOfPlayers; i++ {
			input, err := reader.ReadString('\n')
			if err != nil { // should never happen
				log.Fatal("Terjadi kesalahan. Silakan ulangi permainan.")
			}
			player := strings.ToUpper(strings.TrimSuffix(input, "\n"))
			fmt.Println(player)
			players = append(players, player)
		}
		playersIsValid = pkg.NewScoreboard(players)
		if !playersIsValid {
			fmt.Println("Nama pemain harus berbeda semua. Ulangi masukkan semua yang mau main:")
		}
	}

	pkg.NewDictionary()
}

func validateAndAddScore(playerInput string, scoreInput string) {
	score, err := strconv.Atoi(strings.TrimSuffix(scoreInput, "\n"))
	if err != nil {
		fmt.Println("Masukan skor tidak valid")
		return
	}
	err = pkg.AddScore(playerInput, score)
	if err == nil { // NO error happened
		fmt.Println("Skor", score, "berhasil ditambahkan ke", playerInput)
	} else if err.Error() == "001" { // should not happen
		fmt.Println("Terjadi kesalahan. Mohon ulangi program")
	} else if err.Error() == "002" {
		fmt.Println("Masukan tidak valid. Nama pemain tidak ada dalam daftar")
	} else {
		fmt.Println("Error tidak terduga:", err.Error())
	}
}

func addWord(word string, meaning string) {
	err := pkg.AddNewWord(word, meaning)
	if err == nil {
		fmt.Println("Kata", word, "dengan arti", meaning, "berhasil ditambahkan ke kamus")
	} else if err.Error() == "101" { // should not happen
		fmt.Println("Terjadi kesalahan. Mohon ulangi program")
	} else if err.Error() == "102" {
		fmt.Println("Masukan tidak valid. Kata sudah pernah dimainkan")
	} else {
		fmt.Println("Error tidak terduga:", err.Error())
	}
}

func prompt() {
	reader := bufio.NewReader(os.Stdin)
	initialGame(reader)
	printScoreboard()

	fmt.Println("Masukkan perintah. Ketik 'HELP' untuk melihat daftar perintah.")
	exit := false
	for !exit {
		input, err := reader.ReadString('\n')
		if err != nil { // should never happen
			log.Fatal("Terjadi kesalahan. Silakan ulangi permainan.")
		}
		commands := strings.Split(strings.ToUpper(strings.TrimSuffix(input, "\n")), " ")
		if len(commands) == 1 && commands[0] == "EXIT" {
			exit = true
		} else if len(commands) == 1 && commands[0] == "HELP" {
			listOfCommands()
		} else if len(commands) == 4 && commands[0] == "ADD" && commands[1] == "SCORE" {
			validateAndAddScore(commands[2], commands[3])
		} else if len(commands) == 4 && commands[0] == "ADD" && commands[1] == "WORD" {
			addWord(commands[2], commands[3])
		} else if len(commands) == 2 && commands[0] == "PRINT" && commands[1] == "SCOREBOARD" {
			printScoreboard()
		} else if len(commands) == 2 && commands[0] == "PRINT" && commands[1] == "DICTIONARY" {
			printDictionary()
		} else {
			fmt.Println("Command tidak valid. Masukkan ulang command yang valid")
		}
	}

	// game is over

	printScoreboard()
	printLedger()
	printDictionary()
}
