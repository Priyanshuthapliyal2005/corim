// Copyright 2024 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package tdx

import (
	_ "embed"
	"fmt"
	"log"

	"github.com/veraison/corim/comid"
	"github.com/veraison/corim/corim"
	"github.com/veraison/corim/extensions"
	"github.com/veraison/eat"
	"github.com/veraison/swid"
)

// The below one works
func Example_decode_JSON() {
	profileID, err := eat.NewProfile("http://intel.com/tdx-profile")
	if err != nil {
		panic(err) // will not error, as the hard-coded string above is valid
	}
	profile, found := corim.GetProfile(profileID)
	if !found {
		fmt.Printf("CoRIM Profile NOT FOUND")
		return
	}

	coMID := profile.GetComid()

	if err := coMID.FromJSON([]byte(TDXSeamRefValJSONTemplate)); err != nil {
		panic(err)
	}

	if err := coMID.Valid(); err != nil {
		fmt.Errorf("CoMID is invalid %s", err.Error())

	}
	if coMID.Triples.ReferenceValues == nil {
		fmt.Printf("\n No Reference Value Set \n ")
	}
	if len(coMID.Triples.ReferenceValues.Values[0].Measurements.Values) == 0 {
		fmt.Printf("\n No Measurement Entries Set\n ")
	}
	for _, m := range coMID.Triples.ReferenceValues.Values[0].Measurements.Values {
		decodeMValExtensions(m)
		val, err := m.Val.Extensions.Get("tcbevalnum")
		f, ok := val.(*teeTcbEvalNum)
		if !ok {
			fmt.Printf("val was not pointer to teeTcbEvalNum")
		}
		tcbValNum := *f
		if err != nil {
			fmt.Printf(" \n tcbEvalNum NOT Set: %s \n", err.Error())
		} else {
			fmt.Printf(" \n TcbEvalNum: %d", tcbValNum)
		}
	}

}

func Example_decode_JSON1() {
	profileID, err := eat.NewProfile("http://intel.com/tdx-profile")
	if err != nil {
		panic(err) // will not error, as the hard-coded string above is valid
	}
	profile, found := corim.GetProfile(profileID)
	if !found {
		fmt.Printf("CoRIM Profile NOT FOUND")
		return
	}

	coMID := profile.GetComid()
	if err := coMID.FromJSON([]byte(TDXSeamRefValJSONTemplate)); err != nil {
		panic(err)
	}

	if err := coMID.Valid(); err != nil {
		fmt.Errorf("CoMID is invalid %s", err.Error())

	}
	if coMID.Triples.ReferenceValues == nil {
		fmt.Printf("\n No Reference Value Set \n ")
	}
	if len(coMID.Triples.ReferenceValues.Values[0].Measurements.Values) == 0 {
		fmt.Printf("\n No Measurement Entries Set\n ")
	}
	for _, m := range coMID.Triples.ReferenceValues.Values[0].Measurements.Values {
		decodeMValExtensions(m)
		val, err := m.Val.Extensions.Get("tcbevalnum")
		f, ok := val.(*teeTcbEvalNum)
		if !ok {
			fmt.Printf("val was not pointer to teeTcbEvalNum")
		}
		tcbValNum := *f
		if err != nil {
			fmt.Printf(" \n tcbEvalNum NOT Set: %s \n", err.Error())
		} else {
			fmt.Printf(" \n TcbEvalNum: %d", tcbValNum)
		}
	}
	// output:
	// ImplementationID: 61636d652d696d706c656d656e746174696f6e2d69642d303030303030303031
	// SignerID: acbb11c7e4da217205523ce4ce1a245ae1a239ae3c6bfd9e7871f7e5d8bae86b
	// Label: BL
	// Version: 2.1.0
	// Digest: 87428fc522803d31065e7bce3cf03fe475096631e5e07bbd7a0fde60c4cf25c7
	// SignerID: acbb11c7e4da217205523ce4ce1a245ae1a239ae3c6bfd9e7871f7e5d8bae86b
	// Label: PRoT
	// Version: 1.3.5
	// Digest: 0263829989b6fd954f72baaf2fc64bc2e2f01d692d4de72986ea808f6e99813f
	// SignerID: acbb11c7e4da217205523ce4ce1a245ae1a239ae3c6bfd9e7871f7e5d8bae86b
	// Label: ARoT
	// Version: 0.1.4
	// Digest: a3a5e715f0cc574a73c3f9bebb6bc24f32ffd5b67b387244c2c909da779a1478

}

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

	refVal.Measurements.Add(measurement)
	coMID.Triples.AddReferenceValue(*refVal)
	coMID.RegisterExtensions(extMap)
	// fmt.Printf("len of Measurements = %d ", len(coMID.Triples.ReferenceValues.Values[0].Measurements.Values))
	// Set the Extensions now
	// setMValExtensions(&measurement.Val) ==> this does not work, though
	setMValExtensions(&coMID.Triples.ReferenceValues.Values[0].Measurements.Values[0].Val)
	err := coMID.Valid()
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
	// a301a1005043bbe37f2e614b33aed353cff1428b200281a30065494e54454c01d8207168747470733a2f2f696e74656c2e636f6d028301000204a1008182a100a300d86f4c6086480186f84d01020304050171496e74656c20436f72706f726174696f6e02675444585345414d81a101a638476331323338480a385142010138538282015820e45b72f5c0c0b572db4d8d3ab7e97f368ff74e62347a824decb67a84e5224d7582075830e45b72f5c0c0b572db4d8d3ab7e97f368ff74e62347a824decb67a84e5224d75e45b72f5c0c0b572db4d8d3ab7e97f36385442010138550b
	// {"tag-identity":{"id":"43bbe37f-2e61-4b33-aed3-53cff1428b20"},"entities":[{"name":"INTEL","regid":"https://intel.com","roles":["creator","tagCreator","maintainer"]}],"triples":{"reference-values":[{"environment":{"class":{"id":{"type":"oid","value":"2.16.840.1.113741.1.2.3.4.5"},"vendor":"Intel Corporation","model":"TDXSEAM"}},"measurements":[{"value":{"tcbdate":"123","isvsvn":10,"attributes":"AQE=","mrsigner":["sha-256;5Fty9cDAtXLbTY06t+l/No/3TmI0eoJN7LZ6hOUiTXU=","sha-384;5Fty9cDAtXLbTY06t+l/No/3TmI0eoJN7LZ6hOUiTXXkW3L1wMC1cttNjTq36X82"],"isvprodid":"AQE=","tcbevalnum":11}}]}]}}
}

