package main

import (
	"log"
	"net/http"
	"bufio"
	"os"
	"fmt"
	"strings"
)

// go get gopkg.in/mgo.v2
// go get “github.com/gorilla/mux”
// go get github.com/rs/cors
// go get -u gopkg.in/russross/blackfriday.v2



func main() {


	log.Println("GoApplication READY!!! you can now go to "+getMyIP()+":9191 and interact from there or type in one of the CLI commands here")
	/*
	//println(UserCount())
	//test := MostLinks()
	//test := MostMentions()
	//test := MostHappy()
	//test := MostGrumpy()
	if len(test)>0{
		for i := 0; i<9; i++{
			println(test[i].Name)
			println(test[i].Amount)
		}
	}else {
		println("Uuuhh test was 0 length")
	}

	*/

	go RunCLI()

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":9191", router))

}

func ForLoopPrint(array []UserStats){
	if len(array)>0{
		for i := 0; i<len(array)-1; i++{
			println(array[i].Name)
			println(array[i].Amount)
		}
	}else {
		println("Uuuhh test was 0 length")
	}
}

func RunCLI(){
	CLI := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter Command. \nAvailable Commands: \n userscount     mostlinks         mentioned          mostmactive         mostgrumpy      mosthappy      exit\n")
	for {

		input, _ := CLI.ReadString('\n')
		command := strings.Fields(input)
		if len(command)!=0{

		switch command[0] {
		case "userscount":
			println(UserCount())
		case "mostlinks":
			array := MostLinks()
			ForLoopPrint(array)
		case "mostmentions":
			array := MostMentions()
			ForLoopPrint(array)
		case "mostactive":
			array := MostActive()
			ForLoopPrint(array)
		case "mostgrumpy":
			array := MostGrumpy()
			ForLoopPrint(array)
		case "mosthappy":
			array := MostHappy()
			ForLoopPrint(array)
		case "exit":
			os.Exit(2)

		default:
			println("Did not recognize the command.")
			fmt.Print("\nEnter Command. \nAvailable Commands: \n userscount     mostlinks         mentioned          mostmactive         mostgrumpy      mosthappy      exit\n")
		}
		}else {
			println("Did not recognize the command."
			fmt.Print("\nEnter Command. \nAvailable Commands: \n userscount     mostlinks         mentioned          mostmactive         mostgrumpy      mosthappy      exit\n")
		}
	}
}

