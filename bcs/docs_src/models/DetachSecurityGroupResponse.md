# DetachSecurityGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroup** | [**NetworkInterface**](NetworkInterface.md) | 분리된 보안 그룹 정보 | 

## Methods

### NewDetachSecurityGroupResponse

`func NewDetachSecurityGroupResponse(securityGroup NetworkInterface, ) *DetachSecurityGroupResponse`

NewDetachSecurityGroupResponse instantiates a new DetachSecurityGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetachSecurityGroupResponseWithDefaults

`func NewDetachSecurityGroupResponseWithDefaults() *DetachSecurityGroupResponse`

NewDetachSecurityGroupResponseWithDefaults instantiates a new DetachSecurityGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroup

`func (o *DetachSecurityGroupResponse) GetSecurityGroup() NetworkInterface`

GetSecurityGroup returns the SecurityGroup field if non-nil, zero value otherwise.

### GetSecurityGroupOk

`func (o *DetachSecurityGroupResponse) GetSecurityGroupOk() (*NetworkInterface, bool)`

GetSecurityGroupOk returns a tuple with the SecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroup

`func (o *DetachSecurityGroupResponse) SetSecurityGroup(v NetworkInterface)`

SetSecurityGroup sets SecurityGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


