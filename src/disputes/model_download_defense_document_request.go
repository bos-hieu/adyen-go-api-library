/*
 * Dispute service API
 *
 * This API provides endpoints for dispute management. You can use the API to retrieve defense reasons, supply defense ducoments, delete defense documents, upload chargebacks or accept disputes.  For more information on using the APIs, refer to [Dispute service](https://docs.adyen.com/risk-management/disputes-api).  ## Authentication Each request to the Dispute API must be signed with an API key. For this, obtain an API Key from your Customer Area, as described in [How to get the API key](https://docs.adyen.com/user-management/how-to-get-the-api-key). Then set this key to the `X-API-Key` header value, for example:  ``` curl -H \"Content-Type: application/json\" \\ -H \"X-API-Key: Your_API_key\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ```
 *
 * API version: 1
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package disputes
// DownloadDefenseDocumentRequest struct for DownloadDefenseDocumentRequest
type DownloadDefenseDocumentRequest struct {
	// The specific type of the defense document
	DefenseDocumentType string `json:"defenseDocumentType"`
	// The psp reference of the dispute
	DisputePspReference string `json:"disputePspReference"`
	// The merchant account identifier
	MerchantAccountCode string `json:"merchantAccountCode"`
}