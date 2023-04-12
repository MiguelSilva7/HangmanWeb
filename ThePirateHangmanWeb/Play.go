package hangman

import (
	"fmt"
	"os"
	"bufio"
)

func Play(Pts *HangData){
	var want_play bool
	want_play = false
	respond := bufio.NewScanner(os.Stdin)
	for {
		Game(Pts)
		fmt.Printf("Continuer à joue ? O/N")
		for {
			respond.Scan()
			respond := respond.Text()
			if respond == "O" {
				want_play = true
				break
			} else if respond == "N" {
				want_play = false
				break
			} else if respond != "N" && respond != "O"{
				fmt.Println("Print O/N")
			}
		}
		if !want_play {
			break
		}
	}
}