package resource

import (
	"context"
	"strconv"

	hcppolicyapis "github.com/KETI-Hybrid/hcp-pkg/apis/hcppolicy/v1alpha1"
	hcppolicyv1alpha1 "github.com/KETI-Hybrid/hcp-pkg/client/hcppolicy/v1alpha1/clientset/versioned"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

type WatchingLevel struct {
	Levels []Level
}

type Level struct {
	Type  string
	Value string
}

func GetHCPPolicy(clientset hcppolicyv1alpha1.Clientset, policy_name string) (*hcppolicyapis.HCPPolicy, error) {

	policy, err := clientset.HcpV1alpha1().HCPPolicies("hcp").Get(context.TODO(), policy_name, metav1.GetOptions{})
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	return policy, nil
}

/*
func GetInitialSettingValue(typ string) (int, string) {
	policy, err := GetHCPPolicy("initial-setting")
	if err != nil {
		klog.Info(err)
	}

	policies := policy.Spec.Template.Spec.Policies
	for _, p := range policies {
		println(p.Type)
		if typ == "default_node_option" && p.Type == "default_node_option" {
			var value string = p.Value
			if value == "" {
				fmt.Printf("ERROR: No %s Value\n", typ)
			} else {
				return -1, value
			}
		} else if p.Type == typ {
			var value int
			value, err := strconv.Atoi(p.Value)
			if err != nil {
				fmt.Printf("ERROR: No %s Value\n", typ)
			}
			return value, ""
		}
	}
	fmt.Printf("ERROR: No Such Type %s\n", typ)
	return -1, ""
}
*/

func GetWatchingLevel(clientset hcppolicyv1alpha1.Clientset) WatchingLevel {
	hcppolicy, err := GetHCPPolicy(clientset, "watching-level")
	if err != nil {
		klog.Error(err)
	}

	var watchingLevel WatchingLevel
	watchingLevel.Levels = make([]Level, 5)
	for i, policy := range hcppolicy.Spec.Template.Spec.Policies {
		var level Level = Level{
			Type:  policy.Type,
			Value: policy.Value[0],
		}
		watchingLevel.Levels[i] = level
	}
	return watchingLevel
}

func GetWarningLevel(clientset hcppolicyv1alpha1.Clientset) Level {
	hcppolicy, err := GetHCPPolicy(clientset, "warning-level")
	if err != nil {
		klog.Error(err)
	}

	// for i, policy := range hcppolicy.Spec.Template.Spec.Policies {
	policies := hcppolicy.Spec.Template.Spec.Policies
	if len(policies) == 1 {
		level := Level{
			Type:  hcppolicy.Spec.Template.Spec.Policies[0].Type,
			Value: hcppolicy.Spec.Template.Spec.Policies[0].Value[0],
		}
		return level
	} else {
		level := Level{
			Type: "none",
		}
		return level
	}
}

func GetAlgorithm(clientset hcppolicyv1alpha1.Clientset) ([]string, error) {
	hcppolicy, err := GetHCPPolicy(clientset, "scheduling-policy")
	if err != nil {
		klog.Error(err)
	}

	policies := hcppolicy.Spec.Template.Spec.Policies
	if len(policies) == 1 {
		policy := policies[0]
		if policy.Type == "algorithm" {
			return policy.Value, nil
		}
	}
	return nil, err
}

func GetCycle(clientset hcppolicyv1alpha1.Clientset) float64 {
	hcppolicy, err := GetHCPPolicy(clientset, "weight-calculation-cycle")
	if err != nil {
		klog.Info(err)
	}

	for _, policy := range hcppolicy.Spec.Template.Spec.Policies {
		if policy.Type == "cycle" && len(policy.Value) == 1 {
			cycle, err := strconv.ParseFloat(policy.Value[0], 64)
			if err == nil && cycle > 0 {
				klog.Info("Policy Type : ", "cycle")
				klog.Infof("Policy Value [cycle] : ", cycle)
				return cycle
			}
		}
	}
	return -1
}
