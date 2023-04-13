package referrals

type UploadReferralsStructuredRequest struct {
	AccountCode  string      `json:"accountCode"`
	ReferralType string      `json:"referralType"`
	Action       string      `json:"action"`
	Referrals    []*Referral `json:"referrals"`
	Reason       string      `json:"reason"`
}

type Referral struct {
	ReferralContainer ReferralContainer `json:"referralContainer"`
}

type ReferralContainer struct {
	Referral string `json:"referral"`
}
