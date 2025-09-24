# BnsVpcV1ApiListRouteTablesModelRouteModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 라우팅 경로의 고유 ID | 
**Destination** | Pointer to **NullableString** |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) |  | [optional] 
**TargetType** | Pointer to **NullableString** |  | [optional] 
**TargetName** | Pointer to **NullableString** |  | [optional] 
**IsLocalRoute** | Pointer to **NullableBool** |  | [optional] 
**TargetId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBnsVpcV1ApiListRouteTablesModelRouteModel

`func NewBnsVpcV1ApiListRouteTablesModelRouteModel(id string, ) *BnsVpcV1ApiListRouteTablesModelRouteModel`

NewBnsVpcV1ApiListRouteTablesModelRouteModel instantiates a new BnsVpcV1ApiListRouteTablesModelRouteModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsVpcV1ApiListRouteTablesModelRouteModelWithDefaults

`func NewBnsVpcV1ApiListRouteTablesModelRouteModelWithDefaults() *BnsVpcV1ApiListRouteTablesModelRouteModel`

NewBnsVpcV1ApiListRouteTablesModelRouteModelWithDefaults instantiates a new BnsVpcV1ApiListRouteTablesModelRouteModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) SetId(v string)`

SetId sets Id field to given value.


### GetDestination

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### SetDestinationNil

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) SetDestinationNil(b bool)`

 SetDestinationNil sets the value for Destination to be an explicit nil

### UnsetDestination
`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) UnsetDestination()`

UnsetDestination ensures that no value is present for Destination, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetTargetType

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
### GetTargetName

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### SetTargetNameNil

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) SetTargetNameNil(b bool)`

 SetTargetNameNil sets the value for TargetName to be an explicit nil

### UnsetTargetName
`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) UnsetTargetName()`

UnsetTargetName ensures that no value is present for TargetName, not even an explicit nil
### GetIsLocalRoute

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) GetIsLocalRoute() bool`

GetIsLocalRoute returns the IsLocalRoute field if non-nil, zero value otherwise.

### GetIsLocalRouteOk

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) GetIsLocalRouteOk() (*bool, bool)`

GetIsLocalRouteOk returns a tuple with the IsLocalRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocalRoute

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) SetIsLocalRoute(v bool)`

SetIsLocalRoute sets IsLocalRoute field to given value.

### HasIsLocalRoute

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) HasIsLocalRoute() bool`

HasIsLocalRoute returns a boolean if a field has been set.

### SetIsLocalRouteNil

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) SetIsLocalRouteNil(b bool)`

 SetIsLocalRouteNil sets the value for IsLocalRoute to be an explicit nil

### UnsetIsLocalRoute
`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) UnsetIsLocalRoute()`

UnsetIsLocalRoute ensures that no value is present for IsLocalRoute, not even an explicit nil
### GetTargetId

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *BnsVpcV1ApiListRouteTablesModelRouteModel) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


