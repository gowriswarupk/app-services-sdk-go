/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage Kafka instances.
 *
 * API version: 1.5.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
)

// SupportedKafkaSize Supported Kafka Size
type SupportedKafkaSize struct {

	// Unique identifier of this Kafka instance size.
	Id *string `json:"id,omitempty"`

	IngressThroughputPerSec *SupportedKafkaSizeBytesValueItem `json:"ingress_throughput_per_sec,omitempty"`

	EgressThroughputPerSec *SupportedKafkaSizeBytesValueItem `json:"egress_throughput_per_sec,omitempty"`

	// Maximum amount of total connections available to this Kafka instance size.
	TotalMaxConnections *int32 `json:"total_max_connections,omitempty"`

	MaxDataRetentionSize *SupportedKafkaSizeBytesValueItem `json:"max_data_retention_size,omitempty"`

	// Maximum amount of total partitions available to this Kafka instance size.
	MaxPartitions *int32 `json:"max_partitions,omitempty"`

	// Maximum data retention period available to this Kafka instance size.
	MaxDataRetentionPeriod *string `json:"max_data_retention_period,omitempty"`

	// Maximium connection attempts per second available to this Kafka instance size.
	MaxConnectionAttemptsPerSec *int32 `json:"max_connection_attempts_per_sec,omitempty"`

	// Quota consumed by this Kafka instance size.
	QuotaConsumed *int32 `json:"quota_consumed,omitempty"`

	// Quota type used by this Kafka instance size.
	QuotaType *string `json:"quota_type,omitempty"`

	// Data plane cluster capacity consumed by this Kafka instance size.
	CapacityConsumed *int32 `json:"capacity_consumed,omitempty"`

}

// NewSupportedKafkaSize instantiates a new SupportedKafkaSize object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportedKafkaSize() *SupportedKafkaSize {
	this := SupportedKafkaSize{}
	return &this
}

// NewSupportedKafkaSizeWithDefaults instantiates a new SupportedKafkaSize object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportedKafkaSizeWithDefaults() *SupportedKafkaSize {
	this := SupportedKafkaSize{}












	return &this
}


// GetId returns the Id field value if set, zero value otherwise.
func (o *SupportedKafkaSize) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaSize) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SupportedKafkaSize) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SupportedKafkaSize) SetId(v string) {
	o.Id = &v
}


// GetIngressThroughputPerSec returns the IngressThroughputPerSec field value if set, zero value otherwise.
func (o *SupportedKafkaSize) GetIngressThroughputPerSec() SupportedKafkaSizeBytesValueItem {
	if o == nil || o.IngressThroughputPerSec == nil {
		var ret SupportedKafkaSizeBytesValueItem
		return ret
	}
	return *o.IngressThroughputPerSec
}

// GetIngressThroughputPerSecOk returns a tuple with the IngressThroughputPerSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaSize) GetIngressThroughputPerSecOk() (*SupportedKafkaSizeBytesValueItem, bool) {
	if o == nil || o.IngressThroughputPerSec == nil {
		return nil, false
	}
	return o.IngressThroughputPerSec, true
}

// HasIngressThroughputPerSec returns a boolean if a field has been set.
func (o *SupportedKafkaSize) HasIngressThroughputPerSec() bool {
	if o != nil && o.IngressThroughputPerSec != nil {
		return true
	}

	return false
}

// SetIngressThroughputPerSec gets a reference to the given SupportedKafkaSizeBytesValueItem and assigns it to the IngressThroughputPerSec field.
func (o *SupportedKafkaSize) SetIngressThroughputPerSec(v SupportedKafkaSizeBytesValueItem) {
	o.IngressThroughputPerSec = &v
}


// GetEgressThroughputPerSec returns the EgressThroughputPerSec field value if set, zero value otherwise.
func (o *SupportedKafkaSize) GetEgressThroughputPerSec() SupportedKafkaSizeBytesValueItem {
	if o == nil || o.EgressThroughputPerSec == nil {
		var ret SupportedKafkaSizeBytesValueItem
		return ret
	}
	return *o.EgressThroughputPerSec
}

