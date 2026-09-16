# ListClusterUpgradableVersionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Upgrade** | [**Upgrade**](Upgrade.md) | 클러스터 업그레이드 정보 | 

## Methods

### NewListClusterUpgradableVersionsResponse

`func NewListClusterUpgradableVersionsResponse(upgrade Upgrade, ) *ListClusterUpgradableVersionsResponse`

NewListClusterUpgradableVersionsResponse instantiates a new ListClusterUpgradableVersionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterUpgradableVersionsResponseWithDefaults

`func NewListClusterUpgradableVersionsResponseWithDefaults() *ListClusterUpgradableVersionsResponse`

NewListClusterUpgradableVersionsResponseWithDefaults instantiates a new ListClusterUpgradableVersionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpgrade

`func (o *ListClusterUpgradableVersionsResponse) GetUpgrade() Upgrade`

GetUpgrade returns the Upgrade field if non-nil, zero value otherwise.

### GetUpgradeOk

`func (o *ListClusterUpgradableVersionsResponse) GetUpgradeOk() (*Upgrade, bool)`

GetUpgradeOk returns a tuple with the Upgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrade

`func (o *ListClusterUpgradableVersionsResponse) SetUpgrade(v Upgrade)`

SetUpgrade sets Upgrade field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


