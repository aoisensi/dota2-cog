package main

import (
	"context"
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/aoisensi/dota2-cog/models"
	"github.com/bwmarrin/discordgo"
	"github.com/volatiletech/sqlboiler/boil"
	"go.uber.org/multierr"
)

var fetching = make(map[string]struct{}, 16)

func createGuild(s *discordgo.Session, g *discordgo.Guild) error {
	id, _ := strconv.ParseInt(g.ID, 10, 63)
	guild := models.Guild{
		ID: id,
	}
	if err := guild.Insert(context.Background(), db, boil.Infer()); err != nil {
		log.Println(err)
		return err
	}
	roles, err := createRankRoles(s, g.ID)
	if err != nil {
		return err
	}
	for _, role := range roles {
		if err := role.Insert(context.Background(), db, boil.Infer()); err != nil {
			log.Println(err)
			return err
		}
	}
	role, err := createRegisterdRole(s, g.ID)
	if err := role.Insert(context.Background(), db, boil.Infer()); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func createRankRoles(s *discordgo.Session, guildID string) ([]*models.RankRole, error) {
	roles := make([]*models.RankRole, len(rankNames))
	gid, _ := strconv.ParseInt(guildID, 10, 63)
	var errs error
	for i := range rankNames {
		var err error
		name := "Dota2 " + rankNames[i]
		id, err := createRole(s, guildID, name)
		if err != nil {
			errs = multierr.Append(errs, err)
		}
		roles[i] = &models.RankRole{
			ID:      id,
			GuildID: gid,
			Rank:    i,
		}
	}
	return roles, errs
}

func createRegisterdRole(s *discordgo.Session, guildID string) (*models.RegisterdRole, error) {
	name := "Dota2 Cog Registerd"
	gid, _ := strconv.ParseInt(guildID, 10, 63)
	id, err := createRole(s, guildID, name)
	if err != nil {
		return nil, err
	}
	return &models.RegisterdRole{
		GuildID: gid,
		ID:      id,
	}, nil
}

func createRole(s *discordgo.Session, gid string, name string) (int64, error) {
	role, err := s.GuildRoleCreate(gid)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	role, err = s.GuildRoleEdit(gid, role.ID, name, role.Color, role.Hoist, role.Permissions, false)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	id, _ := strconv.ParseInt(role.ID, 10, 63)
	return id, nil
}

func watchAll(s *discordgo.Session) {
	guilds, err := models.Guilds().All(context.Background(), db)
	if err != nil {
		log.Println(err)
		return
	}
	for _, guild := range guilds {
		watchGuild(s, guild, nil)
	}
}

func watchGuild(s *discordgo.Session, g *models.Guild, p chan progress) (*watchGuildResult, error) {
	guildID := strconv.FormatInt(g.ID, 10)
	fetching[guildID] = struct{}{}
	defer delete(fetching, guildID)
	members, err := s.GuildMembers(guildID, "", 1000)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	roles, err := g.RankRoles().All(context.Background(), db)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	registerdRole, rerr := g.RegisterdRole().One(context.Background(), db)
	if rerr != nil && rerr != sql.ErrNoRows {
		log.Println(err)
		return nil, err
	}
	pa := len(members)
	result := watchGuildResult{}
	for i, member := range members {
		if member.User.Bot {
			continue
		}
		err := watchMember(s, g, member, roles, registerdRole)
		if p != nil {
			p <- progress{Done: i + 1, All: pa}
		}
		if err == sql.ErrNoRows {
			result.NoRegister++
		} else if err != nil {
			result.Error++
		} else {
			result.Success++
		}
		time.Sleep(time.Second)
	}
	return &result, nil
}

func watchMember(
	s *discordgo.Session,
	g *models.Guild,
	member *discordgo.Member,
	roles models.RankRoleSlice,
	registerdRole *models.RegisterdRole,
) error {
	id, _ := strconv.ParseInt(member.User.ID, 10, 63)
	guildID := strconv.FormatInt(g.ID, 10)
	user, err := models.FindUser(context.Background(), db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			// User not registerd
			if registerdRole != nil {
				err := s.GuildMemberRoleRemove(guildID, member.User.ID, strconv.FormatInt(registerdRole.ID, 10))
				if err != nil {
					log.Println(err)
				}
			}
			return nil
		}
		log.Println(err)
		return err
	}
	if registerdRole != nil {
		err := s.GuildMemberRoleAdd(guildID, member.User.ID, strconv.FormatInt(registerdRole.ID, 10))
		if err != nil {
			log.Println(err)
		}
	}
	rank, err := fetchRank(user.SteamID)
	if err != nil {
		log.Println(err)
		return err
	}
	medal := rank / 10

	var errs error
	for _, r := range roles {
		id := strconv.FormatInt(r.ID, 10)
		if containsString(member.Roles, id) {
			if r.Rank != medal {
				if err := s.GuildMemberRoleRemove(guildID, member.User.ID, id); err != nil {
					log.Println(err)
					multierr.Append(errs, err)
				}
			}
		} else {
			if r.Rank == medal {
				if err := s.GuildMemberRoleAdd(guildID, member.User.ID, id); err != nil {
					log.Println(err)
					multierr.Append(errs, err)
				}
			}
		}
	}
	return errs
}

func onGuildCreate(s *discordgo.Session, e *discordgo.GuildCreate) {
	id, _ := strconv.ParseInt(e.ID, 10, 63)
	_, err := models.FindGuild(context.Background(), db, id)
	if err == sql.ErrNoRows {
		if err := createGuild(s, e.Guild); err != nil {
			log.Println(err)
		}
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

type watchGuildResult struct {
	Success, NoRegister, Error int
}

type progress struct {
	Done, All int
}
