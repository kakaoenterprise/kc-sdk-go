# AZPolicies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bcs** | [**AZList**](AZList.md) |  | 
**Vpc** | [**AZList**](AZList.md) |  | 
**Lb** | [**AZList**](AZList.md) |  | 
**K8se** | [**AZList**](AZList.md) |  | 

## Methods

### NewAZPolicies

`func NewAZPolicies(bcs AZList, vpc AZList, lb AZList, k8se AZList, ) *AZPolicies`

NewAZPolicies instantiates a new AZPolicies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAZPoliciesWithDefaults

`func NewAZPoliciesWithDefaults() *AZPolicies`

NewAZPoliciesWithDefaults instantiates a new AZPolicies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBcs

`func (o *AZPolicies) GetBcs() AZList`

GetBcs returns the Bcs field if non-nil, zero value otherwise.

### GetBcsOk

`func (o *AZPolicies) GetBcsOk() (*AZList, bool)`

GetBcsOk returns a tuple with the Bcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcs

`func (o *AZPolicies) SetBcs(v AZList)`

SetBcs sets Bcs field to given value.


### GetVpc

`func (o *AZPolicies) GetVpc() AZList`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *AZPolicies) GetVpcOk() (*AZList, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *AZPolicies) SetVpc(v AZList)`

SetVpc sets Vpc field to given value.


### GetLb

`func (o *AZPolicies) GetLb() AZList`

GetLb returns the Lb field if non-nil, zero value otherwise.

### GetLbOk

`func (o *AZPolicies) GetLbOk() (*AZList, bool)`

GetLbOk returns a tuple with the Lb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLb

`func (o *AZPolicies) SetLb(v AZList)`

SetLb sets Lb field to given value.


### GetK8se

`func (o *AZPolicies) GetK8se() AZList`

GetK8se returns the K8se field if non-nil, zero value otherwise.

### GetK8seOk

`func (o *AZPolicies) GetK8seOk() (*AZList, bool)`

GetK8seOk returns a tuple with the K8se field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8se

`func (o *AZPolicies) SetK8se(v AZList)`

SetK8se sets K8se field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


