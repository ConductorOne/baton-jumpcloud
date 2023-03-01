# CloudInsightsEventsListEventDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addendum** | Pointer to [**CloudInsightsEventsListEventDetailsAddendum**](CloudInsightsEventsListEventDetailsAddendum.md) |  | [optional] 
**AdditionalEventData** | Pointer to [**CloudInsightsEventsListEventDetailsAdditionalEventData**](CloudInsightsEventsListEventDetailsAdditionalEventData.md) |  | [optional] 
**Annotation** | Pointer to **string** |  | [optional] 
**ApiVersion** | Pointer to **string** |  | [optional] 
**AwsRegion** | Pointer to **string** |  | [optional] 
**EdgeDeviceDetails** | Pointer to [**CloudInsightsEventsListEventDetailsEdgeDeviceDetails**](CloudInsightsEventsListEventDetailsEdgeDeviceDetails.md) |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**EventCategory** | Pointer to **string** |  | [optional] 
**EventId** | Pointer to **string** |  | [optional] 
**EventName** | Pointer to **string** |  | [optional] 
**EventSource** | Pointer to **string** |  | [optional] 
**EventTime** | Pointer to [**CloudInsightsEventsListEventDetailsEventTime**](CloudInsightsEventsListEventDetailsEventTime.md) |  | [optional] 
**EventType** | Pointer to **string** |  | [optional] 
**EventVersion** | Pointer to **string** |  | [optional] 
**ManagementEvent** | Pointer to **bool** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**RecipientAccountId** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**RequestParameters** | Pointer to [**CloudInsightsEventsListEventDetailsRequestParameters**](CloudInsightsEventsListEventDetailsRequestParameters.md) |  | [optional] 
**Resources** | Pointer to [**[]CloudInsightsEventsListEventDetailsResourcesInner**](CloudInsightsEventsListEventDetailsResourcesInner.md) |  | [optional] 
**ResponseElements** | Pointer to [**CloudInsightsEventsListEventDetailsResponseElements**](CloudInsightsEventsListEventDetailsResponseElements.md) |  | [optional] 
**ServiceEventDetails** | Pointer to [**CloudInsightsEventsListEventDetailsServiceEventDetails**](CloudInsightsEventsListEventDetailsServiceEventDetails.md) |  | [optional] 
**SessionCredentialFromConsole** | Pointer to **string** |  | [optional] 
**SharedEventID** | Pointer to **string** |  | [optional] 
**SourceIPAddress** | Pointer to **string** |  | [optional] 
**TlsDetails** | Pointer to [**CloudInsightsEventsListEventDetailsTlsDetails**](CloudInsightsEventsListEventDetailsTlsDetails.md) |  | [optional] 
**UserAgent** | Pointer to **string** |  | [optional] 
**UserIdentity** | Pointer to [**CloudInsightsEventsListEventDetailsUserIdentity**](CloudInsightsEventsListEventDetailsUserIdentity.md) |  | [optional] 
**VpcEndpointId** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudInsightsEventsListEventDetails

`func NewCloudInsightsEventsListEventDetails() *CloudInsightsEventsListEventDetails`

NewCloudInsightsEventsListEventDetails instantiates a new CloudInsightsEventsListEventDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudInsightsEventsListEventDetailsWithDefaults

`func NewCloudInsightsEventsListEventDetailsWithDefaults() *CloudInsightsEventsListEventDetails`

NewCloudInsightsEventsListEventDetailsWithDefaults instantiates a new CloudInsightsEventsListEventDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddendum

`func (o *CloudInsightsEventsListEventDetails) GetAddendum() CloudInsightsEventsListEventDetailsAddendum`

GetAddendum returns the Addendum field if non-nil, zero value otherwise.

### GetAddendumOk

`func (o *CloudInsightsEventsListEventDetails) GetAddendumOk() (*CloudInsightsEventsListEventDetailsAddendum, bool)`

GetAddendumOk returns a tuple with the Addendum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddendum

`func (o *CloudInsightsEventsListEventDetails) SetAddendum(v CloudInsightsEventsListEventDetailsAddendum)`

SetAddendum sets Addendum field to given value.

### HasAddendum

`func (o *CloudInsightsEventsListEventDetails) HasAddendum() bool`

