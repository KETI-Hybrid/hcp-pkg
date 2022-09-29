package clientset

import (
	cobrautil "hybridctl/util"

	hcpclusterv1alpha1 "hcp-pkg/client/hcpcluster/v1alpha1/clientset/versioned"
	hcppolicyv1alpha1 "hcp-pkg/client/hcppolicy/v1alpha1/clientset/versioned"
	resourcev1alpha1 "hcp-pkg/client/resource/v1alpha1/clientset/versioned"

	"k8s.io/client-go/kubernetes"
)

var MasterConfig, _ = cobrautil.BuildConfigFromFlags("master", "/root/.kube/config")
var MasterClienset = kubernetes.NewForConfigOrDie(MasterConfig)
var HCPPolicyClientset = hcppolicyv1alpha1.NewForConfigOrDie(MasterConfig)
var HCPClusterClientset = hcpclusterv1alpha1.NewForConfigOrDie(MasterConfig)
var HCPResourceClientset = resourcev1alpha1.NewForConfigOrDie(MasterConfig)
