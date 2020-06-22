/*
 * Adyen for Platforms: Account API
 *
 * The Account API provides endpoints for managing account-related entities on your platform. These related entities include account holders, accounts, bank accounts, shareholders, and KYC-related documents. The management operations include actions such as creation, retrieval, updating, and deletion of them.  For more information, refer to our [documentation](https://docs.adyen.com/marketpay). ## Authentication To connect to the Account API, you must use basic authentication credentials of your web service user. If you don't have one, please contact the [Adyen Support Team](https://support.adyen.com/hc/en-us/requests/new). Then use its credentials to authenticate your request, for example:  ``` curl -U \"ws@MarketPlace.YourMarketPlace\":\"YourWsPassword\" \\ -H \"Content-Type: application/json\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning The Account API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://cal-test.adyen.com/cal/services/Account/v5/createAccountHolder ```
 *
 * API version: 5
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package platformsaccount
// CreateAccountHolderResponse struct for CreateAccountHolderResponse
type CreateAccountHolderResponse struct {
	// The code of a new account created for the account holder.
	AccountCode string `json:"accountCode,omitempty"`
	// The code of the new account holder.
	AccountHolderCode string `json:"accountHolderCode"`
	AccountHolderDetails AccountHolderDetails `json:"accountHolderDetails"`
	AccountHolderStatus AccountHolderStatus `json:"accountHolderStatus"`
	// The description of the new account holder.
	Description string `json:"description,omitempty"`
	// A list of fields that caused the `/createAccountHolder` request to fail.
	InvalidFields *[]ErrorFieldType `json:"invalidFields,omitempty"`
	// The type of legal entity of the new account holder.
	LegalEntity string `json:"legalEntity"`
	// The three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes), with which the prospective account holder primarily deals.
	PrimaryCurrency string `json:"primaryCurrency,omitempty"`
	// The reference of a request.  Can be used to uniquely identify the request.
	PspReference string `json:"pspReference"`
	// The result code.
	ResultCode string `json:"resultCode,omitempty"`
	Verification KYCVerificationResult `json:"verification"`
}
