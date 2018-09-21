package main

import (
    "fmt" 
    "strings"
    "time"
    "os"
    "os/exec"
)

func main() {
    numbers := getNumbers()
    for {
        printTime(time.Now(), numbers)
        time.Sleep(1 * time.Second)
    }
}

func getNumbers() [][]string {
    rawNumbers := `
███  █  ███ ███ █ █ ███ ███ ███ ███ ███    
█ █ ██    █   █ █ █ █   █     █ █ █ █ █ █  
█ █  █  ███  ██ ███ ███ ███   █ ███ ███    
█ █  █  █     █   █   █ █ █   █ █ █   █ █  
███ ███ ███ ███   █ ███ ███   █ ███ ███    
`
    rawNumbers = strings.Trim(rawNumbers, "\n")
    rawNumbersLines := strings.Split(rawNumbers, "\n")

    numbers := [][]string{}
        
    for i := 0; i <= 10; i++ {
        n := extractNumberAtIndex(i, rawNumbersLines)
        numbers = append(numbers, n)
    }

    return numbers 
}

func extractNumberAtIndex(i int, rawNumbersLines []string) []string {
    number := []string {}

    for _, line := range rawNumbersLines {
        runes := []rune(line)
        substr := string(runes[i*4:(i+1)*4])
        number = append(number, substr)
    }

    return number
}

func addNumberToOutput(n int, output []string, numbers [][]string) []string {
    for i, line := range numbers[n] {
        output[i] += line
    }

    return output
}

func printTime(t time.Time, numbers [][]string) {
    const red = "\033[91m"
    const blue = "\033[94m"

    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()

    hours := t.Hour()
    minutes := t.Minute()
    seconds := t.Second()

    output := []string {"", "", "", "", ""}

    // hours
    if hours < 10 {
        output = addNumberToOutput(0, output, numbers)
    } else {
        tens := hours / 10
        output = addNumberToOutput(tens, output, numbers)
    }
    output = addNumberToOutput(hours % 10, output, numbers)

    // :
    output = addNumberToOutput(10, output, numbers)

    // minutes
    if minutes < 10 {
        output = addNumberToOutput(0, output, numbers)
    } else {
        tens := minutes / 10
        output = addNumberToOutput(tens, output, numbers)
    }
    output = addNumberToOutput(minutes % 10, output, numbers)

    // :
    output = addNumberToOutput(10, output, numbers)

    // seconds
    if seconds < 10 {
        output = addNumberToOutput(0, output, numbers)
    } else {
        tens := seconds / 10
        output = addNumberToOutput(tens, output, numbers)
    }
    output = addNumberToOutput(seconds % 10, output, numbers)

    for _, l := range output {
        fmt.Println(blue + l)
    }
}

