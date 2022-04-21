package main

import ( 
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"etri5gc/nfs/amf"
	"github.com/urfave/cli"
)

var flags = []cli.Flag{
	cli.StringFlag{
		Name:  "config, c",
		Usage: "Load configuration from `FILE`",
	},
	cli.StringFlag{
		Name:  "log, l",
		Usage: "Output logs to `FILE`",
	},
}

func main() {
	fmt.Println("Hello world")

	app := cli.NewApp()
	app.Name = "amf"
	app.Usage = "5G Access and Mobility Management Function (AMF)"
	app.Action = action
	app.Flags = flags


	if err := app.Run(os.Args); err != nil {
		//log
		return
	}
}


func action(c *cli.Context) error {
	nf := amf.NewAMF()
	if err := nf.Initialize(c); err != nil {
		return err
	}
	if err := nf.Start(); err != ni {
		return err
	}
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<- sigch
		nf.Terminate()
		os.Exit(0)
	}()

	return nil
}
