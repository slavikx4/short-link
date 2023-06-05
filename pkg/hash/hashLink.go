package hash

import (
	"hash/crc64"
	"strings"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"

const alphabetLen = 63

func HashLink(s string) string {
	//считаем контрольную сумму
	table := crc64.MakeTable(crc64.ISO)
	hash := crc64.Checksum([]byte(s), table)

	//короткая ссылка
	charArray := make([]uint8, 0, 10)

	//каждый следующий символ будет выбран по индексу alphabet,где индекс равен остатку от деления контрольного числа
	for mod := uint64(0); hash != 0 && len(charArray) < 10; {
		mod = hash % alphabetLen
		hash /= alphabetLen
		charArray = append(charArray, alphabet[mod])
	}

	//если осталось место в charArray дополнить нулями
	return strings.Repeat(string(alphabet[0]), 10-len(charArray)) + string(charArray)
}
