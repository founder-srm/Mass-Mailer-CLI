package main

import (
    "fmt"
)

func cliPrompt() {
    var senderEmail string
    var csvFilePath string
    var confirm string
    var subject string
    var htmlTemplatePath string

    asciiText1 := `
___________                      .___                    _________ .__       ___.    
\_   _____/___  __ __  ____    __| _/___________  ______ \_   ___ \|  |  __ _\_ |__  
 |    __)/  _ \|  |  \/    \  / __ |/ __ \_  __ \/  ___/ /    \  \/|  | |  |  \ __ \ 
 |     \(  <_> )  |  /   |  \/ /_/ \  ___/|  | \/\___ \  \     \___|  |_|  |  / \_\ \
 \___  / \____/|____/|___|  /\____ |\___  >__|  /____  >  \______  /____/____/|___  /
     \/                   \/      \/    \/           \/          \/               \/ 

`
    asciiText2 := `
 /$$$$$$$$                                  /$$                                      /$$$$$$  /$$           /$$      
| $$_____/                                 | $$                                     /$$__  $$| $$          | $$      
| $$     /$$$$$$  /$$   /$$ /$$$$$$$   /$$$$$$$  /$$$$$$   /$$$$$$   /$$$$$$$      | $$  \__/| $$ /$$   /$$| $$$$$$$ 
| $$$$$ /$$__  $$| $$  | $$| $$__  $$ /$$__  $$ /$$__  $$ /$$__  $$ /$$_____/      | $$      | $$| $$  | $$| $$__  $$
| $$__/| $$  \ $$| $$  | $$| $$  \ $$| $$  | $$| $$$$$$$$| $$  \__/|  $$$$$$       | $$      | $$| $$  | $$| $$  \ $$
| $$   | $$  | $$| $$  | $$| $$  | $$| $$  | $$| $$_____/| $$       \____  $$      | $$    $$| $$| $$  | $$| $$  | $$
| $$   |  $$$$$$/|  $$$$$$/| $$  | $$|  $$$$$$$|  $$$$$$$| $$       /$$$$$$$/      |  $$$$$$/| $$|  $$$$$$/| $$$$$$$/
|__/    \______/  \______/ |__/  |__/ \_______/ \_______/|__/      |_______/        \______/ |__/ \______/ |_______/ 
                                                                                                                     
`
    fmt.Println(asciiText1)
    fmt.Println(asciiText2)
    fmt.Print("Welcome to mass mailer of Founders Club!")
    fmt.Println()

    fmt.Print("Enter Senders email address: ")
    fmt.Scanf("%s", &senderEmail)

    fmt.Print("Path to csv file: ")
    fmt.Scanf("%s", &csvFilePath)

    fmt.Printf("Path to html template: ")
    fmt.Scanf("%s", &htmlTemplatePath)

    fmt.Print("Subject: ")
    fmt.Scanf("%s", &subject)

    fmt.Print("Confirm? [y/n]: ")
    fmt.Scanf("%s", &confirm)
}

func main() {
    cliPrompt()
}
