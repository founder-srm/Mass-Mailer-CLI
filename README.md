# Go Mass Mailer

## Overview

This project is a **Mass Mailer** written in `GoLang`, designed to send bulk emails using **Gmail's SMTP** service. It reads email addresses from a CSV file, uses an HTML template for email content, and sends the emails with or without attachments using the `gomail` package.

## Features

- **Mass Emailing**: Sends bulk emails to a list of recipients from a CSV file.
- **HTML Template Support**: Sends emails using an HTML template.
- **Gmail SMTP Integration**: Uses Gmail’s SMTP server to send emails.
- **CSV Parsing**: Reads and extracts emails from a CSV file.
- **Attachments**: Supports attaching files to emails.
- **Command Line Interface (CLI)**: Prompts for input like sender email, CSV file path, and email subject via CLI.

