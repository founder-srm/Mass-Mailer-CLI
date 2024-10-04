package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "strings"
)

func main() {
    readCSV("emails.csv")
}

func readCSV(csvFile string) {
    // Open the CSV file
    file, err := os.Open(csvFile)
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

    // Declare the slice of strings for emails
    var emails []string

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
