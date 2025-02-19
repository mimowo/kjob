/*
Copyright 2024 The Kubernetes Authors.

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
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	v1alpha1 "sigs.k8s.io/kjob/apis/v1alpha1"
	scheme "sigs.k8s.io/kjob/client-go/clientset/versioned/scheme"
)

// JobTemplatesGetter has a method to return a JobTemplateInterface.
// A group's client should implement this interface.
type JobTemplatesGetter interface {
	JobTemplates(namespace string) JobTemplateInterface
}

// JobTemplateInterface has methods to work with JobTemplate resources.
type JobTemplateInterface interface {
	Create(ctx context.Context, jobTemplate *v1alpha1.JobTemplate, opts v1.CreateOptions) (*v1alpha1.JobTemplate, error)
	Update(ctx context.Context, jobTemplate *v1alpha1.JobTemplate, opts v1.UpdateOptions) (*v1alpha1.JobTemplate, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.JobTemplate, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.JobTemplateList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.JobTemplate, err error)
	JobTemplateExpansion
}

// jobTemplates implements JobTemplateInterface
type jobTemplates struct {
	*gentype.ClientWithList[*v1alpha1.JobTemplate, *v1alpha1.JobTemplateList]
}

// newJobTemplates returns a JobTemplates
func newJobTemplates(c *KjobctlV1alpha1Client, namespace string) *jobTemplates {
	return &jobTemplates{
		gentype.NewClientWithList[*v1alpha1.JobTemplate, *v1alpha1.JobTemplateList](
			"jobtemplates",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.JobTemplate { return &v1alpha1.JobTemplate{} },
			func() *v1alpha1.JobTemplateList { return &v1alpha1.JobTemplateList{} }),
	}
}
