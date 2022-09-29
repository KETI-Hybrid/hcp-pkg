module hcp-pkg

go 1.18

require (
	admiralty.io/multicluster-controller v0.6.0
	github.com/Jeffail/gabs v1.4.0 // indirect
	github.com/aws/aws-sdk-go v1.44.102 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/googleapis/gnostic v0.5.5 // indirect
	github.com/jinzhu/copier v0.3.5
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/net v0.0.0-20220920203100-d0c6ba3f52d9 // indirect
	golang.org/x/oauth2 v0.0.0-20220909003341-f21342109be1 // indirect
	google.golang.org/api v0.96.0 // indirect
	google.golang.org/genproto v0.0.0-20220920201722-2b89144ce006 // indirect
	google.golang.org/grpc v1.49.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/api v0.25.1
	k8s.io/apimachinery v0.25.1
	k8s.io/autoscaler/vertical-pod-autoscaler v0.12.0
	k8s.io/client-go v0.25.1
	k8s.io/klog v1.0.0
	k8s.io/klog/v2 v2.80.1
	k8s.io/kube-openapi v0.0.0-20220803162953-67bda5d908f1 // indirect
	sigs.k8s.io/aws-iam-authenticator v0.5.9 // indirect
	sigs.k8s.io/controller-runtime v0.13.0
	sigs.k8s.io/kubefed v0.10.0
)

require github.com/pkg/errors v0.9.1 // indirect

require (
	github.com/google/uuid v1.3.0
	hybridctl v0.0.0
)

require (
	cloud.google.com/go/compute v1.7.0 // indirect
	github.com/KETI-Hybrid/hcp-apiserver-v1 v0.0.0-20220921080754-59bf5a1498e9 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/evanphx/json-patch v4.12.0+incompatible // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/gofrs/flock v0.7.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.1.0 // indirect
	github.com/googleapis/gax-go/v2 v2.4.0 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/prometheus/client_golang v1.12.2 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.32.1 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/stretchr/testify v1.8.0 // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/sys v0.0.0-20220728004956-3c1f35247d10 // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20220609170525-579cf78fd858 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/apiextensions-apiserver v0.25.0 // indirect
	k8s.io/utils v0.0.0-20220728103510-ee6ede2d64ed // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

replace (

	admiralty.io/multicluster-controller => admiralty.io/multicluster-controller v0.1.0
	github.com/go-logr/logr => github.com/go-logr/logr v0.4.0
	hcp-apiserver => ../hcp-apiserver
	hybridctl => ../hybridctl
	// k8s.io/client-go => k8s.io/client-go v0.17.2
	k8s.io/api => k8s.io/api v0.23.4

	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.22.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.22.4
	k8s.io/apiserver => k8s.io/apiserver v0.23.0-alpha.4
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.23.0-alpha.4
	// k8s.io/kubernetes/pkg/scheduler/api => k8s.io/kubernetes/pkg/scheduler/api v1.17.17
	// k8s.io/kubernetes/pkg/scheduler/api => k8s.io/kubernetes/pkg/scheduler/api v1.17.17
	// sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.6.0

	//hg
	k8s.io/client-go => k8s.io/client-go v0.22.4

	k8s.io/cloud-provider => k8s.io/cloud-provider v0.23.4
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.23.4
	k8s.io/code-generator => k8s.io/code-generator v0.22.4
	k8s.io/component-base => k8s.io/component-base v0.23.4
	k8s.io/component-helpers => k8s.io/component-helpers v0.23.4
	k8s.io/controller-manager => k8s.io/controller-manager v0.23.4
	k8s.io/cri-api => k8s.io/cri-api v0.23.4
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.23.4
	k8s.io/klog/v2 => k8s.io/klog/v2 v2.9.0
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.23.4
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.23.4
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20211115234752-e816edb12b65
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.23.4
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.23.4
	k8s.io/kubectl => k8s.io/kubectl v0.23.4
	k8s.io/kubelet => k8s.io/kubelet v0.23.4
	// k8s.io/kubernetes/pkg/features => k8s.io/kubernetes/pkg/features v1.22.9
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.23.4
	k8s.io/metrics => k8s.io/metrics v0.22.4
	k8s.io/mount-utils => k8s.io/mount-utils v0.23.4
	k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.23.4
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.23.4
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.10.3
	sigs.k8s.io/kubefed => sigs.k8s.io/kubefed v0.10.0
//sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.5.0
// k8s.io/api => k8s.io/api v0.18.6
)
