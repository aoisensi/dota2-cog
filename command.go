package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/aoisensi/dota2-cog/models"
	"github.com/bwmarrin/discordgo"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func onMessageCreate(s *discordgo.Session, e *discordgo.MessageCreate) {
	switch e.Message.Content {
	case "/dota2-cog debug delete-roles":
		commandDebugDeleteRoles(s, e)
	case "/dota2-cog debug hi":
		commandDebugHi(s, e)
	case "/dota2-cog register":
		commandRegister(s, e)
	case "/dota2-cog force-fetch":
		commandForceFetch(s, e)
	case "/dota2-cog fix-rank-roles":
		commandFixRankRoles(s, e)
	case "/dota2-cog fix-registerd-role":
		commandFixRegisterdRoles(s, e)
	}
}

func commandRegister(s *discordgo.Session, e *discordgo.MessageCreate) {
	id, _ := strconv.ParseInt(e.Author.ID, 10, 63)
	nonce := makeConnectNonce(id)

	mes := fmt.Sprintf(
		"%v\n%v/connect?nonce=%v",
		"Please link your account from this URL.",
		env.Host, nonce,
	)
	s.ChannelMessageSend(e.ChannelID, mes)
}

func commandForceFetch(s *discordgo.Session, e *discordgo.MessageCreate) {
	admin, err := isAdmin(s, e)
	if err != nil {
		return
	}
	if !admin {
		s.ChannelMessageSend(e.ChannelID, "This command is only available to the administrator.")
		return
	}
	if _, ok := fetching[e.GuildID]; ok {
		s.ChannelMessageSend(e.ChannelID, "I'm fetching. Don't spam it.")
		return
	}
	id, _ := strconv.ParseInt(e.GuildID, 10, 63)
	guild, err := models.FindGuild(context.Background(), db, id)
	if err != nil {
		log.Println(err)
		return
	}
	s.ChannelMessageSend(e.ChannelID, "Force fetching...")
	pid, _ := s.ChannelMessageSend(e.ChannelID, "▱▱▱▱▱▱▱▱▱▱ 0%")
	pc := make(chan progress)
	defer close(pc)
	go func() {
		xo := -1
		for {
			p := <-pc
			x := p.Done * 10 / p.All
			if x == xo {
				continue
			}
			msg := fmt.Sprintf(
				"%v%v %v%%",
				strings.Repeat("▰", x),
				strings.Repeat("▱", 10-x),
				p.Done*100/p.All,
			)
			s.ChannelMessageEdit(e.ChannelID, pid.ID, msg)
			if p.Done == p.All {
				break
			}
			xo = x
		}
	}()
	result, err := watchGuild(s, guild, pc)
	if err != nil {
		log.Println(err)
		s.ChannelMessageSend(e.ChannelID, "Failed! Sorry...")
		return
	}
	pc <- progress{Done: 1, All: 1} //force close
	s.ChannelMessageSend(e.ChannelID, fmt.Sprintf("Done!! Success: %v  Error: %v", result.Success, result.Error))

}

func commandFixRankRoles(s *discordgo.Session, e *discordgo.MessageCreate) {
	admin, err := isAdmin(s, e)
	if err != nil {
		log.Println(err)
		return
	}
	if !admin {
		s.ChannelMessageSend(e.ChannelID, "This command is only available to the administrator.")
		return
	}
	s.ChannelMessageSend(e.ChannelID, "Auto fixing...")
	roles, err := models.RankRoles(qm.Where("guild_id = ?", e.GuildID)).All(context.Background(), db)
	if err != nil {
		log.Println(err)
		return
	}
	for _, role := range roles {
		id := strconv.FormatInt(role.ID, 10)
		_, err := s.State.Role(e.GuildID, id)
		if err != nil && err != discordgo.ErrStateNotFound {
			log.Println(err)
			return
		}
		if err == nil {
			continue
		}
		msg := fmt.Sprintf("%v role is missing. Recreating...", rankNames[role.Rank])
		s.ChannelMessageSend(e.ChannelID, msg)
		if _, err := role.Delete(context.Background(), db); err != nil {
			log.Println(err)
			continue
		}
		name := "Dota2 " + rankNames[role.Rank]
		roleID, err := createRole(s, e.GuildID, name)
		if err != nil {
			log.Println(err)
			continue
		}
		gid, _ := strconv.ParseInt(e.GuildID, 10, 63)
		role := &models.RankRole{
			ID:      roleID,
			GuildID: gid,
			Rank:    role.Rank,
		}
		role.Insert(context.Background(), db, boil.Infer())
	}
	s.ChannelMessageSend(e.ChannelID, "Auto fix done!!")
}

func commandFixRegisterdRoles(s *discordgo.Session, e *discordgo.MessageCreate) {
	admin, err := isAdmin(s, e)
	if err != nil {
		log.Println(err)
		return
	}
	if !admin {
		s.ChannelMessageSend(e.ChannelID, "This command is only available to the administrator.")
		return
	}
	role, err := models.RegisterdRoles(qm.Where("guild_id = ?", e.GuildID)).One(context.Background(), db)
	if err != nil {
		if err == sql.ErrNoRows {
			r, err := createRegisterdRole(s, e.GuildID)
			if err != nil {
				log.Println(err)
				return
			}
			r.Insert(context.Background(), db, boil.Infer())
			s.ChannelMessageSend(e.ChannelID, "Added `Dota2 Cog Registerd` role!")
			return
		}
		log.Println(err)
		return
	}
	_, err = s.State.Role(e.GuildID, strconv.FormatInt(role.ID, 10))
	if err != nil {
		if err == discordgo.ErrStateNotFound {
			r, err := createRegisterdRole(s, e.GuildID)
			if err != nil {
				log.Println(err)
				return
			}
			role.ID = r.ID
			role.Update(context.Background(), db, boil.Infer())
			s.ChannelMessageSend(e.ChannelID, "Added `Dota2 Cog Registerd` role!")
			return
		}
		log.Println(err)
		return
	}
	s.ChannelMessageSend(e.ChannelID, "No errors were found.")
}

func commandDebugDeleteRoles(s *discordgo.Session, e *discordgo.MessageCreate) {
	if e.Member.User.ID != aoisensi {
		return
	}
	guild, err := s.Guild(e.GuildID)
	if err != nil {
		log.Println(err)
		return
	}
	for _, role := range guild.Roles {
		if strings.HasPrefix(role.Name, "Dota2 ") && role.Name != "Dota2 Cog" {
			if err := s.GuildRoleDelete(e.GuildID, role.ID); err != nil {
				log.Println(err)
				continue
			}
		}
	}
}

func commandDebugHi(s *discordgo.Session, e *discordgo.MessageCreate) {
	if e.Member.User.ID != aoisensi {
		return
	}
	s.ChannelMessageSend(e.ChannelID, "sup")
}
