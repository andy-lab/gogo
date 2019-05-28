package main

import (
  "os"
  "os/user"
  "bytes"
  "fmt"
  "strings"
  "flag"
  "log"
  "os/exec"
  "path/filepath"
)

func main() {
  act := flag.String("cmd", "", "command name")
  opt := flag.String("opt", "", "option")
  flag.Parse()
  usr, err := user.Current()
  if err != nil {
    fmt.Println("Get current user failed:", err)
  }
  gopath := filepath.Join(usr.HomeDir, "go/src/")
  if *act == "get" && strings.Contains(*opt, "github.com") {
    parts := strings.Split(*opt, "/")
    dirs := strings.Join(parts[:len(parts)-1], "/")
    fmt.Println("Creating ", gopath+dirs)
    os.MkdirAll(dirs, 775)
    cmd := exec.Command("git", "clone", "https://www."+*opt, gopath+dirs)
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
