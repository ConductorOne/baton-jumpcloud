# CustomEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** |  | [optional] 
**Button** | Pointer to **string** |  | [optional] 
**Header** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**NextStepContactInfo** | Pointer to **string** |  | [optional] 
**Subject** | **string** |  | 
**Title** | Pointer to **string** |  | [optional] 
**Type** | [**CustomEmailType**](CustomEmailType.md) |  | 

## Methods

### NewCustomEmail

`func NewCustomEmail(subject string, type_ CustomEmailType, ) *CustomEmail`

NewCustomEmail instantiates a new CustomEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEmailWithDefaults

`func NewCustomEmailWithDefaults() *CustomEmail`

NewCustomEmailWithDefaults instantiates a new CustomEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *CustomEmail) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CustomEmail) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CustomEmail) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *CustomEmail) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetButton

`func (o *CustomEmail) GetButton() string`

GetButton returns the Button field if non-nil, zero value otherwise.

### GetButtonOk

`func (o *CustomEmail) GetButtonOk() (*string, bool)`

GetButtonOk returns a tuple with the Button field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButton

`func (o *CustomEmail) SetButton(v string)`

SetButton sets Button field to given value.

### HasButton

`func (o *CustomEmail) HasButton() bool`

HasButton returns a boolean if a field has been set.

### GetHeader

`func (o *CustomEmail) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *CustomEmail) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *CustomEmail) SetHeader(v string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *CustomEmail) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetId

`func (o *CustomEmail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomEmail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomEmail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomEmail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNextStepContactInfo

`func (o *CustomEmail) GetNextStepContactInfo() string`

GetNextStepContactInfo returns the NextStepContactInfo field if non-nil, zero value otherwise.

### GetNextStepContactInfoOk

`func (o *CustomEmail) GetNextStepContactInfoOk() (*string, bool)`

GetNextStepContactInfoOk returns a tuple with the NextStepContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextStepContactInfo

`func (o *CustomEmail) SetNextStepContactInfo(v string)`

SetNextStepContactInfo sets NextStepContactInfo field to given value.

### HasNextStepContactInfo

`func (o *CustomEmail) HasNextStepContactInfo() bool`

HasNextStepContactInfo returns a boolean if a field has been set.

### GetSubject

`func (o *CustomEmail) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CustomEmail) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CustomEmail) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetTitle

`func (o *CustomEmail) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CustomEmail) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CustomEmail) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CustomEmail) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *CustomEmail) GetType() CustomEmailType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomEmail) GetTypeOk() (*CustomEmailType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomEmail) SetType(v CustomEmailType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


