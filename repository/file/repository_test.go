package file

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//WhiteBox testing. Checks if file was saved on location
func TestSave(t *testing.T){

    // Init
    repo := NewFileRepo()
    context := "test-context"
    id := "987333222"
    code := "abcde"

    // Execution
    filePath, err := repo.Save(context, id, code)
    
    // Validation
    assert.Nil(t, err)
    assert.NotEmpty(filePath)

    f, err2 := os.Open(filePath)
    defer f.Close()

    
}

func TestUniqueContent(t *testing.T)){
    // Init
    repo := NewFileRepo()
    context := "test-context"
    id := "987333222"
    
    code1 := "abcde"
   
    // Execution
    filePath1, err1 := repo.Save(context, id, code1)
    
    f1, err2 := os.Open(filePath1)
    
    //Read content and confirm there is only one line
    ioutil.ReadAll(f1)
    f1.Close()   

    code2 := "12345"
    filePath2, err2 := repo.Save(context, id, code2)
    

}