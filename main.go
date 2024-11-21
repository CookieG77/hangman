package main

import (
	"fmt"
	"os"
	"os/signal"
	"slices"

	hgd "github.com/CookieG77/hangman/functions"
)

func main() {
	// Simple code pour nettoyer le terminal à la fin du programme par Crtl+C.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			if sig == os.Interrupt {
				hgd.ClearCmd()
				os.Exit(1)
			}
		}
	}()

	// Chargement des ressources du jeu
	var Logo string
	var Words []string
	ASCIIart := hgd.ReadASCIIArtFile("docs/hangman.txt")
	ASCIIart_alphabet := hgd.ReadASCIIArtFile("docs/standard.txt")

	// Gestion des arguments et lancement du jeu
	Args := hgd.GetArgs()
	if slices.Contains(Args, "-df") {
		go hgd.PlayLoopMusic("docs/deepfriedmusic.mp3")
		Logo = hgd.ReadLogoFile("docs/deepfriedlogo.txt")
		Words = hgd.ReadWordFile("docs/deepfriedwords.txt")
	} else {
		go hgd.PlayLoopMusic("docs/music.mp3")
		Logo = hgd.ReadLogoFile("docs/logo.txt")
		Words = hgd.ReadWordFile("docs/words.txt")
	}
	if slices.Contains(Args, "--letterFile") {
		ASCIIart_alphabet = hgd.ReadASCIIArtFile(Args[slices.Index(Args, "--letterFile")+1])
	}
	if slices.Contains(Args, "--startWith") {
		if len(Args) > slices.Index(Args, "--startWith")+1 {
			var save map[string]interface{}
			err := hgd.LoadSave(&save, Args[slices.Index(Args, "--startWith")+1])
			if err {
				hgd.ClearCmd()
				fmt.Println("Aucune sauvegarde n'a été trouvée, ou elle est corrompue.\n Assurez-vous que le fichier de sauvegarde soit correct.")
				return
			}
			hgd.Start(Words, ASCIIart, Logo, ASCIIart_alphabet, true, save)
		} else {
			hgd.ClearCmd()
			fmt.Println("Aucune sauvegarde n'a été fournit.")
			return
		}
	} else {
		hgd.Start(Words, ASCIIart, Logo, ASCIIart_alphabet, false, nil)
	}
	hgd.ClearCmd()
}
