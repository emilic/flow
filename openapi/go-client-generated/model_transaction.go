/*
 * Access API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Transaction struct {
	Id string `json:"id"`
	// Base64 encoded Cadence script.
	Script string `json:"script"`
	// Array of Base64 encoded arguments with in [JSON-Cadence interchange format](https://docs.onflow.org/cadence/json-cadence-spec/).
	Arguments []string `json:"arguments"`
	ReferenceBlockId string `json:"reference_block_id"`
	// The limit on the amount of computation a transaction is allowed to preform.
	GasLimit string `json:"gas_limit"`
	Payer string `json:"payer"`
	ProposalKey *ProposalKey `json:"proposal_key"`
	Authorizers []string `json:"authorizers"`
	PayloadSignatures []TransactionSignature `json:"payload_signatures"`
	EnvelopeSignatures []TransactionSignature `json:"envelope_signatures"`
	Result *TransactionResult `json:"result,omitempty"`
	Expandable *TransactionExpandable `json:"_expandable"`
	Links *Links `json:"_links,omitempty"`
}
