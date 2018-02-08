package main



// go get gopkg.in/mgo.v2
// go get “github.com/gorilla/mux”
// go get github.com/rs/cors
// go get -u gopkg.in/russross/blackfriday.v2



func main() {




	//println(UserCount())
	MostLinks()


	/*
	router := NewRouter()
	go log.Fatal(http.ListenAndServe(":8080", router))

	CLI := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter Command. \nAvailable Commands: \n userscount     mostlinks         mentioned          mostmactive         mostgrumpy        exit\n")
	for {

		input, _ := CLI.ReadString('\n')
		command := strings.Fields(input)

		switch command[0] {
		case "userscount":

		case "mostlinks":

		case "mostmentions":

		case "mostactive":

		case "mostgrumpy":

		case "exit":
			os.Exit(2)

		default:
			println("Did not recognize the command.")
			fmt.Print("\nEnter Command. \nAvailable Commands: \n initsdb [sdb name]      savesdb      put [key] [value...]       view [key]       viewall      delete [key]        exit\n")
		}

	}

	*/
}