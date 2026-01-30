# kubearchive-operator

This repository contains the installation operator for KubeArchive.
This operator allows for users to install KubeArchive using a custom resource
called `KubeArchiveInstallation`.


## Installing the operator via YAML

```sh
kubectl apply -f https://github.com/kubearchive/kubearchive-operator/releases/download/<VERSION>/kubearchive-operator.yaml
```

## Install the operator using OLM

**Note**: you need OLM installed in your cluster, see
[their documentation](https://olm.operatorframework.io/) for installation options.

Create a new CatalogSource to pull the KubeArchive Operator from:

```sh
$ cat kubearchive-operator-catalogsource.yaml
---
apiVersion: operators.coreos.com/v1alpha1
kind: CatalogSource
metadata:
  name: kubearchive-opeator-catalog
  namespace: olm
spec:
  sourceType: grpc
  image: quay.io/hemartin/kubearchive-operator-catalog:$VERSION
  displayName: KubeArchive Catalog
  publisher: github.com/kubearchive
$ kubectl apply -f kubearchive-operator-catalogsource.yaml
```

Create the `kubearchive-operator` namespace:

```sh
$ kubectl create ns kubearchive-operator
```

Create an OperatorGroup with an empty `spec`:

```sh
$ cat kubearchive-operatorgroup.yaml
kind: OperatorGroup
apiVersion: operators.coreos.com/v1
metadata:
  name: kubearchive-operator-og
  namespace: kubearchive-operator
spec: {}
$ kubectl apply -f kubearchive-operatorgroup.yaml
```

```sh
$ cat kubearchive-subscription.yaml
---
apiVersion: operators.coreos.com/v1alpha1
kind: Subscription
metadata:
  name: kubearchive
  namespace: kubearchive-operator
spec:
  channel: alpha
  installPlanApproval: Automatic
  name: kubearchive-operator
  source: kubearchive-operator-catalog
  sourceNamespace: olm
$ kubectl apply -f kubearchive-subscription.yaml
```

## Contributing

### Development

To develop the operator use:

```sh
kind create cluster
make install run
```

In a different terminal create a `KubeArchiveInstallation` to kick the Reconciliation loop.

## License

Copyright 2026.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

