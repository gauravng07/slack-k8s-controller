package main

import (
	"context"
	"github.com/spf13/viper"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"slack-k8s-controller/internal"
	"slack-k8s-controller/internal/config"
	controller2 "slack-k8s-controller/internal/controller"
	"slack-k8s-controller/internal/logger"
	"time"
)

var ctx context.Context
const defaultCorrelationId = "0.00.00.0000"

func init() {
	ctx = internal.SetContextWithValue(ctx, internal.ContextKeyCorrelationID, defaultCorrelationId)
}


func main()  {
	env := os.Getenv("env")
	if len(env) == 0 {
		env = "dev"
	}

	err := config.ReadConfig(env)
	if err != nil {
		logger.Fatalf(ctx, "error loading environment config: %v", err)
	}

	cfg, err := clientcmd.BuildConfigFromFlags(viper.GetString(config.K8sURL), viper.GetString(config.KubeConfigPath))
	if err != nil {
		logger.Fatalf(ctx,"error building kube config: %v", err)
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		logger.Fatalf(ctx, "error building k8s client set: %v", err)
	}

	slackClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		logger.Fatalf(ctx, "error building client set: %v", err)
	}

	kubeInformerFactory := informers.NewSharedInformerFactory(kubeClient, 30*time.Second)
	slackInformerFactory := informers.NewSharedInformerFactory(slackClient, 30*time.Second)

	controller := controller2.NewController(kubeClient, slackClient)

}
