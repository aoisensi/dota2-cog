package main

import (
	"context"
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/aoisensi/dota2-auto-role/models"
	"github.com/bwmarrin/discordgo"
	"github.com/volatiletech/sqlboiler/boil"
)

func createGuild(s *discordgo.Session, g *discordgo.Guild) {
	id, _ := strconv.ParseInt(g.ID, 10, 63)
	guild := models.Guild{
		ID: id,
	}
	if err := guild.Insert(context.Background(), db, boil.Infer()); err != nil {
		log.Println(err)
	}
	roles := createRoles(s, g)
	for _, role := range roles {
		role.GuildID = id
		if err := role.Insert(context.Background(), db, boil.Infer()); err != nil {
			log.Println(err)
		}
	}
}

func createRoles(s *discordgo.Session, g *discordgo.Guild) []*models.Role {
	roles := make([]*models.Role, len(rankNames))
	for i, rank := range rankNames {
		name := "Dota2 " + rank
		role, err := s.GuildRoleCreate(g.ID)
		if err != nil {
			log.Println(err)
			continue
		}
		role, err = s.GuildRoleEdit(g.ID, role.ID, name, role.Color, role.Hoist, role.Permissions, false)
		if err != nil {
			log.Println(err)
			continue
		}
		id, _ := strconv.ParseInt(role.ID, 10, 63)
		roles[i] = &models.Role{
			ID:   id,
			Rank: i,
		}
	}
	return roles
}

func watchAll(s *discordgo.Session) {
	guilds, err := models.Guilds().All(context.Background(), db)
	if err != nil {
		log.Println(err)
		return
	}
	for _, guild := range guilds {
		watchGuild(s, guild)
	}
}

func watchGuild(s *discordgo.Session, g *models.Guild) {
	guildID := strconv.FormatInt(g.ID, 10)
	members, err := s.GuildMembers(guildID, "", 1000)
	if err != nil {
		log.Println(err)
		return
	}
	roles, err := g.Roles().All(context.Background(), db)
	if err != nil {
		log.Println(err)
		return
	}
	for _, member := range members {
		time.Sleep(time.Second * 5)
		id, _ := strconv.ParseInt(member.User.ID, 10, 63)
		user, err := models.FindUser(context.Background(), db, id)
		if err != nil {
			if err != sql.ErrNoRows {
				log.Println(err)
			}
			continue
		}
		newRank, err := fetchRank(user.SteamID)
		if err != nil {
			log.Println(err)
			continue
		}
		// remove all cog roles
		for _, roleID := range member.Roles {
			for _, r := range roles {
				if roleID == strconv.FormatInt(r.ID, 10) {
					s.GuildMemberRoleRemove(guildID, member.User.ID, roleID)
					break
				}
			}
		}
		medal := newRank / 10
		for _, role := range roles {
			if role.Rank == medal {
				s.GuildMemberRoleAdd(guildID, member.User.ID, strconv.FormatInt(role.ID, 10))
				break
			}
		}
	}
}

func onGuildCreate(s *discordgo.Session, e *discordgo.GuildCreate) {
	id, _ := strconv.ParseInt(e.ID, 10, 63)
	_, err := models.FindGuild(context.Background(), db, id)
	if err == sql.ErrNoRows {
		createGuild(s, e.Guild)
	} else if err != nil {
		log.Print(err)
	}
}

func onGuildDelete(s *discordgo.Session, e *discordgo.GuildDelete) {
	id, _ := strconv.ParseInt(e.ID, 10, 63)
	guild := &models.Guild{ID: id}
	if _, err := guild.Delete(context.Background(), db); err != nil {
		log.Println(err)
	}
}
