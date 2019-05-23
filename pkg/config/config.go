package config

import (
	"time"

	"github.com/appscode/go/runtime"
	"kmodules.xyz/client-go/meta"
)

var (
	BuiltinTemplates = "/srv/voyager/templates/*.cfg"
)

func init() {
	if meta.PossiblyInCluster() {
		BuiltinTemplates = "/srv/voyager/templates/*.cfg"
	} else {
		BuiltinTemplates = runtime.GOPath() + "/src/github.com/appscode/voyager/hack/docker/voyager/templates/*.cfg"
	}
}

type Config struct {
	Burst                       int
	CloudConfigFile             string
	CloudProvider               string
	HAProxyImage                string
	ExporterImage               string
	IngressClass                string
	MaxNumRequeues              int
	NumThreads                  int
	OperatorNamespace           string
	OperatorService             string
	QPS                         float32
	RestrictToOperatorNamespace bool
	ResyncPeriod                time.Duration
	WatchNamespace              string
	ValidateHAProxyConfig       bool
	EnableValidatingWebhook     bool
}
