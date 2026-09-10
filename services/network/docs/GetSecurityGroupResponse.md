# GetSecurityGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroup** | [**SecurityGroup**](SecurityGroup.md) | 조회된 보안 그룹 정보 | 

## Methods

### NewGetSecurityGroupResponse

`func NewGetSecurityGroupResponse(securityGroup SecurityGroup, ) *GetSecurityGroupResponse`

NewGetSecurityGroupResponse instantiates a new GetSecurityGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSecurityGroupResponseWithDefaults

`func NewGetSecurityGroupResponseWithDefaults() *GetSecurityGroupResponse`

NewGetSecurityGroupResponseWithDefaults instantiates a new GetSecurityGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroup

`func (o *GetSecurityGroupResponse) GetSecurityGroup() SecurityGroup`

GetSecurityGroup returns the SecurityGroup field if non-nil, zero value otherwise.

### GetSecurityGroupOk

`func (o *GetSecurityGroupResponse) GetSecurityGroupOk() (*SecurityGroup, bool)`

GetSecurityGroupOk returns a tuple with the SecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroup

`func (o *GetSecurityGroupResponse) SetSecurityGroup(v SecurityGroup)`

SetSecurityGroup sets SecurityGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