// GetEgressThroughputPerSecOk returns a tuple with the EgressThroughputPerSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaSize) GetEgressThroughputPerSecOk() (*SupportedKafkaSizeBytesValueItem, bool) {
	if o == nil || o.EgressThroughputPerSec == nil {
		return nil, false
	}
	return o.EgressThroughputPerSec, true
}

// HasEgressThroughputPerSec returns a boolean if a field has been set.
func (o *SupportedKafkaSize) HasEgressThroughputPerSec() bool {
	if o != nil && o.EgressThroughputPerSec != nil {
		return true
	}

	return false
}

// SetEgressThroughputPerSec gets a reference to the given SupportedKafkaSizeBytesValueItem and assigns it to the EgressThroughputPerSec field.
func (o *SupportedKafkaSize) SetEgressThroughputPerSec(v SupportedKafkaSizeBytesValueItem) {
	o.EgressThroughputPerSec = &v
}


// GetTotalMaxConnections returns the TotalMaxConnections field value if set, zero value otherwise.
func (o *SupportedKafkaSize) GetTotalMaxConnections() int32 {
	if o == nil || o.TotalMaxConnections == nil {
		var ret int32
		return ret
	}
	return *o.TotalMaxConnections
}

// GetTotalMaxConnectionsOk returns a tuple with the TotalMaxConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaSize) GetTotalMaxConnectionsOk() (*int32, bool) {
	if o == nil || o.TotalMaxConnections == nil {
		return nil, false
	}
	return o.TotalMaxConnections, true
}

// HasTotalMaxConnections returns a boolean if a field has been set.
func (o *SupportedKafkaSize) HasTotalMaxConnections() bool {
	if o != nil && o.TotalMaxConnections != nil {
		return true
	}

	return false
}

// SetTotalMaxConnections gets a reference to the given int32 and assigns it to the TotalMaxConnections field.
func (o *SupportedKafkaSize) SetTotalMaxConnections(v int32) {
	o.TotalMaxConnections = &v
}


// GetMaxDataRetentionSize returns the MaxDataRetentionSize field value if set, zero value otherwise.
func (o *SupportedKafkaSize) GetMaxDataRetentionSize() SupportedKafkaSizeBytesValueItem {
	if o == nil || o.MaxDataRetentionSize == nil {
		var ret SupportedKafkaSizeBytesValueItem
		return ret
	}
	return *o.MaxDataRetentionSize
}

// GetMaxDataRetentionSizeOk returns a tuple with the MaxDataRetentionSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaSize) GetMaxDataRetentionSizeOk() (*SupportedKafkaSizeBytesValueItem, bool) {
	if o == nil || o.MaxDataRetentionSize == nil {
		return nil, false
	}
	return o.MaxDataRetentionSize, true
}

// HasMaxDataRetentionSize returns a boolean if a field has been set.
func (o *SupportedKafkaSize) HasMaxDataRetentionSize() bool {
	if o != nil && o.MaxDataRetentionSize != nil {
		return true
	}

	return false
}

// SetMaxDataRetentionSize gets a reference to the given SupportedKafkaSizeBytesValueItem and assigns it to the MaxDataRetentionSize field.
func (o *SupportedKafkaSize) SetMaxDataRetentionSize(v SupportedKafkaSizeBytesValueItem) {
	o.MaxDataRetentionSize = &v
}


// GetMaxPartitions returns the MaxPartitions field value if set, zero value otherwise.
func (o *SupportedKafkaSize) GetMaxPartitions() int32 {
	if o == nil || o.MaxPartitions == nil {
		var ret int32
		return ret
	}
	return *o.MaxPartitions
}

// GetMaxPartitionsOk returns a tuple with the MaxPartitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaSize) GetMaxPartitionsOk() (*int32, bool) {
	if o == nil || o.MaxPartitions == nil {
		return nil, false
	}
	return o.MaxPartitions, true
}

