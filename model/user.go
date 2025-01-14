package model

import (
	"strings"
	"time"
)

type AccountTierEnum uint

const (
	Basic AccountTierEnum = iota
	Free
	ProTrial
	Pro
	Enterprise
)

type LyridUser struct {
	Id              string    `json:"id" binding:"required"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	EmailVerified   bool      `json:"emailVerified"`
	Roles           []string  `json:"roles"`
	RelatedRoles    []string  `json:"relatedRoles`
	RelatedAccounts []string  `json:"relatedAccounts`
	BetaAccess      bool      `json:"betaAccess"`
	CurrentAccount  string    `json:"currentAccount"`
	DefaultAccount  string    `json:"defaultAccount"`
	LastUpdate      time.Time `json:"lastUpdate"`
}

type Account struct {
	Id        string    			`json:"id"`
	Name      string    			`json:"name" binding:"required"`
	Tier      AccountTierEnum       `json:"tier"`
	TrialUsed bool            		`json:"trialUsed"`
	CreatedOn time.Time 			`json:"createdOn"`
	CreatedBy string    			`json:"createdBy"`
}

func (account *Account) GetBucketName() string {
	return strings.ReplaceAll(account.Id, "-", "")
}

func (account *Account) GetS3BucketName(region string) string {
	return "lyrid-lambda-" + account.GetBucketName() + "-" + region
}
