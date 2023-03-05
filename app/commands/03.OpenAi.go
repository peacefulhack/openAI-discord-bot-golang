package commands

import (
	"github.com/bwmarrin/discordgo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"peacefulhack/discord/app/shared/utils"
	"strings"
)

func ChatGPT3Dot5(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	if args == nil {
		_, err := s.ChannelMessageSend(m.ChannelID, "umm.. Hello there!? please add message after calling me")
		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	}

	joinArgs := strings.Join(args, " ")
	message, err := utils.GPT3TurboChat(joinArgs)
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	_, err = s.ChannelMessageSend(m.ChannelID, message)
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}
	return nil
}

func DallE(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	if args == nil {
		_, err := s.ChannelMessageSend(m.ChannelID, "umm.. Hello there!? please add message after calling me")
		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	}

	joinArgs := strings.Join(args, " ")
	message, err := utils.DallEChat(joinArgs)
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	msgEmbed := &discordgo.MessageEmbed{
		Title: "This is an image of: " + joinArgs,
		Image: &discordgo.MessageEmbedImage{
			URL: message,
		},
	}

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, msgEmbed)
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}
	return nil
}
