package items

import (
	"github.com/dailyhunt/feed/util"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestItemCreation(t *testing.T) {
	item := NewItem("95079741")
	assert.NotNil(t, item, "item should not be null")
}

func TestItemCreationWithIgs(t *testing.T) {
	language := util.HINDI
	igs := strings.Split("HI_X_128_52,HI_X_256_51", ",")
	item := createItem("95079741", igs, language)
	assert.Equal(t, 2, len(item.igs))
}

func TestItemCount(t *testing.T) {
	// given
	item123 := createItem("123", []string{"HI_X_128_52"}, util.HINDI)
	item234 := createItem("234", []string{"HI_X_256_52"}, util.HINDI)
	languageStore := NewLanguageStore(10)
	languageStore.Add(item123)
	languageStore.Add(item234)
	// wait for channel to consume all items
	time.Sleep(100 * time.Millisecond)
	languageStore.Flush()

	// verify
	idToIndexes := languageStore.indexes
	assert.Equal(t, 2, languageStore.Count())
	assert.Equal(t, 2, len(idToIndexes))
}

// todo : actual and expected
func TestItemIgBitSet(t *testing.T) {
	// Given
	item123 := createItem("123", []string{"HI_X_128_52", "HI_X_256_52"}, util.HINDI)
	item234 := createItem("234", []string{"HI_X_256_52"}, util.HINDI)
	languageStore := NewLanguageStore(10)
	languageStore.Add(item123)
	languageStore.Add(item234)
	// wait for channel to consume all items
	time.Sleep(100 * time.Millisecond)
	languageStore.Flush()

	// verify
	idToIndexes := languageStore.indexes
	assert.Equal(t, 2, languageStore.Count())

	itemIndex := idToIndexes["123"]

	shouldPass := languageStore.queryForFilter("HI_X_256_52", itemIndex)
	assert.Equal(t, true, shouldPass)

	shouldNotPass := languageStore.queryForFilter("EN_X_256_52", itemIndex)
	assert.Equal(t, false, shouldNotPass)

}

func createItem(itemId string, igs []string, language util.Language) *Item {
	return NewItem(itemId).
		setIgs(igs).
		setLang(language)
}
