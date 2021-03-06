/*
 * Gaia-Lite for Cosmos
 *
 * A REST interface for state queries, transaction generation and broadcasting.
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// TxQuery struct for TxQuery
type TxQuery struct {
	Code int32 `json:"code,omitempty"`
	Txhash string `json:"txhash,omitempty"`
	Height string `json:"height,omitempty"`
	Tx StdTx `json:"tx,omitempty"`
	Result TxQueryResult `json:"result,omitempty"`
}
