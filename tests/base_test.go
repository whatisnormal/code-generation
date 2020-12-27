package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/whatisnormal/code-generation/app"
)

//TestMain Yo
func TestMain(m *testing.M) {
	//rest.StartMockupServer()
	fmt.Println("now starting application")
	go app.StartApp()
	os.Exit(m.Run())
	fmt.Println("Application STarted")
}