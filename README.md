# set-aws-env-vscode

`set-aws-env-vscode` Allows you to set configs and debug applications that use AWS resources and need to authenticate with aws-okta.

 This tool works just for VSCode and you have to create a launch.json

 This tool simple gets AWS temporary keys generated with aws-okta and sets on launch.json ENV. So you can debug aplications.

## Installing

Clone this repository in your MACOS:

```bash
$ git clone github.com/rv-dtomaz/set-aws-env-vscode

$ make release
```


## Running

To run this tool

```bash
$ aws-okta execute [profilename] -- debug-aws

```

![Alt Text](/img/screen.gif)
