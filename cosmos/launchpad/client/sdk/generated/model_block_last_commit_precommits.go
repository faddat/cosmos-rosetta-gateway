/*
 * Gaia-Lite for Cosmos
 *
 * A REST interface for state queries, transaction generation and broadcasting.
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// BlockLastCommitPrecommits struct for BlockLastCommitPrecommits
type BlockLastCommitPrecommits struct {
	ValidatorAddress string `json:"validator_address,omitempty"`
	ValidatorIndex string `json:"validator_index,omitempty"`
	Height string `json:"height,omitempty"`
	Round string `json:"round,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Type float32 `json:"type,omitempty"`
	BlockId BlockId `json:"block_id,omitempty"`
	Signature string `json:"signature,omitempty"`
}
