package title

import (
    "fmt"
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
