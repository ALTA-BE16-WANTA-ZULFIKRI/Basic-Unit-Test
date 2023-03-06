package main

import (
	"fmt"
	"testing"


	"github.com/stretchr/testify/assert"
	
)

func TestPrintNama(t *testing.T) {
	// your code here
	result := PrintNama("wanta")
	if result != "wanta kobar" {
		t.Error("Result masbi wanta")
	}
	fmt.Println("TestaprintNama wanta")
}

func TestPrintNama2(t *testing.T) {
	// your code here
    result := PrintNama("wanta zulfikri")
    if result!= "wanta kobar" {
        t.Fatal("Result wanta zulfikri")
    }
    fmt.Println("PrintNama wanta")
}

func TestPrintNamaAssert(t *testing.T) {
	// your code here
    result := PrintNama("wanta")
    assert.Equal(t,"wanta kobar", result, "result must be 'wanta kobar'")
    fmt.Println("PrintNama wanta")

}

func TestPrintNama2Assert(t *testing.T) {
	result := PrintNama("WANTA")
	assert.NotEqual(t, "WANTA", result, "result must be 'WANTA'")
	fmt.Println("PrintNama wanta")
}

func TestSubTest(t *testing.T) {
	t.Run("wanta", func(t *testing.T) {
		result := PrintNama("wanta")
	    assert.Equal(t, "wanta kobar", result, "result must be 'wanta kobar'")
    }) 
	t.Run("wanta zulfikri", func(t *testing.T) {
        result := PrintNama("zulfikri")
        assert.Equal(t, "zulfikri kobar", result, "result must be 'wanta kobar'")
    })
}