package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aoisensi/dota2-cog/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/yohcop/openid-go"
)

var oidNonceStore = openid.NewSimpleNonceStore()
var oidDiscoveryCache = openid.NewSimpleDiscoveryCache()

func httpServer() {
	r := gin.Default()
	store := cookie.NewStore([]byte(env.SessionKey))
	r.Use(sessions.Sessions("session", store))
	r.GET("/", handleIndex)
	r.GET("/add-bot", handleAddBot)
	r.GET("/connect", handleConnect)
	r.GET("/connect-ok", handleConnectOK)
	r.Run()
}

func handleIndex(c *gin.Context) {
	f, err := os.Open("./views/index.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	data, _ := ioutil.ReadAll(f)
	c.Writer.Write(data)
}

func handleAddBot(c *gin.Context) {
	url := fmt.Sprintf(
		"https://discordapp.com/oauth2/authorize?client_id=%v&scope=bot&permissions=268438528",
		env.DiscordClientID,
	)
	c.Redirect(http.StatusFound, url)
}

func handleConnect(c *gin.Context) {
	nonce := c.Query("nonce")
	id := findNonce(nonce)
	if id == 0 {
		c.String(http.StatusBadRequest, "")
		return
	}
	session := sessions.Default(c)
	session.Set("discord-id", id)
	url, err := openid.RedirectURL(
		"https://steamcommunity.com/openid/",
		env.Host+"/connect-ok",
		env.Host,
	)
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, "")
		return
	}
	session.Save()
	c.Redirect(http.StatusFound, url)
}

func handleConnectOK(c *gin.Context) {
	session := sessions.Default(c)
	discordID, ok := session.Get("discord-id").(int64)
	if !ok {
		log.Println("session missing")
		c.String(http.StatusBadRequest, "")
		return
	}
	uri := env.Host + c.Request.URL.String()
	url, err := openid.Verify(uri, oidDiscoveryCache, oidNonceStore)
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, "")
		return
	}
	var id int64
	fmt.Sscanf(url, "https://steamcommunity.com/openid/id/%d", &id)
	user := models.User{
		DiscordID: discordID,
		SteamID:   id,
	}
	err = user.Upsert(context.Background(), db, true, []string{"discord_id"}, boil.Infer(), boil.Infer())
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "")
		return
	}
	deleteNonce(discordID)
	c.String(http.StatusOK, "The link has been successfully completed!!\nYou may close this window.")
}
