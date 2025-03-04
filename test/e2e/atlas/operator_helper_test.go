// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//go:build e2e || atlas

package atlas_test

import (
	"context"
	"strings"
	"testing"

	akov1 "github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/util/toptr"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

type operatorHelper struct {
	t         *testing.T
	k8sClient client.Client

	resourcesTracked []client.Object
	deployment       *appsv1.Deployment
}

func newOperatorHelper(t *testing.T) (*operatorHelper, error) {
	t.Helper()

	cfg, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	err = akov1.AddToScheme(scheme.Scheme)
	if err != nil {
		return nil, err
	}

	k8sClient, err := client.New(cfg, client.Options{Scheme: scheme.Scheme})
	if err != nil {
		return nil, err
	}

	return &operatorHelper{
		t:         t,
		k8sClient: k8sClient,
	}, nil
}

func (oh *operatorHelper) getK8sObject(key client.ObjectKey, object client.Object, track bool) error {
	err := oh.k8sClient.Get(context.Background(), key, object, &client.GetOptions{})
	if err != nil {
		return err
	}

	if track {
		oh.resourcesTracked = append(oh.resourcesTracked, object)
	}

	return nil
}

func (oh *operatorHelper) stopOperator() {
	deployment := appsv1.Deployment{}
	err := oh.getK8sObject(
		client.ObjectKey{Name: "mongodb-atlas-operator", Namespace: "mongodb-atlas-system"},
		&deployment,
		false,
	)
	if err != nil {
		oh.t.Errorf("unable to retrieve operator deployment: %v", err)
	}

	deployment.Spec.Replicas = toptr.MakePtr(int32(0))

	err = oh.k8sClient.Update(context.Background(), &deployment, &client.UpdateOptions{})
	if err != nil {
		oh.t.Errorf("unable to stop operator: %v", err)
	}
}

func (oh *operatorHelper) startOperator() {
	deployment := appsv1.Deployment{}
	err := oh.getK8sObject(
		client.ObjectKey{Name: "mongodb-atlas-operator", Namespace: "mongodb-atlas-system"},
		&deployment,
		false,
	)
	if err != nil {
		oh.t.Errorf("unable to retrieve operator deployment: %v", err)
	}

	deployment.Spec.Replicas = toptr.MakePtr(int32(1))

	err = oh.k8sClient.Update(context.Background(), &deployment, &client.UpdateOptions{})
	if err != nil {
		oh.t.Errorf("unable to start operator: %v", err)
	}
}

func (oh *operatorHelper) deleteOperator() {
	deployment := appsv1.Deployment{}
	err := oh.getK8sObject(
		client.ObjectKey{Name: "mongodb-atlas-operator", Namespace: "mongodb-atlas-system"},
		&deployment,
		false,
	)
	if err != nil {
		oh.t.Errorf("unable to retrieve operator deployment: %v", err)
	}

	oh.deployment = &deployment

	err = oh.k8sClient.Delete(context.Background(), &deployment, &client.DeleteOptions{})
	if err != nil {
		oh.t.Errorf("unable to delete operator: %v", err)
	}
}

func (oh *operatorHelper) restoreOperator() {
	if oh.deployment == nil {
		oh.t.Errorf("unable to restore operator. unknown previous state")
	}

	oh.deployment.ResourceVersion = ""
	err := oh.k8sClient.Create(context.Background(), oh.deployment, &client.CreateOptions{})
	if err != nil {
		oh.t.Errorf("unable to restore operator: %v", err)
	}
}

func (oh *operatorHelper) emulateCertifiedOperator() {
	deployment := appsv1.Deployment{}
	err := oh.getK8sObject(
		client.ObjectKey{Name: "mongodb-atlas-operator", Namespace: "mongodb-atlas-system"},
		&deployment,
		false,
	)
	if err != nil {
		oh.t.Errorf("unable to retrieve operator deployment: %v", err)
	}

	container := deployment.Spec.Template.Spec.Containers[0]
	container.Image = "quay.io/" + container.Image
	deployment.Spec.Template.Spec.Containers[0] = container

	err = oh.k8sClient.Update(context.Background(), &deployment, &client.UpdateOptions{})
	if err != nil {
		oh.t.Errorf("unable to emulate certified operator: %v", err)
	}
}

func (oh *operatorHelper) restoreOperatorImage() {
	deployment := appsv1.Deployment{}
	err := oh.getK8sObject(
		client.ObjectKey{Name: "mongodb-atlas-operator", Namespace: "mongodb-atlas-system"},
		&deployment,
		false,
	)
	if err != nil {
		oh.t.Errorf("unable to retrieve operator deployment: %v", err)
	}

	container := deployment.Spec.Template.Spec.Containers[0]
	container.Image = strings.Trim(container.Image, "quay.io/")
	deployment.Spec.Template.Spec.Containers[0] = container

	err = oh.k8sClient.Update(context.Background(), &deployment, &client.UpdateOptions{})
	if err != nil {
		oh.t.Errorf("unable to restore operator image: %v", err)
	}
}

func (oh *operatorHelper) cleanUpResources() {
	for _, object := range oh.resourcesTracked {
		if len(object.GetFinalizers()) > 0 {
			object.SetFinalizers([]string{})

			err := oh.k8sClient.Update(context.Background(), object, &client.UpdateOptions{})
			if err != nil {
				oh.t.Errorf("unable to update k8s resource: %v", err)
			}
		}

		err := oh.k8sClient.Delete(context.Background(), object)
		if err != nil {
			oh.t.Errorf("unable to delete k8s resource: %v", err)
		}
	}
}
