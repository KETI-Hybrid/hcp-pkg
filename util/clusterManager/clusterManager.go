package clusterManager

import (
	// clientV1alpha1 "Hybrid_Cloud/pkg/client/policy/v1alpha1/clientset"

	"context"
	"fmt"

	hcpclusterv1alpha1 "hcp-pkg/client/hcpcluster/v1alpha1/clientset/versioned"
	hcppolicyv1alpha1 "hcp-pkg/client/hcppolicy/v1alpha1/clientset/versioned"
	resourcev1alpha1 "hcp-pkg/client/resource/v1alpha1/clientset/versioned"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	fedv1b1 "sigs.k8s.io/kubefed/pkg/apis/core/v1beta1"
	genericclient "sigs.k8s.io/kubefed/pkg/client/generic"
	util "sigs.k8s.io/kubefed/pkg/controller/util"
)

type ClusterManager struct {
	Fed_namespace   string
	Host_config     *rest.Config
	Host_client     genericclient.Client
	Host_kubeClient *kubernetes.Clientset
	// Crd_client          *clientV1alpha1.ExampleV1Alpha1Client
	Cluster_list        *fedv1b1.KubeFedClusterList
	Node_list           *corev1.NodeList
	Cluster_configs     map[string]*rest.Config
	Cluster_genClients  map[string]genericclient.Client
	Cluster_kubeClients map[string]*kubernetes.Clientset
	HCPCluster_Client   *hcpclusterv1alpha1.Clientset
	HCPPolicy_Client    *hcppolicyv1alpha1.Clientset
	HCPResource_Client  *resourcev1alpha1.Clientset
	// Cluster_dynClients  map[string]dynamic.Interface
	//Mutex	*sync.Mutex
}

func ListKubeFedClusters(genClient genericclient.Client, namespace string) *fedv1b1.KubeFedClusterList {
	tempClusterList := &fedv1b1.KubeFedClusterList{}
	clusterList := &fedv1b1.KubeFedClusterList{}

	err := genClient.List(context.TODO(), tempClusterList, namespace, &client.ListOptions{})

	if err != nil {
		fmt.Printf("Error retrieving list of federated clusters: %+v\n", err)
	}

	/*
		if len(tempClusterList.Items) == 0 {
			fmt.Println("No federated clusters found")
		}
	*/

	// Status Check
	for _, cluster := range tempClusterList.Items {
		status := true

		for _, cond := range cluster.Status.Conditions {
			if cond.Type == "Offline" {
				status = false
				break
			}
		}
		if status {
			clusterList.Items = append(clusterList.Items, cluster)
		}
	}

	return clusterList
}

func KubeFedClusterConfigs(clusterList *fedv1b1.KubeFedClusterList, genClient genericclient.Client, fedNamespace string) map[string]*rest.Config {
	clusterConfigs := make(map[string]*rest.Config)
	for _, cluster := range clusterList.Items {
		config, _ := util.BuildClusterConfig(&cluster, genClient, fedNamespace)
		clusterConfigs[cluster.Name] = config
	}
	return clusterConfigs
}

func KubeFedClusterGenClients(clusterList *fedv1b1.KubeFedClusterList, cluster_configs map[string]*rest.Config) map[string]genericclient.Client {

	cluster_clients := make(map[string]genericclient.Client)
	for _, cluster := range clusterList.Items {
		clusterName := cluster.Name
		cluster_config := cluster_configs[clusterName]
		cluster_client := genericclient.NewForConfigOrDie(cluster_config)
		cluster_clients[clusterName] = cluster_client
	}
	return cluster_clients
}

func KubeFedClusterKubeClients(clusterList *fedv1b1.KubeFedClusterList, cluster_configs map[string]*rest.Config) map[string]*kubernetes.Clientset {

	cluster_clients := make(map[string]*kubernetes.Clientset)
	for _, cluster := range clusterList.Items {
		clusterName := cluster.Name
		cluster_config := cluster_configs[clusterName]
		cluster_client := kubernetes.NewForConfigOrDie(cluster_config)
		cluster_clients[clusterName] = cluster_client
	}
	return cluster_clients
}

func GetNodeList(c *kubernetes.Clientset) (*corev1.NodeList, error) {
	nodeList, err := c.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error retrieving list of Node: %+v\n", err)
		return nodeList, err
	}

	if len(nodeList.Items) == 0 {
		fmt.Println("No Nodes found")
		return nodeList, err
	}
	return nodeList, nil
}

func NewClusterManager() (*ClusterManager, error) {
	fed_namespace := "kube-federation-system"
	host_config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	host_client := genericclient.NewForConfigOrDie(host_config)
	host_kubeclient := kubernetes.NewForConfigOrDie(host_config)

	cluster_list := ListKubeFedClusters(host_client, fed_namespace)
	node_list, _ := GetNodeList(host_kubeclient)

	cluster_configs := KubeFedClusterConfigs(cluster_list, host_client, fed_namespace)
	cluster_gen_clients := KubeFedClusterGenClients(cluster_list, cluster_configs)
	cluster_kube_clients := KubeFedClusterKubeClients(cluster_list, cluster_configs)

	hcppolicy_client := hcppolicyv1alpha1.NewForConfigOrDie(host_config)
	hcpcluster_client := hcpclusterv1alpha1.NewForConfigOrDie(host_config)
	hcpresource_client := resourcev1alpha1.NewForConfigOrDie(host_config)
	cm := &ClusterManager{
		Fed_namespace:       fed_namespace,
		Host_config:         host_config,
		Host_client:         host_client,
		Host_kubeClient:     host_kubeclient,
		Cluster_list:        cluster_list,
		Node_list:           node_list,
		Cluster_configs:     cluster_configs,
		Cluster_genClients:  cluster_gen_clients,
		Cluster_kubeClients: cluster_kube_clients,
		HCPPolicy_Client:    hcppolicy_client,
		HCPCluster_Client:   hcpcluster_client,
		HCPResource_Client:  hcpresource_client,
	}
	return cm, nil
}
