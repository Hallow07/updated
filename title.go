package title

import (
    "fmt"
    "github.com/Hallow07/updated" // Replace "your-username" with your actual GitHub username or the path to where you store the title package
)

func main() {
    // Example of setting the terminal title
    _, err := title.SetTitle("New Title")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Title set successfully.")
}
