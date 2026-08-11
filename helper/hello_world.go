package helper

func HelloWorld(name string) string {
	return "Hello " + name
}
func HelloWorld1(name string) string {
	return "Hello " + name
}
func HelloWorld2(name string) string {
	return "Hello " + name
}
func HelloWorld3(name string) string {
	return "Hello " + name
}
func HelloWorld4(name string) string {
	return "Helloo " + name
}
func HelloWorld5assert(name string) string { // ini yg assertion jika test gagal maka test akan dihentikan tapi tetap lanjut ke test berikutnya
	return "Hello " + name
}
func HelloWorld6Require(name string) string { // ini yg require jika test gagal maka test akan dihentikan dan tidak akan lanjut ke test berikutnya
	return "Hello " + name
}
func TestSkip1(name string) string { // ini yg require jika test gagal maka test akan dihentikan dan tidak akan lanjut ke test berikutnya
	return "nyoba di OS " + name
}

