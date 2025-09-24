# KubernetesEngineV1ApiListClusterNodesModelSubnetResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) | 서브넷이 속한 가용 영역 | 
**CidrBlock** | **string** | 서브넷의 IPv4 CIDR 블록 | 
**Id** | **string** | 서브넷 ID | 

## Methods

### NewKubernetesEngineV1ApiListClusterNodesModelSubnetResponseModel

`func NewKubernetesEngineV1ApiListClusterNodesModelSubnetResponseModel(availabilityZone AvailabilityZone, cidrBlock string, id string, ) *KubernetesEngineV1ApiListClusterNodesModelSubnetResponseModel`

NewKubernetesEngineV1ApiListClusterNodesModelSubnetResponseModel instantiates a new KubernetesEngineV1ApiListClusterNodesModelSubnetResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEngineV1ApiListClusterNodesModelSubnetResponseModelWithDefaults

`func NewKubernetesEngineV1ApiListClusterNodesModelSubnetResponseModelWithDefaults() *KubernetesEngineV1ApiListClusterNodesModelSubnetResponseModel`

NewKubernetesEngineV1ApiListClusterNodesModelSubnetResponseModelWithDefaults instantiates a new KubernetesEngineV1ApiListClusterNodesModelSubnetResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZone

`func (o *KubernetesEngineV1ApiListClusterNodesModelSubnetResponseModel) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *KubernetesEngineV1ApiListClusterNodesModelSubnetResponseModel) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *KubernetesEngineV1ApiListClusterNodesModelSubnetResponseModel) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetCidrBlock

`func (o *KubernetesEngineV1ApiListClusterNodesModelSubnetResponseModel) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *KubernetesEngineV1ApiListClusterNodesModelSubnetResponseModel) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *KubernetesEngineV1ApiListClusterNodesModelSubnetResponseModel) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.


### GetId

`func (o *KubernetesEngineV1ApiListClusterNodesModelSubnetResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesEngineV1ApiListClusterNodesModelSubnetResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesEngineV1ApiListClusterNodesModelSubnetResponseModel) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


