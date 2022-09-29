package hcpcluster

import (
	"context"

	hcpclusterv1alpha1 "hcp-pkg/client/hcpcluster/v1alpha1/clientset/versioned"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
)

func FindHCPClusterList(clientset *hcpclusterv1alpha1.Clientset, cluster string) bool {
	cluster_list, err := clientset.HcpV1alpha1().HCPClusters("hcp").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		klog.Errorln(err)
	} else {
		for _, c := range cluster_list.Items {
			if c.ObjectMeta.Name == cluster {
				klog.Infof("Find %s in HCPClusterList\n", cluster)
				return true
			}
		}
	}
	klog.Infof("Fail to find %s in HCPClusterList\n", cluster)
	klog.Info("You should register your cluster to HCP\n")
	return false
}
