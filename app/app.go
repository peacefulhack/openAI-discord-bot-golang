package app

import (
	"github.com/bwmarrin/discordgo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"os/signal"
	"peacefulhack/discord/app/commands"
	"peacefulhack/discord/app/shared/utils"
	"syscall"
)

func New() error {
	Dclient, err := discordgo.New("Bot " + utils.DiscordToken)
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	Dclient.AddHandler(MessageCreate)

	Dclient.Identify.Intents = discordgo.IntentGuildMessages

	err = Dclient.Open()
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	log.Println("Bot is working like a horse hehehe")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	err = Dclient.Close()
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}
	return nil
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || m.Author.ID != "453895628439814165" {
		return
	}
	pref, args := utils.GetArgs(m.Content)
	if command, ok := commands.CommandList[pref]; ok {
		err := command(s, m, args)
		if err != nil {
			log.Println(err)
		}
	}
}
