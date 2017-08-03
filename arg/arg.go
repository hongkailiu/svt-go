package arg

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/hongkailiu/svt-go/log"
	"github.com/hongkailiu/svt-go/cluster_loader"
	"github.com/hongkailiu/svt-go/http"
	"os"
	"strings"
	"fmt"
)

const AppVersion = "0.0.1-SNAPSHOT"

var (
	httpCommand = kingpin.Command("http", "Start http server.")
	clusterLoaderCommand = kingpin.Command("clusterLoader", "Run cluster loader.")
	versionCommand = kingpin.Command("version", "Show version info.")
	clFile = clusterLoaderCommand.Flag("file", "Config file.").Default("conf/pyconfig.yaml").Short('f').String()
	poolSizeP = clusterLoaderCommand.Flag("pool", "Go routine pool size.").Default("10").Short('p').Int()
)

func init() {

}

func ParseAndRun() {
	kingpin.Version(AppVersion)
	switch kingpin.Parse() {
	case "http":
		log.Debug("aaa")
		http.Server{8080}.Run()
	case "clusterLoader":
		log.Debug("bbb")
		log.Debug("CLFile: " + *clFile)

		configFileString := *clFile
		poolSize := *poolSizeP

		if !strings.HasPrefix(configFileString, string(os.PathSeparator)) &&
			!strings.HasPrefix(configFileString, "~") {
			dir, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			configFileString = dir + string(os.PathSeparator) + *clFile
		}


		log.Debug("ConfigFileString: " + configFileString)
		log.Debug(fmt.Sprintf("Processes: %d", poolSize))

		args := cluster_loader.Args{PoolSize:uint(poolSize), ConfigFile:configFileString}

		log.Debug(args)
		err := cluster_loader.GetClusterLoader().Run(args)

		if err != nil {
			log.Fatal(err)
		}


	}

}