HasAddendum returns a boolean if a field has been set.

### GetAdditionalEventData

`func (o *CloudInsightsEventsListEventDetails) GetAdditionalEventData() CloudInsightsEventsListEventDetailsAdditionalEventData`

GetAdditionalEventData returns the AdditionalEventData field if non-nil, zero value otherwise.

### GetAdditionalEventDataOk

`func (o *CloudInsightsEventsListEventDetails) GetAdditionalEventDataOk() (*CloudInsightsEventsListEventDetailsAdditionalEventData, bool)`

GetAdditionalEventDataOk returns a tuple with the AdditionalEventData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEventData

`func (o *CloudInsightsEventsListEventDetails) SetAdditionalEventData(v CloudInsightsEventsListEventDetailsAdditionalEventData)`

SetAdditionalEventData sets AdditionalEventData field to given value.

### HasAdditionalEventData

`func (o *CloudInsightsEventsListEventDetails) HasAdditionalEventData() bool`

HasAdditionalEventData returns a boolean if a field has been set.

### GetAnnotation

`func (o *CloudInsightsEventsListEventDetails) GetAnnotation() string`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *CloudInsightsEventsListEventDetails) GetAnnotationOk() (*string, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *CloudInsightsEventsListEventDetails) SetAnnotation(v string)`

SetAnnotation sets Annotation field to given value.

### HasAnnotation

`func (o *CloudInsightsEventsListEventDetails) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### GetApiVersion

`func (o *CloudInsightsEventsListEventDetails) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CloudInsightsEventsListEventDetails) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CloudInsightsEventsListEventDetails) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CloudInsightsEventsListEventDetails) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetAwsRegion

`func (o *CloudInsightsEventsListEventDetails) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *CloudInsightsEventsListEventDetails) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *CloudInsightsEventsListEventDetails) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *CloudInsightsEventsListEventDetails) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### GetEdgeDeviceDetails

`func (o *CloudInsightsEventsListEventDetails) GetEdgeDeviceDetails() CloudInsightsEventsListEventDetailsEdgeDeviceDetails`

GetEdgeDeviceDetails returns the EdgeDeviceDetails field if non-nil, zero value otherwise.

### GetEdgeDeviceDetailsOk

`func (o *CloudInsightsEventsListEventDetails) GetEdgeDeviceDetailsOk() (*CloudInsightsEventsListEventDetailsEdgeDeviceDetails, bool)`

GetEdgeDeviceDetailsOk returns a tuple with the EdgeDeviceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeDeviceDetails

`func (o *CloudInsightsEventsListEventDetails) SetEdgeDeviceDetails(v CloudInsightsEventsListEventDetailsEdgeDeviceDetails)`

SetEdgeDeviceDetails sets EdgeDeviceDetails field to given value.

### HasEdgeDeviceDetails

`func (o *CloudInsightsEventsListEventDetails) HasEdgeDeviceDetails() bool`

HasEdgeDeviceDetails returns a boolean if a field has been set.

### GetErrorCode

`func (o *CloudInsightsEventsListEventDetails) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *CloudInsightsEventsListEventDetails) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *CloudInsightsEventsListEventDetails) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *CloudInsightsEventsListEventDetails) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *CloudInsightsEventsListEventDetails) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *CloudInsightsEventsListEventDetails) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *CloudInsightsEventsListEventDetails) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *CloudInsightsEventsListEventDetails) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetEventCategory

`func (o *CloudInsightsEventsListEventDetails) GetEventCategory() string`

GetEventCategory returns the EventCategory field if non-nil, zero value otherwise.

### GetEventCategoryOk

`func (o *CloudInsightsEventsListEventDetails) GetEventCategoryOk() (*string, bool)`

GetEventCategoryOk returns a tuple with the EventCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCategory

`func (o *CloudInsightsEventsListEventDetails) SetEventCategory(v string)`

SetEventCategory sets EventCategory field to given value.

### HasEventCategory

`func (o *CloudInsightsEventsListEventDetails) HasEventCategory() bool`

HasEventCategory returns a boolean if a field has been set.

### GetEventId

`func (o *CloudInsightsEventsListEventDetails) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *CloudInsightsEventsListEventDetails) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *CloudInsightsEventsListEventDetails) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *CloudInsightsEventsListEventDetails) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventName

