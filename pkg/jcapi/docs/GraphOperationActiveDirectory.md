# GraphOperationActiveDirectory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ObjectID of graph object being added or removed as an association. | 
**Op** | **string** | How to modify the graph connection. | 
**Attributes** | Pointer to **map[string]interface{}** | The graph attributes. | [optional] 
**Type** | **string** | Targets which a \&quot;active_directory\&quot; can be associated to. | 

## Methods

### NewGraphOperationActiveDirectory

`func NewGraphOperationActiveDirectory(id string, op string, type_ string, ) *GraphOperationActiveDirectory`

NewGraphOperationActiveDirectory instantiates a new GraphOperationActiveDirectory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphOperationActiveDirectoryWithDefaults

`func NewGraphOperationActiveDirectoryWithDefaults() *GraphOperationActiveDirectory`

NewGraphOperationActiveDirectoryWithDefaults instantiates a new GraphOperationActiveDirectory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GraphOperationActiveDirectory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GraphOperationActiveDirectory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GraphOperationActiveDirectory) SetId(v string)`

SetId sets Id field to given value.


### GetOp

`func (o *GraphOperationActiveDirectory) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *GraphOperationActiveDirectory) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *GraphOperationActiveDirectory) SetOp(v string)`

SetOp sets Op field to given value.


### GetAttributes

`func (o *GraphOperationActiveDirectory) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GraphOperationActiveDirectory) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GraphOperationActiveDirectory) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GraphOperationActiveDirectory) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetType

`func (o *GraphOperationActiveDirectory) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GraphOperationActiveDirectory) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GraphOperationActiveDirectory) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


