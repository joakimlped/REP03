package mycrypt

var ALF_SEM03 []rune = []rune("abcdefghijklmnopqrstuvwxyzæøå0123456789.,:; ")

func Krypter(melding []rune, alphabet []rune, chiffer int) []rune {
	kryptertMelding := make([]rune, len(melding))
	for i := 0; i < len(melding); i++ {
		indeks := sokIAlfabetet(melding[i], alphabet)
		if indeks == -1 {
			// Hvis symbolet ikke finnes i alfabetet, legg det til uendret i kryptertMelding
			kryptertMelding[i] = melding[i]
			continue
		}
		nyIndeks := (indeks + chiffer) % len(alphabet)
		kryptertMelding[i] = alphabet[nyIndeks]
	}
	return kryptertMelding
}

func sokIAlfabetet(symbol rune, alfabet []rune) int {
	for i := 0; i < len(alfabet); i++ {
		if symbol == alfabet[i] {
			return i
			break
		}
	}
	return -1
}
