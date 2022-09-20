package main

import (
	"etri5gc/nfs/udm/config"
	"etri5gc/nfs/udm/service"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
	"os/signal"
	"syscall"
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
	app.Name = "udm"
	app.Usage = "Etri 5G UDM"
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
	var cfg *config.UdmConfig
	filename := c.String("config")
	if cfg, err = config.LoadConfig(filename); err != nil {
		log.Errorf("Fail to parse UDM configuration", err)
		return
	}

	//create the AMF
	if nf, err = service.New(cfg); err != nil {
		log.Errorf("Fail to create UDM", err)
		return
	}

	if err = nf.Start(); err != nil {
		log.Errorf("Fail to start UDM", err)
		return err
	}

	return nil
}
