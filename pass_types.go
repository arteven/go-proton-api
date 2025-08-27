package proton

// Pass API response types matching Proton Pass API

// PassShareResponse represents a Pass share as returned by the API
type PassShareResponse struct {
	ShareID              string `json:"ShareID"`
	VaultID              string `json:"VaultID"`
	AddressID            string `json:"AddressID"`
	Primary              bool   `json:"Primary"`
	Owner                bool   `json:"Owner"`
	TargetType           int    `json:"TargetType"`
	TargetID             string `json:"TargetID"`
	TargetMembers        int    `json:"TargetMembers"`
	TargetMaxMembers     int    `json:"TargetMaxMembers"`
	Shared               bool   `json:"Shared"`
	PendingInvites       int    `json:"PendingInvites"`
	Permission           int    `json:"Permission"`
	Content              string `json:"Content,omitempty"`           // Base64 encrypted content
	ContentKeyRotation   *int   `json:"ContentKeyRotation,omitempty"` // Key rotation used for content
	ContentFormatVersion int    `json:"ContentFormatVersion"`
	CreateTime           int64  `json:"CreateTime"`
	ModifyTime           int64  `json:"ModifyTime"`
	ExpireTime           *int64 `json:"ExpireTime,omitempty"`
}

// PassShareKeyResponse represents a Pass share key as returned by the API
type PassShareKeyResponse struct {
	KeyRotation int    `json:"KeyRotation"`
	Key         string `json:"Key"`         // Base64 PGP-encrypted key
	KeyFlags    int    `json:"KeyFlags"`
	UserKeyID   string `json:"UserKeyID"`   // ID of user key used to encrypt this
	CreateTime  int64  `json:"CreateTime"`
}

// PassSharesResponse represents the response from listing Pass shares
type PassSharesResponse struct {
	Code   int                 `json:"Code"`
	Shares []PassShareResponse `json:"Shares"`
}

// PassShareKeysResponse represents the response from getting Pass share keys
type PassShareKeysResponse struct {
	Code      int `json:"Code"`
	ShareKeys struct {
		Keys []PassShareKeyResponse `json:"Keys"`
	} `json:"ShareKeys"`
}