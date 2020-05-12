package config

import  {
  "fmt"
}

struct Config  {
  Account string
  Key string
  Nodes struct {
    Hostname string
    Account string
  }
}
