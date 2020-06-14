/*
Copyright 2017 The Kubernetes Authors.

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
package provisioner

import (
	"fmt"
	// "path/filepath"
	// "strconv"
	// "strings"

	"github.com/kubernetes-sigs/sig-storage-lib-external-provisioner/controller"
	corev1 "k8s.io/api/core/v1"
	// "k8s.io/apimachinery/pkg/api/errors"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/klog"
)

// NewZvolIscsiProvisioner creates a new glusterfs simple provisioner
func NewZvolISCSIProvisioner(config *rest.Config, client kubernetes.Interface) controller.Provisioner {
	klog.Infof("Creating New ZvolIscsiProvisioner.")
	return newZvolISCSIProvisionerInternal(config, client)
}

func newZvolISCSIProvisionerInternal(config *rest.Config, client kubernetes.Interface) *zvolIscsiProvisioner {
	var identity types.UID

	restClient := client.CoreV1().RESTClient()
	provisioner := &zvolIscsiProvisioner{
		config:     config,
		client:     client,
		restClient: restClient,
		identity:   identity,
	}

	return provisioner
}

type zvolIscsiProvisioner struct {
	client     kubernetes.Interface
	restClient rest.Interface
	config     *rest.Config
	identity   types.UID
}

var _ controller.Provisioner = &zvolIscsiProvisioner{}

func (p *zvolIscsiProvisioner) Provision(options controller.ProvisionOptions) (*corev1.PersistentVolume, error) {
	if options.PVC.Spec.Selector != nil {
		return nil, fmt.Errorf("claim Selector is not supported")
	}
	klog.V(4).Infof("Start Provisioning volume: ProvisionOptions %v", options)
	return nil, nil
}
func (p *zvolIscsiProvisioner) Delete(volume *corev1.PersistentVolume) error {
	return nil
}
