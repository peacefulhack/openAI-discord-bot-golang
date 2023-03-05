package commands

import (
	"github.com/bwmarrin/discordgo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

func ping(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	joinArgs := strings.Join(args, " ")
	_, err := s.ChannelMessageSend(m.ChannelID, "Pong! "+joinArgs)
	if err != nil {
		return err
	}
	return nil
}
func pong(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	joinArgs := strings.Join(args, " ")
	_, err := s.ChannelMessageSend(m.ChannelID, "Ping! "+joinArgs)
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}
	return nil
}
