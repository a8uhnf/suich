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
$ suich
Hello From Suich!!!
```

# Usage

1. Use the ***switch*** command

```
  $ suich switch
```

1. Use the arrow keys on the keyboard to switch contexts when prompted with options

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

1. Hit ***Enter*** key & you have successfully changed context without moving much muscle!

```
  ✔ dev-context
```
