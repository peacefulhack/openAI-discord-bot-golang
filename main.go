package main

import (
	"log"
	"peacefulhack/discord/app"
	"peacefulhack/discord/app/shared/utils"
)

func main() {
	err := utils.NewDC()
	if err != nil {
		log.Println(err)
		return
	}
	err = utils.NewAI()
	if err != nil {
		log.Println(err)
		return
	}
	err = app.New()
	if err != nil {
		log.Println(err)
		return
	}
}
