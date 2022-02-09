/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage Kafka instances.
 *
 * API version: 1.3.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
)

// RegionCapacityListItem schema for a kafka instance type capacity in region
type RegionCapacityListItem struct {

	// kafka instance type
	InstanceType *string `json:"instance_type,omitempty"`

	// flag indicating whether the capacity for the instance type in the region is reached
	MaxCapacityReached bool `json:"max_capacity_reached"`

}

// NewRegionCapacityListItem instantiates a new RegionCapacityListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionCapacityListItem(maxCapacityReached bool) *RegionCapacityListItem {
	this := RegionCapacityListItem{}
	this.MaxCapacityReached = maxCapacityReached
	return &this
}

// NewRegionCapacityListItemWithDefaults instantiates a new RegionCapacityListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionCapacityListItemWithDefaults() *RegionCapacityListItem {
	this := RegionCapacityListItem{}



	return &this
}


// GetInstanceType returns the InstanceType field value if set, zero value otherwise.
func (o *RegionCapacityListItem) GetInstanceType() string {
	if o == nil || o.InstanceType == nil {
		var ret string
		return ret
	}
	return *o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionCapacityListItem) GetInstanceTypeOk() (*string, bool) {
	if o == nil || o.InstanceType == nil {
		return nil, false
	}
	return o.InstanceType, true
}

// HasInstanceType returns a boolean if a field has been set.
func (o *RegionCapacityListItem) HasInstanceType() bool {
	if o != nil && o.InstanceType != nil {
		return true
	}

	return false
}

// SetInstanceType gets a reference to the given string and assigns it to the InstanceType field.
func (o *RegionCapacityListItem) SetInstanceType(v string) {
	o.InstanceType = &v
}


// GetMaxCapacityReached returns the MaxCapacityReached field value
func (o *RegionCapacityListItem) GetMaxCapacityReached() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MaxCapacityReached
}

// GetMaxCapacityReachedOk returns a tuple with the MaxCapacityReached field value
// and a boolean to check if the value has been set.
func (o *RegionCapacityListItem) GetMaxCapacityReachedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MaxCapacityReached, true
}

// SetMaxCapacityReached sets field value
func (o *RegionCapacityListItem) SetMaxCapacityReached(v bool) {
	o.MaxCapacityReached = v
}


func (o RegionCapacityListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.InstanceType != nil {
		toSerialize["instance_type"] = o.InstanceType
	}
    
	if true {
		toSerialize["max_capacity_reached"] = o.MaxCapacityReached
	}
    
	return json.Marshal(toSerialize)
}

type NullableRegionCapacityListItem struct {
	value *RegionCapacityListItem
	isSet bool
}

func (v NullableRegionCapacityListItem) Get() *RegionCapacityListItem {
	return v.value
}

func (v *NullableRegionCapacityListItem) Set(val *RegionCapacityListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionCapacityListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionCapacityListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionCapacityListItem(val *RegionCapacityListItem) *NullableRegionCapacityListItem {
	return &NullableRegionCapacityListItem{value: val, isSet: true}
}

func (v NullableRegionCapacityListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionCapacityListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

