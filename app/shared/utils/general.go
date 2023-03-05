package utils

import (
	"github.com/joho/godotenv"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"strings"
)

func GetTokenEnv(tokenName string) (string, error) {
	err := godotenv.Load("./.env")
	if err != nil {
		return "", status.Errorf(codes.Internal, err.Error())
	}
	token := os.Getenv(tokenName)
	if token == "" {
		return "", status.Errorf(codes.Internal, "Cannot Access Token")
	}
	return token, nil
}

func GetArgs(str string) (string, []string) {
	split := strings.Split(str, " ")
	if len(split) > 1 {
		return split[0], split[1:]
	}
	return str, nil
}