`func (o *CloudInsightsEventsListEventDetails) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *CloudInsightsEventsListEventDetails) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *CloudInsightsEventsListEventDetails) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *CloudInsightsEventsListEventDetails) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetEventSource

`func (o *CloudInsightsEventsListEventDetails) GetEventSource() string`

GetEventSource returns the EventSource field if non-nil, zero value otherwise.

### GetEventSourceOk

`func (o *CloudInsightsEventsListEventDetails) GetEventSourceOk() (*string, bool)`

GetEventSourceOk returns a tuple with the EventSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSource

`func (o *CloudInsightsEventsListEventDetails) SetEventSource(v string)`

SetEventSource sets EventSource field to given value.

### HasEventSource

`func (o *CloudInsightsEventsListEventDetails) HasEventSource() bool`

HasEventSource returns a boolean if a field has been set.

### GetEventTime

`func (o *CloudInsightsEventsListEventDetails) GetEventTime() CloudInsightsEventsListEventDetailsEventTime`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *CloudInsightsEventsListEventDetails) GetEventTimeOk() (*CloudInsightsEventsListEventDetailsEventTime, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *CloudInsightsEventsListEventDetails) SetEventTime(v CloudInsightsEventsListEventDetailsEventTime)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *CloudInsightsEventsListEventDetails) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetEventType

`func (o *CloudInsightsEventsListEventDetails) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *CloudInsightsEventsListEventDetails) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *CloudInsightsEventsListEventDetails) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *CloudInsightsEventsListEventDetails) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEventVersion

`func (o *CloudInsightsEventsListEventDetails) GetEventVersion() string`

GetEventVersion returns the EventVersion field if non-nil, zero value otherwise.

### GetEventVersionOk

`func (o *CloudInsightsEventsListEventDetails) GetEventVersionOk() (*string, bool)`

GetEventVersionOk returns a tuple with the EventVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventVersion

`func (o *CloudInsightsEventsListEventDetails) SetEventVersion(v string)`

SetEventVersion sets EventVersion field to given value.

### HasEventVersion

`func (o *CloudInsightsEventsListEventDetails) HasEventVersion() bool`

HasEventVersion returns a boolean if a field has been set.

### GetManagementEvent

`func (o *CloudInsightsEventsListEventDetails) GetManagementEvent() bool`

GetManagementEvent returns the ManagementEvent field if non-nil, zero value otherwise.

### GetManagementEventOk

`func (o *CloudInsightsEventsListEventDetails) GetManagementEventOk() (*bool, bool)`

GetManagementEventOk returns a tuple with the ManagementEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementEvent

`func (o *CloudInsightsEventsListEventDetails) SetManagementEvent(v bool)`

SetManagementEvent sets ManagementEvent field to given value.

### HasManagementEvent

`func (o *CloudInsightsEventsListEventDetails) HasManagementEvent() bool`

HasManagementEvent returns a boolean if a field has been set.

### GetReadOnly

`func (o *CloudInsightsEventsListEventDetails) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *CloudInsightsEventsListEventDetails) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *CloudInsightsEventsListEventDetails) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *CloudInsightsEventsListEventDetails) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRecipientAccountId

`func (o *CloudInsightsEventsListEventDetails) GetRecipientAccountId() string`

GetRecipientAccountId returns the RecipientAccountId field if non-nil, zero value otherwise.

### GetRecipientAccountIdOk

`func (o *CloudInsightsEventsListEventDetails) GetRecipientAccountIdOk() (*string, bool)`

GetRecipientAccountIdOk returns a tuple with the RecipientAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientAccountId

`func (o *CloudInsightsEventsListEventDetails) SetRecipientAccountId(v string)`

SetRecipientAccountId sets RecipientAccountId field to given value.

### HasRecipientAccountId

`func (o *CloudInsightsEventsListEventDetails) HasRecipientAccountId() bool`

HasRecipientAccountId returns a boolean if a field has been set.

### GetRequestId

`func (o *CloudInsightsEventsListEventDetails) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CloudInsightsEventsListEventDetails) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CloudInsightsEventsListEventDetails) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CloudInsightsEventsListEventDetails) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetRequestParameters

