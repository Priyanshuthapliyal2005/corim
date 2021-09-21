// Copyright 2021 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package corim

import (
	"errors"
	"fmt"
	"time"

	"github.com/veraison/corim/comid"
)

type Signer struct {
	Name string           `cbor:"0,keyasint" json:"name"`
	URI  *comid.TaggedURI `cbor:"1,keyasint,omitempty" json:"uri,omitempty"`
}

func NewSigner() *Signer {
	return &Signer{}
}

func (o *Signer) SetName(name string) *Signer {
	if o != nil {
		if name == "" {
			return nil
		}
		o.Name = name
	}
	return o
}

func (o *Signer) SetURI(uri string) *Signer {
	if o != nil {
		if uri == "" {
			return nil
		}

		taggedURI, err := comid.String2URI(&uri)
		if err != nil {
			return nil
		}

		o.URI = taggedURI
	}
	return o
}

func (o Signer) Valid() error {
	if o.Name == "" {
		return errors.New("empty name")
	}

	if o.URI != nil {
		if *o.URI == "" {
			return errors.New("empty URI")
		}
	}

	return nil
}

type Meta struct {
	Signer   Signer    `cbor:"0,keyasint" json:"signer"`
	Validity *Validity `cbor:"1,keyasint,omitempty" json:"validity,omitempty"`
}

func NewMeta() *Meta {
	return &Meta{}
}

func (o *Meta) SetSigner(name string, uri *string) *Meta {
	if o != nil {
		s := NewSigner().SetName(name)

		if uri != nil {
			s = s.SetURI(*uri)
		}

		if s == nil {
			return nil
		}

		o.Signer = *s
	}
	return o
}

func (o *Meta) SetValidity(notAfter time.Time, notBefore *time.Time) *Meta {
	if o != nil {
		v := NewValidity().Set(notAfter, notBefore)
		if v == nil {
			return nil
		}

		o.Validity = v
	}
	return o
}

func (o Meta) Valid() error {
	if err := o.Signer.Valid(); err != nil {
		return fmt.Errorf("invalid meta: %w", err)
	}

	if o.Validity != nil {
		if err := o.Validity.Valid(); err != nil {
			return fmt.Errorf("invalid meta: %w", err)
		}
	}

	return nil
}

func (o Meta) ToCBOR() ([]byte, error) {
	return em.Marshal(&o)
}

func (o *Meta) FromCBOR(data []byte) error {
	return dm.Unmarshal(data, o)
}
