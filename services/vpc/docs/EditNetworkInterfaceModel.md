# EditNetworkInterfaceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**SecurityGroups** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEditNetworkInterfaceModel

`func NewEditNetworkInterfaceModel() *EditNetworkInterfaceModel`

NewEditNetworkInterfaceModel instantiates a new EditNetworkInterfaceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditNetworkInterfaceModelWithDefaults

`func NewEditNetworkInterfaceModelWithDefaults() *EditNetworkInterfaceModel`

NewEditNetworkInterfaceModelWithDefaults instantiates a new EditNetworkInterfaceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EditNetworkInterfaceModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditNetworkInterfaceModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditNetworkInterfaceModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditNetworkInterfaceModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EditNetworkInterfaceModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EditNetworkInterfaceModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *EditNetworkInterfaceModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditNetworkInterfaceModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditNetworkInterfaceModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EditNetworkInterfaceModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EditNetworkInterfaceModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EditNetworkInterfaceModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSecurityGroups

`func (o *EditNetworkInterfaceModel) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *EditNetworkInterfaceModel) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *EditNetworkInterfaceModel) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *EditNetworkInterfaceModel) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *EditNetworkInterfaceModel) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *EditNetworkInterfaceModel) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


