package util

var (
	Option_file      string
	Option_context   string
	Option_Resource  string
	Option_Namespace string
	Option_Name      string
	CONFIGURATION    string
	PROJECT_ID       string
	ZONE             string
	REGION           string
)

var (
	HCPAPIServer = ""
	ResourceMap  = map[string]string{

		// WORKLOAD APIS
		"cj":       "cronjobs",
		"cronjob":  "cronjobs",
		"cronjobs": "cronjobs",

		"ds":        "daemonsets",
		"damonset":  "daemonsets",
		"damonsets": "daemonsets",

		"deploy":      "deployments",
		"deployment":  "deployments",
		"deployments": "deployments",

		"job":  "jobs",
		"jobs": "jobs",

		"po":   "pods",
		"pod":  "pods",
		"pods": "pods",

		"rs":          "replicasets",
		"replicaset":  "replicasets",
		"replicasets": "replicasets",

		"rc":                     "replicationcontrollers",
		"replicationcontroller":  "replicationcontrollers",
		"replicationcontrollers": "replicationcontrollers",

		"sts":          "statefulsets",
		"statefulset":  "statefulsets",
		"statefulsets": "statefulsets",

		// SERVICE APIS
		"ep":        "endpoints",
		"endpoint":  "endpoints",
		"endpoints": "endpoints",

		"ing":       "ingresses",
		"ingress":   "ingresses",
		"ingresses": "ingresses",

		"svc":      "services",
		"service":  "services",
		"services": "services",

		// CONFIG AND STORAGE APIS
		"cm":         "configmaps",
		"configmap":  "configmaps",
		"configmaps": "configmaps",

		"secret":  "secrets",
		"secrets": "secrets",

		"pvc":                    "persistentvolumeclaims",
		"persistentvolumeclaim":  "persistentvolumeclaims",
		"persistentvolumeclaims": "persistentvolumeclaims",

		"sc":             "storageclasses",
		"storageclasse":  "storageclasses",
		"storageclasses": "storageclasses",

		// METADATA APIS
		"crd":                       "customresourcedefinitions",
		"crds":                      "customresourcedefinitions",
		"customresourcedefinition":  "customresourcedefinitions",
		"customresourcedefinitions": "customresourcedefinitions",

		"ev":     "events",
		"event":  "events",
		"events": "events",

		"hpa":                      "horizontalpodautoscalers",
		"horizontalpodautoscaler":  "horizontalpodautoscalers",
		"horizontalpodautoscalers": "horizontalpodautoscalers",

		"vpa":                    "verticalpodautoscalers",
		"verticalpodautoscaler":  "verticalpodautoscalers",
		"verticalpodautoscalers": "verticalpodautoscalers",

		"pdb":                  "poddisruptionbudgets",
		"poddisruptionbudget":  "poddisruptionbudgets",
		"poddisruptionbudgets": "poddisruptionbudgets",

		// CLUSTER APIS
		"apiservice":  "apiservices",
		"apiservices": "apiservices",

		"clusterrole":  "clusterroles",
		"clusterroles": "clusterroles",

		"clusterrolebinding":  "clusterrolebindings",
		"clusterrolebindings": "clusterrolebindings",

		"ns":         "namespaces",
		"namespace":  "namespaces",
		"namespaces": "namespaces",

		"no":    "nodes",
		"node":  "nodes",
		"nodes": "nodes",

		"pv":                "persistentvolumes",
		"persistentvolume":  "persistentvolumes",
		"persistentvolumes": "persistentvolumes",

		"quota":          "resourcequotas",
		"resourcequota":  "resourcequotas",
		"resourcequotas": "resourcequotas",

		"role":  "roles",
		"roles": "roles",

		"rolebinding":  "rolebindings",
		"rolebindings": "rolebindings",

		"sa":              "serviceaccounts",
		"serviceaccount":  "serviceaccounts",
		"serviceaccounts": "serviceaccounts",

		// ETC APIS
		"cluster":         "kubefedclusters",
		"kubefedcluster":  "kubefedclusters",
		"kubefedclusters": "kubefedclusters",

		"odeploy":        "HCPdeployments",
		"HCPdeployment":  "HCPdeployments",
		"HCPdeployments": "HCPdeployments",

		"osvc":        "HCPservices",
		"HCPservice":  "HCPservices",
		"HCPservices": "HCPservices",

		"oing":        "HCPingresss",
		"HCPingress":  "HCPingresss",
		"HCPingresss": "HCPingresss",

		"ohas":                 "HCPhybridautoscalers",
		"HCPhybridautoscaler":  "HCPhybridautoscalers",
		"HCPhybridautoscalers": "HCPhybridautoscalers",

		"opol":       "HCPpolicys",
		"HCPpolicy":  "HCPpolicys",
		"HCPpolicys": "HCPpolicys",

		"ocm":           "HCPconfigmaps",
		"HCPconfigmap":  "HCPconfigmaps",
		"HCPconfigmaps": "HCPconfigmaps",

		"osec":       "HCPsecrets",
		"HCPsecret":  "HCPsecrets",
		"HCPsecrets": "HCPsecrets",

		"ode":             "HCPdnsendpoints",
		"HCPdnsendpoint":  "HCPdnsendpoints",
		"HCPdnsendpoints": "HCPdnsendpoints",
	}
	apiGroup = map[string]string{
		// WORKLOAD APIS
		"cronjobs":               "/apis/batch/v1beta1",
		"daemonsets":             "/apis/apps/v1",
		"deployments":            "/apis/apps/v1",
		"jobs":                   "/apis/batch/v1",
		"pods":                   "/api/v1",
		"replicasets":            "/apis/apps/v1",
		"replicationcontrollers": "/api/v1",
		"statefulsets":           "/apis/apps/v1",

		// SERVICE APIS
		"endpoints": "/api/v1",
		"ingresses": "/apis/networking.k8s.io/v1beta1",
		"services":  "/api/v1",

		// CONFIG AND STORAGE APIS
		"configmaps":             "/api/v1",
		"secrets":                "/api/v1",
		"persistentvolumeclaims": "/api/v1",
		"storageclasses":         "/apis/storage.k8s.io/v1",

		// METADATA APIS
		"customresourcedefinitions": "/apis/apiextensions.k8s.io/v1",
		"events":                    "/api/v1",
		"horizontalpodautoscalers":  "/apis/autoscaling/v1",
		"verticalpodautoscalers":    "/apis/autoscaling.k8s.io/v1beta2",
		"poddisruptionbudgets":      "/apis/policy/v1beta1",

		// CLUSTER APIS
		"apiservices":         "/apis/apiregistration.k8s.io/v1",
		"clusterroles":        "/apis/rbac.authorization.k8s.io/v1",
		"clusterrolebindings": "/apis/rbac.authorization.k8s.io/v1",
		"namespaces":          "/api/v1",
		"nodes":               "/api/v1",
		"persistentvolumes":   "/api/v1",
		"resourcequotas":      "/api/v1",
		"roles":               "/apis/rbac.authorization.k8s.io/v1",
		"rolebindings":        "/apis/rbac.authorization.k8s.io/v1",
		"serviceaccounts":     "/api/v1",

		// ETC APIS
		"kubefedclusters":      "/apis/core.kubefed.io/v1beta1",
		"hcpdeployments":       "/apis/hcp.crd.com/v1alpha1",
		"hcpservices":          "/apis/hcp.crd.com/v1alpha1",
		"hcpingresss":          "/apis/hcp.crd.com/v1alpha1",
		"hcphybridautoscalers": "/apis/hcp.crd.com/v1alpha1",
		"hcppolicys":           "/apis/hcp.crd.com/v1alpha1",
		"hcpconfigmaps":        "/apis/hcp.crd.com/v1alpha1",
		"hcpsecrets":           "/apis/hcp.crd.com/v1alpha1",
		"hcpdnsendpoints":      "/apis/hcp.crd.com/v1alpha1",
	}

	KindMap = map[string]string{
		// WORKLOAD APIS
		"CronJob":                "cronjobs",
		"DaemonSet":              "daemonsets",
		"Deployment":             "deployments",
		"Job":                    "jobs",
		"Pod":                    "pods",
		"ReplicaSets":            "replicasets",
		"ReplicationControllers": "replicationcontrollers",
		"StatefulSets":           "statefulsets",

		// SERVICE APIS
		"Endpoints": "endpoints",
		"Ingress":   "ingresses",
		"Service":   "services",

		// CONFIG AND STORAGE APIS
		"ConfigMap":             "configmaps",
		"Secret":                "secrets",
		"PersistentVolumeClaim": "persistentvolumeclaims",
		"StorageClass":          "storageclasses",

		// METADATA APIS
		"CustomResourceDefinition": "customresourcedefinitions",
		"Event":                    "events",
		"HorizontalPodAutoscaler":  "horizontalpodautoscalers",
		"VerticalPodAutoscaler":    "verticalpodautoscalers",
		"PodDisruptionBudget":      "poddisruptionbudgets",

		// CLUSTER APIS
		"APIService":         "apiservices",
		"ClusterRole":        "clusterroles",
		"ClusterRoleBinding": "clusterrolebindings",
		"Namespace":          "namespaces",
		"Node":               "nodes",
		"PersistentVolume":   "persistentvolumes",
		"ResourceQuota":      "resourcequotas",
		"Role":               "roles",
		"RoleBinding":        "rolebindings",
		"ServiceAccount":     "serviceaccounts",

		// ETC APIS
		"KubeFedCluster":      "kubefedclusters",
		"HCPDeployment":       "hcpdeployments",
		"HCPService":          "hcpservices",
		"HCPIngress":          "hcpingresss",
		"HCPHybridAutoScaler": "hcphybridautoscalers",
		"HCPPolicy":           "hcppolicys",
		"HCPConfigMap":        "hcpconfigmaps",
		"HCPSecret":           "hcpsecrets",
		"HCPDNSEndpoint":      "hcpdnsendpoints",
	}

	noNamespaceResources = []string{"storageclasses", "customresourcedefinitions", "apiservices", "clusterroles", "clusterrolebindings", "namespaces", "nodes", "persistentvolumes"}
)

type MetaInfo struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata   struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	} `json:"metadata"`
	Message string `json:"message"`
	Reason  string `json:"reason"`
	Details struct {
		Name  string `json:"name"`
		Group string `json:"group"`
		Kind  string `json:"kind"`
	} `json:"details"`
	Code int `json:"code"`
}

func CreateLinkParser(metainfo *MetaInfo) string {
	LINK := "https://" + HCPAPIServer + apiGroup[KindMap[metainfo.Kind]] + "/namespaces/"

	if metainfo.Metadata.Namespace == "" {
		LINK += "default"
	} else {
		LINK += metainfo.Metadata.Namespace
	}
	LINK += "/" + KindMap[metainfo.Kind]

	if Option_context == "" {
		LINK += "?clustername=HCP"
	} else {
		LINK += "?clustername=" + Option_context
	}

	return LINK
}
