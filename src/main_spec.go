package main

import(
	"fmt";
	"./specify";
)

func init() {
	fmt.Println("main.init");
}

func main() {
	specify.Run();
}