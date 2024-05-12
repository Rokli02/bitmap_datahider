package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const HEADER_SIZE = 54
const BITS_USED_FOR_ENCODE = 3
const MAX_LENGTH_OF_WORD_IN_BYTES = 32
const SECRETS_BIT_CLEAR byte = (0b11111111 >> BITS_USED_FOR_ENCODE) << BITS_USED_FOR_ENCODE

func main() {
	err := inputCheck()
	if err != nil {
		fmt.Println(err)
		return
	}

	command, filePath := os.Args[1], os.Args[2]
	var text string
	if len(os.Args) > 3 {
		text = os.Args[3]
	}

	content, err := os.ReadFile(fmt.Sprintf("assets/%s", filePath))
	if err != nil {
		fmt.Println(err)
		return
	}

	switch command {
	case "hide":
		fmt.Println("Adat elrejtés...")
		newMarioByte := hide(content, text)

		f, err := os.Create(fmt.Sprintf("assets/sus/%s", filePath))
		if err != nil {
			return
		}

		w := bufio.NewWriter(f)

		w.Write(newMarioByte)
		w.Flush()
	case "reveal":
		fmt.Println("Felfedés...")

		text = reveal(content)
		if len(text) == 0 {
			fmt.Println("A képfájlban nem volt elrejtett szöveg")

			break
		}

		fmt.Println("Az elrejtett szöveg:", text)
	default:
		fmt.Println("Ismeretlen parancs! Próbáld ki az 'encode' vagy 'decode' parancsot")
	}
}

// GOOD_TO_KNOW: Minden 4. byte (B - G - R - $$$) egy üres, szabadon felhasználható byte
// Ezekben kell elrejteni hogy mennyi byte hosszú az üzenet, a szignifikánsabb 5 bitjében
func hide(picBytesParam []byte, text string) []byte {
	var picBytes []byte = make([]byte, 0, len(picBytesParam))
	picBytes = append(picBytes, picBytesParam...)

	textBytes := []byte(text)
	textLength := len(textBytes)
	textBits := 0
	var picIndex int

	// Előfeldolgozás, el kell rejteni, hogy milyen hosszú a szöveg
	// Első szabadon felhasználható byte terület az 57.
	for picIndex, textLengthOffset := 57, 0; picIndex < len(picBytes) && textLengthOffset < MAX_LENGTH_OF_WORD_IN_BYTES; picIndex, textLengthOffset = picIndex+4, textLengthOffset+(8-BITS_USED_FOR_ENCODE) {
		picBytes[picIndex] |= byte(((textLength & (0b11111 << textLengthOffset)) >> textLengthOffset) << 3)
	}

	for picIndex = HEADER_SIZE; picIndex < len(picBytes) && textBits < len(textBytes)*8; picIndex++ {
		// Az elrejtéshez használt bitek kitakarítása
		picBytes[picIndex] = (picBytes[picIndex] & SECRETS_BIT_CLEAR)
		secretBits := byte(0)

		for i := BITS_USED_FOR_ENCODE - 1; i >= 0 && textBits < len(textBytes)*8; i-- {
			currentByteOffset := textBits - ((textBits >> 3) << 3)

			secretBits |= ((textBytes[textBits>>3] << currentByteOffset) & 0b10000000) >> (7 - i) // A magas értékű bitek kerülnek először bele

			textBits++
		}

		picBytes[picIndex] |= secretBits
	}

	return picBytes
}

func reveal(picBytesParam []byte) string {
	picBytes := picBytesParam[HEADER_SIZE:]
	textBuilder := strings.Builder{}
	textLength := 0

	// Előfeldolgozás, meg kell tudni, hogy milyen hosszú a szöveg
	for picIndex, wordLengthOffset := 57, 0; picIndex < len(picBytesParam) && wordLengthOffset < MAX_LENGTH_OF_WORD_IN_BYTES; picIndex, wordLengthOffset = picIndex+4, wordLengthOffset+(8-BITS_USED_FOR_ENCODE) {
		textLength |= int(picBytesParam[picIndex]>>3) << wordLengthOffset
	}

	var textBits int
	var tempByte byte

	for picIndex := 0; picIndex < len(picBytes) && textBits < textLength*8; picIndex++ {
		// Rejtett bitek kiszedése
		secretBits := picBytes[picIndex] & 0b111
		currentByteOffset := textBits - ((textBits >> 3) << 3)

		// Kiszedett bitek hozzáadása egy ideiglenes byte változóhoz, majd a listához hozzáadni
		bitShiftOffset := 8 - BITS_USED_FOR_ENCODE - currentByteOffset
		if bitShiftOffset > 0 {
			tempByte |= secretBits << bitShiftOffset
		} else {
			tempByte |= secretBits >> abs(bitShiftOffset)
			textBuilder.WriteByte(tempByte)
			tempByte = 0

			if bitShiftOffset != 0 {
				tempByte |= secretBits << (8 + bitShiftOffset)
			}
		}

		textBits += BITS_USED_FOR_ENCODE
	}

	return textBuilder.String()
}
