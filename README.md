# sync-vendor
sync a subset of a vendor tree out to another repo.

This is a play-space right now.


apimachinery
```
tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=origin fromRepo=git@github.com:/openshift/origin.git fromBranch=master fromDir=vendor/k8s.io/kubernetes/staging/src/k8s.io/apimachinery toRepoName=kubernetes-apimachinery toRepo=git@github.com:/openshift/kubernetes-apimachinery toBranch=master hack/sync-vendor.sh
```

client-go
```
tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=origin fromRepo=git@github.com:/openshift/origin.git fromBranch=master fromDir=vendor/k8s.io/kubernetes/staging/src/k8s.io/client-go toRepoName=kubernetes-client-go toRepo=git@github.com:/openshift/kubernetes-client-go toBranch=master hack/sync-vendor.sh
```
