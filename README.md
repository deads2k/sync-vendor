# sync-vendor
sync a subset of a vendor tree out to another repo.

This is a play-space right now.


kubernetes
```
tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=origin fromRepo=git@github.com:/openshift/origin.git fromBranch=master fromDir=vendor/k8s.io/kubernetes toRepoName=kubernetes toRepo=git@github.com:/openshift/kubernetes toBranch=release-1.7.6 hack/sync-vendor.sh
```

apimachinery
```
tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=kubernetes fromRepo=git@github.com:/openshift/kubernetes.git fromBranch=release-1.7.6 fromDir=staging/src/k8s.io/apimachinery toRepoName=kubernetes-apimachinery toRepo=git@github.com:/openshift/kubernetes-apimachinery toBranch=release-1.7.6 hack/sync-vendor.sh
```

client-go
```
tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=kubernetes fromRepo=git@github.com:/openshift/kubernetes.git fromBranch=release-1.7.6 fromDir=staging/src/k8s.io/client-go toRepoName=kubernetes-client-go toRepo=git@github.com:/openshift/kubernetes-client-go toBranch=release-1.7.6 hack/sync-vendor.sh
```

code-generator - because of the multiple renames, we've lost patches so we don't apply cleanly.  Leaving this one like this for now.
```
tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=origin fromRepo=git@github.com:/openshift/origin.git fromBranch=master fromDir=vendor/k8s.io/kubernetes/staging/src/k8s.io/code-generator toRepoName=kubernetes-code-generator toRepo=git@github.com:/openshift/kubernetes-code-generator toBranch=master hack/sync-vendor.sh
```

apiserver
```
tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=kubernetes fromRepo=git@github.com:/openshift/kubernetes.git fromBranch=release-1.7.6 fromDir=staging/src/k8s.io/apiserver toRepoName=kubernetes-apiserver toRepo=git@github.com:/openshift/kubernetes-apiserver toBranch=release-1.7.6 hack/sync-vendor.sh
```

kube-aggregator
```
tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=kubernetes fromRepo=git@github.com:/openshift/kubernetes.git fromBranch=release-1.7.6 fromDir=staging/src/k8s.io/kube-aggregator toRepoName=kubernetes-kube-aggregator toRepo=git@github.com:/openshift/kubernetes-kube-aggregator toBranch=release-1.7.6 hack/sync-vendor.sh
```

apiextensions-apiserver
```
tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=kubernetes fromRepo=git@github.com:/openshift/kubernetes.git fromBranch=release-1.7.6 fromDir=staging/src/k8s.io/apiextensions-apiserver toRepoName=kubernetes-apiextensions-apiserver toRepo=git@github.com:/openshift/kubernetes-apiextensions-apiserver toBranch=release-1.7.6 hack/sync-vendor.sh
```

metrics
```
tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=kubernetes fromRepo=git@github.com:/openshift/kubernetes.git fromBranch=release-1.7.6 fromDir=staging/src/k8s.io/metrics toRepoName=kubernetes-metrics toRepo=git@github.com:/openshift/kubernetes-metrics toBranch=release-1.7.6 hack/sync-vendor.sh
```


docker/distribution
```
tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=origin fromRepo=git@github.com:/openshift/origin.git fromBranch=master fromDir=vendor/github.com/docker/distribution toRepoName=docker-distribution toRepo=git@github.com:/openshift/docker-distribution toBranch=release-2.6.2 hack/sync-vendor.sh
```

emicklei/go-restful-swagger12
```
tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=origin fromRepo=git@github.com:/openshift/origin.git fromBranch=master fromDir=vendor/github.com/emicklei/go-restful-swagger12 toRepoName=go-restful-swagger12 toRepo=git@github.com:/openshift/emicklei-go-restful-swagger12 toBranch=release-1.0.1 hack/sync-vendor.sh
```


