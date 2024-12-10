// Copyright 2024 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package tdx

import (
	"fmt"

	"github.com/veraison/corim/comid"
	"github.com/veraison/corim/corim"
	"github.com/veraison/corim/extensions"
	"github.com/veraison/eat"
)

func Example_tdx_qe_refval() {
	coMID := &comid.Comid{}
	profileID, err := eat.NewProfile("http://intel.com/tdx-profile")
	if err != nil {
		panic(err) // will not error, as the hard-coded string above is valid
	}

	extMap := extensions.NewMap().
		Add(comid.ExtMval, &MvalExtensions{})

	if err := corim.RegisterProfile(profileID, extMap); err != nil {
		// will not error, assuming our profile ID is unique, and we've
		// correctly set up the extensions Map above
		panic(err)
	}
	if err := coMID.FromJSON([]byte(TDXQERefValTemplate)); err != nil {
		fmt.Printf("From JSON Failed %s", err.Error())
	} else {
		fmt.Printf("From JSON Passed \n")
	}
	mVal := coMID.Triples.ReferenceValues.Values[0].Measurements.Values[0].Val
	val, err := mVal.Extensions.Get("tcbevalnum")
	if err != nil {
		fmt.Printf(" \n tcbEvalNum NOT Set: %s \n", err.Error())
	} else {
		fmt.Printf(" \n tcbEvalNum is Set %d", val)
	}
	f, ok := val.(*teeTcbEvalNum)
	if !ok {
		fmt.Printf("val was not pointer to teeTcbEvalNum")
	}
	tcbValNum := *f
	if err != nil {
		fmt.Printf(" \n tcbEvalNum NOT Set: %s \n", err.Error())
	} else {
		fmt.Printf(" \n tcbEvalNum is Set %d", tcbValNum)
	}

	if err := coMID.Valid(); err != nil {
		panic(err)
	}

	// Output:
	//a301a1005043bbe37f2e614b33aed353cff1428b200281a30065494e54454c01d8207168747470733a2f2f696e74656c2e636f6d028301000204a1008182a100a300d86f4c6086480186f84d01020304050171496e74656c20436f72706f726174696f6e02675444585345414d81a101a100a20065312e322e330101

}
