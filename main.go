package main

import (
    "bufio"
    "bytes"
    "encoding/csv"
    "fmt"
    "gopkg.in/gomail.v2"
    "html/template"
    "os"
    "strings"
    "time"

    tea "github.com/charmbracelet/bubbletea"
)

var senderEmail string
var csvFilePath string
var confirm string
var subject string
var htmlTemplatePath string
var emails []string

func sendGoMail(subject string, templatePath string, toEmails []string) {
    println("sending mail...")
    // Get html
    var body bytes.Buffer
    t, err := template.ParseFiles(templatePath)
    if err != nil {
        fmt.Println(err)
        return
    }
    t.Execute(&body, nil)

    //send with go mail
    m := gomail.NewMessage()
    m.SetHeader("From", "vinayakcsurya@gmail.com")
    m.SetHeader("To", toEmails...)
    //m.SetAddressHeader("Cc", "dan@example.com", "Dan")
    m.SetHeader("Subject", subject)
    m.SetBody("text/html", body.String())
    m.Attach("img.png")

    d := gomail.NewDialer("smtp.gmail.com", 587, "vinayakcsurya@gmail.com", "lblvwntpckhbbjds")

    // Send the email to Bob, Cora and Dan.
    if err := d.DialAndSend(m); err != nil {
        panic(err)
    }
    println("mail sent!")
}

func CliPrompt() {
    reader := bufio.NewReader(os.Stdin)

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

    fmt.Println(asciiText2)
    fmt.Printf("Welcome to mass mailer of Founders Club!")
    fmt.Println()

    //fmt.Print("Enter Senders email address: ")
    //fmt.Scanln(&senderEmail)

    fmt.Print("Path to csv file: ")
    csvFilePath, _ = reader.ReadString('\n')
    csvFilePath = strings.TrimSpace(csvFilePath)

    fmt.Print("Path to html template: ")
    htmlTemplatePath, _ = reader.ReadString('\n')
    htmlTemplatePath = strings.TrimSpace(htmlTemplatePath)

    fmt.Print("Subject: ")
    subject, _ = reader.ReadString('\n')
    subject = strings.TrimSpace(subject)

    fmt.Print("Confirm? [y/n]: ")
    confirm, _ = reader.ReadString('\n')
    confirm = strings.TrimSpace(confirm)

    if confirm == "y" {
        ReadCSV(csvFilePath)
        time.Sleep(1 * time.Second)
        sendGoMail(subject, htmlTemplatePath, emails)
    } else {
        fmt.Printf("%s", confirm)
        println("terminated!")
        os.Exit(0)
    }
}

func ReadCSV(csvFilePath string) {
    println("reading CSV file...")
    // Open the CSV file
    file, err := os.Open(csvFilePath)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    reader := csv.NewReader(file)

    records, err := reader.ReadAll()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    for i, record := range records {
        fmt.Printf("Record %d: %v\n", i+1, record)

        // Assuming the emails in each row are separated by commas
        for _, field := range record {
            // Split the emails in the field by commas
            splitEmails := strings.Split(field, ",")
            // Append each email to the emails slice
            emails = append(emails, splitEmails...)
        }
    }

    fmt.Println("Emails:", emails)
}

func main() {
    //subject := "This is a html subject"
    //templatePath := "/Users/vinayak/IdeaProjects/GO-projects/mailer/template.html"
    //csv := "/Users/vinayak/IdeaProjects/GO-projects/mailer/emails.csv"
    //emails := []string{"vinayak.chandra.suryavanshi@gmail.com", "vs9419@srmist.edu.in"}
    //sendGoMail("This is a subject", "/Users/vinayak/IdeaProjects/GO-projects/mailer/template.html", emails)
    CliPrompt()

}
