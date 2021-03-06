/*
 * Gaia-Lite for Cosmos
 *
 * A REST interface for state queries, transaction generation and broadcasting.
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Delegation struct for Delegation
type Delegation struct {
	DelegatorAddress string `json:"delegator_address,omitempty"`
	ValidatorAddress string `json:"validator_address,omitempty"`
	Shares string `json:"shares,omitempty"`
	Balance Coin `json:"balance,omitempty"`
}
