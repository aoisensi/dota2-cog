package main

import (
	"database/sql"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/robfig/cron/v3"
)

var (
	db  *sql.DB
	env Env
)

func main() {
	var err error
	db, err = sql.Open("postgres", env.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		panic(err)
	}

	go httpServer()

	ds, err := discordgo.New("Bot " + env.DiscordBotToken)
	if err != nil {
		panic(err)
	}
	ds.AddHandler(onGuildCreate)
	ds.AddHandler(onGuildDelete)
	ds.AddHandler(onMessageCreate)

	c := cron.New()
	c.AddFunc("* * * * *", deleteExpiredNonce)
	c.AddFunc("0 * * * *", func() { watchAll(ds) })
	c.Start()

	if err := ds.Open(); err != nil {
		panic(err)
	}
	<-make(chan bool)
}

func init() {
	log.SetFlags(log.Lshortfile)
	if err := envconfig.Process("", &env); err != nil {
		panic(err)
	}
}

// Env is environment variable struct for envconfig
type Env struct {
	DiscordClientID string `required:"true" split_words:"true"`
	DiscordBotToken string `required:"true" split_words:"true"`
	DatabaseURL     string `required:"true" split_words:"true"`
	Host            string `required:"true" split_words:"true"`
	SessionKey      string `require:"true" split_words:"true"`
}
