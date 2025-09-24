# ClusterNetworkRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cni** | [**ClusterNetworkCNI**](ClusterNetworkCNI.md) | 사용 중인 CNI(Container Network Interface) 플러그인 &lt;br/&gt; - &#x60;cilium&#x60;: eBPF 기반의 고성능 CNI, 낮은 지연과 강력한 보안 정책 지원  &lt;br/&gt; - &#x60;calico&#x60;: 가장 널리 쓰이는 CNI, 단순한 설정과 다양한 네트워크 모드 제공 | 
**ServiceCidr** | Pointer to **NullableString** |  | [optional] 
**PodCidr** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewClusterNetworkRequestModel

`func NewClusterNetworkRequestModel(cni ClusterNetworkCNI, ) *ClusterNetworkRequestModel`

NewClusterNetworkRequestModel instantiates a new ClusterNetworkRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNetworkRequestModelWithDefaults

`func NewClusterNetworkRequestModelWithDefaults() *ClusterNetworkRequestModel`

NewClusterNetworkRequestModelWithDefaults instantiates a new ClusterNetworkRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCni

`func (o *ClusterNetworkRequestModel) GetCni() ClusterNetworkCNI`

GetCni returns the Cni field if non-nil, zero value otherwise.

### GetCniOk

`func (o *ClusterNetworkRequestModel) GetCniOk() (*ClusterNetworkCNI, bool)`

GetCniOk returns a tuple with the Cni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCni

`func (o *ClusterNetworkRequestModel) SetCni(v ClusterNetworkCNI)`

SetCni sets Cni field to given value.


### GetServiceCidr

`func (o *ClusterNetworkRequestModel) GetServiceCidr() string`

GetServiceCidr returns the ServiceCidr field if non-nil, zero value otherwise.

### GetServiceCidrOk

`func (o *ClusterNetworkRequestModel) GetServiceCidrOk() (*string, bool)`

GetServiceCidrOk returns a tuple with the ServiceCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCidr

`func (o *ClusterNetworkRequestModel) SetServiceCidr(v string)`

SetServiceCidr sets ServiceCidr field to given value.

### HasServiceCidr

`func (o *ClusterNetworkRequestModel) HasServiceCidr() bool`

HasServiceCidr returns a boolean if a field has been set.

### SetServiceCidrNil

`func (o *ClusterNetworkRequestModel) SetServiceCidrNil(b bool)`

 SetServiceCidrNil sets the value for ServiceCidr to be an explicit nil

### UnsetServiceCidr
`func (o *ClusterNetworkRequestModel) UnsetServiceCidr()`

UnsetServiceCidr ensures that no value is present for ServiceCidr, not even an explicit nil
### GetPodCidr

`func (o *ClusterNetworkRequestModel) GetPodCidr() string`

GetPodCidr returns the PodCidr field if non-nil, zero value otherwise.

### GetPodCidrOk

`func (o *ClusterNetworkRequestModel) GetPodCidrOk() (*string, bool)`

GetPodCidrOk returns a tuple with the PodCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCidr

`func (o *ClusterNetworkRequestModel) SetPodCidr(v string)`

SetPodCidr sets PodCidr field to given value.

### HasPodCidr

`func (o *ClusterNetworkRequestModel) HasPodCidr() bool`

HasPodCidr returns a boolean if a field has been set.

### SetPodCidrNil

`func (o *ClusterNetworkRequestModel) SetPodCidrNil(b bool)`

 SetPodCidrNil sets the value for PodCidr to be an explicit nil

### UnsetPodCidr
`func (o *ClusterNetworkRequestModel) UnsetPodCidr()`

UnsetPodCidr ensures that no value is present for PodCidr, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


