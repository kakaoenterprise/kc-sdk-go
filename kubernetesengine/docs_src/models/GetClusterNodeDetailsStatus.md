# GetClusterNodeDetailsStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | [**[]Address**](Address.md) | 노드 주소 목록 | 
**Allocatable** | [**AllocatableResources**](AllocatableResources.md) | 노드에서 사용 가능한 리소스 | 
**Conditions** | [**[]Condition**](Condition.md) | 노드 status 조건 목록(예: Ready, DiskPressure)  &lt;!---status 영문 유지---&gt; | 

## Methods

### NewGetClusterNodeDetailsStatus

`func NewGetClusterNodeDetailsStatus(addresses []Address, allocatable AllocatableResources, conditions []Condition, ) *GetClusterNodeDetailsStatus`

NewGetClusterNodeDetailsStatus instantiates a new GetClusterNodeDetailsStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterNodeDetailsStatusWithDefaults

`func NewGetClusterNodeDetailsStatusWithDefaults() *GetClusterNodeDetailsStatus`

NewGetClusterNodeDetailsStatusWithDefaults instantiates a new GetClusterNodeDetailsStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *GetClusterNodeDetailsStatus) GetAddresses() []Address`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *GetClusterNodeDetailsStatus) GetAddressesOk() (*[]Address, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *GetClusterNodeDetailsStatus) SetAddresses(v []Address)`

SetAddresses sets Addresses field to given value.


### GetAllocatable

`func (o *GetClusterNodeDetailsStatus) GetAllocatable() AllocatableResources`

GetAllocatable returns the Allocatable field if non-nil, zero value otherwise.

### GetAllocatableOk

`func (o *GetClusterNodeDetailsStatus) GetAllocatableOk() (*AllocatableResources, bool)`

GetAllocatableOk returns a tuple with the Allocatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatable

`func (o *GetClusterNodeDetailsStatus) SetAllocatable(v AllocatableResources)`

SetAllocatable sets Allocatable field to given value.


### GetConditions

`func (o *GetClusterNodeDetailsStatus) GetConditions() []Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *GetClusterNodeDetailsStatus) GetConditionsOk() (*[]Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *GetClusterNodeDetailsStatus) SetConditions(v []Condition)`

SetConditions sets Conditions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


