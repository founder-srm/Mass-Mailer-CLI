package main

import (
    "bytes"
    "encoding/csv"
    "fmt"
    "gopkg.in/gomail.v2"
    "html/template"
    "net/smtp"
    "os"
    "strings"
    "time"
)

var senderEmail string
var csvFilePath string
var confirm string
var subject string
var htmlTemplatePath string
var emails []string

func sendMail(subject string, body string, toEmails []string) {
    auth := smtp.PlainAuth("",
        "vinayakcsurya@gmail.com",
        "lblvwntpckhbbjds",
        "smtp.gmail.com")

    msg := "Subject: " + subject + "\n" + body

    err := smtp.SendMail(
        "smtp.gmail.com:587",
        auth,
        "vinayakcsurya@gmail.com",
        toEmails,
        []byte(msg),

    )
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }

}

func sendMailHTML(subject string, templatePath string, toEmails []string) {
    // Get html
    var body bytes.Buffer
    t, err := template.ParseFiles(templatePath)
    t.Execute(&body, struct{ Name string }{Name: "Robby"})

    if err != nil {
        fmt.Println(err)
        return
    }

    auth := smtp.PlainAuth(
        "",
        "vinayakcsurya@gmail.com",
        "lblvwntpckhbbjds",
        "smtp.gmail.com",
    )

    Headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
    msg := "Subject: " + subject + "\n" + Headers + "\n\n" + body.String()

    err = smtp.SendMail(
        "smtp.gmail.com:587",
        auth,
        "vinayakcsurya@gmail.com",
        toEmails,
        []byte(msg),

    )
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }

}

func sendGoMail(subject string, templatePath string, toEmails []string) {
    // Get html
    var body bytes.Buffer
    t, err := template.ParseFiles(templatePath)
    if err != nil {
        fmt.Println(err)
        return
    }
    t.Execute(&body, struct{ Name string }{Name: "vinu"})

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
}

func CliPrompt() {

    //    asciiText1 := `
    //___________                      .___                    _________ .__       ___.
    //\_   _____/___  __ __  ____    __| _/___________  ______ \_   ___ \|  |  __ _\_ |__
    // |    __)/  _ \|  |  \/    \  / __ |/ __ \_  __ \/  ___/ /    \  \/|  | |  |  \ __ \
    // |     \(  <_> )  |  /   |  \/ /_/ \  ___/|  | \/\___ \  \     \___|  |_|  |  / \_\ \
    // \___  / \____/|____/|___|  /\____ |\___  >__|  /____  >  \______  /____/____/|___  /
    //     \/                   \/      \/    \/           \/          \/               \/
    //
    //`
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
    //fmt.Println(asciiText1)
    fmt.Println(asciiText2)
    fmt.Print("Welcome to mass mailer of Founders Club!")
    fmt.Println()

    fmt.Print("Enter Senders email address: ")
    //fmt.Scanf("%s", &senderEmail)

    fmt.Print("Path to csv file: ")
    fmt.Scanf("%s", &csvFilePath)
    //csvFilePath = "/Users/vinayak/IdeaProjects/GO-projects/mailer/emails.csv"

    fmt.Printf("Path to html template: ")
    fmt.Scanf("%s", &htmlTemplatePath)
    //htmlTemplatePath := "/Users/vinayak/IdeaProjects/GO-projects/mailer/template.html"

    fmt.Print("Subject: ")
    fmt.Scanf("%s", &subject)
    //subject = "this is a subject"

    fmt.Print("Confirm? [y/n]: ")
    fmt.Scanf("%s", &confirm)

    ReadCSV(csvFilePath)
    time.Sleep(1 * time.Second)
    sendGoMail(subject, htmlTemplatePath, emails)
}

func ReadCSV(csvFilePath string) {
    // Open the CSV file
    file, err := os.Open(csvFilePath)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    // Create a new CSV reader
    reader := csv.NewReader(file)

    // Read all the rows from the CSV
    records, err := reader.ReadAll()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    // Iterate through each record (row)
    for i, record := range records {
        fmt.Printf("Record %d: %v\n", i+1, record)

        // Assuming the emails in each row are separated by commas
        // Loop through each column in the record
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
