package common

import (
	bcs "github.com/kakaoenterprise/kc-sdk-go/services/bcs"
	iam "github.com/kakaoenterprise/kc-sdk-go/services/iam"
	image "github.com/kakaoenterprise/kc-sdk-go/services/image"
	kubernetesengine "github.com/kakaoenterprise/kc-sdk-go/services/kubernetesengine"
	loadbalancer "github.com/kakaoenterprise/kc-sdk-go/services/loadbalancer"
	network "github.com/kakaoenterprise/kc-sdk-go/services/network"
	volume "github.com/kakaoenterprise/kc-sdk-go/services/volume"
	vpc "github.com/kakaoenterprise/kc-sdk-go/services/vpc"
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

// NewAPIClient : Config(Token/UserAgent 포함)을 받아
// 모든 서비스에 자동 헤더 주입이 적용된 http.Client로 초기화
func NewAPIClient(cfg Config) *APIClient {
	c := &APIClient{cfg: cfg}

	authedClient, authRT := newAuthedHTTPClient(cfg.HTTPClient, cfg.Token, cfg.UserAgent)
	c.authRT = authRT

	{
		cc := iam.NewConfiguration()
		cc.HTTPClient = authedClient
		cc.Servers = iam.ServerConfigurations{{URL: cfg.Endpoints.IAM}}

		cli := iam.NewAPIClient(cc)

		c.IdentityAPI = cli.IdentityAPI
	}

	{
		cc := bcs.NewConfiguration()
		cc.HTTPClient = authedClient
		cc.Servers = bcs.ServerConfigurations{{URL: cfg.Endpoints.BCS}}

		cli := bcs.NewAPIClient(cc)

		c.FlavorAPI = cli.FlavorAPI
		c.InstanceAPI = cli.InstanceAPI
		c.InstanceActionsAPI = cli.InstanceActionsAPI
		c.InstanceAttachedVolumeAPI = cli.InstanceAttachedVolumeAPI
		c.InstanceNetworkInterfaceAPI = cli.InstanceNetworkInterfaceAPI
		c.InstancePasswordAPI = cli.InstancePasswordAPI
		c.InstancePublicIPAPI = cli.InstancePublicIPAPI
		c.InstanceRunAnActionAPI = cli.InstanceRunAnActionAPI
		c.InstanceSecurityGroupAPI = cli.InstanceSecurityGroupAPI
		c.KeypairAPI = cli.KeypairAPI
	}

	{
		cc := image.NewConfiguration()
		cc.HTTPClient = authedClient
		cc.Servers = image.ServerConfigurations{{URL: cfg.Endpoints.Image}}

		cli := image.NewAPIClient(cc)

		c.ImageAPI = cli.ImageAPI
	}

	{
		cc := network.NewConfiguration()
		cc.HTTPClient = authedClient
		cc.Servers = network.ServerConfigurations{{URL: cfg.Endpoints.Network}}

		cli := network.NewAPIClient(cc)

		c.PublicIPAPI = cli.PublicIPAPI
		c.SecurityGroupAPI = cli.SecurityGroupAPI
	}

	{
		cc := vpc.NewConfiguration()
		cc.HTTPClient = authedClient
		cc.Servers = vpc.ServerConfigurations{{URL: cfg.Endpoints.VPC}}

		cli := vpc.NewAPIClient(cc)

		c.NetworkInterfaceAPI = cli.NetworkInterfaceAPI
		c.VPCAPI = cli.VPCAPI
		c.VPCRouteTableAPI = cli.VPCRouteTableAPI
		c.VPCRouteTableAssociationAPI = cli.VPCRouteTableAssociationAPI
		c.VPCRouteTableRouteAPI = cli.VPCRouteTableRouteAPI
		c.VPCSubnetAPI = cli.VPCSubnetAPI
	}

	{
		cc := volume.NewConfiguration()
		cc.HTTPClient = authedClient
		cc.Servers = volume.ServerConfigurations{{URL: cfg.Endpoints.Volume}}

		cli := volume.NewAPIClient(cc)

		c.VolumeAPI = cli.VolumeAPI
		c.VolumeSnapshotAPI = cli.VolumeSnapshotAPI
	}

	{
		cc := loadbalancer.NewConfiguration()
		cc.HTTPClient = authedClient
		cc.Servers = loadbalancer.ServerConfigurations{{URL: cfg.Endpoints.LoadBalancer}}

		cli := loadbalancer.NewAPIClient(cc)

		c.BeyondLoadBalancerAPI = cli.BeyondLoadBalancerAPI
		c.LoadBalancerAPI = cli.LoadBalancerAPI
		c.LoadBalancerEtcAPI = cli.LoadBalancerEtcAPI
		c.LoadBalancerL7PoliciesAPI = cli.LoadBalancerL7PoliciesAPI
		c.LoadBalancerListenerAPI = cli.LoadBalancerListenerAPI
		c.LoadBalancerTargetGroupAPI = cli.LoadBalancerTargetGroupAPI

	}

	{
		cc := kubernetesengine.NewConfiguration()
		cc.HTTPClient = authedClient
		cc.Servers = kubernetesengine.ServerConfigurations{{URL: cfg.Endpoints.KubernetesEngine}}

		cli := kubernetesengine.NewAPIClient(cc)

		c.ClustersAPI = cli.ClustersAPI
		c.ImagesAPI = cli.ImagesAPI
		c.NodePoolsAPI = cli.NodePoolsAPI
		c.NodesAPI = cli.NodesAPI
		c.ScalingAPI = cli.ScalingAPI
		c.ServiceAgentsAPI = cli.ServiceAgentsAPI
		c.UpgradesAPI = cli.UpgradesAPI

	}

	return c
}

// SetToken 토큰 런타임 교체
func (c *APIClient) SetToken(tok string) {
	if c.authRT != nil {
		c.authRT.token.Store(tok)
	}
}

// SetUserAgent User-Agent 런타임 교체(선택)
func (c *APIClient) SetUserAgent(ua string) {
	if c.authRT != nil {
		c.authRT.userAgent.Store(ua)
	}
}
