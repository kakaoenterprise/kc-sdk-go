# ClusterNetworkResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cni** | [**ClusterNetworkCNI**](ClusterNetworkCNI.md) | 사용 중인 CNI(Container Network Interface) 플러그인 &lt;br/&gt; - &#x60;cilium&#x60;: eBPF 기반의 고성능 CNI, 낮은 지연과 강력한 보안 정책 지원  &lt;br/&gt; - &#x60;calico&#x60;: 가장 널리 쓰이는 CNI, 단순한 설정과 다양한 네트워크 모드 제공 | 
**PodCidr** | **string** | 파드가 수신하는 IP | 
**ServiceCidr** | **string** | 클러스터의 서비스 객체가 수신하는 IP | 

## Methods

### NewClusterNetworkResponseModel

`func NewClusterNetworkResponseModel(cni ClusterNetworkCNI, podCidr string, serviceCidr string, ) *ClusterNetworkResponseModel`

NewClusterNetworkResponseModel instantiates a new ClusterNetworkResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNetworkResponseModelWithDefaults

`func NewClusterNetworkResponseModelWithDefaults() *ClusterNetworkResponseModel`

NewClusterNetworkResponseModelWithDefaults instantiates a new ClusterNetworkResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCni

`func (o *ClusterNetworkResponseModel) GetCni() ClusterNetworkCNI`

GetCni returns the Cni field if non-nil, zero value otherwise.

### GetCniOk

`func (o *ClusterNetworkResponseModel) GetCniOk() (*ClusterNetworkCNI, bool)`

GetCniOk returns a tuple with the Cni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCni

`func (o *ClusterNetworkResponseModel) SetCni(v ClusterNetworkCNI)`

SetCni sets Cni field to given value.


### GetPodCidr

`func (o *ClusterNetworkResponseModel) GetPodCidr() string`

GetPodCidr returns the PodCidr field if non-nil, zero value otherwise.

### GetPodCidrOk

`func (o *ClusterNetworkResponseModel) GetPodCidrOk() (*string, bool)`

GetPodCidrOk returns a tuple with the PodCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCidr

`func (o *ClusterNetworkResponseModel) SetPodCidr(v string)`

SetPodCidr sets PodCidr field to given value.


### GetServiceCidr

`func (o *ClusterNetworkResponseModel) GetServiceCidr() string`

GetServiceCidr returns the ServiceCidr field if non-nil, zero value otherwise.

### GetServiceCidrOk

`func (o *ClusterNetworkResponseModel) GetServiceCidrOk() (*string, bool)`

GetServiceCidrOk returns a tuple with the ServiceCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCidr

`func (o *ClusterNetworkResponseModel) SetServiceCidr(v string)`

SetServiceCidr sets ServiceCidr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


