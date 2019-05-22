# suich
Switch k8s context easily

# Install Suich

### prerequisite

- Need to set $GOPATH in your PC
- Add $GOPATH/bin in your $PATH variable

### install

```
go get -u github.com/a8uhnf/suich
```
### check installation
```
$ suich help

Root command for switch context in k8s config

Usage:
  suich [flags]
  suich [command]

Available Commands:
  gcp         Command to switch gcloud config[IN PROGRESS]
  help        Help about any command
  kubectl     Update to provided kubectl version.
  pf          port-forward kubernetes pod.[IN PROGRESS]
  rm          Remove context and cluster from kubeconfig
  switch      To switch context use this command

Flags:
  -h, --help   help for suich

Use "suich [command] --help" for more information about a command.
```

# Usage

### suich switch [switch k8s context]

1. Use the ***switch*** command

```
  $ suich switch
```

2. Use the arrow keys on the keyboard to switch contexts when prompted with options

```
----------
2019/05/12 14:24:18 Starting reading config file....
2019/05/12 14:24:18 Successfully read kube-config...
Use the arrow keys to navigate: ↓ ↑ → ←
? Select context:
▸ dev-context
  prod-context
  minikube

```

3. Hit ***Enter*** key & you have successfully changed context without moving much muscle!

```
  ✔ dev-context
```
### suich rm [remove k8s context from config]

1. `suich rm `

2. select cluster with up/down arrow key 

```
2019/05/22 23:51:00 Successfully read kube-config...
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select context: 
  ▸ cluster-1
    cluster-2
    cluster-3
↓   cluster-4
```
3. Select clluster's kubeconfig will be removed from config file

### Download specific version of kubectl

1. use `suich kubectl` command

```
Update to provided kubectl version. Kubectl version must be provided. now by default machine type set to amd64

Usage:
  suich kubectl [flags]

Flags:
  -h, --help             help for kubectl
  -v, --version string   kubectl valid version (default "v1.9.0")
```
2. `suich kubectl -v v1.10.0` this command will download kubectl `v1.10.0` and change your local `kubectl` version