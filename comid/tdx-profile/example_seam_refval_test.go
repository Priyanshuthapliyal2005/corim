// Copyright 2024 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package tdx

import (
	"fmt"

	"github.com/veraison/corim/comid"
	"github.com/veraison/corim/extensions"
	"github.com/veraison/swid"
)

func Example_tdx_seam_refval() {
	comid := comid.Comid{}

	if err := comid.FromJSON([]byte(TDXSeamRefValJSONTemplate)); err != nil {
		panic(err)
	}

	if err := comid.Valid(); err != nil {
		panic(err)
	}

	// Decode individual Elements
}

func Example_encode_tdx_seam_refval() {

	var refVal comid.ValueTriple
	var measurement *comid.Measurement
	measurement = &comid.Measurement{}
	refVal.Environment = comid.Environment{
		Class: comid.NewClassOID(TestOID).
			SetVendor("Intel Corporation").
			SetModel("TDXSEAM"),
	}

	extMap := extensions.NewMap().
		Add(comid.ExtReferenceValue, &MvalExtensions{})

	coMID := comid.NewComid().
		SetTagIdentity("43BBE37F-2E61-4B33-AED3-53CFF1428B20", 0).
		AddEntity("INTEL", &TestRegID, comid.RoleCreator, comid.RoleTagCreator, comid.RoleMaintainer)

	coMID.Triples.ReferenceValues.RegisterExtensions(extMap)

	// Bug: Needs Mandatory setting of a minimum of one value, apart from Extensions
	measurement.Val.Ver = comid.NewVersion()
	measurement.Val.Ver.SetVersion("1.2.3")
	measurement.Val.Ver.SetScheme(1)
	err := measurement.Val.Ver.Valid()
	if err != nil {
		fmt.Printf("\n Measurement Validation Failed: %s \n", err.Error())
	}
	refVal.Measurements.Add(measurement)
	// Set the Extensions now
	measurement.Val.Extensions.Set("tcbdate", "123")
	measurement.Val.Extensions.Set("isvprodid", 1)
	measurement.Val.Extensions.Set("isvsvn", 10)
	measurement.Val.Extensions.Set("tcbEvalNum", 11)
	measurement.Val.Extensions.Set("attributes", []byte{0x01, 0x01})

	d := comid.NewDigests()
	d.AddDigest(swid.Sha256, comid.MustHexDecode(nil, "e45b72f5c0c0b572db4d8d3ab7e97f368ff74e62347a824decb67a84e5224d75"))
	d.AddDigest(swid.Sha256, comid.MustHexDecode(nil, "e45b72f5c0c0b572db4d8d3ab7e97f368ff74e62347a824decb67a84e5224d75e45b72f5c0c0b572db4d8d3ab7e97f36"))

	measurement.Val.Extensions.Set("mrsigner", d)

	coMID.Triples.AddReferenceValue(refVal)
	err = coMID.Valid()
	if err != nil {
		fmt.Printf("coMID is not Valid :%s", err.Error())
	}

	cbor, err := coMID.ToCBOR()
	if err == nil {
		fmt.Printf("%x\n", cbor)
	} else {
		fmt.Printf("\n To CBOR Failed: %s \n", err.Error())
	}

	json, err := coMID.ToJSON()
	if err == nil {
		fmt.Printf("%s\n", string(json))
	} else {
		fmt.Printf("\n To JSON Failed \n")
	}

	// Output:
	//{"tag-identity":{"id":"my-ns:acme-roadrunner-supplement"},"entities":[{"name":"ACME Ltd.","regid":"https://acme.example","roles":["creator","tagCreator","maintainer"]}],"triples":{"reference-values":[{"environment":{"class":{"id":{"type":"psa.impl-id","value":"YWNtZS1pbXBsZW1lbnRhdGlvbi1pZC0wMDAwMDAwMDE="},"vendor":"ACME Ltd.","model":"RoadRunner 2.0"}},"measurements":[{"key":{"type":"psa.refval-id","value":{"label":"BL","version":"5.0.5","signer-id":"rLsRx+TaIXIFUjzkzhokWuGiOa48a/2eeHH35di66Gs="}},"value":{"digests":["sha-256-32;q83vAA=="]}},{"key":{"type":"psa.refval-id","value":{"label":"PRoT","version":"1.3.5","signer-id":"rLsRx+TaIXIFUjzkzhokWuGiOa48a/2eeHH35di66Gs="}},"value":{"digests":["sha-256-32;q83vAA=="]}}]}],"attester-verification-keys":[{"environment":{"instance":{"type":"ueid","value":"At6tvu/erQ=="}},"verification-keys":[{"type":"pkix-base64-key","value":"-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEW1BvqF+/ry8BWa7ZEMU1xYYHEQ8B\nlLT4MFHOaO+ICTtIvrEeEpr/sfTAP66H2hCHdb5HEXKtRKod6QLcOLPA1Q==\n-----END PUBLIC KEY-----"}]}]}}

}
