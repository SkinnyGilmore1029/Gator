package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/SkinnyGilmore1029/gator/internal/config"
	"github.com/SkinnyGilmore1029/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	//Read the config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("Error reading config file", err)
	}

	//Now I can use the loaded config file
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatal("Error opening database connection", err)
	}

	dbQueries := database.New(db)

	s := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	c := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	c.register("login", handlerLogin)
	c.register("register", handlerRegister)
	c.register("reset", handlerReset)
	c.register("users", handlerListUsers)
	c.register("agg", handlerAgg)

	if len(os.Args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}

	commandName := os.Args[1]
	arguments := os.Args[2:]

	cmd := command{
		name: commandName,
		Args: arguments,
	}

	if err := c.run(s, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
