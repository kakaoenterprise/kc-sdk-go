# MysqlV1ApiCreateMysqlInstanceGroupModelSubnetInfoRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Replicas** | **int32** | 해당 서브넷에 생성할 인스턴스 수(1 ~6) | 
**SubnetId** | **string** | 인스턴스를 생성할 서브넷 ID | 

## Methods

### NewMysqlV1ApiCreateMysqlInstanceGroupModelSubnetInfoRequestModel

`func NewMysqlV1ApiCreateMysqlInstanceGroupModelSubnetInfoRequestModel(replicas int32, subnetId string, ) *MysqlV1ApiCreateMysqlInstanceGroupModelSubnetInfoRequestModel`

NewMysqlV1ApiCreateMysqlInstanceGroupModelSubnetInfoRequestModel instantiates a new MysqlV1ApiCreateMysqlInstanceGroupModelSubnetInfoRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlV1ApiCreateMysqlInstanceGroupModelSubnetInfoRequestModelWithDefaults

`func NewMysqlV1ApiCreateMysqlInstanceGroupModelSubnetInfoRequestModelWithDefaults() *MysqlV1ApiCreateMysqlInstanceGroupModelSubnetInfoRequestModel`

NewMysqlV1ApiCreateMysqlInstanceGroupModelSubnetInfoRequestModelWithDefaults instantiates a new MysqlV1ApiCreateMysqlInstanceGroupModelSubnetInfoRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplicas

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelSubnetInfoRequestModel) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelSubnetInfoRequestModel) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelSubnetInfoRequestModel) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.


### GetSubnetId

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelSubnetInfoRequestModel) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelSubnetInfoRequestModel) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelSubnetInfoRequestModel) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


