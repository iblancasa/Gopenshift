package main

import (
	"fmt"
	"os"
	"path/filepath"

	is_openshift "github.com/iblancasa/gopenshift/pkg/is-openshift"
	configV1 "github.com/openshift/api/config/v1"
	"github.com/spf13/pflag"
	k8sruntime "k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var scheme *k8sruntime.Scheme

func init() {
	scheme = k8sruntime.NewScheme()
	utilruntime.Must(configV1.AddToScheme(scheme))
}

func main() {
	var kubeconfigPath string

	defaultKubeconfigPath := filepath.Join(homedir.HomeDir(), ".kube", "config")

	pflag.StringVar(&kubeconfigPath, "kubeconfig-path", defaultKubeconfigPath, "Absolute path to the KubeconfigPath file")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		fmt.Printf("Error reading the kubeconfig: %s\n", err.Error())
		os.Exit(1)
	}

	isOpenShift, err := is_openshift.IsOpenShift(config)
	if err != nil {
		fmt.Printf("Error checking if running against an OpenShif cluster: %s\n", err.Error())
		os.Exit(1)
	}

	var returnCode int
	if isOpenShift {
		returnCode = 0
	} else {
		returnCode = 1
	}

	os.Exit(returnCode)

}
