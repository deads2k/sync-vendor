# sync-vendor
sync a subset of a vendor tree out to another repo.

This is a play-space right now.


kubernetes
```
tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=origin fromRepo=git@github.com:/openshift/origin.git fromBranch=master fromDir=vendor/k8s.io/kubernetes/staging/src/k8s.io/kubernetes toRepoName=kubernetes toRepo=git@github.com:/openshift/kubernetes toBranch=release-1.7.0 hack/sync-vendor.sh
```

apimachinery
```
tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=origin fromRepo=git@github.com:/openshift/origin.git fromBranch=master fromDir=vendor/k8s.io/kubernetes/staging/src/k8s.io/apimachinery toRepoName=kubernetes-apimachinery toRepo=git@github.com:/openshift/kubernetes-apimachinery toBranch=master hack/sync-vendor.sh
```

client-go
```
tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=origin fromRepo=git@github.com:/openshift/origin.git fromBranch=master fromDir=vendor/k8s.io/kubernetes/staging/src/k8s.io/client-go toRepoName=kubernetes-client-go toRepo=git@github.com:/openshift/kubernetes-client-go toBranch=master hack/sync-vendor.sh
```

code-generator
```
tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=origin fromRepo=git@github.com:/openshift/origin.git fromBranch=master fromDir=vendor/k8s.io/kubernetes/staging/src/k8s.io/code-generator toRepoName=kubernetes-code-generator toRepo=git@github.com:/openshift/kubernetes-code-generator toBranch=master hack/sync-vendor.sh
```
apiserver
```
tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=origin fromRepo=git@github.com:/openshift/origin.git fromBranch=master fromDir=vendor/k8s.io/kubernetes/staging/src/k8s.io/apiserver toRepoName=kubernetes-apiserver toRepo=git@github.com:/openshift/kubernetes-apiserver toBranch=master hack/sync-vendor.sh
```
