package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Check if a random seq with a given argument
//generates unique elements all the times
func TestRandSeq(t *testing.T){
    //Run 10 times, generate 10 unique codes.
    set := map[string]bool{}

    for i:=0 ; i<20;i++{        
        code := randSeq(10)
        if(!assert.Equal(t,10, len(code))){
            break
        }
        
        set[code] = true
    }

    assert.Equal(t, 20, len(set))    
}