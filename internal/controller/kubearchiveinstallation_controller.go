/*
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
*/

package controller

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	kubearchiveorgv1 "github.com/kubearchive/kubearchive-operator/api/v1"
	"gopkg.in/yaml.v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

// KubeArchiveInstallationReconciler reconciles a KubeArchiveInstallation object
type KubeArchiveInstallationReconciler struct {
	client.Client
	Scheme        *runtime.Scheme
	DynamicClient *dynamic.DynamicClient
}

// +kubebuilder:rbac:groups=kubearchive.org,resources=kubearchiveinstallations,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=kubearchive.org,resources=kubearchiveinstallations/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=kubearchive.org,resources=kubearchiveinstallations/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *KubeArchiveInstallationReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logf.FromContext(ctx)
	log.Info("Reconciling...")

	kaInstallation := kubearchiveorgv1.KubeArchiveInstallation{}
	if err := r.Client.Get(ctx, req.NamespacedName, &kaInstallation); err != nil {
		// Ignore not-found errors, since they can't be fixed by an immediate requeue (we need
		// to wait for a new notification), and we can get them on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	installSpec := kaInstallation.Spec
	log.Info("Installing KubeArchive version", "version", installSpec.Version)

	httpClient := http.Client{}
	downloadURL := fmt.Sprintf("https://github.com/kubearchive/kubearchive/releases/download/%s/kubearchive.yaml", installSpec.Version)
	resp, err := httpClient.Get(downloadURL)
	if err != nil {
		return ctrl.Result{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ctrl.Result{}, err
	}

	var resources []unstructured.Unstructured
	for resourceStr := range strings.SplitSeq(string(body), "---") {
		var res map[string]any
		err = yaml.Unmarshal([]byte(resourceStr), &res)
		if res == nil {
			continue
		}
		resource := unstructured.Unstructured{Object: res}
		resources = append(resources, resource)
	}

	for _, resource := range resources {
		parts := strings.Split(resource.GetAPIVersion(), "/")
		group := ""
		version := "v1"
		res := ""
		if len(parts) > 1 {
			group = parts[0]
			version = parts[1]
		}

		switch kind := resource.GetKind(); kind {
		case "Namespace":
			res = "namespaces"
		case "CustomResourceDefinition":
			res = "customresourcedefinitions"
		case "ServiceAccount":
			res = "serviceaccounts"
		case "ClusterRole":
			res = "clusterroles"
		case "Role":
			res = "roles"
		case "RoleBinding":
			res = "rolebindings"
		case "ClusterRoleBinding":
			res = "clusterrolebindings"
		case "Deployment":
			res = "deployments"
		case "Service":
			res = "services"
		case "Secret":
			res = "secrets"
		case "ConfigMap":
			res = "configmaps"
		case "CronJob":
			res = "cronjobs"
		case "Job":
			res = "jobs"
		case "ValidatingWebhookConfiguration":
			res = "validatingwebhookconfigurations"
		case "MutatingWebhookConfiguration":
			res = "mutatingwebhookconfigurations"
		case "Certificate":
			res = "certificates"
		case "Issuer":
			res = "issuers"
		}

		var cProxy dynamic.ResourceInterface
		if resource.GetNamespace() != "" {
			cProxy = r.DynamicClient.Resource(schema.GroupVersionResource{Group: group, Version: version, Resource: res}).Namespace(resource.GetNamespace())
		} else {
			cProxy = r.DynamicClient.Resource(schema.GroupVersionResource{Group: group, Version: version, Resource: res})
		}

		_, err := cProxy.Apply(ctx, resource.GetName(), &resource, v1.ApplyOptions{FieldManager: "kubearchive-installer"})
		if err != nil {
			log.Error(err, "failed to apply resource", "resource", resource)
			_, err := cProxy.Create(ctx, &resource, v1.CreateOptions{})
			if err != nil {
				log.Error(err, "failed to create resource", "resource", resource)
				return ctrl.Result{}, err
			}
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *KubeArchiveInstallationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&kubearchiveorgv1.KubeArchiveInstallation{}).
		Named("kubearchiveinstallation").
		Complete(r)
}
