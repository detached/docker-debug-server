package main

import (
  "fmt"
  "strings"
  "net/http"
  "net/http/httputil"
)

var(
  m = "<html><head></head><body>%s</body></html>"
  re = strings.NewReplacer("\n", "<br>")
)

func handler(w http.ResponseWriter, r *http.Request) {
  d, e := httputil.DumpRequest(r, true)

  var msg string
  if e != nil {
    msg = fmt.Sprintf(m, e)
    fmt.Println(e)
  } else {
    s := string(d)
    msg = fmt.Sprintf(m, re.Replace(s))
    fmt.Println(s)
  }

  fmt.Fprint(w, msg)
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
