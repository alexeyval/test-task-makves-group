package models

type User struct {
	UserID                    string  `json:"#"`
	ID                        int64   `json:"id"`
	UID                       string  `json:"uid"`
	Domain                    string  `json:"domain"`
	CN                        string  `json:"cn"`
	Department                string  `json:"department"`
	Title                     string  `json:"title"`
	Who                       string  `json:"who"`
	LogonCount                string  `json:"logon_count"`
	NumLogons7                string  `json:"num_logons7"`
	NumShare7                 string  `json:"num_share7"`
	NumFile7                  string  `json:"num_file7"`
	NumAd7                    string  `json:"num_ad7"`
	NumN7                     string  `json:"num_n7"`
	NumLogons14               string  `json:"num_logons14"`
	NumShare14                string  `json:"num_share14"`
	NumFile14                 string  `json:"num_file14"`
	NumAd14                   string  `json:"num_ad14"`
	NumN14                    string  `json:"num_n14"`
	NumLogons30               string  `json:"num_logons30"`
	NumShare30                string  `json:"num_share30"`
	NumFile30                 string  `json:"num_file30"`
	NumAd30                   string  `json:"num_ad30"`
	NumN30                    string  `json:"num_n30"`
	NumLogons150              string  `json:"num_logons150"`
	NumShare150               string  `json:"num_share150"`
	NumFile150                string  `json:"num_file150"`
	NumAd150                  string  `json:"num_ad150"`
	NumN150                   string  `json:"num_n150"`
	NumLogons365              string  `json:"num_logons365"`
	NumShare365               string  `json:"num_share365"`
	NumFile365                string  `json:"num_file365"`
	NumAd365                  string  `json:"num_ad365"`
	NumN365                   string  `json:"num_n365"`
	HasUserPrincipalName      string  `json:"has_user_principal_name"`
	HasMail                   string  `json:"has_mail"`
	HasPhone                  string  `json:"has_phone"`
	FlagDisabled              string  `json:"flag_disabled"`
	FlagLockout               string  `json:"flag_lockout"`
	FlagPasswordNotRequired   string  `json:"flag_password_not_required"`
	FlagPasswordCantChange    string  `json:"flag_password_cant_change"`
	FlagDontExpirePassword    string  `json:"flag_dont_expire_password"`
	OwnedFiles                string  `json:"owned_files"`
	NumMailboxes              string  `json:"num_mailboxes"`
	NumMemberOfGroups         string  `json:"num_member_of_groups"`
	NumMemberOfIndirectGroups string  `json:"num_member_of_indirect_groups"`
	MemberOfIndirectGroupsIds []int64 `json:"member_of_indirect_groups_ids,omitempty"`
	MemberOfGroupsIds         []int64 `json:"member_of_groups_ids,omitempty"`
	IsAdmin                   string  `json:"is_admin,omitempty"`
	IsService                 string  `json:"is_service"`
}

const (
	UserID = iota
	ID
	UID
	Domain
	CN
	Department
	Title
	Who
	LogonCount
	NumLogons7
	NumShare7
	NumFile7
	NumAd7
	NumN7
	NumLogons14
	NumShare14
	NumFile14
	NumAd14
	NumN14
	NumLogons30
	NumShare30
	NumFile30
	NumAd30
	NumN30
	NumLogons150
	NumShare150
	NumFile150
	NumAd150
	NumN150
	NumLogons365
	NumShare365
	NumFile365
	NumAd365
	NumN365
	HasUserPrincipalName
	HasMail
	HasPhone
	FlagDisabled
	FlagLockout
	FlagPasswordNotRequired
	FlagPasswordCantChange
	FlagDontExpirePassword
	OwnedFiles
	NumMailboxes
	NumMemberOfGroups
	NumMemberOfIndirectGroups
	MemberOfIndirectGroupsIds
	MemberOfGroupsIds
	IsAdmin
	IsService
	CountField
)
