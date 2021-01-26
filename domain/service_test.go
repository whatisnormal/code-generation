package domain

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

//Check if a random seq with a given argument
//generates unique elements all the times
func TestRandSeq(t *testing.T) {
	//Run 10 times, generate 10 unique codes.
	set := map[string]bool{}

	for i := 0; i < 20; i++ {
		code := randSeq(10)
		if !assert.Equal(t, 10, len(code)) {
			break
		}

		set[code] = true
	}

	assert.Equal(t, 20, len(set))
}

func TestGenerate(t *testing.T) {
	//Given
	repo := &inMemRepository{}
	notifier := &dummyNotifier{}
	codeManagerService := NewCodeManagerService(repo, notifier)

	r := &GenerateReq{
		APIKey:  "1234",
		Context: "sms",
		ID:      "923459988",
	}

	//When
	code, err := codeManagerService.Generate(r)

	log.Printf("Generated code %v", code)

	//Then
	assert.Nil(t, err)
	assert.NotNil(t, code)
}
