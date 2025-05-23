/*


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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

// AWSRolesRefApplyConfiguration represents a declarative configuration of the AWSRolesRef type for use
// with apply.
type AWSRolesRefApplyConfiguration struct {
	IngressARN              *string `json:"ingressARN,omitempty"`
	ImageRegistryARN        *string `json:"imageRegistryARN,omitempty"`
	StorageARN              *string `json:"storageARN,omitempty"`
	NetworkARN              *string `json:"networkARN,omitempty"`
	KubeCloudControllerARN  *string `json:"kubeCloudControllerARN,omitempty"`
	NodePoolManagementARN   *string `json:"nodePoolManagementARN,omitempty"`
	ControlPlaneOperatorARN *string `json:"controlPlaneOperatorARN,omitempty"`
}

// AWSRolesRefApplyConfiguration constructs a declarative configuration of the AWSRolesRef type for use with
// apply.
func AWSRolesRef() *AWSRolesRefApplyConfiguration {
	return &AWSRolesRefApplyConfiguration{}
}

// WithIngressARN sets the IngressARN field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IngressARN field is set to the value of the last call.
func (b *AWSRolesRefApplyConfiguration) WithIngressARN(value string) *AWSRolesRefApplyConfiguration {
	b.IngressARN = &value
	return b
}

// WithImageRegistryARN sets the ImageRegistryARN field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ImageRegistryARN field is set to the value of the last call.
func (b *AWSRolesRefApplyConfiguration) WithImageRegistryARN(value string) *AWSRolesRefApplyConfiguration {
	b.ImageRegistryARN = &value
	return b
}

// WithStorageARN sets the StorageARN field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StorageARN field is set to the value of the last call.
func (b *AWSRolesRefApplyConfiguration) WithStorageARN(value string) *AWSRolesRefApplyConfiguration {
	b.StorageARN = &value
	return b
}

// WithNetworkARN sets the NetworkARN field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NetworkARN field is set to the value of the last call.
func (b *AWSRolesRefApplyConfiguration) WithNetworkARN(value string) *AWSRolesRefApplyConfiguration {
	b.NetworkARN = &value
	return b
}

// WithKubeCloudControllerARN sets the KubeCloudControllerARN field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KubeCloudControllerARN field is set to the value of the last call.
func (b *AWSRolesRefApplyConfiguration) WithKubeCloudControllerARN(value string) *AWSRolesRefApplyConfiguration {
	b.KubeCloudControllerARN = &value
	return b
}

// WithNodePoolManagementARN sets the NodePoolManagementARN field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodePoolManagementARN field is set to the value of the last call.
func (b *AWSRolesRefApplyConfiguration) WithNodePoolManagementARN(value string) *AWSRolesRefApplyConfiguration {
	b.NodePoolManagementARN = &value
	return b
}

// WithControlPlaneOperatorARN sets the ControlPlaneOperatorARN field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ControlPlaneOperatorARN field is set to the value of the last call.
func (b *AWSRolesRefApplyConfiguration) WithControlPlaneOperatorARN(value string) *AWSRolesRefApplyConfiguration {
	b.ControlPlaneOperatorARN = &value
	return b
}
