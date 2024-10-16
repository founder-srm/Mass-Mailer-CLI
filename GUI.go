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
)

// ANSI escape codes for coloring text
const (
    Red      = "\033[31m"
    Green    = "\033[32m"
    Yellow   = "\033[33m"
    Blue     = "\033[34m"
    Magenta  = "\033[35m"
    Cyan     = "\033[36m"
    White    = "\033[37m"
    Bold     = "\033[1m"
    Underline = "\033[4m"
    Reset    = "\033[0m"
)

// Command map to handle different commands
var commands = map[string]func(args []string){
    "echo":     echoCommand,
    "exit":     exitCommand,
    "sendmail": sendMailCommand,
}

func echoCommand(args []string) {
    if len(args) > 0 {
        fmt.Println(strings.Join(args, " "))
    } else {
        fmt.Println("")
    }
}

func exitCommand(args []string) {
    os.Exit(0)
}


func sendMailCommand(args []string) {
    if len(args) < 3 {
        printError("Usage: sendmail <csv_file_path> <html_template_path> <subject>", Yellow)
        return
    }
    csvFilePath := args[0]
    htmlTemplatePath := args[1]
    subject := strings.Join(args[2:], " ")

    emails, err := ReadCSV(csvFilePath)
    if err != nil {
        printError(fmt.Sprintf("Error reading CSV: %v", err), Red)
        return
    }

    result := sendGoMail(subject, htmlTemplatePath, emails)
    fmt.Println(result)
}

func ReadCSV(csvFilePath string) ([]string, error) {
    var emails []string
    file, err := os.Open(csvFilePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return nil, err
    }

    for _, record := range records {
        for _, field := range record {
            splitEmails := strings.Split(field, ",")
            emails = append(emails, splitEmails...)
        }
    }
    return emails, nil
}

func sendGoMail(subject string, templatePath string, toEmails []string) string {
    var body bytes.Buffer
    t, err := template.ParseFiles(templatePath)
    if err != nil {
        return fmt.Sprintf("Error parsing template: %v", err)
    }
    t.Execute(&body, nil)

    m := gomail.NewMessage()
    m.SetHeader("From", "your-email@gmail.com")
    m.SetHeader("To", toEmails...)
    m.SetHeader("Subject", subject)
    m.SetBody("text/html", body.String())

    d := gomail.NewDialer("smtp.gmail.com", 587, "your-email@gmail.com", "your-app-password")

    if err := d.DialAndSend(m); err != nil {
        return fmt.Sprintf("Error sending email: %v", err)
    }
    return "Mail sent successfully!"
}

// Function to print an error message with a specified color
func printError(message string, color string) {
    fmt.Printf("%s%s%s\n", color, message, Reset)
}


func main() {
    fmt.Println(Cyan, `
 /$$$$$$$$                                  /$$                                      /$$$$$$  /$$           /$$      
| $$_____/                                 | $$                                     /$$__  $$| $$          | $$      
| $$     /$$$$$$  /$$   /$$ /$$$$$$$   /$$$$$$$  /$$$$$$   /$$$$$$   /$$$$$$$      | $$  \__/| $$ /$$   /$$| $$$$$$$ 
| $$$$$ /$$__  $$| $$  | $$| $$__  $$ /$$__  $$ /$$__  $$ /$$__  $$ /$$_____/      | $$      | $$| $$  | $$| $$__  $$
| $$__/| $$  \ $$| $$  | $$| $$  \ $$| $$  | $$| $$$$$$$$| $$  \__/|  $$$$$$       | $$      | $$| $$  | $$| $$  \ $$
| $$   | $$  | $$| $$  | $$| $$  | $$| $$  | $$| $$_____/| $$       \____  $$      | $$    $$| $$| $$  | $$| $$  | $$
| $$   |  $$$$$$/|  $$$$$$/| $$  | $$|  $$$$$$$|  $$$$$$$| $$       /$$$$$$$/      |  $$$$$$/| $$|  $$$$$$/| $$$$$$$/
|__/    \______/  \______/ |__/  |__/ \_______/ \_______/|__/      |_______/        \______/ |__/ \______/ |_______/ 
`,)
    fmt.Println(Green, "Welcome to mass mailer of Founders Club!")

    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print(Reset,"$~ ")

        input, err := reader.ReadString('\n')
        if err != nil {
            printError(fmt.Sprintf("Error reading input: %v", err), Red)
            continue
        }

        input = strings.TrimSpace(input)
        if input == "" {
            continue
        }

        parts := strings.Fields(input)
        command := parts[0]
        args := parts[1:]

        if cmdFunc, exists := commands[command]; exists {
            cmdFunc(args)
        } else {
            printError(fmt.Sprintf("%s: command not found", command), Red)
        }
    }
}
