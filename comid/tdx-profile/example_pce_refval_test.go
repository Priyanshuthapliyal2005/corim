package tdx

import "github.com/veraison/corim/comid"

func Example_tdx_pce_refval() {
	comid := comid.Comid{}

	if err := comid.FromJSON([]byte(TDXPCERefValTemplate)); err != nil {
		panic(err)
	}

	if err := comid.Valid(); err != nil {
		panic(err)
	}

}
