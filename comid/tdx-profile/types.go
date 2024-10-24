package tdx

import "github.com/veraison/corim/comid"

type teeModel string
type teeVendor string

type pceID string

type teeSVN int

type maskType []byte

type tdate string

type teeDigest comid.Digests

type epochSeconds int

type setType any

type epochExpression struct {
	grace_period epochSeconds
	epochID      *tdate
}

type teeTcbStatus setType

type teeTcbEvalNum uint

type teeTcbCompSvn [16][16]teeSVN

type teeMiscSelect maskType

// TO DO Align later, this is with real TD profile
type teeIsvProdID comid.UUID

type teeInstanceID []byte

type teeCryptoKey comid.CryptoKey

type teeAdvisoryID setType
