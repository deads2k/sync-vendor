# sync-vendor
sync a subset of a vendor tree out to another repo.

This is a play-space right now.


kubernetes - 1.9.x
```
tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=origin fromRepo=git@github.com:/openshift/origin.git fromBranch=master fromDir=vendor/github.com/opencontainers/runc toRepoName=opencontainers-runc toRepo=git@github.com:/openshift/opencontainers-runc toBranch=openshift-3.9 hack/sync-vendor.sh

tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=origin fromRepo=git@github.com:/openshift/origin.git fromBranch=master fromDir=vendor/k8s.io/kubernetes toRepoName=kubernetes toRepo=git@github.com:/openshift/kubernetes toBranch=release-1.9.0-beta.1 hack/sync-vendor.sh

tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=kubernetes fromRepo=git@github.com:/openshift/kubernetes.git fromBranch=release-1.9.0-beta.1 fromDir=staging/src/k8s.io/code-generator toRepoName=kubernetes-code-generator toRepo=git@github.com:/openshift/kubernetes-code-generator toBranch=release-1.9.0-beta.1 hack/sync-vendor.sh

tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=kubernetes fromRepo=git@github.com:/openshift/kubernetes.git fromBranch=release-1.9.0-beta.1 fromDir=staging/src/k8s.io/apimachinery toRepoName=kubernetes-apimachinery toRepo=git@github.com:/openshift/kubernetes-apimachinery toBranch=release-1.9.0-beta.1 hack/sync-vendor.sh

tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=kubernetes fromRepo=git@github.com:/openshift/kubernetes.git fromBranch=release-1.9.0-beta.1 fromDir=staging/src/k8s.io/api toRepoName=kubernetes-api toRepo=git@github.com:/openshift/kubernetes-api toBranch=release-1.9.0-beta.1 hack/sync-vendor.sh

tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=kubernetes fromRepo=git@github.com:/openshift/kubernetes.git fromBranch=release-1.9.0-beta.1 fromDir=staging/src/k8s.io/client-go toRepoName=kubernetes-client-go toRepo=git@github.com:/openshift/kubernetes-client-go toBranch=release-1.9.0-beta.1 hack/sync-vendor.sh

tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=kubernetes fromRepo=git@github.com:/openshift/kubernetes.git fromBranch=release-1.9.0-beta.1 fromDir=staging/src/k8s.io/apiserver toRepoName=kubernetes-apiserver toRepo=git@github.com:/openshift/kubernetes-apiserver toBranch=release-1.9.0-beta.1 hack/sync-vendor.sh

tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=kubernetes fromRepo=git@github.com:/openshift/kubernetes.git fromBranch=release-1.9.0-beta.1 fromDir=staging/src/k8s.io/kube-aggregator toRepoName=kubernetes-kube-aggregator toRepo=git@github.com:/openshift/kubernetes-kube-aggregator toBranch=release-1.9.0-beta.1 hack/sync-vendor.sh

tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=kubernetes fromRepo=git@github.com:/openshift/kubernetes.git fromBranch=release-1.9.0-beta.1 fromDir=staging/src/k8s.io/apiextensions-apiserver toRepoName=kubernetes-apiextensions-apiserver toRepo=git@github.com:/openshift/kubernetes-apiextensions-apiserver toBranch=release-1.9.0-beta.1 hack/sync-vendor.sh

tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=kubernetes fromRepo=git@github.com:/openshift/kubernetes.git fromBranch=release-1.9.0-beta.1 fromDir=staging/src/k8s.io/metrics toRepoName=kubernetes-metrics toRepo=git@github.com:/openshift/kubernetes-metrics toBranch=release-1.9.0-beta.1 hack/sync-vendor.sh

```

