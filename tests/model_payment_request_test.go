/*
 * Adyen Checkout API
 *
 * Adyen Checkout API provides a simple and flexible way to initiate and authorise online payments. You can use the same integration for payments made with cards (including 3D Secure), mobile wallets, and local payment methods (for example, iDEAL and Sofort).  This API reference provides information on available endpoints and how to interact with them. To learn more about the API, visit [Checkout documentation](https://docs.adyen.com/online-payments).  ## Authentication Each request to the Checkout API must be signed with an API key. For this, obtain an API Key from your Customer Area, as described in [How to get the API key](https://docs.adyen.com/development-resources/api-credentials#generate-api-key). Then set this key to the `X-API-Key` header value, for example:  ``` curl -H \"Content-Type: application/json\" \\ -H \"X-API-Key: Your_Checkout_API_key\" \\ ... ``` Note that when going live, you need to generate a new API Key to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning Checkout API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://checkout-test.adyen.com/v67/payments ```
 *
 * API version: 67
 * Contact: developer-experience@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tests

import (
	"encoding/json"
	"testing"

	"github.com/adyen/adyen-go-api-library/v6/src/checkout"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPaymentRequest_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		req     checkout.PaymentRequest
		json    string
		wantErr bool
		wantFn  func(pm checkout.PaymentRequest, t *testing.T)
	}{
		{
			"unmarshalls an empty payment request correctly",
			checkout.PaymentRequest{},
			"{}",
			false,
			func(got checkout.PaymentRequest, t *testing.T) {
				require.NotNil(t, got)
				require.Nil(t, got.PaymentMethod)
			},
		},
		{
			"unmarshalls a payment request without payment method correctly",
			checkout.PaymentRequest{},
			`{
                "riskData":{"clientData":"eyJ2ZXJzaW9uIjoiMS4w"},
                "browserInfo":{"acceptHeader":"*/*","colorDepth":24,"language":"en-US","javaEnabled":false,"screenHeight":1080,"screenWidth":1920,"userAgent":"Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:85.0) Gecko/20100101 Firefox/85.0","timeZoneOffset":-60}
            }`,
			false,
			func(got checkout.PaymentRequest, t *testing.T) {
				require.NotNil(t, got)
				require.Nil(t, got.PaymentMethod)
				require.NotNil(t, got.RiskData)
				require.NotNil(t, got.BrowserInfo)
			},
		},
		{
			"unmarshalls a payment request with ideal correctly",
			checkout.PaymentRequest{},
			`{
                "riskData":{"clientData":"eyJ2ZXJzaW9uIjoiMS4"},
                "paymentMethod":{"type":"ideal","issuer":"1121"}
            }`,
			false,
			func(got checkout.PaymentRequest, t *testing.T) {
				require.NotNil(t, got)
				require.NotNil(t, got.PaymentMethod)
				require.NotNil(t, got.RiskData)
				assert.Equal(t, "ideal", got.PaymentMethod.(*checkout.IdealDetails).Type)
				assert.Equal(t, "1121", got.PaymentMethod.(*checkout.IdealDetails).Issuer)
			},
		},
		{
			"unmarshalls a payment request with scheme correctly",
			checkout.PaymentRequest{},
			`{
                "riskData":{"clientData":"eyJ2ZXJzaW9uIjoiMS4w"},
                "paymentMethod":{"type":"scheme","holderName":"d","encryptedCardNumber":"adyenjs_0_1_25$J8/5xp5l6DjYVPokO6FwAQj","encryptedExpiryMonth":"adyenjs_0_1_25$bLCWe/ZHR37Okz0d28bzrDBYXw","encryptedExpiryYear":"adyenjs_0_1_25$nqasksbOSfn0grzrmna2vpWkQMhOHT6Cd","encryptedSecurityCode":"adyenjs_0_1_25$TbomjrfaGwHFfxpPuf","brand":"amex"},
                "browserInfo":{"acceptHeader":"*/*","colorDepth":24,"language":"en-US","javaEnabled":false,"screenHeight":1080,"screenWidth":1920,"userAgent":"Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:85.0) Gecko/20100101 Firefox/85.0","timeZoneOffset":-60}
            }`,
			false,
			func(got checkout.PaymentRequest, t *testing.T) {
				require.NotNil(t, got)
				require.NotNil(t, got.PaymentMethod)
				require.NotNil(t, got.RiskData)
				require.NotNil(t, got.BrowserInfo)
				assert.Equal(t, "scheme", got.PaymentMethod.(*checkout.CardDetails).Type)
				assert.Equal(t, "adyenjs_0_1_25$J8/5xp5l6DjYVPokO6FwAQj", got.PaymentMethod.(*checkout.CardDetails).EncryptedCardNumber)

				jsonString, err := json.Marshal(got)
				assert.Nil(t, err)
				assert.Equal(t, `{"amount":{"currency":"","value":0},"browserInfo":{"acceptHeader":"*/*","colorDepth":24,"javaEnabled":false,"language":"en-US","screenHeight":1080,"screenWidth":1920,"timeZoneOffset":-60,"userAgent":"Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:85.0) Gecko/20100101 Firefox/85.0"},"merchantAccount":"","paymentMethod":{"brand":"amex","encryptedCardNumber":"adyenjs_0_1_25$J8/5xp5l6DjYVPokO6FwAQj","encryptedExpiryMonth":"adyenjs_0_1_25$bLCWe/ZHR37Okz0d28bzrDBYXw","encryptedExpiryYear":"adyenjs_0_1_25$nqasksbOSfn0grzrmna2vpWkQMhOHT6Cd","encryptedSecurityCode":"adyenjs_0_1_25$TbomjrfaGwHFfxpPuf","holderName":"d","type":"scheme"},"reference":"","returnUrl":"","riskData":{"clientData":"eyJ2ZXJzaW9uIjoiMS4w"}}`, string(jsonString))
			},
		},
		{
			"unmarshalls a payment request with ach correctly",
			checkout.PaymentRequest{},
			`{
                "riskData":{"clientData":"eyJ2ZX"},
                "paymentMethod":{"type":"ach","encryptedBankAccountNumber":"adyenjs_0_1_25","encryptedBankLocationId":"adyenjs_0_1_25","ownerName":"test"},
                "billingAddress":{"street":"test","houseNumberOrName":"2","postalCode":"123456","city":"porto rico","stateOrProvince":"N/A","country":"PR"}
            }`,
			false,
			func(got checkout.PaymentRequest, t *testing.T) {
				require.NotNil(t, got)
				require.NotNil(t, got.PaymentMethod)
				require.NotNil(t, got.RiskData)
				require.NotNil(t, got.BillingAddress)
				assert.Equal(t, "test", got.BillingAddress.Street)
				assert.Equal(t, "ach", got.PaymentMethod.(*checkout.AchDetails).Type)
				assert.Equal(t, "adyenjs_0_1_25", got.PaymentMethod.(*checkout.AchDetails).EncryptedBankAccountNumber)
			},
		},
		{
			"unmarshalls a payment request with klarna correctly",
			checkout.PaymentRequest{},
			`{
                "riskData":{"clientData":"eyJ2ZX"},
                "paymentMethod":{"type":"klarna","billingAddress":"adyenjs_0_1_25"},
                "billingAddress":{"street":"test","houseNumberOrName":"2","postalCode":"123456","city":"porto rico","stateOrProvince":"N/A","country":"PR"}
            }`,
			false,
			func(got checkout.PaymentRequest, t *testing.T) {
				require.NotNil(t, got)
				require.NotNil(t, got.PaymentMethod)
				require.NotNil(t, got.RiskData)
				require.NotNil(t, got.BillingAddress)
				assert.Equal(t, "test", got.BillingAddress.Street)
				assert.Equal(t, "klarna", got.PaymentMethod.(*checkout.KlarnaDetails).Type)
				assert.Equal(t, "adyenjs_0_1_25", got.PaymentMethod.(*checkout.KlarnaDetails).BillingAddress)
			},
		},
		{
			"unmarshalls a payment request with klarna_paynow correctly",
			checkout.PaymentRequest{},
			`{
                "riskData":{"clientData":"eyJ2ZX"},
                "paymentMethod":{"type":"klarna_paynow","billingAddress":"adyenjs_0_1_25"},
                "billingAddress":{"street":"test","houseNumberOrName":"2","postalCode":"123456","city":"porto rico","stateOrProvince":"N/A","country":"PR"}
            }`,
			false,
			func(got checkout.PaymentRequest, t *testing.T) {
				require.NotNil(t, got)
				require.NotNil(t, got.PaymentMethod)
				require.NotNil(t, got.RiskData)
				require.NotNil(t, got.BillingAddress)
				assert.Equal(t, "test", got.BillingAddress.Street)
				assert.Equal(t, "klarna_paynow", got.PaymentMethod.(*checkout.KlarnaDetails).Type)
				assert.Equal(t, "adyenjs_0_1_25", got.PaymentMethod.(*checkout.KlarnaDetails).BillingAddress)
			},
		},
		{
			"unmarshalls a payment request with a payment type not defined with concrete struct correctly",
			checkout.PaymentRequest{},
			`{
                "riskData":{"clientData":"eyJ2ZXJzaW9uIjoiMS4w"},
                "paymentMethod":{"type":"mypay","holderName":"d","number":"1234", "brand":"amex"},
                "browserInfo":{"acceptHeader":"*/*","colorDepth":24,"language":"en-US","javaEnabled":false,"screenHeight":1080,"screenWidth":1920,"userAgent":"Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:85.0) Gecko/20100101 Firefox/85.0","timeZoneOffset":-60}
            }`,
			false,
			func(got checkout.PaymentRequest, t *testing.T) {
				require.NotNil(t, got)
				require.NotNil(t, got.PaymentMethod)
				require.NotNil(t, got.RiskData)
				require.NotNil(t, got.BrowserInfo)
				assert.Equal(t, "mypay", got.PaymentMethod.(map[string]interface{})["type"].(string))
				assert.Equal(t, "1234", got.PaymentMethod.(map[string]interface{})["number"])
				assert.Equal(t, "amex", got.PaymentMethod.(map[string]interface{})["brand"])

				jsonString, err := json.Marshal(got)
				assert.Nil(t, err)
				assert.Equal(t, `{"amount":{"currency":"","value":0},"browserInfo":{"acceptHeader":"*/*","colorDepth":24,"javaEnabled":false,"language":"en-US","screenHeight":1080,"screenWidth":1920,"timeZoneOffset":-60,"userAgent":"Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:85.0) Gecko/20100101 Firefox/85.0"},"merchantAccount":"","paymentMethod":{"brand":"amex","holderName":"d","number":"1234","type":"mypay"},"reference":"","returnUrl":"","riskData":{"clientData":"eyJ2ZXJzaW9uIjoiMS4w"}}`, string(jsonString))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pm := checkout.PaymentRequest{}
			err := json.Unmarshal([]byte(tt.json), &pm)
			if (err != nil) != tt.wantErr {
				t.Errorf("PaymentRequest.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Nil(t, err)
			tt.wantFn(pm, t)
		})
	}
}
