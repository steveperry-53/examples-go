package main

import "fmt"

type Command struct {
  Use string
  Short string
  Long string
  Example string
  Run func(code string)
  Otherfield1 string
  Otherfield2 string
  Etc string
}

type Rectangle struct {
  width int
  length int
}

func main() {
  fmt.Println("Hello structs")

  r := Rectangle{3, 5}

  var s = Rectangle{7, 10}

  var t = Rectangle{width: 12, length: 15}

  var u = Rectangle{width: 20}

  var v = &Rectangle{length: 30}

  w := &Rectangle{width: 40}

  // Similar to https://github.com/kubernetes/kubernetes/blob/master/pkg/kubectl/cmd/create_secret.go

  cmd := &Command{
    Use: "docker-registry NAME --docker-username=user --docker-password=password --docker-email=email [--docker-server=string] [--from-literal=key1=value1] [--dry-run]",
    Short: "Create a secret for use with a Docker registry.",
    Long: "Long description about creating a secret for use with a Docker registry.",
    Example: "Code example of creating a secret for use with a Docker registry",
    Run: func(code string) {
           fmt.Println("That was error ", code)
         },
  }

  fmt.Println(r.width, r.length)
  fmt.Println(s.width, s.length)
  fmt.Println(t.width, t.length)
  fmt.Println(u.width, u.length)
  fmt.Println( (*v).width, (*v).length)
  fmt.Println(v.width, v.length)
  fmt.Println(w.width, w.length)

  fmt.Println(cmd.Short)
  fmt.Println( (*cmd).Long )
  cmd.Run("BadThing")
  (*cmd).Run("SillyThing")
  fmt.Println(cmd.Use)
}
