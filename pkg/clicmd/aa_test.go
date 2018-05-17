package clicmd

import (
	"fmt"
	"testing"
)

func initIt() KSConfig {

	return KSConfig{
		// ks
		rootPath: "/tmp/newTest",
		appName:  "max-kf",
		envName:  "env_def2",
		// k8s
		namespace: "def_2",
		//specFlag:  "version:v1.9.1",
		specFlag:  "file:///tmp/swagger.json",
		serverUrl: "https://localhost:8443",
	}
}

//func initIt2() KSConfig {
//
//	return KSConfig{
//		// ks
//		rootPath: "/tmp/newTest",
//		appName:  "max-kf2",
//		envName:  "env_def2",
//		// k8s
//		namespace: "def_2",
//		specFlag:  "version:v1.9.1",
//		serverUrl: "https://localhost:8443",
//	}
//}

func Test_EnvList(t *testing.T) {
	cfg := initIt()
	kstool := NewKSTool()
	err := kstool.Init(cfg)
	if err != nil {
		fmt.Println(err)
	}
	kapp, err := kstool.LoadApp()
	if err != nil {
		fmt.Println(err)
	}
	kapp.EnvList()
}

func Test_showYaml(t *testing.T) {
	// ""
	//kfVersion := "v0.1.2"
	//registryName := "kubeflow"

	cfg := initIt()
	kstool := NewKSTool()
	err := kstool.Init(cfg)
	if err != nil {
		fmt.Println(err)
	}
	kapp, err := kstool.LoadApp()
	if err != nil {
		fmt.Println("LoadApp", err)
		return
	}
	kapp.EnvList()
	r, err := kapp.GetYaml(cfg.envName, []string{cfg.appName})
	if err != nil {
		fmt.Println("yaml", err)
		return
	}
	_ = r
	fmt.Println(r)
}

func Test_PrintYaml(t *testing.T) {
	// ""
	kfVersion := "v0.1.2"
	registryName := "kubeflow"

	cfg := initIt()
	kstool := NewKSTool()
	err := kstool.Init(cfg)
	if err != nil {
		fmt.Println(err)
	}
	kapp, err := kstool.LoadApp()
	if err != nil {
		fmt.Println("LoadApp", err)
		return
	}
	kapp.EnvList()
	err = kapp.AddKubeFlowRegistry(registryName)
	if err != nil {
		fmt.Println("reg", err)
		return
	}
	err = kapp.InstallKubeFlowCorePkg(registryName, kfVersion)
	if err != nil {
		fmt.Println("pkg", err)
		return
	}
	err = kapp.GenKubeFlowCoreComponent(cfg.appName)
	if err != nil {
		fmt.Println("comp", err)
		return
	}
	r, err := kapp.GetYaml(cfg.envName, []string{cfg.appName})
	if err != nil {
		fmt.Println("yaml", err)
		return
	}
	_ = r
	fmt.Println(r)

}
