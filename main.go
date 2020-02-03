package main

import (
	"bufio"
	"fmt"
	"github.com/ahmdrz/goinsta"
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

func likeUser(inst *goinsta.Instagram, username string, counter int) error {
	user, err := inst.Profiles.ByName(username)
	if err != nil {
		return err
	}
	media := user.Feed()
	for media.Next() {
		for _, item := range media.Items {
			if counter == 0 {
				break
			}
			item.Like()
			time.Sleep(5 * time.Second)
			fmt.Print("liked " + item.Code+"\n")
			counter--
		}
	}
	return nil
}

func main() () {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	os.Getenv("PASSWORD")
	inst := goinsta.New(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
	if err := inst.Login(); err != nil {
		fmt.Print("Login error")
		return
	}
	fmt.Print("Login as "+os.Getenv("USERNAME"))

	inst.Export("~/.goinsta")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome")
	fmt.Println("---------------------")

	likeRegexp := "like[ ]([a-zA-Z0-9._!-@]+)"
	for {
		fmt.Print("your command ✏️ ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("exit", text) == 0 {
			inst.Logout()
			fmt.Println("Goodbye, have a nice day ️🌿")
			break
		}

		if m, _ := regexp.MatchString(likeRegexp, text); m == true {
			username := strings.Replace(text, "like ", "", 1)
			if username == "" {
				fmt.Print("❌ invalid username\n")
			}
			err = likeUser(inst, username, -1)
			if err != nil {
				fmt.Print(err)
			}
		}

	}
}
