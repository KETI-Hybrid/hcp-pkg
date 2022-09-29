package util

import (
	"context"
	"fmt"

	"gopkg.in/yaml.v2"

	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	genericclient "sigs.k8s.io/kubefed/pkg/client/generic"

	"strconv"
	"strings"
	"time"
)

func GetDuration(createionTime, completionTime time.Time) string {
	duration := completionTime.Sub(createionTime)

	durationTime := ""
	if duration.Hours() >= 240 {
		durationTime = strconv.Itoa(int(duration.Hours()/24)) + "d"
	} else if duration.Hours() >= 24 {
		durationTime = strconv.Itoa(int(duration.Hours()/24)) + "d" + strconv.Itoa(int(duration.Hours())%24) + "h"
	} else if duration.Hours() >= 10 {
		durationTime = strconv.Itoa(int(duration.Hours())) + "h"
	} else if duration.Hours() >= 1 {
		durationTime = strconv.Itoa(int(duration.Hours())) + "h" + strconv.Itoa(int(duration.Minutes())%60) + "m"
	} else if duration.Minutes() >= 10 {
		durationTime = strconv.Itoa(int(duration.Minutes())) + "m"
	} else if duration.Minutes() >= 1 {
		durationTime = strconv.Itoa(int(duration.Minutes())) + "m" + strconv.Itoa(int(duration.Seconds())%60) + "s"
	} else {
		durationTime = strconv.Itoa(int(duration.Seconds())) + "s"
	}
	return durationTime
}

func GetAge(createionTime time.Time) string {
	duration := time.Since(createionTime)
	age := ""
	if duration.Hours() >= 240 {
		age = strconv.Itoa(int(duration.Hours()/24)) + "d"
	} else if duration.Hours() >= 24 {
		age = strconv.Itoa(int(duration.Hours()/24)) + "d" + strconv.Itoa(int(duration.Hours())%24) + "h"
	} else if duration.Hours() >= 10 {
		age = strconv.Itoa(int(duration.Hours())) + "h"
	} else if duration.Hours() >= 1 {
		age = strconv.Itoa(int(duration.Hours())) + "h" + strconv.Itoa(int(duration.Minutes())%60) + "m"
	} else if duration.Minutes() >= 10 {
		age = strconv.Itoa(int(duration.Minutes())) + "m"
	} else if duration.Minutes() >= 1 {
		age = strconv.Itoa(int(duration.Minutes())) + "m" + strconv.Itoa(int(duration.Seconds())%60) + "s"
	} else {
		age = strconv.Itoa(int(duration.Seconds())) + "s"
	}
	return age
}
func GetFileNameList() []string {
	fileOrDirname, _ := filepath.Abs(Option_file)
	filenameList := []string{}

	fi, err := os.Stat(fileOrDirname)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		// do directory stuff

		files, err := ioutil.ReadDir(fileOrDirname)
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range files {
			if err != nil {
				fmt.Println(err)
			}
			if filepath.Ext(f.Name()) == ".yaml" || filepath.Ext(f.Name()) == ".yml" {
				filenameList = append(filenameList, f.Name())
			}
		}
	case mode.IsRegular():
		// do file stuff

		filenameList = append(filenameList, fileOrDirname)
	}
	return filenameList
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}

type ClusterConfiguration struct {
	ControlPlaneEndpoint string `yaml:"controlPlaneEndpoint"`
}

func GetEndpointIP() string {
	kubeconfig, _ := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	genClient := genericclient.NewForConfigOrDie(kubeconfig)
	obj := &v1.ConfigMap{}
	err := genClient.Get(context.TODO(), obj, "kube-system", "kubeadm-config")
	if err != nil {
		log.Fatal(err)
	}
	ip := ""
	if clusterConfigYaml, ok := obj.Data["ClusterConfiguration"]; ok {
		//do something here
		cc := ClusterConfiguration{}
		err := yaml.Unmarshal([]byte(clusterConfigYaml), &cc)
		if err != nil {
			panic(err)
		}
		if cc.ControlPlaneEndpoint != "" {
			endpoints := strings.Split(cc.ControlPlaneEndpoint, ":")
			ip = endpoints[0]
		} else {
			ip = GetOutboundIP()
		}

	} else {
		ip = GetOutboundIP()
	}

	return ip

}