func Example_encode_tdx_seam_refval_with_profile() {
	profileID, err := eat.NewProfile("http://intel.com/tdx-profile")
	if err != nil {
		panic(err) // will not error, as the hard-coded string above is valid
	}
	profile, found := corim.GetProfile(profileID)
	if !found {
		fmt.Printf("CoRIM Profile NOT FOUND")
		return
	}

	coMID := profile.GetComid()
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

	refVal.Measurements.Add(measurement)
	coMID.Triples.AddReferenceValue(*refVal)

	setMValExtensions(&coMID.Triples.ReferenceValues.Values[0].Measurements.Values[0].Val)
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
	//726174696f6e02675444585345414d81a101a638476331323338480a385142010138538282015820e45b72f5c0c0b572db4d8d3ab7e97f368ff74e62347a824decb67a84e5224d7582075830e45b72f5c0c0b572db4d8d3ab7e97f368ff74e62347a824decb67a84e5224d75e45b72f5c0c0b572db4d8d3ab7e97f36385442010138550b
	// {"tag-identity":{"id":"43bbe37f-2e61-4b33-aed3-53cff1428b20"},"entities":[{"name":"INTEL","regid":"https://intel.com","roles":["creator","tagCreator","maintainer"]}],"triples":{"reference-values":[{"environment":{"class":{"id":{"type":"oid","value":"2.16.840.1.113741.1.2.3.4.5"},"vendor":"Intel Corporation","model":"TDXSEAM"}},"measurements":[{"value":{"tcbdate":"123","isvsvn":10,"attributes":"AQE=","mrsigner":["sha-256;5Fty9cDAtXLbTY06t+l/No/3TmI0eoJN7LZ6hOUiTXU=","sha-384;5Fty9cDAtXLbTY06t+l/No/3TmI0eoJN7LZ6hOUiTXXkW3L1wMC1cttNjTq36X82"],"isvprodid":"AQE=","tcbevalnum":11}}]}]}}
}

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
		log.Fatal("could not register mval extensions")
	}

	setMValExtensions(&measurement.Val)

	refVal.Measurements.Add(measurement)
	coMID.Triples.AddReferenceValue(*refVal)

	err := coMID.Valid()
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
	//a301a1005043bbe37f2e614b33aed353cff1428b200281a30065494e54454c01d8207168747470733a2f2f696e74656c2e636f6d028301000204a1008182a100a300d86f4c6086480186f84d01020304050171496e74656c20436f72706f726174696f6e02675444585345414d81a101a638476331323338480a385142010138538282015820e45b72f5c0c0b572db4d8d3ab7e97f368ff74e62347a824decb67a84e5224d7582075830e45b72f5c0c0b572db4d8d3ab7e97f368ff74e62347a824decb67a84e5224d75e45b72f5c0c0b572db4d8d3ab7e97f36385442010138550b
	// {"tag-identity":{"id":"43bbe37f-2e61-4b33-aed3-53cff1428b20"},"entities":[{"name":"INTEL","regid":"https://intel.com","roles":["creator","tagCreator","maintainer"]}],"triples":{"reference-values":[{"environment":{"class":{"id":{"type":"oid","value":"2.16.840.1.113741.1.2.3.4.5"},"vendor":"Intel Corporation","model":"TDXSEAM"}},"measurements":[{"value":{"tcbdate":"123","isvsvn":10,"attributes":"AQE=","mrsigner":["sha-256;5Fty9cDAtXLbTY06t+l/No/3TmI0eoJN7LZ6hOUiTXU=","sha-384;5Fty9cDAtXLbTY06t+l/No/3TmI0eoJN7LZ6hOUiTXXkW3L1wMC1cttNjTq36X82"],"isvprodid":"AQE=","tcbevalnum":11}}]}]}}
}

