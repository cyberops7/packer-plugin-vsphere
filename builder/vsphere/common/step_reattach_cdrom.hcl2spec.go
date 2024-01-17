// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package common

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatReattachCDRomConfig is an auto-generated flat version of ReattachCDRomConfig.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatReattachCDRomConfig struct {
	ReattachCDRom *int `mapstructure:"reattach_cdroms" cty:"reattach_cdroms" hcl:"reattach_cdroms"`
}

// FlatMapstructure returns a new FlatReattachCDRomConfig.
// FlatReattachCDRomConfig is an auto-generated flat version of ReattachCDRomConfig.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*ReattachCDRomConfig) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatReattachCDRomConfig)
}

// HCL2Spec returns the hcl spec of a ReattachCDRomConfig.
// This spec is used by HCL to read the fields of ReattachCDRomConfig.
// The decoded values from this spec will then be applied to a FlatReattachCDRomConfig.
func (*FlatReattachCDRomConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"reattach_cdroms": &hcldec.AttrSpec{Name: "reattach_cdroms", Type: cty.Number, Required: false},
	}
	return s
}