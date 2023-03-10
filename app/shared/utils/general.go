package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"path/filepath"
	"strconv"
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

func CheckFolder(folderName string) bool {
	_, err := os.Stat("temp")
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateFolder(folderName string) error {
	err := os.Mkdir(folderName, 0755)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFolder(folderName string) error {
	err := os.Remove(folderName)
	if err != nil {
		return err
	}
	return nil
}

func NameQueFolder(folderName, prefix string) (string, error) {
	var lastNum, nextNum, num int
	var prefixd, prefixs, nextFileName string
	prefixd = strings.Replace(prefix, "*", "%03d", 1)
	prefixs = strings.Replace(prefix, "*", "%03s", 1)

	files, err := filepath.Glob(filepath.Join(folderName, prefix))

	if err != nil {
		return "", err
	}

	for _, file := range files {
		fileName := filepath.Base(file)
		_, err := fmt.Sscanf(fileName, prefixd, &num)
		if err == nil && num > lastNum {
			lastNum = num
		}
	}

	nextNum = lastNum + 1
	nextFileName = fmt.Sprintf(prefixs, strconv.Itoa(nextNum))
	return nextFileName, nil

}
