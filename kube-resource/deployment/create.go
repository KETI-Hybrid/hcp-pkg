package deployment

import (
	"context"
	"strconv"

	"github.com/KETI-Hybrid/hcp-pkg/util/clusterManager"
	"github.com/google/uuid"

	resourcev1alpha1apis "github.com/KETI-Hybrid/hcp-pkg/apis/resource/v1alpha1"

	ns "github.com/KETI-Hybrid/hcp-pkg/kube-resource/namespace"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"

	v1 "k8s.io/api/apps/v1"
)

func CreateDeployment(clientset *kubernetes.Clientset, node string, deployment *v1.Deployment) error {
	deployment.Spec.Template.Spec.NodeName = node
	deployment.ResourceVersion = ""

	// Namespace 생성
	namespace := deployment.ObjectMeta.Namespace
	ns.CreateNamespace(clientset, namespace)

	// Deployment 배포
	new_dep, err := clientset.AppsV1().Deployments(namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})

	if err != nil {
		klog.Error(err)
		return err
	} else {
		klog.Info("success to create %s [replicas : %d]\n", new_dep.Name, *deployment.Spec.Replicas)
	}

	return nil
}

func DeployDeploymentFromHCPDeployment(hcp_resource *resourcev1alpha1apis.HCPDeployment) (int, bool) {

	// uid 생성
	uid := uuid.ClockSequence()
	cm, _ := clusterManager.NewClusterManager()
	targets := hcp_resource.Spec.SchedulingResult.Targets

	// hcp_resource uid 설정
	hcp_resource.Spec.RealDeploymentMetadata.Labels["uuid"] = strconv.Itoa(uid)
	hcp_resource.Spec.RealDeploymentSpec.Selector.MatchLabels["uuid"] = strconv.Itoa(uid)
	spec := hcp_resource.Spec.RealDeploymentSpec
	metadata := hcp_resource.Spec.RealDeploymentMetadata

	if metadata.Namespace == "" {
		metadata.Namespace = "default"
	}

	// HCPDeployment SchedulingResult에 따라 Deployment 배포
	for _, target := range targets {
		// spec 값 재설정하기
		spec.Replicas = target.Replicas

		// 배포할 Deployment resource 정의
		kube_resource := appsv1.Deployment{
			ObjectMeta: metadata,
			Spec:       spec,
		}

		// Deployment 배포
		clientset := cm.Cluster_kubeClients[target.Cluster]
		r, err := clientset.AppsV1().Deployments(metadata.Namespace).Create(context.TODO(), &kube_resource, metav1.CreateOptions{})

		if err != nil {
			klog.Error(err)
			return -1, false
		} else {
			klog.Info("succeed to deploy deployment %s in %s\n", r.ObjectMeta.Name, target.Cluster)
		}
	}
	return uid, true
}
