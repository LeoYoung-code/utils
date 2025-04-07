package encrypt

import "errors"

var (
	InvalidInput = errors.New("invalid input")
)

/* 128 common bigrams. */
const bigrams = "intherreheanonesorteattistenntartondalitseediseangoulecomeneriroderaioicliofasetvetasihamaecomceelllcaurlachhidihofonsotacnarssoprrtsassusnoiltsemctgeloeebetrnipeiepancpooldaadviunamutwimoshyoaiewowosfiepttmiopiaweagsuiddoooirspplscaywaigeirylytuulivimabty"

/* 256 common English words of length four letters or more. */
var words = []string{
	"that", "this", "with", "from", "your", "have", "more", "will", "home",
	"about", "page", "search", "free", "other", "information", "time", "they",
	"what", "which", "their", "news", "there", "only", "when", "contact", "here",
	"business", "also", "help", "view", "online", "first", "been", "would", "were",
	"some", "these", "click", "like", "service", "than", "find", "date", "back",
	"people", "list", "name", "just", "over", "year", "into", "email", "health",
	"world", "next", "used", "work", "last", "most", "music", "data", "make",
	"them", "should", "product", "post", "city", "policy", "number", "such",
	"please", "available", "copyright", "support", "message", "after", "best",
	"software", "then", "good", "video", "well", "where", "info", "right", "public",
	"high", "school", "through", "each", "order", "very", "privacy", "book", "item",
	"company", "read", "group", "need", "many", "user", "said", "does", "under",
	"general", "research", "university", "january", "mail", "full", "review",
	"program", "life", "know", "days", "management", "part", "could", "great",
	"united", "real", "international", "center", "ebay", "must", "store", "travel",
	"comment", "made", "development", "report", "detail", "line", "term", "before",
	"hotel", "send", "type", "because", "local", "those", "using", "result",
	"office", "education", "national", "design", "take", "posted", "internet",
	"address", "community", "within", "state", "area", "want", "phone", "shipping",
	"reserved", "subject", "between", "forum", "family", "long", "based", "code",
	"show", "even", "black", "check", "special", "price", "website", "index",
	"being", "women", "much", "sign", "file", "link", "open", "today", "technology",
	"south", "case", "project", "same", "version", "section", "found", "sport",
	"house", "related", "security", "both", "county", "american", "game", "member",
	"power", "while", "care", "network", "down", "computer", "system", "three",
	"total", "place", "following", "download", "without", "access", "think",
	"north", "resource", "current", "media", "control", "water", "history",
	"picture", "size", "personal", "since", "including", "guide", "shop",
	"directory", "board", "location", "change", "white", "text", "small", "rating",
	"rate", "government", "child", "during", "return", "student", "shopping",
	"account", "site", "level", "digital", "profile", "previous", "form", "event",
	"love", "main", "another", "class", "still",
}

// CompressBytes compresses a byte slice. it returns the length of compressed bytes.
func CompressBytes(s []byte) []byte {
	dst := make([]byte, 0)
	verbatimLen := 0
	sb := s

	for len(sb) > 0 {
		var i int
		var word string
		wordMatch := false

		if len(sb) >= 4 {
			for i, word = range words {
				if len(sb) >= len(word) && string(sb[:len(word)]) == word {
					wordMatch = true
					break
				}
			}
		}

		if wordMatch {
			switch {
			case sb[0] == ' ':
				dst = append(dst, 8) // Space + word escape.
				dst = append(dst, byte(i))
				sb = sb[1:]
			case len(sb) > len(word) && sb[len(word)] == ' ':
				dst = append(dst, 7) // Word + space escape.
				dst = append(dst, byte(i))
				sb = sb[1:]
			default:
				dst = append(dst, 6) // Just word escape.
				dst = append(dst, byte(i))
			}

			sb = sb[len(word):]
			verbatimLen = 0
			continue
		}

		if len(sb) >= 2 {
			for i = 0; i < len(bigrams); i += 2 {
				if string(sb[:2]) == bigrams[i:i+2] {
					break
				}
			}

			if i < len(bigrams) {
				dst = append(dst, 1<<7|byte(i/2))
				sb = sb[2:]
				verbatimLen = 0
				continue
			}
		}

		if 0 < sb[0] && sb[0] < 9 && sb[0] < 128 {
			dst = append(dst, sb[0])
			sb = sb[1:]
			verbatimLen = 0
			continue
		}

		verbatimLen++
		if verbatimLen == 1 {
			dst = append(dst, byte(verbatimLen), sb[0])
		} else {
			dst = append(dst, sb[0])
			dst[len(dst)-(verbatimLen+1)] = byte(verbatimLen)
			if verbatimLen == 5 {
				verbatimLen = 0
			}
		}
		sb = sb[1:]
	}
	return dst
}

// Compress compresses a string.return the compressed bytes.
func Compress(s string) []byte {
	sb := []byte(s)
	return CompressBytes(sb)
}

// DecompressToBytes decompresses a byte slice to the original bytes.
func DecompressToBytes(c []byte) ([]byte, error) {
	if len(c) == 0 {
		return nil, InvalidInput
	}
	i := 0
	res := make([]byte, 0)
	for i < len(c) {
		switch {
		case c[i]&128 != 0:
			idx := c[i] & 127
			res = append(res, bigrams[idx*2:idx*2+2]...)
			i++
		case 0 < c[i] && c[i] < 6:
			res = append(res, c[i+1:i+1+int(c[i])]...)
			i += 1 + int(c[i])
		case 5 < c[i] && c[i] < 9:
			if c[i] == 8 {
				res = append(res, ' ')
			}
			res = append(res, words[c[i+1]]...)
			if c[i] == 7 {
				res = append(res, ' ')
			}
			i += 2
		default:
			res = append(res, c[i])
			i++
		}
	}
	return res, nil
}

// Decompress decompresses a byte slice to the original string.
func Decompress(c []byte) (string, error) {
	origin, err := DecompressToBytes(c)
	if err != nil {
		return "", err
	}
	return string(origin), nil
}
