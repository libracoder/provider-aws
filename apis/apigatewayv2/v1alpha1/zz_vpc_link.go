/*
Copyright 2020 The Crossplane Authors.

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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// VPCLinkParameters defines the desired state of VPCLink
type VPCLinkParameters struct {
	// Region is which region the VPCLink will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`

	// +kubebuilder:validation:Required
	Name *string `json:"name"`

	Tags                    map[string]*string `json:"tags,omitempty"`
	CustomVPCLinkParameters `json:",inline"`
}

// VPCLinkSpec defines the desired state of VPCLink
type VPCLinkSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  VPCLinkParameters `json:"forProvider"`
}

// VPCLinkObservation defines the observed state of VPCLink
type VPCLinkObservation struct {
	CreatedDate *metav1.Time `json:"createdDate,omitempty"`

	SecurityGroupIDs []*string `json:"securityGroupIDs,omitempty"`

	SubnetIDs []*string `json:"subnetIDs,omitempty"`

	VPCLinkID *string `json:"vpcLinkID,omitempty"`

	VPCLinkStatus *string `json:"vpcLinkStatus,omitempty"`

	VPCLinkStatusMessage *string `json:"vpcLinkStatusMessage,omitempty"`

	VPCLinkVersion *string `json:"vpcLinkVersion,omitempty"`
}

// VPCLinkStatus defines the observed state of VPCLink.
type VPCLinkStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     VPCLinkObservation `json:"atProvider"`
}

// +kubebuilder:object:root=true

// VPCLink is the Schema for the VPCLinks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCLink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCLinkSpec   `json:"spec,omitempty"`
	Status            VPCLinkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCLinkList contains a list of VPCLinks
type VPCLinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCLink `json:"items"`
}

// Repository type metadata.
var (
	VPCLinkKind             = "VPCLink"
	VPCLinkGroupKind        = schema.GroupKind{Group: Group, Kind: VPCLinkKind}.String()
	VPCLinkKindAPIVersion   = VPCLinkKind + "." + GroupVersion.String()
	VPCLinkGroupVersionKind = GroupVersion.WithKind(VPCLinkKind)
)

func init() {
	SchemeBuilder.Register(&VPCLink{}, &VPCLinkList{})
}
