package items

import (
	"github.com/willf/bitset"
)

const (
	defaultExpectedSize      = 10000
	minBitsetExpansionRation = 0.7
)

type LanguageItems struct {
	items             []*Item
	filterSize        int
	queryFilters      map[string]*bitset.BitSet
	indexes           map[string]uint // Item id to index
	itemChan          chan *Item
	scoredItemIndexes []uint
	done              chan interface{}
}

func (l *LanguageItems) Add(item *Item) {
	if item != nil {
		l.itemChan <- item
	}
}

func (l *LanguageItems) Count() int {
	return len(l.items)
}

func (l *LanguageItems) Flush() {
	close(l.done)
	close(l.itemChan)
}

func (l *LanguageItems) insert() {
	for {
		select {
		case item, ok := <-l.itemChan:
			if ok {
				l.insertItem(item)
			}
		case <-l.done:
			//fmt.Println("Done ..flushing this store")
			return
		}
	}
}

func (l *LanguageItems) insertItem(item *Item) {

	// Add item to list
	l.items = append(l.items, item)
	index := uint(len(l.items) - 1)

	// Update item to index mapping
	if _, ok := l.indexes[item.id]; !ok {
		l.indexes[item.id] = index
	}

	// Prepare filters
	for _, ig := range item.igs {
		l.updateFilter(ig, index)
	}

	// Todo : Prepare other filters

}

func (l *LanguageItems) itemIndexes() *map[string]uint {
	return &l.indexes
}

func (l *LanguageItems) updateFilter(filterKey string, itemIndex uint) {
	if filterBitSet, ok := l.queryFilters[filterKey]; ok {
		// Note : Extending bitset is handled by bitset module
		filterBitSet.Set(itemIndex)
		return
	}

	filterBitSet := bitset.New(itemIndex + 100) // Todo : should be from configuration
	filterBitSet.Set(itemIndex)
	l.queryFilters[filterKey] = filterBitSet
	//fmt.Println("Key ", filterKey, " => ", filterBitSet.String())
}

func (l *LanguageItems) queryForFilter(filterKey string, index uint) bool {
	if bs, ok := l.queryFilters[filterKey]; ok {
		return bs.Test(index)
	}
	return false
}

func NewLanguageStore(expectedSize int) *LanguageItems {
	newSize := defaultExpectedSize
	if expectedSize > 0 {
		newSize = expectedSize
	}

	languageItems := &LanguageItems{
		items:        make([]*Item, 0, newSize),
		filterSize:   newSize,
		queryFilters: make(map[string]*bitset.BitSet),
		indexes:      make(map[string]uint),
		itemChan:     make(chan *Item),
		done:         make(chan interface{}),
	}

	go languageItems.insert()

	return languageItems

}
