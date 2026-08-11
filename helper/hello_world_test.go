package helper

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert" // library untuk assertion jika gagal test maka test akan dihentikan
	"github.com/stretchr/testify/require"
	"runtime"
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

	if result != "Helloo diosatu coba4" { // kalo hasil nya beda dari yang di inginkan maka test akan gagal
		//unit test failed
		t.Fatal("Result should be 'Hello diosatu coba4'")
	}
	fmt.Println("TestHelloWorldcoba sukses 4")
	
}


// function buat testing assertion
// ini function ke 5   yg bagian 	"github.com/stretchr/testify/assert"
func TestHelloWorld5(t *testing.T) {
	result := HelloWorld5assert("diosatu coba5")
	assert.Equal(t, "Hello diosatu coba5", result, "Result harusnya 'Hello diosatu coba5") // assertion adalah 
	// function yang digunakan untuk mengecek apakah test yang kita buat berhasil atau gagal

	fmt.Println("HelloWorld5assert sukses")
}

// ini function ke 6 
func TestHelloWorld6(t *testing.T) {
	result := HelloWorld6Require("diosatu coba6")
	require.Equal(t, "Hello diosatu coba6", result, "Result harusnya 'Hello diosatu coba6") // assertion adalah 
	// function yang digunakan untuk mengecek apakah test yang kita buat berhasil atau gagal

	fmt.Println("TestHelloWorld6 sukses")
}

func TestSkipDulu(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can't run on Windows") // skip test jika OS adalah windows
	}
	result := TestSkip1("mana aja")
	require.Equal(t, "nyoba di OS mana aja", result, "Result harusnya 'nyoba di OS mana aja'") // ini assertion untuk 
	// mengecek apakah test berhasil atau gagal 

}


