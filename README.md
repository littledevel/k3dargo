# k3dargo
A small demo repository for GitOps demo using k3d and Argocd


## Introduction


## Installing k3d & tools

#### Run the following command in a terminal:
`wget -q -O - https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash`

#### Verify it was installed properly
`k3d version`

#### Install kubectl 
`curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" `

#### Verify it works
`kubectl version`

#### Install Helm
`wget https://get.helm.sh/helm-v3.13.2-linux-amd64.tar.gz`

`tar zxvf helm-v3.13.2-linux-amd64.tar.gz`

`cp linux-amd64/helm /usr/local/bin/helm`

#### Verify helm works
`helm version`

#### Cleanup helm downloads
`rm zxvf helm-v3.13.2-linux-amd64.tar.gz`

`rm -rf linux-amd64`


## Create cluster

#### Simple cluster with defaults
`k3d cluster create demo`

#### HA Cluster with more controllers
`k3d cluster create demo --servers 3`

#### Scaled HA cluster with more controllers and workers
`k3d cluster create demo --servers 3 --agents 5`

#### Verify all nodes are up
`kubectl get nodes`

#### Expose API port to a specific port
`k3d cluster create demo --api-port 6443`

#### Expose ingress controller to a specific port to use as our entrypoint
`k3d cluster create demo --api-port 6443 -p 8080:80@loadbalancer -p 8443:443@loadbalancer`

#### Cluster management
`k3d cluster list`

`k3d cluster delete demo`

-----------------------------------------

## ArgoCD installation

#### Create namespace
`kubectl create ns argocd`

#### Install with yaml
`kubectl apply -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml -n argocd`

#### Download argocd cli
`wget https://github.com/argoproj/argo-cd/releases/latest/download/argocd-linux-amd64`

`chmod +x argocd-linux-amd64`

`sudo mv argocd-linux-amd4 /usr/local/bin/argocd`

#### Port forward to argocd installation
`kubectl port-forward -n argocd svc/argocd-server 8000:80 &`

#### Login to argocd
`kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo`

## Basic usage

#### Install a public repository
`kubectl apply -f repository.yaml`
