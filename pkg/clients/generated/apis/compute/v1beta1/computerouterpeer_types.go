// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RouterpeerAdvertisedIpRanges struct {
	/* User-specified description for the IP range. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* The IP range to advertise. The value must be a
	CIDR-formatted string. */
	Range string `json:"range"`
}

type RouterpeerBfd struct {
	/* The minimum interval, in milliseconds, between BFD control packets
	received from the peer router. The actual value is negotiated
	between the two routers and is equal to the greater of this value
	and the transmit interval of the other router. If set, this value
	must be between 1000 and 30000. */
	// +optional
	MinReceiveInterval *int `json:"minReceiveInterval,omitempty"`

	/* The minimum interval, in milliseconds, between BFD control packets
	transmitted to the peer router. The actual value is negotiated
	between the two routers and is equal to the greater of this value
	and the corresponding receive interval of the other router. If set,
	this value must be between 1000 and 30000. */
	// +optional
	MinTransmitInterval *int `json:"minTransmitInterval,omitempty"`

	/* The number of consecutive BFD packets that must be missed before
	BFD declares that a peer is unavailable. If set, the value must
	be a value between 5 and 16. */
	// +optional
	Multiplier *int `json:"multiplier,omitempty"`

	/* The BFD session initialization mode for this BGP peer.
	If set to 'ACTIVE', the Cloud Router will initiate the BFD session
	for this BGP peer. If set to 'PASSIVE', the Cloud Router will wait
	for the peer router to initiate the BFD session for this BGP peer.
	If set to 'DISABLED', BFD is disabled for this BGP peer. Possible values: ["ACTIVE", "DISABLED", "PASSIVE"]. */
	SessionInitializationMode string `json:"sessionInitializationMode"`
}

type RouterpeerIpAddress struct {
	/*  */
	// +optional
	External *string `json:"external,omitempty"`
}

type ComputeRouterPeerSpec struct {
	/* User-specified flag to indicate which mode to use for advertisement.
	Valid values of this enum field are: 'DEFAULT', 'CUSTOM' Default value: "DEFAULT" Possible values: ["DEFAULT", "CUSTOM"]. */
	// +optional
	AdvertiseMode *string `json:"advertiseMode,omitempty"`

	/* User-specified list of prefix groups to advertise in custom
	mode, which can take one of the following options:

	* 'ALL_SUBNETS': Advertises all available subnets, including peer VPC subnets.
	* 'ALL_VPC_SUBNETS': Advertises the router's own VPC subnets.
	* 'ALL_PEER_VPC_SUBNETS': Advertises peer subnets of the router's VPC network.


	Note that this field can only be populated if advertiseMode is 'CUSTOM'
	and overrides the list defined for the router (in the "bgp" message).
	These groups are advertised in addition to any specified prefixes.
	Leave this field blank to advertise no custom groups. */
	// +optional
	AdvertisedGroups []string `json:"advertisedGroups,omitempty"`

	/* User-specified list of individual IP ranges to advertise in
	custom mode. This field can only be populated if advertiseMode
	is 'CUSTOM' and is advertised to all peers of the router. These IP
	ranges will be advertised in addition to any specified groups.
	Leave this field blank to advertise no custom IP ranges. */
	// +optional
	AdvertisedIpRanges []RouterpeerAdvertisedIpRanges `json:"advertisedIpRanges,omitempty"`

	/* The priority of routes advertised to this BGP peer.
	Where there is more than one matching route of maximum
	length, the routes with the lowest priority value win. */
	// +optional
	AdvertisedRoutePriority *int `json:"advertisedRoutePriority,omitempty"`

	/* BFD configuration for the BGP peering. */
	// +optional
	Bfd *RouterpeerBfd `json:"bfd,omitempty"`

	/* The status of the BGP peer connection. If set to false, any active session
	with the peer is terminated and all associated routing information is removed.
	If set to true, the peer connection can be established with routing information.
	The default is true. */
	// +optional
	Enable *bool `json:"enable,omitempty"`

	/* IP address of the interface inside Google Cloud Platform.
	Only IPv4 is supported. */
	// +optional
	IpAddress *RouterpeerIpAddress `json:"ipAddress,omitempty"`

	/* Peer BGP Autonomous System Number (ASN).
	Each BGP interface may use a different value. */
	PeerAsn int `json:"peerAsn"`

	/* IP address of the BGP interface outside Google Cloud Platform.
	Only IPv4 is supported. */
	PeerIpAddress string `json:"peerIpAddress"`

	/* Immutable. Region where the router and BgpPeer reside.
	If it is not provided, the provider region is used. */
	Region string `json:"region"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* The URI of the VM instance that is used as third-party router
	appliances such as Next Gen Firewalls, Virtual Routers, or Router
	Appliances. The VM instance must be located in zones contained in
	the same region as this Cloud Router. The VM instance is the peer
	side of the BGP session. */
	// +optional
	RouterApplianceInstanceRef *v1alpha1.ResourceRef `json:"routerApplianceInstanceRef,omitempty"`

	/* The interface the BGP peer is associated with. */
	RouterInterfaceRef v1alpha1.ResourceRef `json:"routerInterfaceRef"`

	/* The Cloud Router in which this BGP peer will be configured. */
	RouterRef v1alpha1.ResourceRef `json:"routerRef"`
}

type ComputeRouterPeerStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeRouterPeer's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The resource that configures and manages this BGP peer.

	* 'MANAGED_BY_USER' is the default value and can be managed by
	you or other users
	* 'MANAGED_BY_ATTACHMENT' is a BGP peer that is configured and
	managed by Cloud Interconnect, specifically by an
	InterconnectAttachment of type PARTNER. Google automatically
	creates, updates, and deletes this type of BGP peer when the
	PARTNER InterconnectAttachment is created, updated,
	or deleted. */
	ManagementType string `json:"managementType,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeRouterPeer is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeRouterPeer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeRouterPeerSpec   `json:"spec,omitempty"`
	Status ComputeRouterPeerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeRouterPeerList contains a list of ComputeRouterPeer
type ComputeRouterPeerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeRouterPeer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeRouterPeer{}, &ComputeRouterPeerList{})
}
