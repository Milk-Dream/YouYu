package main

import (
	"log"
	"os"
)

func init() {
	//惯例:前缀都是大写

	log.SetPrefix("[小黑日志]")
	file, _ := os.OpenFile("youyu.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY,0666)

	log.SetOutput(file)//将日志打印到文件中
	//Ldate Ltime Lmicroseconds Llongfile Lshortfile
	//LUTC Lmsgprefix LstdFlags
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile)
}

func main() {
	/*
	log.Println("这是一条日志")
	log.Panicln("这是一个恐慌")
	log.Fatalln("这是一个严重错误")
	*/
}