package controller

import (
	coreinformers "k8s.io/client-go/informers/core/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type Controller struct {
	name string

	queue      workqueue.RateLimitingInterface
	handleFunc func(obj interface{}) error
	enqueueObj func(obj interface{})

	podIndexer         cache.Indexer
	deploymentIndexer  cache.Indexer
	statefulsetIndexer cache.Indexer
	// A store of pods, populated by the shared informer passed to NewReplicaSetController
	podLister corelisters.PodLister
	// podListerSynced returns true if the pod store has been synced at least once.
	// Added as a member to the struct to allow injection for testing.
	podListerSynced         cache.InformerSynced
	deploymentListerSynced  cache.InformerSynced
	statefulsetListerSynced cache.InformerSynced
}

func NewController(podInformer *coreinformers.PodInformer) *Controller {

	return &Controller{
		name: "controller-demo",
	}

}