// HasMaxPartitions returns a boolean if a field has been set.
func (o *SupportedKafkaSize) HasMaxPartitions() bool {
	if o != nil && o.MaxPartitions != nil {
		return true
	}

	return false
}

// SetMaxPartitions gets a reference to the given int32 and assigns it to the MaxPartitions field.
func (o *SupportedKafkaSize) SetMaxPartitions(v int32) {
	o.MaxPartitions = &v
}


// GetMaxDataRetentionPeriod returns the MaxDataRetentionPeriod field value if set, zero value otherwise.
func (o *SupportedKafkaSize) GetMaxDataRetentionPeriod() string {
	if o == nil || o.MaxDataRetentionPeriod == nil {
		var ret string
		return ret
	}
	return *o.MaxDataRetentionPeriod
}

// GetMaxDataRetentionPeriodOk returns a tuple with the MaxDataRetentionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaSize) GetMaxDataRetentionPeriodOk() (*string, bool) {
	if o == nil || o.MaxDataRetentionPeriod == nil {
		return nil, false
	}
	return o.MaxDataRetentionPeriod, true
}

// HasMaxDataRetentionPeriod returns a boolean if a field has been set.
func (o *SupportedKafkaSize) HasMaxDataRetentionPeriod() bool {
	if o != nil && o.MaxDataRetentionPeriod != nil {
		return true
	}

	return false
}

// SetMaxDataRetentionPeriod gets a reference to the given string and assigns it to the MaxDataRetentionPeriod field.
func (o *SupportedKafkaSize) SetMaxDataRetentionPeriod(v string) {
	o.MaxDataRetentionPeriod = &v
}


// GetMaxConnectionAttemptsPerSec returns the MaxConnectionAttemptsPerSec field value if set, zero value otherwise.
func (o *SupportedKafkaSize) GetMaxConnectionAttemptsPerSec() int32 {
	if o == nil || o.MaxConnectionAttemptsPerSec == nil {
		var ret int32
		return ret
	}
	return *o.MaxConnectionAttemptsPerSec
}

// GetMaxConnectionAttemptsPerSecOk returns a tuple with the MaxConnectionAttemptsPerSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaSize) GetMaxConnectionAttemptsPerSecOk() (*int32, bool) {
	if o == nil || o.MaxConnectionAttemptsPerSec == nil {
		return nil, false
	}
	return o.MaxConnectionAttemptsPerSec, true
}

// HasMaxConnectionAttemptsPerSec returns a boolean if a field has been set.
func (o *SupportedKafkaSize) HasMaxConnectionAttemptsPerSec() bool {
	if o != nil && o.MaxConnectionAttemptsPerSec != nil {
		return true
	}

	return false
}

// SetMaxConnectionAttemptsPerSec gets a reference to the given int32 and assigns it to the MaxConnectionAttemptsPerSec field.
func (o *SupportedKafkaSize) SetMaxConnectionAttemptsPerSec(v int32) {
	o.MaxConnectionAttemptsPerSec = &v
}


// GetQuotaConsumed returns the QuotaConsumed field value if set, zero value otherwise.
func (o *SupportedKafkaSize) GetQuotaConsumed() int32 {
	if o == nil || o.QuotaConsumed == nil {
		var ret int32
		return ret
	}
	return *o.QuotaConsumed
}

// GetQuotaConsumedOk returns a tuple with the QuotaConsumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaSize) GetQuotaConsumedOk() (*int32, bool) {
	if o == nil || o.QuotaConsumed == nil {
		return nil, false
	}
	return o.QuotaConsumed, true
}

// HasQuotaConsumed returns a boolean if a field has been set.
func (o *SupportedKafkaSize) HasQuotaConsumed() bool {
	if o != nil && o.QuotaConsumed != nil {
		return true
	}

	return false
}

// SetQuotaConsumed gets a reference to the given int32 and assigns it to the QuotaConsumed field.
func (o *SupportedKafkaSize) SetQuotaConsumed(v int32) {
	o.QuotaConsumed = &v
}


