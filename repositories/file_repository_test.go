package repositories

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

//WhiteBox testing. Checks if file was saved on location
func TestSave(t *testing.T) {

	// Init
	repo := NewFileRepo()
	context := "test-context"
	id := "987333222"
	code := "abcde"

	// Execution
	err := repo.Save(context, id, code)

	// Validation
	assert.Nil(t, err)

}

//Tests that only one distinct record is persisted.
func TestUniqueContent(t *testing.T) {
	// Init
	repo := NewFileRepo()
	context := "test-context"
	id := "987333222"

	code1 := "abcde"

	// Execution
	saveErr1 := repo.Save(context, id, code1)
	assert.Nil(t, saveErr1)

	content1, findErr1 := repo.Find(context, id)
	assert.Nil(t, findErr1)
	assert.NotNil(t, content1)

	assert.Equal(t, code1, content1)

	log.Printf("For Code1 %v received content 1 %v", code1, content1)

	code2 := "12345"

	saveErr2 := repo.Save(context, id, code2)
	assert.Nil(t, saveErr2)

	content2, findErr2 := repo.Find(context, id)
	assert.Nil(t, findErr2)
	assert.NotNil(t, content2)

	assert.Equal(t, code2, content2)

	log.Printf("For Code2 %v received content 2 %v", code2, content2)

}
