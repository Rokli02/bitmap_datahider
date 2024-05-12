package main

import (
	"fmt"
	"os"
)

const hideCommand = "hide \"...assests mappához relatív elérés/../neve.bmp\" \"...valami titkosítandó szöveg...\""
const revealCommand = "reveal \"...assests mappához relatív elérés/../neve.bmp\""

func inputCheck() error {
	if BITS_USED_FOR_ENCODE > 7 {
		return fmt.Errorf("a képben való adatelrejtéshez legfeljebb 7 bit használható fel")
	}

	if len(os.Args) < 3 {
		return fmt.Errorf("a program használatához szükséges megadni a végrehajtandó műveletet:\n\tDekódolás= '%s'\n\tTitkosítás= '%s'", revealCommand, hideCommand)
	}

	if len(os.Args) < 4 && os.Args[1] != "reveal" {
		return fmt.Errorf("a titkosítás kötelezően követendő szintaktikája: '%s'", hideCommand)
	}

	return nil
}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}
