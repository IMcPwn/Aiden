/* Aiden by IMcPwn.
 * Copyright 2016 Carleton Stuberg

 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at

 * http://www.apache.org/licenses/LICENSE-2.0

 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.

 * For the latest code and contact information visit: http://imcpwn.com
 */

package main

import (
	"fmt"
	"os"
        "strings"
        "strconv"
        "math/rand"

	"github.com/bwmarrin/discordgo"
)

func main() {
        
    TOKEN := os.Getenv("TOKEN")
    /* Alternative using username/password instead of a token
    USERNAME := os.Getenv("USERNAME")
    PASSWORD := os.Getenv("PASSWORD")
    dg, err := discordogo.New(USERNAME, PASSWORD)
    */
    dg, err := discordgo.New(TOKEN)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Register messageCreate as a callback for the messageCreate events.
    dg.AddHandler(messageCreate)

    // Open the websocket and begin listening.
    err = dg.Open()
    if err != nil {
        fmt.Println(err)
        return
    }
    // Make sure we're logged in successfully
    prefix, err := dg.User("@me")
    if err != nil {
        fmt.Println(err)
        fmt.Println("Make sure TOKEN is defined and valid.\nWindows: set TOKEN=change_to_token\nLinux: export TOKEN=change_to_token")
        return
    } else {
        fmt.Println("Logged in as " + prefix.Username)
    }

    fmt.Println("Welcome to Aiden! Press enter to quit.")
    var input string
    fmt.Scanln(&input)
    return
}

func sendMessage(s *discordgo.Session, m *discordgo.MessageCreate, content string) (msg *discordgo.Message, err error) {
    msg, err = s.ChannelMessageSend(m.ChannelID, "@" + m.Author.Username + m.Author.Discriminator + " " + content)
    return
}

func is_prime(n int) (bool) {
    if n <= 1 {
        return false
    } else if n <= 3 {
        return true
    } else if n % 2 == 0 || n % 3 == 0 {
        return false
    }
    i := 5
    for i * i <= n {
        if n % i == 0 || n % (i + 2) == 0 {
            return false
        }
        i = i + 6
   }
   return true
}

func reverse(s string) (result string) {
  for _,v := range s {
    result = string(v) + result
  }
  return 
}


func is_palindrome(n string) (bool) {
    if len(n) < 2 {
        return false
    }
    if reverse(n) == n {
        return true
    } else {
        return false
    }
}

func printUsage(s *discordgo.Session, m *discordgo.MessageCreate) (msg *discordgo.Message, err error) {
    prefix, err := s.User("@me")
    if err != nil {
        fmt.Println(err)
        return
    }
    desc := fmt.Sprintf("```Hi, I'm a bot. Ask me stuff!\nAll commands are prefixed with @%s.\n" +
    "Find my source code at imcpwn.com\n" +
    "Replace what's inside the brackets.\n\n" + 
    "help --> Responds with this message.\n" +
    "add [num1] [num2] --> Adds two whole numbers.\n" +
    "prime [num] --> States if the number is prime or not.\n" +
    "pal [input] --> States if the input is a palindrome or not.\n" +
    "choose [comma list] --> Chooses a value from a list.\n" +
    "answer [question] --> Answers a yes or no question.\n" +
    "[question]? --> Answers a yes or no question.\n" +
    "```", prefix.Username + prefix.Discriminator)
    msg, err = s.ChannelMessageSend(m.ChannelID, "@" + m.Author.Username + m.Author.Discriminator + " " + desc)
    return
}


func handleCommands(s *discordgo.Session, m *discordgo.MessageCreate) {
    if strings.Contains(m.Content, "add") {
        if len(strings.Split(m.Content, " ")) < 4 {
            printUsage(s, m)
            return
        }
        num1, err := strconv.Atoi(strings.Split(m.Content, " ")[2])
        if err != nil {
            fmt.Println(err)
            sendMessage(s, m, "Invalid number.")
            return
        }
        num2, err := strconv.Atoi(strings.Split(m.Content, " ")[3])
        if err != nil {
            sendMessage(s, m, "Invalid number.")
            fmt.Println(err)
            return
        }
        sendMessage(s, m, strconv.Itoa(num1 + num2))
    } else if strings.Contains(m.Content, "answer") || strings.HasSuffix(m.Content, "?") {
        if rand.Intn(2) == 1 {
            sendMessage(s, m, "Yes")
        } else {
            sendMessage(s, m, "No")
        }
    } else if strings.Contains(m.Content, "prime") {
        if len(strings.Split(m.Content, " ")) < 3 {
            printUsage(s, m)
            return
        }
        n, err := strconv.Atoi(strings.Split(m.Content, " ")[2])
        if err != nil {
           fmt.Println(err)
           sendMessage(s, m, "Invalid number.")
           return
        }
        if is_prime(n) {
            sendMessage(s, m, "Yes")
        } else {
            sendMessage(s, m, "No")
        }
    } else if strings.Contains(m.Content, "pal") || strings.Contains(m.Content, "palindrome") {
        if len(strings.Split(m.Content, " ")) < 3 {
            printUsage(s, m)
            return
        }
        n := strings.Split(m.Content, " ")[2]
        if is_palindrome(n) {
            sendMessage(s, m, "Yes")
        } else {
            sendMessage(s, m, "No")
        }
    } else if strings.Contains(m.Content, "choose") {
        if len(strings.Split(m.Content, " ")) < 3 {
            printUsage(s, m)
            return
        }
        list := strings.Split(strings.Split(m.Content, " ")[2], ",")
        sendMessage(s, m, list[rand.Intn(len(list))])
    } else {
        printUsage(s, m)
    }
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated user has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
        if len(m.Mentions) < 1 {
            return
        }
        prefix, err := s.User("@me")
        if err != nil {
            fmt.Println(err)
            return
        } 
        if m.Mentions[0].ID == prefix.ID  {
            fmt.Println("Mentioned. Handling commands.")
            handleCommands(s, m)
        }
}
