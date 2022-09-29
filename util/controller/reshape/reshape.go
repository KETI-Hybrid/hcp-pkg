package reshape

import (
	"context"
	"fmt"

	"hcp-pkg/util/clusterManager"

	"admiralty.io/multicluster-controller/pkg/cluster"
	"admiralty.io/multicluster-controller/pkg/controller"
	"admiralty.io/multicluster-controller/pkg/reconcile"
	"github.com/jinzhu/copier"
	"k8s.io/klog"

	"os"

	"sigs.k8s.io/controller-runtime/pkg/client"
	fedapis "sigs.k8s.io/kubefed/pkg/apis"
	"sigs.k8s.io/kubefed/pkg/apis/core/v1beta1"
)

var c chan string

var prev_length int = 0

var cm *clusterManager.ClusterManager

func NewController(live *cluster.Cluster, ghosts []*cluster.Cluster, ghostNamespace string, myClusterManager *clusterManager.ClusterManager) (*controller.Controller, error) {
	fmt.Println("Reshape New Controller")

	c = make(chan string)
	cm = myClusterManager

	liveclient, err := live.GetDelegatingClient()
	if err != nil {
		return nil, fmt.Errorf("getting delegating client for live cluster: %v", err)
	}
	ghostclients := []client.Client{}

	/*
		for _, ghost := range ghosts {
			ghostclient, err := ghost.GetDelegatingClient()
			if err != nil {
				return nil, fmt.Errorf("getting delegating client for ghost cluster: %v", err)
			}
			ghostclients = append(ghostclients, ghostclient)
		}
	*/

	co := controller.New(&reconciler{live: liveclient, ghosts: ghostclients, ghostNamespace: ghostNamespace}, controller.Options{})
	if err := fedapis.AddToScheme(live.GetScheme()); err != nil {
		return nil, fmt.Errorf("adding APIs to live cluster's scheme: %v", err)
	}

	if err := co.WatchResourceReconcileObject(live, &v1beta1.KubeFedCluster{}, controller.WatchOptions{}); err != nil {
		return nil, fmt.Errorf("setting up Pod watch in live cluster: %v", err)
	}

	return co, nil
}

type reconciler struct {
	live           client.Client
	ghosts         []client.Client
	ghostNamespace string
}

//var i int = 0

func (r *reconciler) Reconcile(req reconcile.Request) (reconcile.Result, error) {
	//i += 1

	// Fetch the instance
	kubeFedClusterList := &v1beta1.KubeFedClusterList{}
	err := r.live.List(context.TODO(), kubeFedClusterList, &client.ListOptions{})
	if err != nil {
		klog.V(0).Info(err)
	}

	//fmt.Println("Reshape Cluster [len(kubeFedClusterList.Items)] -> ", len(kubeFedClusterList.Items))
	//fmt.Println("Reshape Cluster [prev_length] -> ", prev_length)

	if len(kubeFedClusterList.Items) != prev_length {
		fmt.Println("Reshape Cluster")
		//i = 0

		newCm, _ := clusterManager.NewClusterManager()
		//cm.Mutex.Lock()
		copier.Copy(cm, newCm)
		//cm.Mutex.Unlock()

		prev_length = len(kubeFedClusterList.Items)
		//c <- "reshape"
	}

	return reconcile.Result{}, nil // err
}

func SetupSignalHandler() (stopCh <-chan struct{}) {

	stop := make(chan struct{})

	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}
