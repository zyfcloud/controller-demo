package main

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

var defaultResync = time.Duration(time.Second * 60)

func main() {

	stopCh := make(chan struct{})
	defer close(stopCh)

	masterUrl := "172.27.32.110:8080"
	config, err := clientcmd.BuildConfigFromFlags(masterUrl, "")
	if err != nil {
		klog.Errorf("BuildConfigFromFlags err, err: %v", err)
	}
	clientset, err := kubernetes.NewForConfig(config) // 问题：哪里用到了clientset和reflector？？
	if err != nil {
		klog.Errorf("Get clientset err, err: %v", err)
	}
	sharedInformers := informers.NewSharedInformerFactory(clientset, defaultResync)

	// lister的使用
	// sharedInformers.Core().V1().Pods().Informer() // 注册了一个informer,f.informers[informerType] = informer
	// sharedInformers.Start(stopCh)                 // start all informers
	// time.Sleep(3 * time.Second)
	// podLister := sharedInformers.Core().V1().Pods().Lister()
	// pods, err := podLister.List(labels.Everything())
	// if err != nil {
	// 	klog.Infof("err: %v", err)
	// }
	// klog.Infof("len(pods), %d", len(pods))
	// for _, v := range pods {
	// 	klog.Infof("pod: %s", v.Name)
	// }

	podInformer := sharedInformers.Core().V1().Pods().Informer()
	podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			mObj := obj.(v1.Object) //为什么不直接用corev1.Pod？？
			klog.Infof("Get new obj: %v", mObj)
			klog.Infof("Get new obj name: %s", mObj.GetName())
		},
	})
	podInformer.Run(stopCh)
}
