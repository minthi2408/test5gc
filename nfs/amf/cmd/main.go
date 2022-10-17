package main

import (
	"etrib5gc/nfs/amf/config"
	"etrib5gc/nfs/amf/service"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
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

var nf *service.AMF

func main() {
	log.Println("Hello beautiful world")

	app := cli.NewApp()
	app.Name = "amf"
	app.Usage = "5G Access and Mobility Management Function (AMF)"
	app.Action = action
	app.Flags = flags

	if err := app.Run(os.Args); err != nil {
		//log
		log.Fatal("Fail to start application", err)
	} else {
		quit := make(chan struct{})
		sigch := make(chan os.Signal, 1)
		signal.Notify(sigch, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-sigch
			if nf != nil {
				nf.Terminate()
			}
			log.Info("Received a kill signal")
			quit <- struct{}{}
		}()
		<-quit
		log.Info("Good bye the world")
	}
}

func action(c *cli.Context) (err error) {
	log.SetLevel(log.InfoLevel)
	//read config
	var cfg *config.AmfConfig
	filename := c.String("config")
	if cfg, err = config.LoadConfig(filename); err != nil {
		log.Errorf("Fail to parse AMF configuration", err)
		return
	}
	log.Info(cfg.Agent.NfType)
	//create the AMF
	if nf, err = service.New(cfg); err != nil {
		log.Errorf("Fail to create AMF", err)
		return
	}

	if err = nf.Start(); err != nil {
		log.Errorf("Fail to start AMF", err)
		return err
	}

	return nil
}
