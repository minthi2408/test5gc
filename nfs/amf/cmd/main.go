package main

import ( 
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"etri5gc/nfs/amf/service"
	"etri5gc/nfs/amf/config"
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
	}
}


func action(c *cli.Context) err error {

	//read config
	var *cfg config.Config
	if cfg, err = config.LoadConfig("dumpconfig.yaml"); err != nill {
		return 
	}

	//create the AMF
	var nf *service.AMF
	if nf, err = service.CreateAMF(cfg); err != nil {
		return
	}


	if err = nf.Start(); err != ni {
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
