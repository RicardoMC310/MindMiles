package models

type ContactModel struct {
    ID      int    `json:"id"`
    Name    string `json:"name"`
    Email   string `json:"email"`
    Content string `json:"content"`
}
