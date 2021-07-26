/*
 * Kafka Admin REST API
 *
 * An API to provide REST endpoints for query Kafka for admin operations
 *
 * API version: 0.3.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
)

// NewTopicInput Input object to create a new topic.
type NewTopicInput struct {

	// The topic name, this value must be unique.
	Name string `json:"name"`

	Settings TopicSettings `json:"settings"`

}

// NewNewTopicInput instantiates a new NewTopicInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewTopicInput(name string, settings TopicSettings) *NewTopicInput {
	this := NewTopicInput{}
	this.Name = name
	this.Settings = settings
	return &this
}

// NewNewTopicInputWithDefaults instantiates a new NewTopicInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewTopicInputWithDefaults() *NewTopicInput {
	this := NewTopicInput{}



	return &this
}


// GetName returns the Name field value
func (o *NewTopicInput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NewTopicInput) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NewTopicInput) SetName(v string) {
	o.Name = v
}


// GetSettings returns the Settings field value
func (o *NewTopicInput) GetSettings() TopicSettings {
	if o == nil {
		var ret TopicSettings
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *NewTopicInput) GetSettingsOk() (*TopicSettings, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *NewTopicInput) SetSettings(v TopicSettings) {
	o.Settings = v
}


func (o NewTopicInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if true {
		toSerialize["name"] = o.Name
	}
    
	if true {
		toSerialize["settings"] = o.Settings
	}
    
	return json.Marshal(toSerialize)
}

type NullableNewTopicInput struct {
	value *NewTopicInput
	isSet bool
}

func (v NullableNewTopicInput) Get() *NewTopicInput {
	return v.value
}

func (v *NullableNewTopicInput) Set(val *NewTopicInput) {
	v.value = val
	v.isSet = true
}

func (v NullableNewTopicInput) IsSet() bool {
	return v.isSet
}

func (v *NullableNewTopicInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewTopicInput(val *NewTopicInput) *NullableNewTopicInput {
	return &NullableNewTopicInput{value: val, isSet: true}
}

func (v NullableNewTopicInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewTopicInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

