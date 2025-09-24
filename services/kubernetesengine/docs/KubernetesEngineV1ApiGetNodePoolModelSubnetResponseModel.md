# KubernetesEngineV1ApiGetNodePoolModelSubnetResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) | 서브넷이 속한 가용 영역 | 
**CidrBlock** | **string** | 서브넷의 IPv4 CIDR 블록 | 
**Id** | **string** | 서브넷의 고유 ID | 

## Methods

### NewKubernetesEngineV1ApiGetNodePoolModelSubnetResponseModel

`func NewKubernetesEngineV1ApiGetNodePoolModelSubnetResponseModel(availabilityZone AvailabilityZone, cidrBlock string, id string, ) *KubernetesEngineV1ApiGetNodePoolModelSubnetResponseModel`

NewKubernetesEngineV1ApiGetNodePoolModelSubnetResponseModel instantiates a new KubernetesEngineV1ApiGetNodePoolModelSubnetResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEngineV1ApiGetNodePoolModelSubnetResponseModelWithDefaults

`func NewKubernetesEngineV1ApiGetNodePoolModelSubnetResponseModelWithDefaults() *KubernetesEngineV1ApiGetNodePoolModelSubnetResponseModel`

NewKubernetesEngineV1ApiGetNodePoolModelSubnetResponseModelWithDefaults instantiates a new KubernetesEngineV1ApiGetNodePoolModelSubnetResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZone

`func (o *KubernetesEngineV1ApiGetNodePoolModelSubnetResponseModel) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *KubernetesEngineV1ApiGetNodePoolModelSubnetResponseModel) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *KubernetesEngineV1ApiGetNodePoolModelSubnetResponseModel) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetCidrBlock

`func (o *KubernetesEngineV1ApiGetNodePoolModelSubnetResponseModel) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *KubernetesEngineV1ApiGetNodePoolModelSubnetResponseModel) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *KubernetesEngineV1ApiGetNodePoolModelSubnetResponseModel) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.


### GetId

`func (o *KubernetesEngineV1ApiGetNodePoolModelSubnetResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesEngineV1ApiGetNodePoolModelSubnetResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesEngineV1ApiGetNodePoolModelSubnetResponseModel) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


