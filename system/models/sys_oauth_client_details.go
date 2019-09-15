package models

type SysOauthClientDetails struct {
	ClientId              string `xorm:"not null pk VARCHAR(64)"`
	ResourceIds           string `xorm:"default 'NULL' VARCHAR(256)"`
	ClientSecret          string `xorm:"default 'NULL' VARCHAR(256)"`
	Scope                 string `xorm:"default 'NULL' VARCHAR(256)"`
	AuthorizedGrantTypes  string `xorm:"default 'NULL' VARCHAR(256)"`
	WebServerRedirectUri  string `xorm:"default 'NULL' VARCHAR(256)"`
	Authorities           string `xorm:"default 'NULL' VARCHAR(256)"`
	AccessTokenValidity   int    `xorm:"default NULL INT(11)"`
	RefreshTokenValidity  int    `xorm:"default NULL INT(11)"`
	AdditionalInformation string `xorm:"default 'NULL' VARCHAR(4096)"`
	Autoapprove           string `xorm:"default 'NULL' VARCHAR(256)"`
}
