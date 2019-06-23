package main
import (
    "fmt"
    "net/http"
    //"html/template"
    "os"
)
func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello. you are in this path: %s ", r.URL.Path)
}

func main() {
    http.HandleFunc("/", hello)

    if err := http.ListenAndServe(":1234", nil); err != nil {
        fmt.Println(err)
        os.Exit(0)
    }
    fmt.Println("server runinig on :1234")


}
