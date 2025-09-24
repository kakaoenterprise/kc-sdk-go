# HealthCheckSubnetModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 서브넷 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**CidrBlock** | Pointer to **NullableString** |  | [optional] 
**AvailabilityZone** | Pointer to [**NullableAvailabilityZone**](AvailabilityZone.md) |  | [optional] 
**HealthCheckIps** | Pointer to **[]string** |  | [optional] 

## Methods

### NewHealthCheckSubnetModel

`func NewHealthCheckSubnetModel(id string, ) *HealthCheckSubnetModel`

NewHealthCheckSubnetModel instantiates a new HealthCheckSubnetModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckSubnetModelWithDefaults

`func NewHealthCheckSubnetModelWithDefaults() *HealthCheckSubnetModel`

NewHealthCheckSubnetModelWithDefaults instantiates a new HealthCheckSubnetModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HealthCheckSubnetModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HealthCheckSubnetModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HealthCheckSubnetModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *HealthCheckSubnetModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HealthCheckSubnetModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HealthCheckSubnetModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HealthCheckSubnetModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HealthCheckSubnetModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HealthCheckSubnetModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCidrBlock

`func (o *HealthCheckSubnetModel) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *HealthCheckSubnetModel) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *HealthCheckSubnetModel) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *HealthCheckSubnetModel) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.

### SetCidrBlockNil

`func (o *HealthCheckSubnetModel) SetCidrBlockNil(b bool)`

 SetCidrBlockNil sets the value for CidrBlock to be an explicit nil

### UnsetCidrBlock
`func (o *HealthCheckSubnetModel) UnsetCidrBlock()`

UnsetCidrBlock ensures that no value is present for CidrBlock, not even an explicit nil
### GetAvailabilityZone

`func (o *HealthCheckSubnetModel) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *HealthCheckSubnetModel) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *HealthCheckSubnetModel) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *HealthCheckSubnetModel) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *HealthCheckSubnetModel) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *HealthCheckSubnetModel) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetHealthCheckIps

`func (o *HealthCheckSubnetModel) GetHealthCheckIps() []string`

GetHealthCheckIps returns the HealthCheckIps field if non-nil, zero value otherwise.

### GetHealthCheckIpsOk

`func (o *HealthCheckSubnetModel) GetHealthCheckIpsOk() (*[]string, bool)`

GetHealthCheckIpsOk returns a tuple with the HealthCheckIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckIps

`func (o *HealthCheckSubnetModel) SetHealthCheckIps(v []string)`

SetHealthCheckIps sets HealthCheckIps field to given value.

### HasHealthCheckIps

`func (o *HealthCheckSubnetModel) HasHealthCheckIps() bool`

HasHealthCheckIps returns a boolean if a field has been set.

### SetHealthCheckIpsNil

`func (o *HealthCheckSubnetModel) SetHealthCheckIpsNil(b bool)`

 SetHealthCheckIpsNil sets the value for HealthCheckIps to be an explicit nil

### UnsetHealthCheckIps
`func (o *HealthCheckSubnetModel) UnsetHealthCheckIps()`

UnsetHealthCheckIps ensures that no value is present for HealthCheckIps, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


