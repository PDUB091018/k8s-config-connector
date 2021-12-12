// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/compute/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeComputeV1beta1 struct {
	*testing.Fake
}

func (c *FakeComputeV1beta1) ComputeAddresses(namespace string) v1beta1.ComputeAddressInterface {
	return &FakeComputeAddresses{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeBackendBuckets(namespace string) v1beta1.ComputeBackendBucketInterface {
	return &FakeComputeBackendBuckets{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeBackendServices(namespace string) v1beta1.ComputeBackendServiceInterface {
	return &FakeComputeBackendServices{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeDisks(namespace string) v1beta1.ComputeDiskInterface {
	return &FakeComputeDisks{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeExternalVPNGateways(namespace string) v1beta1.ComputeExternalVPNGatewayInterface {
	return &FakeComputeExternalVPNGateways{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeFirewalls(namespace string) v1beta1.ComputeFirewallInterface {
	return &FakeComputeFirewalls{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeFirewallPolicies(namespace string) v1beta1.ComputeFirewallPolicyInterface {
	return &FakeComputeFirewallPolicies{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeFirewallPolicyRules(namespace string) v1beta1.ComputeFirewallPolicyRuleInterface {
	return &FakeComputeFirewallPolicyRules{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeForwardingRules(namespace string) v1beta1.ComputeForwardingRuleInterface {
	return &FakeComputeForwardingRules{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeHTTPHealthChecks(namespace string) v1beta1.ComputeHTTPHealthCheckInterface {
	return &FakeComputeHTTPHealthChecks{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeHTTPSHealthChecks(namespace string) v1beta1.ComputeHTTPSHealthCheckInterface {
	return &FakeComputeHTTPSHealthChecks{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeHealthChecks(namespace string) v1beta1.ComputeHealthCheckInterface {
	return &FakeComputeHealthChecks{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeImages(namespace string) v1beta1.ComputeImageInterface {
	return &FakeComputeImages{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeInstances(namespace string) v1beta1.ComputeInstanceInterface {
	return &FakeComputeInstances{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeInstanceGroups(namespace string) v1beta1.ComputeInstanceGroupInterface {
	return &FakeComputeInstanceGroups{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeInstanceGroupManagers(namespace string) v1beta1.ComputeInstanceGroupManagerInterface {
	return &FakeComputeInstanceGroupManagers{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeInstanceTemplates(namespace string) v1beta1.ComputeInstanceTemplateInterface {
	return &FakeComputeInstanceTemplates{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeInterconnectAttachments(namespace string) v1beta1.ComputeInterconnectAttachmentInterface {
	return &FakeComputeInterconnectAttachments{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeNetworks(namespace string) v1beta1.ComputeNetworkInterface {
	return &FakeComputeNetworks{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeNetworkEndpointGroups(namespace string) v1beta1.ComputeNetworkEndpointGroupInterface {
	return &FakeComputeNetworkEndpointGroups{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeNetworkPeerings(namespace string) v1beta1.ComputeNetworkPeeringInterface {
	return &FakeComputeNetworkPeerings{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeNodeGroups(namespace string) v1beta1.ComputeNodeGroupInterface {
	return &FakeComputeNodeGroups{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeNodeTemplates(namespace string) v1beta1.ComputeNodeTemplateInterface {
	return &FakeComputeNodeTemplates{c, namespace}
}

func (c *FakeComputeV1beta1) ComputePacketMirrorings(namespace string) v1beta1.ComputePacketMirroringInterface {
	return &FakeComputePacketMirrorings{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeProjectMetadatas(namespace string) v1beta1.ComputeProjectMetadataInterface {
	return &FakeComputeProjectMetadatas{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeReservations(namespace string) v1beta1.ComputeReservationInterface {
	return &FakeComputeReservations{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeResourcePolicies(namespace string) v1beta1.ComputeResourcePolicyInterface {
	return &FakeComputeResourcePolicies{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeRoutes(namespace string) v1beta1.ComputeRouteInterface {
	return &FakeComputeRoutes{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeRouters(namespace string) v1beta1.ComputeRouterInterface {
	return &FakeComputeRouters{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeRouterInterfaces(namespace string) v1beta1.ComputeRouterInterfaceInterface {
	return &FakeComputeRouterInterfaces{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeRouterNATs(namespace string) v1beta1.ComputeRouterNATInterface {
	return &FakeComputeRouterNATs{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeRouterPeers(namespace string) v1beta1.ComputeRouterPeerInterface {
	return &FakeComputeRouterPeers{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeSSLCertificates(namespace string) v1beta1.ComputeSSLCertificateInterface {
	return &FakeComputeSSLCertificates{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeSSLPolicies(namespace string) v1beta1.ComputeSSLPolicyInterface {
	return &FakeComputeSSLPolicies{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeSecurityPolicies(namespace string) v1beta1.ComputeSecurityPolicyInterface {
	return &FakeComputeSecurityPolicies{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeServiceAttachments(namespace string) v1beta1.ComputeServiceAttachmentInterface {
	return &FakeComputeServiceAttachments{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeSharedVPCHostProjects(namespace string) v1beta1.ComputeSharedVPCHostProjectInterface {
	return &FakeComputeSharedVPCHostProjects{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeSharedVPCServiceProjects(namespace string) v1beta1.ComputeSharedVPCServiceProjectInterface {
	return &FakeComputeSharedVPCServiceProjects{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeSnapshots(namespace string) v1beta1.ComputeSnapshotInterface {
	return &FakeComputeSnapshots{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeSubnetworks(namespace string) v1beta1.ComputeSubnetworkInterface {
	return &FakeComputeSubnetworks{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeTargetGRPCProxies(namespace string) v1beta1.ComputeTargetGRPCProxyInterface {
	return &FakeComputeTargetGRPCProxies{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeTargetHTTPProxies(namespace string) v1beta1.ComputeTargetHTTPProxyInterface {
	return &FakeComputeTargetHTTPProxies{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeTargetHTTPSProxies(namespace string) v1beta1.ComputeTargetHTTPSProxyInterface {
	return &FakeComputeTargetHTTPSProxies{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeTargetInstances(namespace string) v1beta1.ComputeTargetInstanceInterface {
	return &FakeComputeTargetInstances{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeTargetPools(namespace string) v1beta1.ComputeTargetPoolInterface {
	return &FakeComputeTargetPools{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeTargetSSLProxies(namespace string) v1beta1.ComputeTargetSSLProxyInterface {
	return &FakeComputeTargetSSLProxies{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeTargetTCPProxies(namespace string) v1beta1.ComputeTargetTCPProxyInterface {
	return &FakeComputeTargetTCPProxies{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeTargetVPNGateways(namespace string) v1beta1.ComputeTargetVPNGatewayInterface {
	return &FakeComputeTargetVPNGateways{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeURLMaps(namespace string) v1beta1.ComputeURLMapInterface {
	return &FakeComputeURLMaps{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeVPNGateways(namespace string) v1beta1.ComputeVPNGatewayInterface {
	return &FakeComputeVPNGateways{c, namespace}
}

func (c *FakeComputeV1beta1) ComputeVPNTunnels(namespace string) v1beta1.ComputeVPNTunnelInterface {
	return &FakeComputeVPNTunnels{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeComputeV1beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
