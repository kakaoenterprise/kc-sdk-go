package common

import (
	bcs "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
	"github.com/kakaoenterprise/kc-sdk-go/v2/services/config"
	iam "github.com/kakaoenterprise/kc-sdk-go/v2/services/iam"
	image "github.com/kakaoenterprise/kc-sdk-go/v2/services/image"
	kms "github.com/kakaoenterprise/kc-sdk-go/v2/services/kms"
	kubernetesengine "github.com/kakaoenterprise/kc-sdk-go/v2/services/kubernetesengine"
	loadbalancer "github.com/kakaoenterprise/kc-sdk-go/v2/services/loadbalancer"
	mysql "github.com/kakaoenterprise/kc-sdk-go/v2/services/mysql"
	network "github.com/kakaoenterprise/kc-sdk-go/v2/services/network"
	secretsmanager "github.com/kakaoenterprise/kc-sdk-go/v2/services/secretsmanager"
	tgw "github.com/kakaoenterprise/kc-sdk-go/v2/services/tgw"
	volume "github.com/kakaoenterprise/kc-sdk-go/v2/services/volume"
	vpc "github.com/kakaoenterprise/kc-sdk-go/v2/services/vpc"
)

type APIClient struct {
	cfg    Config
	authRT *xAuthTokenTransport

	// auth
	IdentityAPI iam.IdentityAPI

	// config
	ConfigAPI config.PublicAPI

	// bcs
	InstanceAPI bcs.InstanceAPI
	KeypairAPI  bcs.KeypairAPI

	// image
	ImageAPI image.ImageAPI

	// loadbalancer
	LoadBalancerHealthMonitorAPI         loadbalancer.HealthMonitorAPI
	LoadBalancerHighAvailabilityGroupAPI loadbalancer.HighAvailabilityGroupAPI
	LoadBalancerL7PolicyAPI              loadbalancer.L7PolicyAPI
	LoadBalancerL7RuleAPI                loadbalancer.L7RuleAPI
	LoadBalancerListenerAPI              loadbalancer.ListenerAPI
	LoadBalancerAPI                      loadbalancer.LoadBalancerAPI
	LoadBalancerTargetGroupAPI           loadbalancer.TargetGroupAPI

	// network
	NetworkInterfaceAPI     network.NetworkInterfaceAPI
	NetworkPublicIPAPI      network.PublicIPAPI
	NetworkSecurityGroupAPI network.SecurityGroupAPI

	// volume
	VolumeImageAPI    volume.ImageAPI
	VolumeSnapshotAPI volume.SnapshotAPI
	VolumeAPI         volume.VolumeAPI

	// vpc
	VPCInternetGatewayAPI  vpc.InternetGatewayAPI
	VPCNetworkInterfaceAPI vpc.NetworkInterfaceAPI
	VPCSubnetAPI           vpc.SubnetAPI
	VPCAPI                 vpc.VPCAPI
	VPCRouteTableAPI       vpc.VPCRouteTableAPI

	// kubernetes engine
	KubernetesEngineAPI         kubernetesengine.KubernetesEngineAPI
	KubernetesEngineClusterAPI  kubernetesengine.KubernetesEngineClusterAPI
	KubernetesEngineNodeAPI     kubernetesengine.NodeAPI
	KubernetesEngineNodePoolAPI kubernetesengine.NodePoolAPI

	// kms
	KMSAPI        kms.KMSAPI
	KeyAPI        kms.KeyAPI
	KeyVersionAPI kms.KeyVersionAPI
	PublicKeyAPI  kms.PublicKeyAPI

	// tgw
	TransitGatewayAPI           tgw.TransitGatewayAPI
	TransitGatewayAttachmentAPI tgw.TransitGatewayAttachmentAPI
	TransitGatewayRouteTableAPI tgw.TransitGatewayRouteTableAPI

	// mysql
	MySQLBackupAPI                mysql.BackupAPI
	MySQLCustomParameterGroupAPI  mysql.CustomParameterGroupAPI
	MySQLDefaultParameterGroupAPI mysql.DefaultParameterGroupAPI
	MySQLEngineVersionAPI         mysql.EngineVersionAPI
	MySQLInstanceAPI              mysql.InstanceAPI
	MySQLInstanceGroupAPI         mysql.InstanceGroupAPI

	// secrets manager
	SecretAPI         secretsmanager.SecretAPI
	SecretVersionAPI  secretsmanager.SecretVersionAPI
	SecretsManagerAPI secretsmanager.SecretsManagerAPI
}

