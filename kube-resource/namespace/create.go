package namespace

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"
)

func CreateNamespace(clientset *kubernetes.Clientset, namespace string) (*corev1.Namespace, error) {

	temp := GetNamespace(clientset, namespace)
	if temp == nil {
		Namespace := corev1.Namespace{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Namespace",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: namespace,
			},
		}

		ns, err := clientset.CoreV1().Namespaces().Create(context.TODO(), &Namespace, metav1.CreateOptions{})

		if err != nil {
			klog.Error(err)
			return nil, err
		} else {
			klog.Infof("success to create namespace %s\n", namespace)
			return ns, nil
		}

	} else {
		return temp, nil
	}
}
