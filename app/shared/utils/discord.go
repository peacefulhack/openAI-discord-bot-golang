package utils

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var DiscordToken string

func NewDC() error {
	res, err := GetTokenEnv("DISCORD_BOT_TOKEN")
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}
	DiscordToken = res
	return nil
}
