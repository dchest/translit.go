// Copyright (c) 2011 Dmitry Chestnykh
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package translit implements non-standard one-way string transliteration from Cyrillic to Latin.
package translit

import (
	"unicode"
	"strings"
)

// Mapping of Russian Cyrillic characters to ASCII Latin.
var RussianASCII = map[int]string{
	'а': "a",
	'б': "b",
	'в': "v",
	'г': "g",
	'д': "d",
	'е': "e",
	'ё': "yo",
	'ж': "zh",
	'з': "z",
	'и': "i",
	'й': "j",
	'к': "k",
	'л': "l",
	'м': "m",
	'н': "n",
	'о': "o",
	'п': "p",
	'р': "r",
	'с': "s",
	'т': "t",
	'у': "u",
	'ф': "f",
	'х': "h",
	'ц': "c",
	'ч': "ch",
	'ш': "sh",
	'щ': "sch",
	'ъ': "'",
	'ы': "y",
	'ь': "",
	'э': "e",
	'ю': "ju",
	'я': "ja",
}

// Mapping of Cyrillic characters to (Czech/Serbian) Latin.
var CyrillicLatin = map[int] string{
	'а': "a",
	'б': "b",
	'в': "v",
	'г': "g",
	'ґ': "g",
	'ѓ': "ǵ",
	'д': "d",
	'ђ': "đ",
	'е': "e",
	'ё': "ë",
	'ж': "ž",
	'з': "z",
	'ѕ': "ẑ",
	'и': "i",
	'і': "ì",
	'ї': "ï",
	'й': "j",
	'ј': "j",
	'к': "k",
	'л': "l",
	'љ': "lj",
	'м': "m",
	'н': "n",
	'њ': "nj",
	'о': "o",
	'п': "p",
	'р': "r",
	'с': "s",
	'т': "t",
	'ћ': "ć",
	'у': "u",
	'ф': "f",
	'х': "h",
	'ц': "c",
	'ч': "č",
	'џ': "dž",
	'ш': "š",
	'щ': "šč",
	'ъ': "'",
	'ы': "y",
	'ь': "",
	'ѣ': "ě",
	'э': "e",
	'є': "ê",
	'ю': "ju",
	'я': "ja",	
}

// ToLatin returns a string transliterated using provided mapping table.
// Runes that cannot be transliterated inserted as-is.
func ToLatin(s string, table map[int]string) string {
	runes := []int(s)
	out := make([]int, 0, len(s))
	for i, rune := range runes {
		if tr, ok := table[unicode.ToLower(rune)]; ok {
			if tr == "" {
				continue
			}
			if unicode.IsUpper(rune) {
				// Correctly translate case of successive characters:
				// ЩИ -> SCHI
				// Щи -> Schi
				if i+1 < len(runes) && !unicode.IsUpper(runes[i+1]) {
					t := []int(tr)
					t[0] = unicode.ToUpper(t[0])
					out = append(out, t...)
					continue
				}
				out = append(out, []int(strings.ToUpper(tr))...)
				continue
			}
			out = append(out, []int(tr)...)
		} else {
			out = append(out, rune)
		}
	}
	return string(out)
}