`func (o *CloudInsightsEventsListEventDetails) GetRequestParameters() CloudInsightsEventsListEventDetailsRequestParameters`

GetRequestParameters returns the RequestParameters field if non-nil, zero value otherwise.

### GetRequestParametersOk

`func (o *CloudInsightsEventsListEventDetails) GetRequestParametersOk() (*CloudInsightsEventsListEventDetailsRequestParameters, bool)`

GetRequestParametersOk returns a tuple with the RequestParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParameters

`func (o *CloudInsightsEventsListEventDetails) SetRequestParameters(v CloudInsightsEventsListEventDetailsRequestParameters)`

SetRequestParameters sets RequestParameters field to given value.

### HasRequestParameters

`func (o *CloudInsightsEventsListEventDetails) HasRequestParameters() bool`

HasRequestParameters returns a boolean if a field has been set.

### GetResources

`func (o *CloudInsightsEventsListEventDetails) GetResources() []CloudInsightsEventsListEventDetailsResourcesInner`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CloudInsightsEventsListEventDetails) GetResourcesOk() (*[]CloudInsightsEventsListEventDetailsResourcesInner, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CloudInsightsEventsListEventDetails) SetResources(v []CloudInsightsEventsListEventDetailsResourcesInner)`

SetResources sets Resources field to given value.

### HasResources

`func (o *CloudInsightsEventsListEventDetails) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetResponseElements

`func (o *CloudInsightsEventsListEventDetails) GetResponseElements() CloudInsightsEventsListEventDetailsResponseElements`

GetResponseElements returns the ResponseElements field if non-nil, zero value otherwise.

### GetResponseElementsOk

`func (o *CloudInsightsEventsListEventDetails) GetResponseElementsOk() (*CloudInsightsEventsListEventDetailsResponseElements, bool)`

GetResponseElementsOk returns a tuple with the ResponseElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseElements

`func (o *CloudInsightsEventsListEventDetails) SetResponseElements(v CloudInsightsEventsListEventDetailsResponseElements)`

SetResponseElements sets ResponseElements field to given value.

### HasResponseElements

`func (o *CloudInsightsEventsListEventDetails) HasResponseElements() bool`

HasResponseElements returns a boolean if a field has been set.

### GetServiceEventDetails

`func (o *CloudInsightsEventsListEventDetails) GetServiceEventDetails() CloudInsightsEventsListEventDetailsServiceEventDetails`

GetServiceEventDetails returns the ServiceEventDetails field if non-nil, zero value otherwise.

### GetServiceEventDetailsOk

`func (o *CloudInsightsEventsListEventDetails) GetServiceEventDetailsOk() (*CloudInsightsEventsListEventDetailsServiceEventDetails, bool)`

GetServiceEventDetailsOk returns a tuple with the ServiceEventDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEventDetails

`func (o *CloudInsightsEventsListEventDetails) SetServiceEventDetails(v CloudInsightsEventsListEventDetailsServiceEventDetails)`

SetServiceEventDetails sets ServiceEventDetails field to given value.

### HasServiceEventDetails

`func (o *CloudInsightsEventsListEventDetails) HasServiceEventDetails() bool`

HasServiceEventDetails returns a boolean if a field has been set.

### GetSessionCredentialFromConsole

`func (o *CloudInsightsEventsListEventDetails) GetSessionCredentialFromConsole() string`

GetSessionCredentialFromConsole returns the SessionCredentialFromConsole field if non-nil, zero value otherwise.

### GetSessionCredentialFromConsoleOk

`func (o *CloudInsightsEventsListEventDetails) GetSessionCredentialFromConsoleOk() (*string, bool)`

GetSessionCredentialFromConsoleOk returns a tuple with the SessionCredentialFromConsole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCredentialFromConsole

`func (o *CloudInsightsEventsListEventDetails) SetSessionCredentialFromConsole(v string)`

SetSessionCredentialFromConsole sets SessionCredentialFromConsole field to given value.

### HasSessionCredentialFromConsole

`func (o *CloudInsightsEventsListEventDetails) HasSessionCredentialFromConsole() bool`

HasSessionCredentialFromConsole returns a boolean if a field has been set.

### GetSharedEventID

