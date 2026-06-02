package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("diosatu")

	if result != "Hello diosatu" { // kalo hasil nya beda dari yang di inginkan maka test akan gagal
		//unit test failed
		panic("Result is not Hello diosatu")
	}
	
}
// kalo mau run per function bisa pake  
// go test -v //
// go test -v -run=TestHelloWorld // jika function nya ada unsur nama yang sama dia bakal ikut ke run, sama kaya like '%abcd%'

