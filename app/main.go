package main

import (
	// "fmt"
	"github.com/joho/godotenv"
  "github.com/binhnt-teko/test_blockchain/app/cmd"
  "github.com/urfave/cli/v2"
	"os"
	"log"
	"path/filepath"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	 dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	 if err != nil {
			 log.Fatal(err)
	 }
	 environmentPath := filepath.Join(dir, ".env")
	 err = godotenv.Load(environmentPath)
	 godotenv.Load()
}

func main() {
		app := &cli.App{
		    Commands: []*cli.Command{
		      {
		        Name:    "server",
		        Aliases: []string{"s"},
		        Usage:   "server",
		        Action:  cmd.MigrateFlags(cmd.NodeServer),
		      },
				},
		  }

	  err := app.Run(os.Args)
	  if err != nil {
	    log.Fatal(err)
	  }
}
