package util

type Language uint8

const (
	UNKNOWN Language = iota
	HINDI
	ENGLISH
	TAMIL
)

func (lang Language) String() string {
	languages := [...]string{
		"UNKNOWN",
		"HINDI",
		"ENGLISH",
	}
	return languages[lang]
}
