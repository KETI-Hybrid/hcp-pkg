package deployment

import (
	"context"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"
)

func UpdateDeployment(clientset *kubernetes.Clientset, deployment *v1.Deployment, replicas int32) error {

	if deployment.Status.Replicas == deployment.Status.ReadyReplicas && deployment.Status.Replicas == deployment.Status.AvailableReplicas {

		klog.Info("Deployment status is OK")
		deployment.Spec.Replicas = &replicas
		new_dep, err := clientset.AppsV1().Deployments("default").Update(context.TODO(), deployment, metav1.UpdateOptions{})

		if err != nil {
			klog.Error(err)
			return err
		} else {
			klog.Infof("success to update %s [replicas : %d]\n", new_dep.Name, *deployment.Spec.Replicas)
		}
	} else {
		klog.Info("deployment status is not stable")
	}

	return nil
}