// GetQuotaType returns the QuotaType field value if set, zero value otherwise.
func (o *SupportedKafkaSize) GetQuotaType() string {
	if o == nil || o.QuotaType == nil {
		var ret string
		return ret
	}
	return *o.QuotaType
}

// GetQuotaTypeOk returns a tuple with the QuotaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaSize) GetQuotaTypeOk() (*string, bool) {
	if o == nil || o.QuotaType == nil {
		return nil, false
	}
	return o.QuotaType, true
}

// HasQuotaType returns a boolean if a field has been set.
func (o *SupportedKafkaSize) HasQuotaType() bool {
	if o != nil && o.QuotaType != nil {
		return true
	}

	return false
}

// SetQuotaType gets a reference to the given string and assigns it to the QuotaType field.
func (o *SupportedKafkaSize) SetQuotaType(v string) {
	o.QuotaType = &v
}


// GetCapacityConsumed returns the CapacityConsumed field value if set, zero value otherwise.
func (o *SupportedKafkaSize) GetCapacityConsumed() int32 {
	if o == nil || o.CapacityConsumed == nil {
		var ret int32
		return ret
	}
	return *o.CapacityConsumed
}

// GetCapacityConsumedOk returns a tuple with the CapacityConsumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaSize) GetCapacityConsumedOk() (*int32, bool) {
	if o == nil || o.CapacityConsumed == nil {
		return nil, false
	}
	return o.CapacityConsumed, true
}

// HasCapacityConsumed returns a boolean if a field has been set.
func (o *SupportedKafkaSize) HasCapacityConsumed() bool {
	if o != nil && o.CapacityConsumed != nil {
		return true
	}

	return false
}

// SetCapacityConsumed gets a reference to the given int32 and assigns it to the CapacityConsumed field.
func (o *SupportedKafkaSize) SetCapacityConsumed(v int32) {
	o.CapacityConsumed = &v
}


func (o SupportedKafkaSize) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
    
	if o.IngressThroughputPerSec != nil {
		toSerialize["ingress_throughput_per_sec"] = o.IngressThroughputPerSec
	}
    
	if o.EgressThroughputPerSec != nil {
		toSerialize["egress_throughput_per_sec"] = o.EgressThroughputPerSec
	}
    
	if o.TotalMaxConnections != nil {
		toSerialize["total_max_connections"] = o.TotalMaxConnections
	}
    
	if o.MaxDataRetentionSize != nil {
		toSerialize["max_data_retention_size"] = o.MaxDataRetentionSize
	}
    
	if o.MaxPartitions != nil {
		toSerialize["max_partitions"] = o.MaxPartitions
	}
    
	if o.MaxDataRetentionPeriod != nil {
		toSerialize["max_data_retention_period"] = o.MaxDataRetentionPeriod
	}
    
	if o.MaxConnectionAttemptsPerSec != nil {
		toSerialize["max_connection_attempts_per_sec"] = o.MaxConnectionAttemptsPerSec
	}
    
	if o.QuotaConsumed != nil {
		toSerialize["quota_consumed"] = o.QuotaConsumed
	}
    
	if o.QuotaType != nil {
		toSerialize["quota_type"] = o.QuotaType
	}
    
	if o.CapacityConsumed != nil {
		toSerialize["capacity_consumed"] = o.CapacityConsumed
	}
    
	return json.Marshal(toSerialize)
}

type NullableSupportedKafkaSize struct {
	value *SupportedKafkaSize
	isSet bool
}

func (v NullableSupportedKafkaSize) Get() *SupportedKafkaSize {
	return v.value
}

func (v *NullableSupportedKafkaSize) Set(val *SupportedKafkaSize) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportedKafkaSize) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportedKafkaSize) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportedKafkaSize(val *SupportedKafkaSize) *NullableSupportedKafkaSize {
	return &NullableSupportedKafkaSize{value: val, isSet: true}
}

func (v NullableSupportedKafkaSize) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportedKafkaSize) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

