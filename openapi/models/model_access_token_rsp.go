/*
 * NRF OAuth2
 *
 * NRF OAuth2 Authorization
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type AccessTokenRsp struct {
	// JWS Compact Serialized representation of JWS signed JSON object (AccessTokenClaims)
	AccessToken string `json:"access_token" yaml:"access_token" bson:"access_token" mapstructure:"AccessToken"`
	TokenType   string `json:"token_type" yaml:"token_type" bson:"token_type" mapstructure:"TokenType"`
	ExpiresIn   int32  `json:"expires_in,omitempty" yaml:"expires_in" bson:"expires_in" mapstructure:"ExpiresIn"`
	Scope       string `json:"scope,omitempty" yaml:"scope" bson:"scope" mapstructure:"Scope"`
}
