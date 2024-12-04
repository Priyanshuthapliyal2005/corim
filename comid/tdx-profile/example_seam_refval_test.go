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
		fmt.Errorf("CoMID is invalid %s", err.Error())
	}

	// Decode individual Elements
}

func Example_encode_tdx_seam_refval() {

	refVal := &comid.ValueTriple{}
	measurement := &comid.Measurement{}
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
	refVal.Measurements.Add(measurement)
	coMID.Triples.AddReferenceValue(*refVal)
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
	//a301a1005043bbe37f2e614b33aed353cff1428b200281a30065494e54454c01d8207168747470733a2f2f696e74656c2e636f6d028301000204a1008182a100a300d86f4c6086480186f84d01020304050171496e74656c20436f72706f726174696f6e02675444585345414d81a101a100a20065312e322e330101
}
