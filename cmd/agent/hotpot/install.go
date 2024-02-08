package hotpot

import "fmt"

const installScriptUrl = "https://raw.githubusercontent.com/zcubbs/hotpot/%s/scripts/install.sh"

func Install() error {
	version := "main"
	installScriptUrlWithVersion := fmt.Sprintf(installScriptUrl, version)
	_ = installScriptUrlWithVersion
	// TODO: implement
	return nil
}
