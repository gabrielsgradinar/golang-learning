package main

import "github.com/gabrielsgradinar/golang-learning/apis/configs"

func main(){
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}