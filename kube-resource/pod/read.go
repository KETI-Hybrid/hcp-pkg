package pod

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/klog"
)

func GetPod(clientset *kubernetes.Clientset, pod string, pod_namespace string) (*corev1.Pod, error) {
	p, err := clientset.CoreV1().Pods(pod_namespace).Get(context.TODO(), pod, metav1.GetOptions{})
	if err != nil {
		klog.Error(err)
		return p, err
	} else {
		klog.Infof("success to get pod %s [cluster %s, node %s]\n", p.Name, p.ClusterName, p.Spec.NodeName)
		return p, err
	}
}
