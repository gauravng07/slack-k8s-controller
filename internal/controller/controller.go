package controller

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type Controller struct {
	clientSet 			kubernetes.Interface
	sampleClientSet		kubernetes.Interface
	queue 				workqueue.RateLimitingInterface
	informer 			cache.SharedIndexInformer
}

func NewController(clientSet kubernetes.Interface,  sampleClientSet kubernetes.Interface) *Controller {
	return &Controller{
		clientSet: clientSet,
		queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "foos"),
		sampleClientSet: sampleClientSet,
	}
}