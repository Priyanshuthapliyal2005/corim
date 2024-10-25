package tdx

import "github.com/veraison/corim/comid"

// TO DO, check change this later to more expandable type
type numericType uint

type teeModel string
type teeVendor string

type pceID string

type teeSVN numericType

type maskType []byte

type tdate string

type teeDigest comid.Digests

type epochSeconds int

type setType any

type epochExpression struct {
	gracePeriod epochSeconds
	epochID     *tdate
}

type teeTcbStatus setType

type teeTcbEvalNum uint

type teeTcbCompSvn [16][16]teeSVN

type teeMiscSelect maskType

type teeAtttributes maskType

// TO DO Check with Ned, why it is NOT UUID but either an Integer or Bstr in the Profile Document
type teeIsvProdID comid.UUID

// TO DO Change this Instance ID to be a type choice with expression for a []byte
type teeInstanceID uint

type teeCryptoKey comid.CryptoKey

type teeAdvisoryID setType

type epochTimeStamp tdate

// TO DO Set of Set Type: Where it is used and is it needed, for this profile ...?

// TO DO Check with Ned, What is time? in the CDDL Document, not defined???
