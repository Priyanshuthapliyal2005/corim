// Copyright 2024 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package tdx

import (
	"fmt"
	"log"

	"github.com/veraison/corim/comid"
	"github.com/veraison/corim/corim"
	"github.com/veraison/corim/extensions"
	"github.com/veraison/eat"
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

// In Example: Example_encode_tdx_seam_refval_without_profile() the Extensions are NOT Encoded in CBOR AND JSON Correctly!
// This example is WITHOUT PROFILE
func Example_encode_tdx_seam_refval_without_profile() {
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
	setMValExtensions(measurement.Val)

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
	// {"tag-identity":{"id":"43bbe37f-2e61-4b33-aed3-53cff1428b20"},"entities":[{"name":"INTEL","regid":"https://intel.com","roles":["creator","tagCreator","maintainer"]}],"triples":{"reference-values":[{"environment":{"class":{"id":{"type":"oid","value":"2.16.840.1.113741.1.2.3.4.5"},"vendor":"Intel Corporation","model":"TDXSEAM"}},"measurements":[{"value":{"version":{"value":"1.2.3","scheme":"multipartnumeric"}}}]}]}}
}

// Same Effect of Failure Unable to Register Extensions with Profile as well!!!
// In Example: Example_encode_tdx_seam_refval_with_profile() the Extensions are NOT Encoded in CBOR AND JSON Correctly!
// This example is ONE WITH PROFILE
func Example_encode_tdx_seam_refval_with_profile() {

	profID, err := eat.NewProfile("http://intel.com/tdx-profile")
	if err != nil {
		fmt.Printf("Unable to get new Profile")
	}

	extMap := extensions.NewMap().
		Add(comid.ExtReferenceValue, &MvalExtensions{})
	err = corim.RegisterProfile(profID, extMap)

	myprofile, found := corim.GetProfile(profID)
	if !found {
		fmt.Printf("Profile NOT Found")
		return
	}
	coMID := myprofile.GetComid()
	if coMID == nil {
		fmt.Printf("\n CoMID is NIL\n")
	}
	coMID.SetTagIdentity("43BBE37F-2E61-4B33-AED3-53CFF1428B20", 0).
		AddEntity("INTEL", &TestRegID, comid.RoleCreator, comid.RoleTagCreator, comid.RoleMaintainer)

	refVal := &comid.ValueTriple{}
	measurement := &comid.Measurement{}
	refVal.Environment = comid.Environment{
		Class: comid.NewClassOID(TestOID).
			SetVendor("Intel Corporation").
			SetModel("TDXSEAM"),
	}

	// Bug: Needs Mandatory setting of a minimum of one value, apart from Extensions
	measurement.Val.Ver = comid.NewVersion()
	measurement.Val.Ver.SetVersion("1.2.3")
	measurement.Val.Ver.SetScheme(1)
	err = measurement.Val.Ver.Valid()
	if err != nil {
		fmt.Printf("\n Measurement Validation Failed: %s \n", err.Error())
	}
	setMValExtensions(measurement.Val)
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
	// {"tag-identity":{"id":"43bbe37f-2e61-4b33-aed3-53cff1428b20"},"entities":[{"name":"INTEL","regid":"https://intel.com","roles":["creator","tagCreator","maintainer"]}],"triples":{"reference-values":[{"environment":{"class":{"id":{"type":"oid","value":"2.16.840.1.113741.1.2.3.4.5"},"vendor":"Intel Corporation","model":"TDXSEAM"}},"measurements":[{"value":{"version":{"value":"1.2.3","scheme":"multipartnumeric"}}}]}]}}
}

// In Example: Example_encode_tdx_seam_refval_without_profile() the Extensions are NOT Encoded in CBOR AND JSON Correctly!
// This example is WITHOUT PROFILE
func Example_encode_tdx_seam_refval_direct() {
	refVal := &comid.ValueTriple{}
	measurement := &comid.Measurement{}
	refVal.Environment = comid.Environment{
		Class: comid.NewClassOID(TestOID).
			SetVendor("Intel Corporation").
			SetModel("TDXSEAM"),
	}

	extMap := extensions.NewMap().Add(comid.ExtMval, &MvalExtensions{})
	coMID := comid.NewComid().
		SetTagIdentity("43BBE37F-2E61-4B33-AED3-53CFF1428B20", 0).
		AddEntity("INTEL", &TestRegID, comid.RoleCreator, comid.RoleTagCreator, comid.RoleMaintainer)

	if err := measurement.Val.RegisterExtensions(extMap); err != nil {
		log.Fatal("could not register refval extensions")
	}

	// Bug: Needs Mandatory setting of a minimum of one value, apart from Extensions
	measurement.Val.Ver = comid.NewVersion()
	measurement.Val.Ver.SetVersion("1.2.3")
	measurement.Val.Ver.SetScheme(1)
	err := measurement.Val.Ver.Valid()
	if err != nil {
		fmt.Printf("\n Measurement Validation Failed: %s \n", err.Error())
	}
	setMValExtensions(measurement.Val)

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
	// {"tag-identity":{"id":"43bbe37f-2e61-4b33-aed3-53cff1428b20"},"entities":[{"name":"INTEL","regid":"https://intel.com","roles":["creator","tagCreator","maintainer"]}],"triples":{"reference-values":[{"environment":{"class":{"id":{"type":"oid","value":"2.16.840.1.113741.1.2.3.4.5"},"vendor":"Intel Corporation","model":"TDXSEAM"}},"measurements":[{"value":{"version":{"value":"1.2.3","scheme":"multipartnumeric"}}}]}]}}
}

func setMValExtensions(val comid.Mval) {
	tcbDate := tdate("123")
	isvProdID := teeIsvProdID([]byte{0x01, 0x01})
	svn := teeSVN(10)
	teeTcbEvalNum := teeTcbEvalNum(11)
	teeAttr := teeAttributes([]byte{0x01, 0x01})
	val.Extensions.Set("tcbdate", &tcbDate)
	val.Extensions.Extensions.Set("isvprodid", &isvProdID)
	val.Extensions.Extensions.Set("isvsvn", &svn)
	val.Extensions.Extensions.Set("tcbevalnum", &teeTcbEvalNum)
	val.Extensions.Extensions.Set("attributes", &teeAttr)

	d := comid.NewDigests()
	d.AddDigest(swid.Sha256, comid.MustHexDecode(nil, "e45b72f5c0c0b572db4d8d3ab7e97f368ff74e62347a824decb67a84e5224d75"))
	d.AddDigest(swid.Sha384, comid.MustHexDecode(nil, "e45b72f5c0c0b572db4d8d3ab7e97f368ff74e62347a824decb67a84e5224d75e45b72f5c0c0b572db4d8d3ab7e97f36"))

	val.Extensions.Set("mrsigner", d)
}
