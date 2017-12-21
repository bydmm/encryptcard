package main

import (
	"encryptcard/share"
	"fmt"
	"os"
)

func saveBlock(block share.CardBlock) {
	path := fmt.Sprintf("./saves/%s.json", block.CardID())
	fmt.Printf("save to :%s\n", path)
	f, err := os.Create(path)
	f.WriteString(block.JSON())
	if err != nil {
		panic(err)
	}
}

func showCardInfo(block share.CardBlock) {
	card, err := block.Card()
	if err != nil {
		log.Infof("error card: %d\n", block.CardID())
	}
	log.Infof("CardBlockReceive: %s\n", block.CardID())
	log.Infof("id: %d\n", card.ID)
	log.Infof("attack: %d\n", card.Attack)
	log.Infof("defense: %d\n", card.Defense)
}
