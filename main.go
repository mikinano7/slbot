package main

import (
    "fmt"
    "log"
    "strings"
    "time"
    "github.com/joho/godotenv"
    "github.com/mattn/go-haiku"
    "github.com/ikawaha/slackbot"
    "math/rand"
    "os"
)

const botResponseSleepTime = 3 * time.Second

func main() {
    err := godotenv.Load("slack_auth.env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    bot, err := slackbot.New(os.Getenv("SLACK_TOKEN"))
    if err != nil {
        log.Fatal(err)
    }
    defer bot.Close()

    r575 := []int{5, 7, 5}
    for {
        msg, err := bot.GetMessage()

        if err != nil {
            log.Printf("receive error, %v", err)
        }

        if msg.Type != "message" {
            continue
        }

        go func(m slackbot.Message) {
            t := m.TextBody()
            hs := haiku.Find(t, r575)
            if len(hs) < 1 {
                if strings.Contains(t, "を？") || strings.Contains(t, "を！？") {
                    if (rand.Intn(2) == 1) {
                        m.Text = "できらぁ！"
                    } else {
                        m.Text = "できねえーー！"
                    }

                    time.Sleep(botResponseSleepTime)
                    bot.PostMessage(m)
                } else {
                    return
                }
            } else {
                var tmp []string
                for _, h := range hs {
                    tmp = append(tmp, fmt.Sprintf("`%v`", h))
                }

                m.Text = strings.Join(tmp, "\n") + " 575だ"
                time.Sleep(botResponseSleepTime)
                bot.PostMessage(m)
            }
        }(msg)
    }
}
