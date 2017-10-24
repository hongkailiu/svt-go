package arg

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/hongkailiu/svt-go/log"
	"github.com/hongkailiu/svt-go/cluster_loader"
	"github.com/hongkailiu/svt-go/http"
	"os"
	"strings"
	"fmt"
	"github.com/hongkailiu/svt-go/version"
	"github.com/hongkailiu/svt-go/inv_gen"
)

var (
	httpCommand = kingpin.Command("http", "Start http server.")
	clusterLoaderCommand = kingpin.Command("clusterLoader", "Run cluster loader.")
	clFile = clusterLoaderCommand.Flag("file", "Config file.").Default("conf/pyconfig.yaml").Short('f').String()
	poolSizeP = clusterLoaderCommand.Flag("pool", "Go routine pool size.").Default("10").Short('p').Int()

	invGenCommand = kingpin.Command("invGen", "Generate inventory file for openshift-ansible.")
	igFile = invGenCommand.Flag("file", "Config file.").Default("conf/invGenConfig.yaml").Short('f').String()
)

func init() {

}

func ParseAndRun() {
	kingpin.Version(version.GetVersion())
	switch kingpin.Parse() {
	case "http":
		log.Debug("aaa")
		http.Server{Port:8080}.Run()
	case "clusterLoader":
		log.Debug("bbb")
		log.Debug("CLFile: " + *clFile)

		configFileString := *clFile
		poolSize := *poolSizeP

		configFileString, err := handleRelativePath(configFileString)
		if err != nil {
			log.Fatal(err)
		}

		log.Debug("ConfigFileString: " + configFileString)
		log.Debug(fmt.Sprintf("Processes: %d", poolSize))

		args := cluster_loader.Args{PoolSize:uint(poolSize), ConfigFile:configFileString}

		log.Debug(args)
		err = cluster_loader.GetClusterLoader().Run(args)

		if err != nil {
			log.Fatal(err)
		}

	case "invGen":
		log.Debug("ccc")
		log.Debug("igFile: " + *igFile)
		configFileString := *igFile

		configFileString, err := handleRelativePath(configFileString)
		if err != nil {
			log.Fatal(err)
		}

		args := inv_gen.Args{ConfigFile:configFileString}

		log.Debug(args)
		err = inv_gen.GetInventoryGenerator().Run(args)

		if err != nil {
			log.Fatal(err)
		}
	}

}


func handleRelativePath(configFileString string) (string, error) {
	if !strings.HasPrefix(configFileString, string(os.PathSeparator)) &&
		!strings.HasPrefix(configFileString, "~") {
		dir, err := os.Getwd()
		if err != nil {
			return "", err
		}
		return dir + string(os.PathSeparator) + configFileString, nil
	}
	return configFileString, nil
}
