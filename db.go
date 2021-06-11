package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/bwmarrin/discordgo"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Id string   `json:"id"`
	Db Database `json:"database"`
}

type Database struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
	Name string `json:"name"`
}

func loadUsers(jsonFile *os.File) (*Users, error) {
	// read json as a byte array
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var users Users

	// unmarshal byte array
	json.Unmarshal(byteValue, &users)

	return &users, nil
}

func (u *Users) GetUser(id string) *User {
	// lookup user by discord id
	for i := 0; i < len(users.Users); i++ {
		if users.Users[i].Id == id {
			return &users.Users[i]
		}
	}

	return nil
}

func GenerateDBEmbed(u *User) *discordgo.MessageEmbed {
	// create message embed fields
	fields := make([]*discordgo.MessageEmbedField, 0)
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "Host",
		Value:  u.Db.Host,
		Inline: false,
	})
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "User",
		Value:  u.Db.User,
		Inline: true,
	})
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "Database Name",
		Value:  u.Db.Name,
		Inline: true,
	})
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "Pass",
		Value:  u.Db.Pass,
		Inline: true,
	})
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "Port",
		Value:  u.Db.Port,
		Inline: true,
	})

	// create instance of message embed
	return &discordgo.MessageEmbed{
		Title:       "Your database information",
		Description: "Use MySQL Workbench or any other MySQL interface to connect to your database.",
		Fields:      fields,
		Color:       5570730,
	}
}
