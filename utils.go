package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

const aoisensi = "135617831864762368"

func containsString(s []string, e string) bool {
	for _, s := range s {
		if s == e {
			return true
		}
	}
	return false
}

func isAdmin(s *discordgo.Session, e *discordgo.MessageCreate) (bool, error) {
	if e.GuildID == "" {
		return false, nil
	}
	dGuild, err := s.Guild(e.GuildID)
	if err != nil {
		log.Println(err)
		return false, err
	}
	if dGuild.OwnerID == e.Author.ID {
		return true, err
	}
	for _, roleID := range e.Member.Roles {
		for _, role := range dGuild.Roles {
			if roleID == role.ID {
				if (role.Permissions & 0x8) > 0 {
					return true, nil
				}
			}
		}
	}
	return false, nil
}