kubernetes - 1.8.x
```
tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=origin fromRepo=git@github.com:/openshift/origin.git fromBranch=release-3.8 fromDir=vendor/k8s.io/gengo toRepoName=kubernetes-gengo toRepo=git@github.com:/openshift/kubernetes-gengo toBranch=openshift-3.8 hack/sync-vendor.sh

tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=origin fromRepo=git@github.com:/openshift/origin.git fromBranch=release-3.8 fromDir=vendor/k8s.io/kubernetes toRepoName=kubernetes toRepo=git@github.com:/openshift/kubernetes toBranch=release-1.8.1 hack/sync-vendor.sh

tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=kubernetes fromRepo=git@github.com:/openshift/kubernetes.git fromBranch=release-1.8.1 fromDir=staging/src/k8s.io/code-generator toRepoName=kubernetes-code-generator toRepo=git@github.com:/openshift/kubernetes-code-generator toBranch=release-1.8.1 hack/sync-vendor.sh

tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=kubernetes fromRepo=git@github.com:/openshift/kubernetes.git fromBranch=release-1.8.1 fromDir=staging/src/k8s.io/apimachinery toRepoName=kubernetes-apimachinery toRepo=git@github.com:/openshift/kubernetes-apimachinery toBranch=release-1.8.1 hack/sync-vendor.sh

tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=kubernetes fromRepo=git@github.com:/openshift/kubernetes.git fromBranch=release-1.8.1 fromDir=staging/src/k8s.io/api toRepoName=kubernetes-api toRepo=git@github.com:/openshift/kubernetes-api toBranch=release-1.8.1 hack/sync-vendor.sh

tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=kubernetes fromRepo=git@github.com:/openshift/kubernetes.git fromBranch=release-1.8.1 fromDir=staging/src/k8s.io/client-go toRepoName=kubernetes-client-go toRepo=git@github.com:/openshift/kubernetes-client-go toBranch=release-1.8.1 hack/sync-vendor.sh

tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=kubernetes fromRepo=git@github.com:/openshift/kubernetes.git fromBranch=release-1.8.1 fromDir=staging/src/k8s.io/apiserver toRepoName=kubernetes-apiserver toRepo=git@github.com:/openshift/kubernetes-apiserver toBranch=release-1.8.1 hack/sync-vendor.sh

tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=kubernetes fromRepo=git@github.com:/openshift/kubernetes.git fromBranch=release-1.8.1 fromDir=staging/src/k8s.io/kube-aggregator toRepoName=kubernetes-kube-aggregator toRepo=git@github.com:/openshift/kubernetes-kube-aggregator toBranch=release-1.8.1 hack/sync-vendor.sh

tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=kubernetes fromRepo=git@github.com:/openshift/kubernetes.git fromBranch=release-1.8.1 fromDir=staging/src/k8s.io/apiextensions-apiserver toRepoName=kubernetes-apiextensions-apiserver toRepo=git@github.com:/openshift/kubernetes-apiextensions-apiserver toBranch=release-1.8.1 hack/sync-vendor.sh

tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=kubernetes fromRepo=git@github.com:/openshift/kubernetes.git fromBranch=release-1.8.1 fromDir=staging/src/k8s.io/metrics toRepoName=kubernetes-metrics toRepo=git@github.com:/openshift/kubernetes-metrics toBranch=release-1.8.1 hack/sync-vendor.sh

```


kubernetes - 1.7.x
```
tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=origin fromRepo=git@github.com:/openshift/origin.git fromBranch=master fromDir=vendor/k8s.io/kubernetes toRepoName=kubernetes toRepo=git@github.com:/openshift/kubernetes toBranch=release-1.7.6 hack/sync-vendor.sh

tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=kubernetes fromRepo=git@github.com:/openshift/kubernetes.git fromBranch=release-1.7.6 fromDir=staging/src/k8s.io/apimachinery toRepoName=kubernetes-apimachinery toRepo=git@github.com:/openshift/kubernetes-apimachinery toBranch=release-1.7.6 hack/sync-vendor.sh

tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=kubernetes fromRepo=git@github.com:/openshift/kubernetes.git fromBranch=release-1.7.6 fromDir=staging/src/k8s.io/client-go toRepoName=kubernetes-client-go toRepo=git@github.com:/openshift/kubernetes-client-go toBranch=release-1.7.6 hack/sync-vendor.sh

tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=origin fromRepo=git@github.com:/openshift/origin.git fromBranch=master fromDir=vendor/k8s.io/kubernetes/staging/src/k8s.io/code-generator toRepoName=kubernetes-code-generator toRepo=git@github.com:/openshift/kubernetes-code-generator toBranch=master hack/sync-vendor.sh

tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=kubernetes fromRepo=git@github.com:/openshift/kubernetes.git fromBranch=release-1.7.6 fromDir=staging/src/k8s.io/apiserver toRepoName=kubernetes-apiserver toRepo=git@github.com:/openshift/kubernetes-apiserver toBranch=release-1.7.6 hack/sync-vendor.sh

tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=kubernetes fromRepo=git@github.com:/openshift/kubernetes.git fromBranch=release-1.7.6 fromDir=staging/src/k8s.io/kube-aggregator toRepoName=kubernetes-kube-aggregator toRepo=git@github.com:/openshift/kubernetes-kube-aggregator toBranch=release-1.7.6 hack/sync-vendor.sh

tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=kubernetes fromRepo=git@github.com:/openshift/kubernetes.git fromBranch=release-1.7.6 fromDir=staging/src/k8s.io/apiextensions-apiserver toRepoName=kubernetes-apiextensions-apiserver toRepo=git@github.com:/openshift/kubernetes-apiextensions-apiserver toBranch=release-1.7.6 hack/sync-vendor.sh

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

cadvisor - for 3.7
```
tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=origin fromRepo=git@github.com:/openshift/origin.git fromBranch=release-3.7 fromDir=vendor/github.com/google/cadvisor toRepoName=google-cadvisor toRepo=git@github.com:/openshift/google-cadvisor toBranch=release-0.26.1 hack/sync-vendor.sh
```
cadvisor - for 3.8-master
```
tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=origin fromRepo=git@github.com:/openshift/origin.git fromBranch=master fromDir=vendor/github.com/google/cadvisor toRepoName=google-cadvisor toRepo=git@github.com:/openshift/google-cadvisor toBranch=release-0.27.1 hack/sync-vendor.sh
```

container/images - for 3.8-master
```
tmpDir=/home/deads/workspaces/sync-vendor/src/github.com/deads2k/working-dir-02 fromRepoName=origin fromRepo=git@github.com:/openshift/origin.git fromBranch=master fromDir=vendor/github.com/containers/image toRepoName=containers-image toRepo=git@github.com:/openshift/containers-image toBranch=openshift-3.8 hack/sync-vendor.sh
```

