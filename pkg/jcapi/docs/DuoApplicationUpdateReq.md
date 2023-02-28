# DuoApplicationUpdateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiHost** | **string** |  | 
**IntegrationKey** | **string** |  | 
**Name** | **string** |  | 
**SecretKey** | Pointer to **string** |  | [optional] 

## Methods

### NewDuoApplicationUpdateReq

`func NewDuoApplicationUpdateReq(apiHost string, integrationKey string, name string, ) *DuoApplicationUpdateReq`

NewDuoApplicationUpdateReq instantiates a new DuoApplicationUpdateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDuoApplicationUpdateReqWithDefaults

`func NewDuoApplicationUpdateReqWithDefaults() *DuoApplicationUpdateReq`

NewDuoApplicationUpdateReqWithDefaults instantiates a new DuoApplicationUpdateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiHost

`func (o *DuoApplicationUpdateReq) GetApiHost() string`

GetApiHost returns the ApiHost field if non-nil, zero value otherwise.

### GetApiHostOk

`func (o *DuoApplicationUpdateReq) GetApiHostOk() (*string, bool)`

GetApiHostOk returns a tuple with the ApiHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiHost

`func (o *DuoApplicationUpdateReq) SetApiHost(v string)`

SetApiHost sets ApiHost field to given value.


### GetIntegrationKey

`func (o *DuoApplicationUpdateReq) GetIntegrationKey() string`

GetIntegrationKey returns the IntegrationKey field if non-nil, zero value otherwise.

### GetIntegrationKeyOk

`func (o *DuoApplicationUpdateReq) GetIntegrationKeyOk() (*string, bool)`

GetIntegrationKeyOk returns a tuple with the IntegrationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationKey

`func (o *DuoApplicationUpdateReq) SetIntegrationKey(v string)`

SetIntegrationKey sets IntegrationKey field to given value.


### GetName

`func (o *DuoApplicationUpdateReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DuoApplicationUpdateReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DuoApplicationUpdateReq) SetName(v string)`

SetName sets Name field to given value.


### GetSecretKey

`func (o *DuoApplicationUpdateReq) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *DuoApplicationUpdateReq) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *DuoApplicationUpdateReq) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *DuoApplicationUpdateReq) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


