package referrals

type UploadReferralsStructuredResponse struct {
	ReferralServiceResult struct {
		Success bool `json:"success"`
	} `json:"referralServiceResult"`
	SkippedReferrals []interface{} `json:"skippedReferrals"`
}
