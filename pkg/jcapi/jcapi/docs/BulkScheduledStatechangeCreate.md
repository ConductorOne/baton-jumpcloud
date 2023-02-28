# BulkScheduledStatechangeCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationEmailOverride** | Pointer to **string** | Send the activation or welcome email to the specified email address upon activation. Can only be used with a single user_id and scheduled activation. This field will be ignored if &#x60;send_activation_emails&#x60; is explicitly set to false. | [optional] 
**SendActivationEmails** | Pointer to **bool** | Set to true to send activation or welcome email(s) to each user_id upon activation. Set to false to suppress emails. Can only be used with scheduled activation(s). | [optional] 
**StartDate** | **time.Time** | Date and time that scheduled action should occur | 
**State** | **string** | The state to move the user(s) to | 
**UserIds** | **[]string** | Array of system user ids to schedule for a state change | 

## Methods

### NewBulkScheduledStatechangeCreate

`func NewBulkScheduledStatechangeCreate(startDate time.Time, state string, userIds []string, ) *BulkScheduledStatechangeCreate`

NewBulkScheduledStatechangeCreate instantiates a new BulkScheduledStatechangeCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkScheduledStatechangeCreateWithDefaults

`func NewBulkScheduledStatechangeCreateWithDefaults() *BulkScheduledStatechangeCreate`

NewBulkScheduledStatechangeCreateWithDefaults instantiates a new BulkScheduledStatechangeCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationEmailOverride

`func (o *BulkScheduledStatechangeCreate) GetActivationEmailOverride() string`

GetActivationEmailOverride returns the ActivationEmailOverride field if non-nil, zero value otherwise.

### GetActivationEmailOverrideOk

`func (o *BulkScheduledStatechangeCreate) GetActivationEmailOverrideOk() (*string, bool)`

GetActivationEmailOverrideOk returns a tuple with the ActivationEmailOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationEmailOverride

`func (o *BulkScheduledStatechangeCreate) SetActivationEmailOverride(v string)`

SetActivationEmailOverride sets ActivationEmailOverride field to given value.

### HasActivationEmailOverride

`func (o *BulkScheduledStatechangeCreate) HasActivationEmailOverride() bool`

HasActivationEmailOverride returns a boolean if a field has been set.

### GetSendActivationEmails

`func (o *BulkScheduledStatechangeCreate) GetSendActivationEmails() bool`

GetSendActivationEmails returns the SendActivationEmails field if non-nil, zero value otherwise.

### GetSendActivationEmailsOk

`func (o *BulkScheduledStatechangeCreate) GetSendActivationEmailsOk() (*bool, bool)`

GetSendActivationEmailsOk returns a tuple with the SendActivationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendActivationEmails

`func (o *BulkScheduledStatechangeCreate) SetSendActivationEmails(v bool)`

SetSendActivationEmails sets SendActivationEmails field to given value.

### HasSendActivationEmails

`func (o *BulkScheduledStatechangeCreate) HasSendActivationEmails() bool`

HasSendActivationEmails returns a boolean if a field has been set.

### GetStartDate

`func (o *BulkScheduledStatechangeCreate) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BulkScheduledStatechangeCreate) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BulkScheduledStatechangeCreate) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetState

`func (o *BulkScheduledStatechangeCreate) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BulkScheduledStatechangeCreate) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BulkScheduledStatechangeCreate) SetState(v string)`

SetState sets State field to given value.


### GetUserIds

`func (o *BulkScheduledStatechangeCreate) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *BulkScheduledStatechangeCreate) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *BulkScheduledStatechangeCreate) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


