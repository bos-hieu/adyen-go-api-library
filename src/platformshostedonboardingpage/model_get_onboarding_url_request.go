/*
 * Adyen for Platforms: Hosted Onboarding Page API
 *
 * The Hosted Onboarding Page (HOP) API provides endpoints for using the Hosted Onboarding Page. The related entities include account holders only.  For more information, refer to our [documentation](https://docs.adyen.com/marketpay/onboarding-and-verification/hosted-onboarding-page). ## Authentication To connect to the HOP API, you must use basic authentication credentials of your web service user. If you don't have one, please contact the [Adyen Support Team](https://support.adyen.com/hc/en-us/requests/new). Then use its credentials to authenticate your request, for example:  ``` curl -U \"ws@MarketPlace.YourMarketPlace\":\"YourWsPassword\" \\ -H \"Content-Type: application/json\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning The HOP API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://cal-test.adyen.com/cal/services/Hop/v1/getOnboardingUrl ```
 *
 * API version: 1
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package platformshostedonboardingpage
// GetOnboardingUrlRequest struct for GetOnboardingUrlRequest
type GetOnboardingUrlRequest struct {
	// The account holder code you provided when you created the account holder.
	AccountHolderCode string `json:"accountHolderCode"`
	// The platform name which will show up in the welcome page.
	PlatformName string `json:"platformName,omitempty"`
	// The URL where the sub-merchant will be redirected back to after they complete the onboarding, or if their session times out. Maximum length of 500 characters. If you don't provide this, the sub-merchant will be redirected back to the default return URL configured in your platform account.
	ReturnUrl string `json:"returnUrl,omitempty"`
}
