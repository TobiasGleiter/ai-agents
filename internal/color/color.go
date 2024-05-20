package color

import "fmt"


const (
    Reset  = "\033[0m"
    Red    = "\033[31m"
    Green  = "\033[32m"
    Yellow = "\033[33m"
    Blue   = "\033[34m"
    Purple = "\033[35m"
    Cyan   = "\033[36m"
    Gray   = "\033[37m"
    White  = "\033[97m"
)

type ChatColor string


func WrapColor(color ChatColor, text string) string {
    return string(color) + text + Reset
}

func PrintColor(color ChatColor, text string) {
    fmt.Println(WrapColor(color, text))
}
