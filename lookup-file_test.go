package envir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepareName(t *testing.T) {
	is := assert.New(t)

	type testCase struct {
		in   string
		inst *lookupFile
		exp  string
	}

	cases := []testCase{
		{
			in: "PG_HOST",
			inst: &lookupFile{
				cutPrefix: "",
			},
			exp: "PG_HOST",
		},
		{
			in: "PG_HOST",
			inst: &lookupFile{
				cutPrefix: "PG_",
			},
			exp: "HOST",
		},
		{
			in: "PG",
			inst: &lookupFile{
				cutPrefix: "PG_",
			},
			exp: "PG",
		},
		{
			in: "PG_USER",
			inst: &lookupFile{
				cutPrefix: "APPNAME_",
			},
			exp: "PG_USER",
		},
		{
			in: "APPNAME_PG_USER",
			inst: &lookupFile{
				cutPrefix: "APPNAME_",
			},
			exp: "PG_USER",
		},
		{
			in: "APPNAMEPG_USER",
			inst: &lookupFile{
				cutPrefix: "APPNAME_",
			},
			exp: "APPNAMEPG_USER",
		},
	}

	for _, tc := range cases {
		res := tc.inst.prepareName(tc.in)
		is.Equal(tc.exp, res)
	}
}
