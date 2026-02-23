# Console

A simple and lightweight Go library for creating console user interfaces effectively. It provides tools for prompting user input and rendering structured data as tables.

## Installation

```bash
go get github.com/dtylman/console
```

## Usage

### Prompts

`console` provides several functions to interact with the user via the command line.

#### Simple Prompt

Ask the user for input with a simple prompt.

```go
package main

import (
	"fmt"
	"log"

	"github.com/dtylman/console"
)

func main() {
	name, err := console.Ask("What is your name?")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Hello, %s!\n", name)
}
```

#### Password Prompt

Ask for sensitive information like passwords. The input will be masked.

```go
password, err := console.AskPassword("Enter your password")
if err != nil {
	log.Fatal(err)
}
fmt.Println("Password received.")
```

#### Prompt with Options

Ask with more control, such as setting a default value, making the input required, or masking it.

```go
// Prompt: "Enter your mood (happy): "
// Default: "happy"
// Required: true
// Mask: false
mood, err := console.AskOptions("Enter your mood", "happy", true, false)
```

#### String Array Prompt

Ask for a comma-separated list of strings.

```go
// Prompt: "Enter tags (tag1,tag2): "
tags, err := console.AskStringArray("Enter tags", []string{"tag1", "tag2"}, true)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Tags: %v\n", tags)
```

### Tables

Render structured data (structs, maps, or slices) as formatted ASCII tables.

#### Table from Struct or Map

```go
type User struct {
	ID    int
	Name  string
	Email string
}

user := User{ID: 1, Name: "Alice", Email: "alice@example.com"}

// Output:
// +-------+-------------------+
// | FIELD |       VALUE       |
// +-------+-------------------+
// | ID    | 1                 |
// | Name  | Alice             |
// | Email | alice@example.com |
// +-------+-------------------+
console.WriteTable(user)
```

#### Table from Slice of Structs

```go
users := []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com"},
	{ID: 2, Name: "Bob", Email: "bob@example.com"},
}

// Output:
// +----+-------+-------------------+
// | ID | NAME  |       EMAIL       |
// +----+-------+-------------------+
// |  1 | Alice | alice@example.com |
// |  2 | Bob   | bob@example.com   |
// +----+-------+-------------------+
console.WriteTable(users)
```

You can also specify headers explicitly:

```go
console.WriteTable(users, "User ID", "Full Name", "Contact Email")
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
