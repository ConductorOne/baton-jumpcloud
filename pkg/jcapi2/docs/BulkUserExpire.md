# BulkUserExpire

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to **[]map[string]interface{}** | Map of additional attributes. | [optional] 
**Id** | Pointer to **string** | Object ID of the systemuser to expire | [optional] 
**Organization** | Pointer to **string** | The identifier for an organization to link this systemuser to | [optional] 

## Methods

### NewBulkUserExpire

`func NewBulkUserExpire() *BulkUserExpire`

NewBulkUserExpire instantiates a new BulkUserExpire object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUserExpireWithDefaults

`func NewBulkUserExpireWithDefaults() *BulkUserExpire`

NewBulkUserExpireWithDefaults instantiates a new BulkUserExpire object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *BulkUserExpire) GetAttributes() []map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BulkUserExpire) GetAttributesOk() (*[]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BulkUserExpire) SetAttributes(v []map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *BulkUserExpire) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *BulkUserExpire) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkUserExpire) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkUserExpire) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BulkUserExpire) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganization

`func (o *BulkUserExpire) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *BulkUserExpire) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *BulkUserExpire) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *BulkUserExpire) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


