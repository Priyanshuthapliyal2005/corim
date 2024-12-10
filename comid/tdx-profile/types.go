package tdx

import "github.com/veraison/corim/comid"

type numericType uint

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

type teeTcbCompSvn [16]teeSVN

type teeMiscSelect maskType

type teeAttributes maskType

type teeIsvProdID []byte

type teeInstanceID uint

type teeCryptoKey comid.CryptoKey

type teeAdvisoryID setType

type epochTimeStamp tdate
