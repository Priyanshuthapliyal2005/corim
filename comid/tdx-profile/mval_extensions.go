package tdx

import (
	"github.com/veraison/corim/comid"
	"github.com/veraison/corim/corim"
	"github.com/veraison/corim/extensions"
	"github.com/veraison/eat"
)

//"github.com/veraison/eat"

// MvalExtensions contains the Intel TDX profile extensions which can appear in
// both Reference Values and Endorsed Values
type MvalExtensions struct {
	// a string field extension
	TcbDate     *tdate         `cbor:"-72,keyasint,omitempty" json:"tcbdate,omitempty"`
	IsvSVN      *teeSVN        `cbor:"-73,keyasint,omitempty" json:"isvsvn,omitempty"`
	PCEID       *pceID         `cbor:"-80,keyasint,omitempty" json:"pceid,omitempty"`
	MiscSelect  *teeMiscSelect `cbor:"-81,keyasint,omitempty" json:"miscselect,omitempty"`
	Attributes  *teeAttributes `cbor:"-82,keyasint,omitempty" json:"attributes,omitempty"`
	MrSigner    *teeDigest     `cbor:"-84,keyasint,omitempty" json:"mrsigner,omitempty"`
	IsvProdID   *teeIsvProdID  `cbor:"-85,keyasint,omitempty" json:"isvprodid,omitempty"`
	TcbEvalNum  *teeTcbEvalNum `cbor:"-86,keyasint,omitempty" json:"tcbevalnum,omitempty"`
	TcbStatus   *teeTcbStatus  `cbor:"-88,keyasint,omitempty" json:"tcbstatus,omitempty"`
	AdvisoryIDs *teeAdvisoryID `cbor:"-89,keyasint,omitempty" json:"advisoryids,omitempty"`
	Epoch       *epochSeconds  `cbor:"-90, keyasint,omitempty" json:"epoch,omitempty"`

	TeeCryptoKeys *[]teeCryptoKey `cbor:"-91, keyasint,omitempty" json:"teecryptokeys,omitempty"`
	TeeTCBCompSvn *teeTcbCompSvn  `cbor:"-125, keyasint,omitempty" json:"teetcbcompsvn,omitempty"`
}

// Registering the profile inside init() in the same file where it is defined
// ensures that the profile will always be available, and you don't need to
// remember to register it at the time you want to use it. The only potential
// danger with that is if the your profile ID clashes with another profile,
// which should not happen if it a registered PEN or a URL containing a domain
// that you own.

func init() {
	profileID, err := eat.NewProfile("http://intel.com/tdx-profile")
	if err != nil {
		panic(err) // will not error, as the hard-coded string above is valid
	}

	extMap := extensions.NewMap().
		Add(comid.ExtReferenceValue, &MvalExtensions{}).
		Add(comid.ExtEndorsedValue, &MvalExtensions{})

	if err := corim.RegisterProfile(profileID, extMap); err != nil {
		// will not error, assuming our profile ID is unique, and we've
		// correctly set up the extensions Map above
		panic(err)
	}
}
