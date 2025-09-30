package common

import (
	bcs "github.kakaoenterprise.in/kic2/kc-sdk-go/services/bcs"
	iam "github.kakaoenterprise.in/kic2/kc-sdk-go/services/iam"
	image "github.kakaoenterprise.in/kic2/kc-sdk-go/services/image"
	kubernetesengine "github.kakaoenterprise.in/kic2/kc-sdk-go/services/kubernetesengine"
	loadbalancer "github.kakaoenterprise.in/kic2/kc-sdk-go/services/loadbalancer"
	network "github.kakaoenterprise.in/kic2/kc-sdk-go/services/network"
	volume "github.kakaoenterprise.in/kic2/kc-sdk-go/services/volume"
	vpc "github.kakaoenterprise.in/kic2/kc-sdk-go/services/vpc"
)

type APIClient struct {
	cfg    Config
	authRT *xAuthTokenTransport

	// auth
	IdentityAPI iam.IdentityAPI

	// bcs
	FlavorAPI                   bcs.FlavorAPI
	InstanceAPI                 bcs.InstanceAPI
	InstanceActionsAPI          bcs.InstanceActionsAPI
	InstanceAttachedVolumeAPI   bcs.InstanceAttachedVolumeAPI
	InstanceNetworkInterfaceAPI bcs.InstanceNetworkInterfaceAPI
	InstancePasswordAPI         bcs.InstancePasswordAPI
	InstancePublicIPAPI         bcs.InstancePublicIPAPI
	InstanceRunAnActionAPI      bcs.InstanceRunAnActionAPI
	InstanceSecurityGroupAPI    bcs.InstanceSecurityGroupAPI
	KeypairAPI                  bcs.KeypairAPI

	// image
	ImageAPI image.ImageAPI

	// loadbalancer
	BeyondLoadBalancerAPI      loadbalancer.BeyondLoadBalancerAPI
	LoadBalancerAPI            loadbalancer.LoadBalancerAPI
	LoadBalancerEtcAPI         loadbalancer.LoadBalancerEtcAPI
	LoadBalancerL7PoliciesAPI  loadbalancer.LoadBalancerL7PoliciesAPI
	LoadBalancerListenerAPI    loadbalancer.LoadBalancerListenerAPI
	LoadBalancerTargetGroupAPI loadbalancer.LoadBalancerTargetGroupAPI

	// network
	PublicIPAPI      network.PublicIPAPI
	SecurityGroupAPI network.SecurityGroupAPI

	// volume
	VolumeAPI         volume.VolumeAPI
	VolumeSnapshotAPI volume.VolumeSnapshotAPI

	// vpc
	NetworkInterfaceAPI         vpc.NetworkInterfaceAPI
	VPCAPI                      vpc.VPCAPI
	VPCRouteTableAPI            vpc.VPCRouteTableAPI
	VPCRouteTableAssociationAPI vpc.VPCRouteTableAssociationAPI
	VPCRouteTableRouteAPI       vpc.VPCRouteTableRouteAPI
	VPCSubnetAPI                vpc.VPCSubnetAPI

	// kubernetes engine
	ClustersAPI      kubernetesengine.ClustersAPI
	ImagesAPI        kubernetesengine.ImagesAPI
	NodePoolsAPI     kubernetesengine.NodePoolsAPI
	NodesAPI         kubernetesengine.NodesAPI
	ScalingAPI       kubernetesengine.ScalingAPI
	ServiceAgentsAPI kubernetesengine.ServiceAgentsAPI
	UpgradesAPI      kubernetesengine.UpgradesAPI
}
