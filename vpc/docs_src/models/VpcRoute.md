# VpcRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | 라우팅 경로의 고유 ID | [optional] 
**Destination** | Pointer to **NullableString** | 목적지 네트워크 주소 (CIDR 형식) | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 | [optional] 
**TargetName** | Pointer to **NullableString** | 트래픽 전달 대상 리소스의 이름 | [optional] 
**IsLocalRoute** | Pointer to **NullableBool** | 로컬 통신용 자동 생성 라우트인지 여부 - &#x60;true&#x60;인 경우 VPC 내 통신을 위한 기본 경로 | [optional] 
**TargetId** | Pointer to **NullableString** | 트래픽 전달 대상 리소스의 ID | [optional] 
**TargetType** | Pointer to **NullableString** | 트래픽 전달 대상의 유형 | [optional] 

## Methods

### NewVpcRoute

`func NewVpcRoute() *VpcRoute`

NewVpcRoute instantiates a new VpcRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcRouteWithDefaults

`func NewVpcRouteWithDefaults() *VpcRoute`

NewVpcRouteWithDefaults instantiates a new VpcRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VpcRoute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpcRoute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpcRoute) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VpcRoute) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *VpcRoute) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *VpcRoute) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDestination

`func (o *VpcRoute) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *VpcRoute) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *VpcRoute) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *VpcRoute) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### SetDestinationNil

`func (o *VpcRoute) SetDestinationNil(b bool)`

 SetDestinationNil sets the value for Destination to be an explicit nil

### UnsetDestination
`func (o *VpcRoute) UnsetDestination()`

UnsetDestination ensures that no value is present for Destination, not even an explicit nil
### GetProvisioningStatus

`func (o *VpcRoute) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *VpcRoute) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *VpcRoute) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *VpcRoute) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *VpcRoute) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *VpcRoute) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetTargetName

`func (o *VpcRoute) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *VpcRoute) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *VpcRoute) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *VpcRoute) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### SetTargetNameNil

`func (o *VpcRoute) SetTargetNameNil(b bool)`

 SetTargetNameNil sets the value for TargetName to be an explicit nil

### UnsetTargetName
`func (o *VpcRoute) UnsetTargetName()`

UnsetTargetName ensures that no value is present for TargetName, not even an explicit nil
### GetIsLocalRoute

`func (o *VpcRoute) GetIsLocalRoute() bool`

GetIsLocalRoute returns the IsLocalRoute field if non-nil, zero value otherwise.

### GetIsLocalRouteOk

`func (o *VpcRoute) GetIsLocalRouteOk() (*bool, bool)`

GetIsLocalRouteOk returns a tuple with the IsLocalRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocalRoute

`func (o *VpcRoute) SetIsLocalRoute(v bool)`

SetIsLocalRoute sets IsLocalRoute field to given value.

### HasIsLocalRoute

`func (o *VpcRoute) HasIsLocalRoute() bool`

HasIsLocalRoute returns a boolean if a field has been set.

### SetIsLocalRouteNil

`func (o *VpcRoute) SetIsLocalRouteNil(b bool)`

 SetIsLocalRouteNil sets the value for IsLocalRoute to be an explicit nil

### UnsetIsLocalRoute
`func (o *VpcRoute) UnsetIsLocalRoute()`

UnsetIsLocalRoute ensures that no value is present for IsLocalRoute, not even an explicit nil
### GetTargetId

`func (o *VpcRoute) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *VpcRoute) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *VpcRoute) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *VpcRoute) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *VpcRoute) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *VpcRoute) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetTargetType

`func (o *VpcRoute) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *VpcRoute) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *VpcRoute) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *VpcRoute) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *VpcRoute) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *VpcRoute) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


