package util

import (
	"context"
	"io/ioutil"
	"log"

	hcpclusterv1alpha1 "hcp-pkg/client/hcpcluster/v1alpha1/clientset/versioned"

	"gopkg.in/yaml.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var master_config, _ = BuildConfigFromFlags("master", "/root/.kube/config")

type KubeConfig struct {
	APIVersion string `yaml:"apiVersion"`
	Clusters   []struct {
		Cluster struct {
			CertificateAuthorityData string `yaml:"certificate-authority-data"`
			Server                   string `yaml:"server"`
		} `yaml:"cluster"`
		Name string `yaml:"name"`
	} `yaml:"clusters"`
	Contexts []struct {
		Context struct {
			Cluster string `yaml:"cluster"`
			User    string `yaml:"user"`
		} `yaml:"context"`
		Name string `yaml:"name"`
	} `yaml:"contexts"`
	CurrentContext string `yaml:"current-context"`
	Kind           string `yaml:"kind"`
	Preferences    struct {
	} `yaml:"preferences"`
	Users []struct {
		Name string `yaml:"name"`
		User struct {
			ClientCertificateData string `yaml:"client-certificate-data,omitempty"`
			ClientKeyData         string `yaml:"client-key-data,omitempty"`
			Token                 string `yaml:"token,omitempty"`
			AuthProvider          struct {
				Config struct {
					AccessToken string `yaml:"access-token,omitempty"`
					CmdArgs     string `yaml:"cmd-args,omitempty"`
					CmdPath     string `yaml:"cmd-path,omitempty"`
					Expiry      string `yaml:"expiry,omitempty"`
					ExpiryKey   string `yaml:"expiry-key,omitempty"`
					TokenKey    string `yaml:"token-key,omitempty"`
				} `yaml:"config,omitempty"`
				Name string `yaml:"name,omitempty"`
			} `yaml:"auth-provider,omitempty"`
			Exec struct {
				APIVersion string      `yaml:"apiVersion,omitempty"`
				Args       []string    `yaml:"args,omitempty"`
				Command    string      `yaml:"command,omitempty"`
				Env        interface{} `yaml:"env,omitempty"`
			} `yaml:"exec,omitempty"`
		} `yaml:"user"`
	} `yaml:"users"`
}

func BuildConfigFromFlags(context, kubeconfigPath string) (*rest.Config, error) {
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfigPath},
		&clientcmd.ConfigOverrides{
			CurrentContext: context,
		}).ClientConfig()
}

func UnMarshalKubeConfig(data []byte) (KubeConfig, error) {
	var kubeconfig KubeConfig
	err := yaml.Unmarshal(data, &kubeconfig)
	return kubeconfig, err
}

func GetKubeConfig(kubeconfigPath string) *KubeConfig {
	c := &KubeConfig{}
	yamlFile, err := ioutil.ReadFile(kubeconfigPath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

func WriteKubeConfig(c *KubeConfig, filepath string) error {
	d, err := yaml.Marshal(&c)
	err = ioutil.WriteFile(filepath, d, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (c *KubeConfig) GetContextList() []string {
	res := []string{}

	for _, context := range c.Contexts {
		res = append(res, context.Name)
	}
	return res

}

func ChangeConfigClusterName(platform string, clustername string) error {

	clientset, err := hcpclusterv1alpha1.NewForConfig(master_config)
	if err != nil {
		return err
	}

	cluster, err := clientset.HcpV1alpha1().HCPClusters(platform).Get(context.TODO(), clustername, metav1.GetOptions{})
	if err != nil {
		return err
	}
	hcpconfig, err := UnMarshalKubeConfig(cluster.Spec.KubeconfigInfo)
	if err != nil {
		return err
	}

	if len(hcpconfig.Clusters) > 0 && len(hcpconfig.Contexts) > 0 && len(hcpconfig.Users) > 0 {
		hcpconfig.Clusters[0].Name = clustername
		hcpconfig.Contexts[0].Name = clustername
		hcpconfig.Contexts[0].Context.Cluster = clustername
		hcpconfig.Contexts[0].Context.User = clustername
		hcpconfig.Users[0].Name = clustername
	} else {
		return nil
	}
	data, err := yaml.Marshal(hcpconfig)
	if err != nil {
		return err
	}
	cluster.Spec.KubeconfigInfo = data
	_, err = clientset.HcpV1alpha1().HCPClusters(platform).Update(context.TODO(), cluster, metav1.UpdateOptions{})
	if err != nil {
		return err
	}

	// add this config to .kube/config

	bytes, err := ioutil.ReadFile("/root/.kube/config")
	if err != nil {
		return err
	}
	kubeconfig, err := UnMarshalKubeConfig(bytes)
	if err != nil {
		return err
	}
	exist := false
	for _, c := range kubeconfig.Clusters {
		if c.Name == clustername {
			exist = true
			break
		}
	}
	if !exist {
		kubeconfig.Clusters = append(kubeconfig.Clusters, hcpconfig.Clusters...)
		kubeconfig.Contexts = append(kubeconfig.Contexts, hcpconfig.Contexts...)
		kubeconfig.Users = append(kubeconfig.Users, hcpconfig.Users...)
	}
	data, err = yaml.Marshal(&kubeconfig)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("/root/.kube/config", data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func DeleteConfig(platform string, clustername string) error {

	hcp_cluster, err := hcpclusterv1alpha1.NewForConfig(master_config)
	if err != nil {
		return err
	}

	cluster, err := hcp_cluster.HcpV1alpha1().HCPClusters("hcp").Get(context.TODO(), clustername, metav1.GetOptions{})
	if err != nil {
		return err
	}
	hcpconfig, err := UnMarshalKubeConfig(cluster.Spec.KubeconfigInfo)
	if err != nil {
		return err
	}

	if len(hcpconfig.Clusters) > 0 && len(hcpconfig.Contexts) > 0 && len(hcpconfig.Users) > 0 {
		bytes, err := ioutil.ReadFile("/root/.kube/config")
		if err != nil {
			return err
		}
		kubeconfig, err := UnMarshalKubeConfig(bytes)
		if err != nil {
			return err
		}

		for i, c := range kubeconfig.Clusters {
			if c.Name == clustername {
				kubeconfig.Clusters = append(kubeconfig.Clusters[:i], kubeconfig.Clusters[i+1:]...)
				break
			}
		}

		for i, c := range kubeconfig.Contexts {
			if c.Name == clustername {
				kubeconfig.Contexts = append(kubeconfig.Contexts[:i], kubeconfig.Contexts[i+1:]...)
				break
			}
		}

		for i, c := range kubeconfig.Users {
			if c.Name == clustername {
				kubeconfig.Users = append(kubeconfig.Users[:i], kubeconfig.Users[i+1:]...)
				break
			}
		}

		data, err := yaml.Marshal(&kubeconfig)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile("/root/.kube/config", data, 0644)
		if err != nil {
			return err
		}
	} else {
		return nil
	}

	return nil
}
