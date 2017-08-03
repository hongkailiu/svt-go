package cluster_loader

type ClustreLoader interface {
	run() error
}

type Args struct {
	PoolSize int
	ConfigFile string
}

type MyClusterLoader struct {

}