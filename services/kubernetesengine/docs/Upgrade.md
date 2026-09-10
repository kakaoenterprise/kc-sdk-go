# Upgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Upgradable** | Pointer to [**[]UpgradableVersion**](UpgradableVersion.md) |  | [optional] 
**Current** | [**CurrentOMT**](CurrentOMT.md) | 현재 클러스터 버전 | 

## Methods

### NewUpgrade

`func NewUpgrade(current CurrentOMT, ) *Upgrade`

NewUpgrade instantiates a new Upgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeWithDefaults

`func NewUpgradeWithDefaults() *Upgrade`

NewUpgradeWithDefaults instantiates a new Upgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpgradable

`func (o *Upgrade) GetUpgradable() []UpgradableVersion`

GetUpgradable returns the Upgradable field if non-nil, zero value otherwise.

### GetUpgradableOk

`func (o *Upgrade) GetUpgradableOk() (*[]UpgradableVersion, bool)`

GetUpgradableOk returns a tuple with the Upgradable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradable

`func (o *Upgrade) SetUpgradable(v []UpgradableVersion)`

SetUpgradable sets Upgradable field to given value.

### HasUpgradable

`func (o *Upgrade) HasUpgradable() bool`

HasUpgradable returns a boolean if a field has been set.

### SetUpgradableNil

`func (o *Upgrade) SetUpgradableNil(b bool)`

 SetUpgradableNil sets the value for Upgradable to be an explicit nil

### UnsetUpgradable
`func (o *Upgrade) UnsetUpgradable()`

UnsetUpgradable ensures that no value is present for Upgradable, not even an explicit nil
### GetCurrent

`func (o *Upgrade) GetCurrent() CurrentOMT`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *Upgrade) GetCurrentOk() (*CurrentOMT, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *Upgrade) SetCurrent(v CurrentOMT)`

SetCurrent sets Current field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


