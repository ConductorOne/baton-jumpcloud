# MfaEnrollment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OverallStatus** | Pointer to [**MfaEnrollmentStatus**](MfaEnrollmentStatus.md) |  | [optional] 
**PushStatus** | Pointer to [**MfaEnrollmentStatus**](MfaEnrollmentStatus.md) |  | [optional] 
**TotpStatus** | Pointer to [**MfaEnrollmentStatus**](MfaEnrollmentStatus.md) |  | [optional] 
**WebAuthnStatus** | Pointer to [**MfaEnrollmentStatus**](MfaEnrollmentStatus.md) |  | [optional] 

## Methods

### NewMfaEnrollment

`func NewMfaEnrollment() *MfaEnrollment`

NewMfaEnrollment instantiates a new MfaEnrollment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMfaEnrollmentWithDefaults

`func NewMfaEnrollmentWithDefaults() *MfaEnrollment`

NewMfaEnrollmentWithDefaults instantiates a new MfaEnrollment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverallStatus

`func (o *MfaEnrollment) GetOverallStatus() MfaEnrollmentStatus`

GetOverallStatus returns the OverallStatus field if non-nil, zero value otherwise.

### GetOverallStatusOk

`func (o *MfaEnrollment) GetOverallStatusOk() (*MfaEnrollmentStatus, bool)`

GetOverallStatusOk returns a tuple with the OverallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStatus

`func (o *MfaEnrollment) SetOverallStatus(v MfaEnrollmentStatus)`

SetOverallStatus sets OverallStatus field to given value.

### HasOverallStatus

`func (o *MfaEnrollment) HasOverallStatus() bool`

HasOverallStatus returns a boolean if a field has been set.

### GetPushStatus

`func (o *MfaEnrollment) GetPushStatus() MfaEnrollmentStatus`

GetPushStatus returns the PushStatus field if non-nil, zero value otherwise.

### GetPushStatusOk

`func (o *MfaEnrollment) GetPushStatusOk() (*MfaEnrollmentStatus, bool)`

GetPushStatusOk returns a tuple with the PushStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushStatus

`func (o *MfaEnrollment) SetPushStatus(v MfaEnrollmentStatus)`

SetPushStatus sets PushStatus field to given value.

### HasPushStatus

`func (o *MfaEnrollment) HasPushStatus() bool`

HasPushStatus returns a boolean if a field has been set.

### GetTotpStatus

`func (o *MfaEnrollment) GetTotpStatus() MfaEnrollmentStatus`

GetTotpStatus returns the TotpStatus field if non-nil, zero value otherwise.

### GetTotpStatusOk

`func (o *MfaEnrollment) GetTotpStatusOk() (*MfaEnrollmentStatus, bool)`

GetTotpStatusOk returns a tuple with the TotpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpStatus

`func (o *MfaEnrollment) SetTotpStatus(v MfaEnrollmentStatus)`

SetTotpStatus sets TotpStatus field to given value.

### HasTotpStatus

`func (o *MfaEnrollment) HasTotpStatus() bool`

HasTotpStatus returns a boolean if a field has been set.

### GetWebAuthnStatus

`func (o *MfaEnrollment) GetWebAuthnStatus() MfaEnrollmentStatus`

GetWebAuthnStatus returns the WebAuthnStatus field if non-nil, zero value otherwise.

### GetWebAuthnStatusOk

`func (o *MfaEnrollment) GetWebAuthnStatusOk() (*MfaEnrollmentStatus, bool)`

GetWebAuthnStatusOk returns a tuple with the WebAuthnStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAuthnStatus

`func (o *MfaEnrollment) SetWebAuthnStatus(v MfaEnrollmentStatus)`

SetWebAuthnStatus sets WebAuthnStatus field to given value.

### HasWebAuthnStatus

`func (o *MfaEnrollment) HasWebAuthnStatus() bool`

HasWebAuthnStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


