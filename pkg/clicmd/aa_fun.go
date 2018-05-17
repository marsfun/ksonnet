package clicmd

import (
	"bytes"
	"fmt"
	"github.com/marsfun/ksonnet/pkg/app"
	"github.com/marsfun/ksonnet/pkg/actions"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

const (
	KubeFlowUri     = "file:///Users/max/github/kubeflow/kubeflow"
	KubeFlowPkgCore = "core"
)

type KSConfig struct {
	// ks
	rootPath string
	appName  string
	envName  string
	// k8s
	namespace string
	specFlag  string
	serverUrl string
}

type KSTool struct {
	fs     afero.Fs
	config KSConfig
}

func NewKSTool() *KSTool {
	return &KSTool{
		fs: afero.NewOsFs(),
		//fs: afero.NewMemMapFs(),
	}
}

func (kst *KSTool) Init(ksconfig KSConfig) error {

	//TODO validate

	//
	kst.config = ksconfig

	//appName := app_name
	wd := ksconfig.rootPath

	initDir := ""

	appRoot, err := genKsRoot(ksconfig.appName, wd, initDir)
	if err != nil {
		return err
	}
	// api-server url , namesapce
	//server, namespace := ksconfig.serverUrl, ksconfig.namespace
	//specFlag := ksconfig.specFlag

	m := map[string]interface{}{
		actions.OptionFs:                    kst.fs,
		actions.OptionName:                  kst.config.appName,
		actions.OptionRootPath:              appRoot,
		actions.OptionEnvName:               kst.config.envName,
		actions.OptionSpecFlag:              kst.config.specFlag,
		actions.OptionServer:                kst.config.serverUrl,
		actions.OptionNamespace:             kst.config.namespace,
		actions.OptionSkipDefaultRegistries: true,
	}

	return runAction(actionInit, m)
}

func (kst *KSTool) LoadApp() (*KSApp, error) {
	//if kst.config == nil {
	//	return nil, fmt.Errorf("KSTool config not found.")
	//}

	ka, err := app.Load(kst.fs, fmt.Sprintf("%s/%s", kst.config.rootPath, kst.config.appName), false)

	if err != nil {
		return nil, fmt.Errorf("load app error %v", err)
	}
	return &KSApp{&ka}, nil
}

//=========================================

type KSApp struct {
	app *app.App
}

func (kapp *KSApp) EnvList() error {
	m := map[string]interface{}{
		actions.OptionApp:    *kapp.app,
		actions.OptionOutput: "",
	}

	return runAction(actionEnvList, m)
}

func (kapp *KSApp) addRegistry(name, url string) error {
	m := map[string]interface{}{
		actions.OptionApp:      *kapp.app,
		actions.OptionName:     name,
		actions.OptionURI:      url,
		actions.OptionVersion:  viper.GetString(vRegistryAddVersion),
		actions.OptionOverride: viper.GetBool(vRegistryAddOverride),
	}
	err := runAction(actionRegistryAdd, m)
	if err != nil && err != app.ErrRegistryExists {
		return err
	}
	return nil
}

func (kapp *KSApp) AddKubeFlowRegistry(name string) error {
	return kapp.addRegistry(name, KubeFlowUri)
}

func (kapp *KSApp) installPkg(registry, pkg, version string) error {
	m := map[string]interface{}{
		actions.OptionApp:     *kapp.app,
		actions.OptionLibName: fmt.Sprintf("%s/%s@%s", registry, pkg, version),
		actions.OptionName:    viper.GetString(vPkgInstallName),
	}

	err := runAction(actionPkgInstall, m)
	if err != nil && err.Error() != fmt.Sprintf("package '%s' already exists. Use the --name flag to install this package with a unique identifier",
		pkg) {
		return err
	}

	return nil
}

func (kapp *KSApp) InstallKubeFlowCorePkg(registry, version string) error {
	return kapp.installPkg(registry, KubeFlowPkgCore, version)
}

func (kapp *KSApp) GenComponent(args []string) error {
	m := map[string]interface{}{
		actions.OptionApp:       *kapp.app,
		actions.OptionArguments: args,
	}

	return runAction(actionPrototypeUse, m)

}

func (kapp *KSApp) GenKubeFlowCoreComponent(name string) error {
	compnameSlice := []string{"kubeflow-core", name}
	err := kapp.GenComponent(compnameSlice)
	if err != nil && err.Error() != fmt.Sprintf("create component: component with name %q in module %q already exists", name, "/") {
		return err
	}
	return nil
}

func (kapp *KSApp) GetYaml(envName string, componentsName []string) (string, error) {
	buf := new(bytes.Buffer)
	m := map[string]interface{}{
		actions.OptionApp:            *kapp.app,
		actions.OptionComponentNames: componentsName,
		actions.OptionEnvName:        envName,
		actions.OptionFormat:         viper.GetString(vShowFormat),
		actions.OptionBuffer:         buf,
	}

	if err := extractJsonnetFlags("show"); err != nil {
		return "", errors.Wrap(err, "handle jsonnet flags")
	}

	//return runAction(actionShow, m)
	runErr := runAction(actionShow, m)
	return buf.String(), runErr

}
