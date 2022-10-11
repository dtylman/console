package console

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/term"
)

//Ask asks the user something
func Ask(prompt string) (string, error) {
	return AskOptions(prompt, "", true, false)
}

//AskPassword asks a password
func AskPassword(prompt string) (string, error) {
	return AskOptions(prompt, "", true, true)
}

//AskOptions ...
func AskOptions(prompt string, def string, required bool, mask bool) (string, error) {
	for {
		fmt.Print(prompt)
		if def != "" {
			fmt.Printf(" (%v)", def)
		}
		fmt.Printf(": ")
		answer, err := readline(mask)
		if err != nil {
			return answer, err
		}
		answer = trimSuffix(answer)
		if answer == "" {
			answer = def
		}
		if !required || answer != "" {
			return answer, nil
		}

	}
}

//AskStringArray convert string array to string and asks
func AskStringArray(prompt string, def []string, required bool) ([]string, error) {
	value, err := AskOptions(prompt, strings.Join(def, ","), required, false)
	if err != nil {
		return nil, err
	}
	return strings.Split(value, ","), nil
}

func trimSuffix(s string) string {
	var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}
	stop := len(s)
	for ; stop > 0; stop-- {
		c := s[stop-1]
		if c >= utf8.RuneSelf {
			return strings.TrimFunc(s[0:stop], unicode.IsSpace)
		}
		if asciiSpace[c] == 0 {
			break
		}
	}
	return s[:stop]
}

func readline(mask bool) (string, error) {
	intChan := make(chan os.Signal, 1)
	signal.Notify(intChan, os.Interrupt)

	defer func() {
		signal.Stop(intChan)
	}()

	var answer string
	var err error

	go func() {
		defer close(intChan)
		if mask {
			data, err1 := term.ReadPassword(int(os.Stdin.Fd()))
			fmt.Println()
			answer = string(data)
			err = err1
		} else {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			answer = scanner.Text()
			err = scanner.Err()
		}
	}()
	sig := <-intChan
	if sig != nil {
		return "", fmt.Errorf("got signal %v", sig)
	}
	return answer, err
}

//AskSite asks for a site
func AskSite(def string) (string, error) {
	if def != "" {
		return def, nil
	}
	site, err := Ask("site")
	if err != nil {
		return "", err
	}
	if site == "" {
		return "", errors.New("site cannot be empty")
	}
	return site, nil
}

//AskShield asks for shield id and allows empty
func AskShield(def string) (string, error) {
	if def != "" {
		return def, nil
	}
	fmt.Println("leave shield id empty for all shields")
	return AskOptions("shield id", def, false, false)
}

//AskString if current is empty, asks for prompt, stores the answer in target
func AskString(prompt string, target *string, current string) error {
	if current != "" {
		*target = current
		return nil
	}
	answer, err := AskOptions(prompt, current, true, false)
	if err != nil {
		return err
	}
	*target = answer
	return nil
}

//AskBool asks for a boolean value
func AskBool(prompt string, target *bool, current bool) error {
	hint := "(y/N)"
	if current {
		hint = "(Y/n)"
	}
	for {
		answer, err := AskOptions(prompt+" "+hint, "", false, false)
		if err != nil {
			return err
		}
		answer = strings.ToLower(strings.TrimSpace(answer))
		if answer == "" {
			if current {
				fmt.Println("Yes.")
			} else {
				fmt.Println("No.")
			}
			target = &current
			return nil
		}
		if answer == "y" || answer == "yes" {
			*target = true
			return nil
		} else if answer == "n" || answer == "no" {
			*target = false
			return nil
		} else {
			fmt.Println("Please answer y or n")
		}
	}
}
