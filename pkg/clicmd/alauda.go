package clicmd

//-------------------------

func RunAction(name initName, args map[string]interface{}) error {
	return runAction(name, args)
}

func GenKsRoot(appName, ksDir, wd string) (string, error) {
	return genKsRoot(appName, ksDir, wd)
}

const (
	ActionComponentList     = actionComponentList
	ActionComponentRm       = actionComponentRm
	ActionDelete            = actionDelete
	ActionDiff              = actionDiff
	ActionEnvAdd            = actionEnvAdd
	ActionEnvCurrent        = actionEnvCurrent
	ActionEnvDescribe       = actionEnvDescribe
	ActionEnvList           = actionEnvList
	ActionEnvRm             = actionEnvRm
	ActionEnvSet            = actionEnvSet
	ActionEnvTargets        = actionEnvTargets
	ActionEnvUpdate         = actionEnvUpdate
	ActionImport            = actionImport
	ActionInit              = actionInit
	ActionModuleCreate      = actionModuleCreate
	ActionModuleList        = actionModuleList
	ActionParamDelete       = actionParamDelete
	ActionParamDiff         = actionParamDiff
	ActionParamList         = actionParamList
	ActionParamSet          = actionParamSet
	ActionParamUnset        = actionParamUnset
	ActionPkgDescribe       = actionPkgDescribe
	ActionPkgInstall        = actionPkgInstall
	ActionPkgList           = actionPkgList
	ActionPrototypeDescribe = actionPrototypeDescribe
	ActionPrototypeList     = actionPrototypeList
	ActionPrototypePreview  = actionPrototypePreview
	ActionPrototypeSearch   = actionPrototypeSearch
	ActionPrototypeUse      = actionPrototypeUse
	ActionRegistryAdd       = actionRegistryAdd
	ActionRegistryDescribe  = actionRegistryDescribe
	ActionRegistryList      = actionRegistryList
	ActionShow              = actionShow
	ActionUpgrade           = actionUpgrade
	ActionValidate          = actionValidate
)
