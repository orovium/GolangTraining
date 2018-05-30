package main

func main() {
	myBool := (true && false) || (false && true) || !(false && false)
	println(myBool)
}
