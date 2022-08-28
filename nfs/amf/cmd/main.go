package main

import ( 
	"os"
	"os/signal"
	"syscall"
	"etri5gc/nfs/amf/service"
	"etri5gc/nfs/amf/config"
	"github.com/urfave/cli"
	log "github.com/sirupsen/logrus"
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
			<- sigch
			if nf != nil {
				nf.Terminate()
			}
			log.Info("Received a kill signal")
			quit <- struct{}{}
		}()
		<- quit
		log.Info("Good bye the world")
	}
}


func action(c *cli.Context) (err error) {
	log.SetLevel(log.InfoLevel)
	//read config
	var cfg *config.Config
	filename := c.String("config")
	if cfg, err = config.LoadConfig(filename); err != nil {
		log.Errorf("Fail to parse AMF configuration", err)
		return
	}

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
