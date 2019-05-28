package main

import (
  "os"
  "bytes"
  "fmt"
  "strings"
  "flag"
  "log"
  "os/exec"
)

func main() {
  act := flag.String("cmd", "", "command name")
  opt := flag.String("opt", "", "option")
  fmt.Println(*act, *opt)
  if *act == "get" && strings.Contains(*opt, "github.com") {
    parts := strings.Split(*opt, "/")
    dirs := strings.Join(parts[:len(parts)-1], "/")
    fmt.Println("Creating ", dirs)
    os.MkdirAll(dirs, 775)
    cmd := exec.Command("git clone https://www."+*opt)
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Start()
    if err != nil {
      log.Fatal(err)
    }
    log.Println(cmd.Args)
    err = cmd.Wait()
    if err != nil {
      log.Printf("Command finished with error: %v", err)
    }
    fmt.Println(out.String())
  }
}
