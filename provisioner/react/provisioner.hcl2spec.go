// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package react

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	DistSource       *string `mapstructure:"distSource" required:"true" cty:"distSource" hcl:"distSource"`
	SslCertBase64    *string `mapstructure:"sslCertBase64" required:"true" cty:"sslCertBase64" hcl:"sslCertBase64"`
	SslCertKeyBase64 *string `mapstructure:"sslCertKeyBase64" required:"true" cty:"sslCertKeyBase64" hcl:"sslCertKeyBase64"`
	AppDomain        *string `mapstructure:"appDomain" required:"true" cty:"appDomain" hcl:"appDomain"`
	HomeDir          *string `mapstructure:"homeDir" required:"false" cty:"homeDir" hcl:"homeDir"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"distSource":       &hcldec.AttrSpec{Name: "distSource", Type: cty.String, Required: false},
		"sslCertBase64":    &hcldec.AttrSpec{Name: "sslCertBase64", Type: cty.String, Required: false},
		"sslCertKeyBase64": &hcldec.AttrSpec{Name: "sslCertKeyBase64", Type: cty.String, Required: false},
		"appDomain":        &hcldec.AttrSpec{Name: "appDomain", Type: cty.String, Required: false},
		"homeDir":          &hcldec.AttrSpec{Name: "homeDir", Type: cty.String, Required: false},
	}
	return s
}
