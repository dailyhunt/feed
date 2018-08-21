package items

import (
	"github.com/dailyhunt/feed/util"
)

type Item struct {
	id       string
	igs      []string
	language util.Language
}

func (i *Item) setIgs(igs []string) *Item {
	i.igs = igs
	return i
}

func (i *Item) setLang(lang util.Language) *Item {
	i.language = lang
	return i

}

func NewItem(id string) *Item {
	return &Item{id: id}
}