func NewAPIClient(cfg Config) *APIClient {
	c := &APIClient{cfg: cfg}

	authedClient, authRT := newAuthedHTTPClient(cfg.HTTPClient, cfg.Token, cfg.UserAgent, cfg.Version)
	c.authRT = authRT

	{
		cc := iam.NewConfiguration()
		cc.HTTPClient = authedClient
		cc.Servers = iam.ServerConfigurations{{URL: cfg.Endpoints.IAM}}

		cli := iam.NewAPIClient(cc)

		c.IdentityAPI = cli.IdentityAPI
	}

	{
		cc := config.NewConfiguration()
		cc.HTTPClient = authedClient
		cc.Servers = config.ServerConfigurations{{URL: cfg.Endpoints.Config}}

		cli := config.NewAPIClient(cc)

		c.ConfigAPI = cli.PublicAPI
	}

	{
		cc := bcs.NewConfiguration()
		cc.HTTPClient = authedClient
		cc.Servers = bcs.ServerConfigurations{{URL: cfg.Endpoints.BCS}}

		cli := bcs.NewAPIClient(cc)

		c.InstanceAPI = cli.InstanceAPI
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

		c.NetworkInterfaceAPI = cli.NetworkInterfaceAPI
		c.NetworkPublicIPAPI = cli.PublicIPAPI
		c.NetworkSecurityGroupAPI = cli.SecurityGroupAPI
	}

	{
		cc := vpc.NewConfiguration()
		cc.HTTPClient = authedClient
		cc.Servers = vpc.ServerConfigurations{{URL: cfg.Endpoints.VPC}}

		cli := vpc.NewAPIClient(cc)

		c.VPCInternetGatewayAPI = cli.InternetGatewayAPI
		c.VPCNetworkInterfaceAPI = cli.NetworkInterfaceAPI
		c.VPCSubnetAPI = cli.SubnetAPI
		c.VPCAPI = cli.VPCAPI
		c.VPCRouteTableAPI = cli.VPCRouteTableAPI
	}

	{
		cc := volume.NewConfiguration()
		cc.HTTPClient = authedClient
		cc.Servers = volume.ServerConfigurations{{URL: cfg.Endpoints.Volume}}

		cli := volume.NewAPIClient(cc)

		c.VolumeImageAPI = cli.ImageAPI
		c.VolumeSnapshotAPI = cli.SnapshotAPI
		c.VolumeAPI = cli.VolumeAPI
	}

	{
		cc := loadbalancer.NewConfiguration()
		cc.HTTPClient = authedClient
		cc.Servers = loadbalancer.ServerConfigurations{{URL: cfg.Endpoints.LoadBalancer}}

		cli := loadbalancer.NewAPIClient(cc)

		c.LoadBalancerHealthMonitorAPI = cli.HealthMonitorAPI
		c.LoadBalancerHighAvailabilityGroupAPI = cli.HighAvailabilityGroupAPI
		c.LoadBalancerL7PolicyAPI = cli.L7PolicyAPI
		c.LoadBalancerL7RuleAPI = cli.L7RuleAPI
		c.LoadBalancerListenerAPI = cli.ListenerAPI
		c.LoadBalancerAPI = cli.LoadBalancerAPI
		c.LoadBalancerTargetGroupAPI = cli.TargetGroupAPI
	}

	{
		cc := kubernetesengine.NewConfiguration()
		cc.HTTPClient = authedClient
		cc.Servers = kubernetesengine.ServerConfigurations{{URL: cfg.Endpoints.KubernetesEngine}}

		cli := kubernetesengine.NewAPIClient(cc)

		c.KubernetesEngineAPI = cli.KubernetesEngineAPI
		c.KubernetesEngineClusterAPI = cli.KubernetesEngineClusterAPI
		c.KubernetesEngineNodeAPI = cli.NodeAPI
		c.KubernetesEngineNodePoolAPI = cli.NodePoolAPI
	}

	{
		cc := kms.NewConfiguration()
		cc.HTTPClient = authedClient
		cc.Servers = kms.ServerConfigurations{{URL: cfg.Endpoints.KMS}}

		cli := kms.NewAPIClient(cc)

		c.KMSAPI = cli.KMSAPI
		c.KeyAPI = cli.KeyAPI
		c.KeyVersionAPI = cli.KeyVersionAPI
		c.PublicKeyAPI = cli.PublicKeyAPI
	}

	{
		cc := tgw.NewConfiguration()
		cc.HTTPClient = authedClient
		cc.Servers = tgw.ServerConfigurations{{URL: cfg.Endpoints.TGW}}

		cli := tgw.NewAPIClient(cc)

		c.TransitGatewayAPI = cli.TransitGatewayAPI
		c.TransitGatewayAttachmentAPI = cli.TransitGatewayAttachmentAPI
		c.TransitGatewayRouteTableAPI = cli.TransitGatewayRouteTableAPI
	}

	{
		cc := mysql.NewConfiguration()
		cc.HTTPClient = authedClient
		cc.Servers = mysql.ServerConfigurations{{URL: cfg.Endpoints.MySQL}}

		cli := mysql.NewAPIClient(cc)

		c.MySQLBackupAPI = cli.BackupAPI
		c.MySQLCustomParameterGroupAPI = cli.CustomParameterGroupAPI
		c.MySQLDefaultParameterGroupAPI = cli.DefaultParameterGroupAPI
		c.MySQLEngineVersionAPI = cli.EngineVersionAPI
		c.MySQLInstanceAPI = cli.InstanceAPI
		c.MySQLInstanceGroupAPI = cli.InstanceGroupAPI
	}

	{
		cc := secretsmanager.NewConfiguration()
		cc.HTTPClient = authedClient
		cc.Servers = secretsmanager.ServerConfigurations{{URL: cfg.Endpoints.SecretsManager}}

		cli := secretsmanager.NewAPIClient(cc)

		c.SecretAPI = cli.SecretAPI
		c.SecretVersionAPI = cli.SecretVersionAPI
		c.SecretsManagerAPI = cli.SecretsManagerAPI
	}

	return c
}

func (c *APIClient) SetToken(tok string) {
	if c.authRT != nil {
		c.authRT.token.Store(tok)
	}
}

func (c *APIClient) SetUserAgent(ua string) {
	if c.authRT != nil {
		c.authRT.userAgent.Store(ua)
	}
}
