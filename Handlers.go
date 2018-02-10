package main

import (
	"net/http"
	"strconv"
	"bytes"
	"os"
	"log"
	"net"
	"gopkg.in/russross/blackfriday.v2"
)

func setheader (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin ,Accept, Content-Type, Content-Length, Accept-Encoding")
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	} else { w.Header().Set("Access-Control-Allow-Origin", "*")}
}

func Index(w http.ResponseWriter, r *http.Request) {
	myip := "http://"+getMyIP()+":9191"

	var buffer bytes.Buffer
	buffer.WriteString("# Database Excersise 2!\n")
	buffer.WriteString("Document oriented Databases [Feat. MongoDB]. Link To project description here: [Link](https://github.com/datsoftlyngby/soft2018spring-databases-teaching-material/blob/master/lecture_notes/02-Intro_to_MongoDB.ipynb). Link to github: [Here](https://github.com/Games-of-Threads/DBEX2-DFH)\n\n")
	buffer.WriteString("## Routes!\n\n")
	buffer.WriteString("#### Question 1: \n")
	buffer.WriteString("- How many unique users is there?\n\n")
	buffer.WriteString("/usercount [CLICKME]("+myip+"/usercount)\n\n")
	buffer.WriteString("#### Question 2: \n")
	buffer.WriteString("- Which top 10 users are linking the most to other users?\n\n")
	buffer.WriteString("/mostlinks [CLICKME]("+myip+"/mostlinks)\n\n")
	buffer.WriteString("#### Question 3: \n")
	buffer.WriteString("- Which top 10 users are being linked too the most?\n\n")
	buffer.WriteString("/mentioned [CLICKME]("+myip+"/mentioned)\n\n")
	buffer.WriteString("#### Question 4: \n")
	buffer.WriteString("- Which top 10 users have posted the most (Active)?\n\n")
	buffer.WriteString("/mostactive [CLICKME]("+myip+"/mostactive)\n\n")
	buffer.WriteString("#### Question 5a: \n")
	buffer.WriteString("- Which top 10 users are among the most grumpy?\n\n")
	buffer.WriteString("/mostgrumpy [CLICKME]("+myip+"/mostgrumpy)\n\n")
	buffer.WriteString("#### Question 5b: \n")
	buffer.WriteString("- Which top 10 users are among the most happy?\n\n")
	buffer.WriteString("/mosthappy [CLICKME]("+myip+"/mosthappy)\n\n")

	w.Write(blackfriday.Run([]byte(buffer.String())))
}

func APIusercount(w http.ResponseWriter, r *http.Request){
	setheader(w, r)
	w.Write([]byte("Question 1: How many unique users is there?: \n\n"+strconv.Itoa(UserCount())))
}

func APImostlinks(w http.ResponseWriter, r *http.Request){
	setheader(w, r)
	array := MostLinks()
	w.Write([]byte("Question 2: Which top 10 users are linking the most to other users?\n\n "+arrayToOneString(array)))

}

func APImentioned(w http.ResponseWriter, r *http.Request){
	setheader(w, r)
	array := MostMentions()
	w.Write([]byte("Question 3: Which top 10 users are being linked too the most?\n\n"+arrayToOneString(array)))
}

func APImostactive(w http.ResponseWriter, r *http.Request){
	setheader(w, r)
	array := MostActive()
	w.Write([]byte("Question 4: Which top 10 users have posted the most (Active)?\n\n"+arrayToOneString(array)))
}

func APImostgrumpy(w http.ResponseWriter, r *http.Request){
	setheader(w, r)
	array := MostGrumpy()
	w.Write([]byte("Question 5a: Which top 10 users are among the most grumpy?"+arrayToOneString(array)))
}

func APImosthappy(w http.ResponseWriter, r *http.Request){
	setheader(w, r)
	array := MostHappy()
	w.Write([]byte("Question 5b: Which top 10 users are among the most happy?"+arrayToOneString(array)))
}

func arrayToOneString(array []UserStats)string{
	var buffer bytes.Buffer
	if len(array)>0{
	for i := 0; i<len(array)-1; i++{
		buffer.WriteString("Name: ")
		buffer.WriteString(array[i].Name)
		buffer.WriteString(" - Value: ")
		buffer.WriteString(strconv.Itoa(array[i].Amount))
		buffer.WriteString("\n")
	}
	}else {
		buffer.WriteString("Array was empty")
	}

	return buffer.String()

}


func getMyIP() string {
	name, err := os.Hostname()
	if err != nil {log.Println(err.Error())}
	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {log.Println(err.Error())}
	return addr.String()

}