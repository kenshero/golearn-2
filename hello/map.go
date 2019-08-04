package main
import "fmt"

func main() {
    x := make(map[string] string)
    x["TH"] = "THAILAND"
    x["JP"] = "JAPAN"
    x["EN"] = "ENGLAND"
    fmt.Println(x)

    y := map[string] string {
        "TH": "THAILAND",
        "JP": "THAILAND",
    }
    fmt.Println(y)

}
