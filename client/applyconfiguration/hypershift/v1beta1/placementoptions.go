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

// PlacementOptionsApplyConfiguration represents a declarative configuration of the PlacementOptions type for use
// with apply.
type PlacementOptionsApplyConfiguration struct {
	Tenancy *string `json:"tenancy,omitempty"`
}

// PlacementOptionsApplyConfiguration constructs a declarative configuration of the PlacementOptions type for use with
// apply.
func PlacementOptions() *PlacementOptionsApplyConfiguration {
	return &PlacementOptionsApplyConfiguration{}
}

// WithTenancy sets the Tenancy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Tenancy field is set to the value of the last call.
func (b *PlacementOptionsApplyConfiguration) WithTenancy(value string) *PlacementOptionsApplyConfiguration {
	b.Tenancy = &value
	return b
}
