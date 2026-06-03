package helper

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

// func TestHelloWorld(t *testing.T) {
// 	result := HelloWorld("diosatu")

// 	if result != "Hello diosatu" { // kalo hasil nya beda dari yang di inginkan maka test akan gagal
// 		//unit test failed
// 		panic("Result is not Hello diosatu")  // kalo pake panic test nya bakal berenti semua
// 	}
	
// }

// kalo mau run per function bisa pake  
// go test -v //
// go test -v -run=TestHelloWorld // jika function nya ada unsur nama yang sama dia bakal ikut ke run, sama kaya like '%abcd%'

// buat ngegagalin test ada function fail(),FailNow()
// ,Eror(),Errorf(),Fatal(),Fatalf() 






// ini function ke 1
func TestHelloWorld1(t *testing.T) {
	result := HelloWorld1("diosatu coba")

	if result != "Hello diosatu coba" { // kalo hasil nya beda dari yang di inginkan maka test akan gagal
		//unit test failed
		t.Fail()
	}
	fmt.Println("TestHelloWorldcoba sukses")
	
}


// ini function ke 2
func TestHelloWorld2(t *testing.T) {
	result := HelloWorld2("diosatu coba2")

	if result != "Hello diosatu coba2" { // kalo hasil nya beda dari yang di inginkan maka test akan gagal
		//unit test failed
		t.FailNow()
	}
	fmt.Println("TestHelloWorldcoba sukses 2")
	
}


// ini function ke 3
func TestHelloWorld3(t *testing.T) {
	result := HelloWorld3("diosatu coba3")

	if result != "Hello diosatu coba3" { // kalo hasil nya beda dari yang di inginkan maka test akan gagal
		//unit test failed
		t.Error("Result should be 'Hello diosatu coba3'")
	}
	fmt.Println("TestHelloWorldcoba sukses 3")
	
}


// ini function ke 4
func TestHelloWorld4(t *testing.T) {
	result := HelloWorld4("diosatu coba4")

	if result != "Hello diosatu coba4" { // kalo hasil nya beda dari yang di inginkan maka test akan gagal
		//unit test failed
		t.Fatal("Result should be 'Hello diosatu coba4'")
	}
	fmt.Println("TestHelloWorldcoba sukses 4")
	
}


// function buat testing assertion
// ini function ke 5
func TestHelloWorld5(t *testing.T) {
	result := HelloWorld5("diosatu coba5")
	assert.Equal(t, "Hello diosatu coba5", result, "Result harus nya 'Hello diosatu coba5'")
	fmt.Println("TestHelloWorld5 sukses")
}