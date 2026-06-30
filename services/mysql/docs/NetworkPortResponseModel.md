# NetworkPortResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubnetId** | Pointer to **NullableString** |  | [optional] 
**SecurityGroupIds** | **[]string** | 보안 그룹 ID 목록 | 

## Methods

### NewNetworkPortResponseModel

`func NewNetworkPortResponseModel(securityGroupIds []string, ) *NetworkPortResponseModel`

NewNetworkPortResponseModel instantiates a new NetworkPortResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPortResponseModelWithDefaults

`func NewNetworkPortResponseModelWithDefaults() *NetworkPortResponseModel`

NewNetworkPortResponseModelWithDefaults instantiates a new NetworkPortResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnetId

`func (o *NetworkPortResponseModel) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *NetworkPortResponseModel) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *NetworkPortResponseModel) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *NetworkPortResponseModel) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *NetworkPortResponseModel) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *NetworkPortResponseModel) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetSecurityGroupIds

`func (o *NetworkPortResponseModel) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *NetworkPortResponseModel) GetSecurityGroupIdsOk() (*[]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIds

`func (o *NetworkPortResponseModel) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds sets SecurityGroupIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


