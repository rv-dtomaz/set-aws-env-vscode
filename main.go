package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {

	mapK := map[string]string{
		"AWS_ACCESS_KEY_ID":     "",
		"AWS_SECRET_ACCESS_KEY": "",
		"AWS_OKTA_PROFILE":      "",
		"AWS_SESSION_TOKEN":     "",
		"AWS_SECURITY_TOKEN":    "",
	}
	env := os.Environ()
	for _, e := range env {
		kv := strings.SplitN(e, "=", 2)
		if _, ok := mapK[kv[0]]; ok {
			mapK[kv[0]] = kv[1]
		}
	}

	basepath, _ := os.Getwd()
	var fullPath = fmt.Sprintf("%s/.vscode/launch.json", basepath)
	if _, err := os.Stat(fullPath); err == nil {
		CreateKeys(fullPath, mapK)
	} else if os.IsNotExist(err) {
		fmt.Println("Create launch.json file: Menu Debug-->AddConfiguration")

	} else {
		fmt.Println(err)
	}

}
func CreateKeys(pathFile string, keys map[string]string) {

	bstr := make([]string, 0)
	envStr := make([]string, 0)
	for k, v := range keys {
		bstr = append(bstr, fmt.Sprintf("\t\t\t\t\t\"%s\":\"%s\"", k, v))
		envStr = append(envStr, fmt.Sprintf("%s=%s", k, v))
	}
	data, _ := ioutil.ReadFile(pathFile)
	fullFileStr := string(data)

	r, _ := regexp.Compile(`\"env\":\s{([^}]+)}`)
	strInterfaceRepo := ""
	finalStr := "\"env\": {\n" + strings.Join(bstr, ",\n") + "\n\t\t\t\t\t}"

	if r.MatchString(fullFileStr) {
		strInterfaceRepo = r.ReplaceAllString(fullFileStr, finalStr)
	} else {
		r, _ = regexp.Compile(`\"env\":\s{}`)
		strInterfaceRepo = r.ReplaceAllString(fullFileStr, finalStr)

	}

	baseEnvPath, _ := os.Getwd()
	var envPath = fmt.Sprintf("%s/.env", baseEnvPath)
	ioutil.WriteFile(envPath, []byte(strings.Join(envStr, "\n")), 0666)
	ioutil.WriteFile(pathFile, []byte(strInterfaceRepo), 0666)
}
