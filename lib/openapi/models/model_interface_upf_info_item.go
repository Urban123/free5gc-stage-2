/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type InterfaceUpfInfoItem struct {
	InterfaceType UpInterfaceType `json:"interfaceType" bson:"interfaceType"`

	Ipv4EndpointAddresses []string `json:"ipv4EndpointAddresses,omitempty" bson:"ipv4EndpointAddresses"`

	Ipv6EndpointAddresses []string `json:"ipv6EndpointAddresses,omitempty" bson:"ipv6EndpointAddresses"`

	EndpointFqdn string `json:"endpointFqdn,omitempty" bson:"endpointFqdn"`

	NetworkInstance string `json:"networkInstance,omitempty" bson:"networkInstance"`
}
