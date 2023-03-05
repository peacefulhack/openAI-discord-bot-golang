package commands

import "github.com/bwmarrin/discordgo"

var CommandList = map[string]func(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error{
	"ping":    ping,
	"pong":    pong,
	"chatgpt": ChatGPT3Dot5,
	"dall-e":  DallE,
}