func setMValExtensions(val *comid.Mval) {
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

func decodeMValExtensions(m comid.Measurement) error {
	val, err := m.Val.Extensions.Get("tcbevalnum")
	if err != nil {
		return fmt.Errorf("failed to decode tcbevalnum from measurement extensions")
	}
	f, ok := val.(*teeTcbEvalNum)
	if !ok {
		fmt.Printf("val was not pointer to teeTcbEvalNum")
	}
	tcbValNum := *f
	fmt.Printf(" \n tcbEvalNum: %d", tcbValNum)

	val, err = m.Val.Extensions.Get("isvprodid")
	if err != nil {
		return fmt.Errorf("failed to decode isvprodid from measurement extensions")
	}
	tS, ok := val.(*teeIsvProdID)
	if !ok {
		fmt.Printf("val was not pointer to teeIsvProdID")
	}

	fmt.Printf(" \n IsvProdID: %d", *tS)

	val, err = m.Val.Extensions.Get("isvsvn")
	if err != nil {
		return fmt.Errorf("failed to decode isvsvn from measurement extensions")
	}
	tSV, ok := val.(*teeSVN)
	if !ok {
		fmt.Printf("val was not pointer to tee svn")
	}

	fmt.Printf(" \n ISVSVN: %d", *tSV)

	return nil
}

var (
	// test cases are based on diag files here:
	// https://github.com/ietf-rats-wg/draft-ietf-rats-corim/tree/main/cddl/examples

	//go:embed testcases/comid_seam_refval.cbor
	testComid1 []byte
)

func Example_decode_CBOR() {
	profileID, err := eat.NewProfile("http://intel.com/tdx-profile")
	if err != nil {
		panic(err) // will not error, as the hard-coded string above is valid
	}
	profile, found := corim.GetProfile(profileID)
	if !found {
		fmt.Printf("CoRIM Profile NOT FOUND")
		return
	}

	coMID := profile.GetComid()

	if err := coMID.FromCBOR(testComid1); err != nil {
		panic(err)
	}
	if err := coMID.Valid(); err != nil {
		fmt.Errorf("CoMID is invalid %s", err.Error())

	}
	if coMID.Triples.ReferenceValues == nil {
		fmt.Printf("\n No Reference Value Set \n ")
	}
	if len(coMID.Triples.ReferenceValues.Values[0].Measurements.Values) == 0 {
		fmt.Printf("\n No Measurement Entries Set\n ")
	}
	for _, m := range coMID.Triples.ReferenceValues.Values[0].Measurements.Values {
		decodeMValExtensions(m)
		val, err := m.Val.Extensions.Get("tcbevalnum")
		f, ok := val.(*teeTcbEvalNum)
		if !ok {
			fmt.Printf("val was not pointer to teeTcbEvalNum")
		}
		tcbValNum := *f
		if err != nil {
			fmt.Printf(" \n tcbEvalNum NOT Set: %s \n", err.Error())
		} else {
			fmt.Printf(" \n TcbEvalNum: %d", tcbValNum)
		}
	}
	// Output: OK
}