`func (o *CloudInsightsEventsListEventDetails) GetSharedEventID() string`

GetSharedEventID returns the SharedEventID field if non-nil, zero value otherwise.

### GetSharedEventIDOk

`func (o *CloudInsightsEventsListEventDetails) GetSharedEventIDOk() (*string, bool)`

GetSharedEventIDOk returns a tuple with the SharedEventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedEventID

`func (o *CloudInsightsEventsListEventDetails) SetSharedEventID(v string)`

SetSharedEventID sets SharedEventID field to given value.

### HasSharedEventID

`func (o *CloudInsightsEventsListEventDetails) HasSharedEventID() bool`

HasSharedEventID returns a boolean if a field has been set.

### GetSourceIPAddress

`func (o *CloudInsightsEventsListEventDetails) GetSourceIPAddress() string`

GetSourceIPAddress returns the SourceIPAddress field if non-nil, zero value otherwise.

### GetSourceIPAddressOk

`func (o *CloudInsightsEventsListEventDetails) GetSourceIPAddressOk() (*string, bool)`

GetSourceIPAddressOk returns a tuple with the SourceIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIPAddress

`func (o *CloudInsightsEventsListEventDetails) SetSourceIPAddress(v string)`

SetSourceIPAddress sets SourceIPAddress field to given value.

### HasSourceIPAddress

`func (o *CloudInsightsEventsListEventDetails) HasSourceIPAddress() bool`

HasSourceIPAddress returns a boolean if a field has been set.

### GetTlsDetails

`func (o *CloudInsightsEventsListEventDetails) GetTlsDetails() CloudInsightsEventsListEventDetailsTlsDetails`

GetTlsDetails returns the TlsDetails field if non-nil, zero value otherwise.

### GetTlsDetailsOk

`func (o *CloudInsightsEventsListEventDetails) GetTlsDetailsOk() (*CloudInsightsEventsListEventDetailsTlsDetails, bool)`

GetTlsDetailsOk returns a tuple with the TlsDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsDetails

`func (o *CloudInsightsEventsListEventDetails) SetTlsDetails(v CloudInsightsEventsListEventDetailsTlsDetails)`

SetTlsDetails sets TlsDetails field to given value.

### HasTlsDetails

`func (o *CloudInsightsEventsListEventDetails) HasTlsDetails() bool`

HasTlsDetails returns a boolean if a field has been set.

### GetUserAgent

`func (o *CloudInsightsEventsListEventDetails) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *CloudInsightsEventsListEventDetails) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *CloudInsightsEventsListEventDetails) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *CloudInsightsEventsListEventDetails) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetUserIdentity

`func (o *CloudInsightsEventsListEventDetails) GetUserIdentity() CloudInsightsEventsListEventDetailsUserIdentity`

GetUserIdentity returns the UserIdentity field if non-nil, zero value otherwise.

### GetUserIdentityOk

`func (o *CloudInsightsEventsListEventDetails) GetUserIdentityOk() (*CloudInsightsEventsListEventDetailsUserIdentity, bool)`

GetUserIdentityOk returns a tuple with the UserIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdentity

`func (o *CloudInsightsEventsListEventDetails) SetUserIdentity(v CloudInsightsEventsListEventDetailsUserIdentity)`

SetUserIdentity sets UserIdentity field to given value.

### HasUserIdentity

`func (o *CloudInsightsEventsListEventDetails) HasUserIdentity() bool`

HasUserIdentity returns a boolean if a field has been set.

### GetVpcEndpointId

`func (o *CloudInsightsEventsListEventDetails) GetVpcEndpointId() string`

GetVpcEndpointId returns the VpcEndpointId field if non-nil, zero value otherwise.

### GetVpcEndpointIdOk

`func (o *CloudInsightsEventsListEventDetails) GetVpcEndpointIdOk() (*string, bool)`

GetVpcEndpointIdOk returns a tuple with the VpcEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcEndpointId

`func (o *CloudInsightsEventsListEventDetails) SetVpcEndpointId(v string)`

SetVpcEndpointId sets VpcEndpointId field to given value.

### HasVpcEndpointId

`func (o *CloudInsightsEventsListEventDetails) HasVpcEndpointId() bool`

HasVpcEndpointId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


