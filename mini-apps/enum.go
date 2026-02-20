package main


import "fmt"

type Status uint

const (
  PENDING Status = iota
  ACCEPTED
  REFUSED
)

func (s Status) String() string {
  return [...]string{"Pending", "Accepted", "Refused"}[s]
}


func main() {

  fmt.Println(REFUSED)

}
