package main

import (
	"io/ioutil"

	"gopkg.in/macaron.v1"
)

func main() {
	m := macaron.Classic()

	// mainpage route
	// will show list of user games
	// how can we recognize user? generate user id to cookie on any request

	// get user stat * after db

	// get all stat * after db

	// create game
	// will return game id
	// will start goroutine with game id
	// goroutine will die after timeout without state update while we keep it in memory
	// we will move game states to db and call them by id and kill after week * after db

	// communication with game:

	// show game by id

	// end game ? by id

	// post word guess by id

	//

	m.Get("/", func() []byte {
		page, err := ioutil.ReadFile("public/index.html")
		if err != nil {
			return []byte("Duh")
		}

		return page
	})

	m.Run()
}
