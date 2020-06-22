/*
 * Adyen for Platforms: Fund API
 *
 * The Fund API provides endpoints for managing the funds in the accounts on your platform. These management operations include actions such as the transfer of funds from one account to another, the payout of funds to an account holder, and the retrieval of balances in an account.  For more information, refer to our [documentation](https://docs.adyen.com/marketpay). ## Authentication To connect to the Fund API, you must use basic authentication credentials of your web service user. If you don't have one, please contact the [Adyen Support Team](https://support.adyen.com/hc/en-us/requests/new). Then use its credentials to authenticate your request, for example:  ``` curl -U \"ws@MarketPlace.YourMarketPlace\":\"YourWsPassword\" \\ -H \"Content-Type: application/json\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning The Fund API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://cal-test.adyen.com/cal/services/Fund/v5/accountHolderBalance ```
 *
 * API version: 5
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package platformsfund
import (
	"time"
)
// Transaction struct for Transaction
type Transaction struct {
	Amount *Amount `json:"amount,omitempty"`
	BankAccountDetail *BankAccountDetail `json:"bankAccountDetail,omitempty"`
	// The merchant reference of a related capture.
	CaptureMerchantReference string `json:"captureMerchantReference,omitempty"`
	// The psp reference of a related capture.
	CapturePspReference string `json:"capturePspReference,omitempty"`
	// The date on which the transaction was performed.
	CreationDate *time.Time `json:"creationDate,omitempty"`
	// A description of the transaction.
	Description string `json:"description,omitempty"`
	// The code of the account to which funds were credited during an outgoing fund transfer.
	DestinationAccountCode string `json:"destinationAccountCode,omitempty"`
	// The psp reference of the related dispute.
	DisputePspReference string `json:"disputePspReference,omitempty"`
	// The reason code of a dispute.
	DisputeReasonCode string `json:"disputeReasonCode,omitempty"`
	// The merchant reference of a transaction.
	MerchantReference string `json:"merchantReference,omitempty"`
	// The psp reference of the related authorisation or transfer.
	PaymentPspReference string `json:"paymentPspReference,omitempty"`
	// The psp reference of the related payout.
	PayoutPspReference string `json:"payoutPspReference,omitempty"`
	// The psp reference of a transaction.
	PspReference string `json:"pspReference,omitempty"`
	// The code of the account from which funds were debited during an incoming fund transfer.
	SourceAccountCode string `json:"sourceAccountCode,omitempty"`
	// The status of the transaction. >Permitted values: `PendingCredit`, `CreditFailed`, `Credited`, `Converted`, `PendingDebit`, `DebitFailed`, `Debited`, `DebitReversedReceived`, `DebitedReversed`, `ChargebackReceived`, `Chargeback`, `ChargebackReversedReceived`, `ChargebackReversed`, `Payout`, `PayoutReversed`, `FundTransfer`, `PendingFundTransfer`, `ManualCorrected`.
	TransactionStatus string `json:"transactionStatus,omitempty"`
	// The transfer code of the transaction.
	TransferCode string `json:"transferCode,omitempty"`
}
