# ansi-formatter

## Example

```go
func main() {
    fmt.Println(ansi.Red("This is red text"))
    fmt.Println(ansi.FormatTextWithStyles("Bold and Underlined", ColorRed, Bold, Underline))
    fmt.Println(ansi.CustomFormat("\033[38;5;82m", "This is green text with 256 color"))
}
```