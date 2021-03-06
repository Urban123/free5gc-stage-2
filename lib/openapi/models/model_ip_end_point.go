/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type IpEndPoint struct {
	Ipv4Address string `json:"ipv4Address,omitempty" bson:"ipv4Address"`

	Ipv6Address string `json:"ipv6Address,omitempty" bson:"ipv6Address"`

	Transport TransportProtocol `json:"transport,omitempty" bson:"transport"`

	Port int32 `json:"port,omitempty" bson:"port"`
}
