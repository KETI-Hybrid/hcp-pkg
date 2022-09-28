package deployment

import (
	"context"
	"strings"

	resourcev1alpha1 "github.com/KETI-Hybrid/hcp-pkg/apis/resource/v1alpha1"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"
)

func GetDeploymentName(pod *corev1.Pod) string {
	str := pod.GenerateName
	list := strings.SplitAfterN(str, "-", -1)
	length := len(list[len(list)-2]) + 1
	name := str[:len(str)-length]
	return name
}

func GetDeployment(clientset *kubernetes.Clientset, deployment_name string, namespace string) (*v1.Deployment, error) {
	d, err := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), deployment_name, metav1.GetOptions{})
	if err != nil {
		klog.Error(err)
	} else {
		klog.Infof("success to get deployment %s [replicas : %d]\n", d.Name, *d.Spec.Replicas)
	}
	return d, err
}

func GetDeploymentFromPod(clientset *kubernetes.Clientset, pod *corev1.Pod) (*v1.Deployment, error) {
	deploymentName := GetDeploymentName(pod)
	d, err := clientset.AppsV1().Deployments(pod.Namespace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		klog.Error(err)
	} else {
		klog.Infof("success to get deployment %s [replicas : %d]\n", d.Name, *d.Spec.Replicas)
	}
	return d, err
}

func HCPDeploymentToDeployment(hcp_resource *resourcev1alpha1.HCPDeployment) appsv1.Deployment {
	kube_resource := appsv1.Deployment{}
	metadata := hcp_resource.Spec.RealDeploymentMetadata
	if metadata.Namespace == "" {
		metadata.Namespace = "default"
	}
	spec := hcp_resource.Spec.RealDeploymentSpec

	kube_resource.ObjectMeta = metadata
	kube_resource.Spec = spec

	return kube_resource
}
