# StackSciences CLI

## Setup


You can create a .stctl.yaml file into your home directory and specify the platform hostname and token

```
platform: app.stacksciences.com
token: xxxx
```

Or you can specify the platform you want to connect to and your token using the CLI options

```
--platform app.stacksciences.com --token xxxxx
```

## Exec

```
Usage:
  stctl [flags]
  stctl [command]

Available Commands:
  cluster     manage your clusters
  help        Help about any command
  version     get cli version

Flags:
      --config string     config file (default is $HOME/.stctl/config)
  -h, --help              help for stctl
      --platform string   Platform hostname
      --token string      Platform authentication token

Use "stctl [command] --help" for more information about a command.
```
