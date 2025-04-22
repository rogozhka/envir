package envir

import (
	"strconv"
	"testing"
	"time"

	"github.com/rogozhka/envir/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestOinment_Bool(t *testing.T) {
	is := assert.New(t)

	type args struct {
		unprefixed   string
		prefix       string
		defaultValue bool
	}

	type resStruct struct {
		res bool
	}

	type testCase struct {
		name  string
		inst  *oinment
		ctrlr *gomock.Controller
		args  args
		exp   resStruct
	}

	cases := []testCase{
		func() testCase {
			const name = "no prefix exists true"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const envName = "TEST_ENV_NAME"

			mockOs.EXPECT().LookupEnv(envName).Return("true", true).AnyTimes()

			in := args{
				unprefixed:   envName,
				prefix:       "",
				defaultValue: false,
			}
			exp := resStruct{
				res: true,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
		func() testCase {
			const name = "no prefix exists false"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const envName = "TEST_ENV_NAME"

			mockOs.EXPECT().LookupEnv(envName).Return("false", true).AnyTimes()

			in := args{
				unprefixed:   envName,
				prefix:       "",
				defaultValue: true,
			}
			exp := resStruct{
				res: false,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
		func() testCase {
			const name = "no prefix does not exist true default"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const envName = "TEST_ENV_NAME"

			mockOs.EXPECT().LookupEnv(envName).Return("", false)

			in := args{
				unprefixed:   envName,
				prefix:       "",
				defaultValue: true,
			}
			exp := resStruct{
				res: true,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
		func() testCase {
			const name = "no prefix does not exist false default"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const envName = "TEST_ENV_NAME"

			mockOs.EXPECT().LookupEnv(envName).Return("", false)

			in := args{
				unprefixed:   envName,
				prefix:       "",
				defaultValue: true,
			}
			exp := resStruct{
				res: true,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
		func() testCase {
			const name = "prefix exists true"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const envName = "TEST_ENV_NAME"
			const prefix = "PREFIX"

			mockOs.EXPECT().LookupEnv(prefix+"_"+envName).Return("true", true).AnyTimes()

			in := args{
				unprefixed:   envName,
				prefix:       prefix,
				defaultValue: false,
			}
			exp := resStruct{
				res: true,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
		func() testCase {
			const name = "prefix exists false"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const envName = "TEST_ENV_NAME"
			const prefix = "PREFIX"

			mockOs.EXPECT().LookupEnv(prefix+"_"+envName).Return("false", true).AnyTimes()

			in := args{
				unprefixed:   envName,
				prefix:       prefix,
				defaultValue: true,
			}
			exp := resStruct{
				res: false,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
		func() testCase {
			const name = "prefix does not exist true default"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const envName = "TEST_ENV_NAME"
			const prefix = "PREFIX"

			mockOs.EXPECT().LookupEnv(prefix+"_"+envName).Return("", false)

			in := args{
				unprefixed:   envName,
				prefix:       prefix,
				defaultValue: true,
			}
			exp := resStruct{
				res: true,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
		func() testCase {
			const name = "prefix does not exist false default"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const envName = "TEST_ENV_NAME"
			const prefix = "PREFIX"

			mockOs.EXPECT().LookupEnv(prefix+"_"+envName).Return("", false)

			in := args{
				unprefixed:   envName,
				prefix:       prefix,
				defaultValue: true,
			}
			exp := resStruct{
				res: true,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
	}

	t.Parallel()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res := resStruct{}
			res.res = tc.inst.Bool(tc.args.unprefixed, tc.args.defaultValue)
			is.Equal(tc.exp, res)

			if tc.ctrlr != nil {
				tc.ctrlr.Finish()
			}
		})
	}
}

func TestOinment_Float64(t *testing.T) {
	is := assert.New(t)

	type args struct {
		unprefixed   string
		prefix       string
		defaultValue float64
	}

	type resStruct struct {
		res float64
	}

	type testCase struct {
		name  string
		inst  *oinment
		ctrlr *gomock.Controller
		args  args
		exp   resStruct
	}

	cases := []testCase{
		func() testCase {
			const name = "no prefix exists"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const (
				envName = "TEST_ENV_NAME"
				v       = float64(0.12345)
			)

			mockOs.EXPECT().LookupEnv(envName).Return("0.12345", true).AnyTimes()

			in := args{
				unprefixed:   envName,
				prefix:       "",
				defaultValue: 5.4321,
			}
			exp := resStruct{
				res: v,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
		func() testCase {
			const name = "no prefix does not exist"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const (
				envName = "TEST_ENV_NAME"
				v       = float64(0.12345)
			)

			mockOs.EXPECT().LookupEnv(envName).Return("", false)

			in := args{
				unprefixed:   envName,
				prefix:       "",
				defaultValue: 5.4321,
			}
			exp := resStruct{
				res: in.defaultValue,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
		func() testCase {
			const name = "prefix exists"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const (
				envName = "TEST_ENV_NAME"
				prefix  = "ABCDE"
				v       = float64(0.12345)
			)

			mockOs.EXPECT().LookupEnv(prefix+"_"+envName).Return("0.12345", true).AnyTimes()

			in := args{
				unprefixed:   envName,
				prefix:       prefix,
				defaultValue: 5.4321,
			}
			exp := resStruct{
				res: v,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
		func() testCase {
			const name = "prefix does not exist"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const (
				envName = "TEST_ENV_NAME"
				prefix  = "ABCDE"
				v       = float64(0.12345)
			)

			mockOs.EXPECT().LookupEnv(prefix+"_"+envName).Return("", false)

			in := args{
				unprefixed:   envName,
				prefix:       prefix,
				defaultValue: 5.4321,
			}
			exp := resStruct{
				res: in.defaultValue,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
	}

	t.Parallel()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res := resStruct{}
			res.res = tc.inst.Float64(tc.args.unprefixed, tc.args.defaultValue)
			is.Equal(tc.exp, res)

			if tc.ctrlr != nil {
				tc.ctrlr.Finish()
			}
		})
	}
}

func TestOinment_String(t *testing.T) {
	is := assert.New(t)

	type args struct {
		unprefixed   string
		prefix       string
		defaultValue string
	}

	type resStruct struct {
		res string
	}

	type testCase struct {
		name  string
		inst  *oinment
		ctrlr *gomock.Controller
		args  args
		exp   resStruct
	}

	cases := []testCase{
		func() testCase {
			const name = "no prefix exists"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const (
				envName = "TEST_ENV_NAME"
				v       = "testValue"
			)

			mockOs.EXPECT().LookupEnv(envName).Return(v, true).AnyTimes()

			in := args{
				unprefixed:   envName,
				prefix:       "",
				defaultValue: "default",
			}
			exp := resStruct{
				res: v,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
		func() testCase {
			const name = "no prefix does not exist"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const (
				envName = "TEST_ENV_NAME"
			)

			mockOs.EXPECT().LookupEnv(envName).Return("", false)

			in := args{
				unprefixed:   envName,
				prefix:       "",
				defaultValue: "default",
			}
			exp := resStruct{
				res: in.defaultValue,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
		func() testCase {
			const name = "prefix exists"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const (
				envName = "TEST_ENV_NAME"
				prefix  = "ABCDE"
				v       = "testValue"
			)

			mockOs.EXPECT().LookupEnv(prefix+"_"+envName).Return(v, true).AnyTimes()

			in := args{
				unprefixed:   envName,
				prefix:       prefix,
				defaultValue: "default",
			}
			exp := resStruct{
				res: v,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
		func() testCase {
			const name = "prefix does not exist"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const (
				envName = "TEST_ENV_NAME"
				prefix  = "ABCDE"
			)

			mockOs.EXPECT().LookupEnv(prefix+"_"+envName).Return("", false)

			in := args{
				unprefixed:   envName,
				prefix:       prefix,
				defaultValue: "default",
			}
			exp := resStruct{
				res: in.defaultValue,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
	}

	t.Parallel()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res := resStruct{}
			res.res = tc.inst.String(tc.args.unprefixed, tc.args.defaultValue)
			is.Equal(tc.exp, res)

			if tc.ctrlr != nil {
				tc.ctrlr.Finish()
			}
		})
	}
}

func TestOinment_Int(t *testing.T) {
	is := assert.New(t)

	type args struct {
		unprefixed   string
		prefix       string
		defaultValue int
	}

	type resStruct struct {
		res int
	}

	type testCase struct {
		name  string
		inst  *oinment
		ctrlr *gomock.Controller
		args  args
		exp   resStruct
	}

	cases := []testCase{
		func() testCase {
			const name = "no prefix exists"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const (
				envName = "TEST_ENV_NAME"
				v       = 42
			)

			mockOs.EXPECT().LookupEnv(envName).Return(strconv.Itoa(v), true).AnyTimes()

			in := args{
				unprefixed:   envName,
				prefix:       "",
				defaultValue: 0,
			}
			exp := resStruct{
				res: v,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
		func() testCase {
			const name = "no prefix does not exist"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const (
				envName = "TEST_ENV_NAME"
			)

			mockOs.EXPECT().LookupEnv(envName).Return("", false)

			in := args{
				unprefixed:   envName,
				prefix:       "",
				defaultValue: 0,
			}
			exp := resStruct{
				res: in.defaultValue,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
		func() testCase {
			const name = "prefix exists"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const (
				envName = "TEST_ENV_NAME"
				prefix  = "ABCDE"
				v       = 42
			)

			mockOs.EXPECT().LookupEnv(prefix+"_"+envName).Return(strconv.Itoa(v), true).AnyTimes()

			in := args{
				unprefixed:   envName,
				prefix:       prefix,
				defaultValue: 0,
			}
			exp := resStruct{
				res: v,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
		func() testCase {
			const name = "prefix does not exist"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const (
				envName = "TEST_ENV_NAME"
				prefix  = "ABCDE"
			)

			mockOs.EXPECT().LookupEnv(prefix+"_"+envName).Return("", false)

			in := args{
				unprefixed:   envName,
				prefix:       prefix,
				defaultValue: 0,
			}
			exp := resStruct{
				res: in.defaultValue,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
	}

	t.Parallel()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res := resStruct{}
			res.res = tc.inst.Int(tc.args.unprefixed, tc.args.defaultValue)
			is.Equal(tc.exp, res)

			if tc.ctrlr != nil {
				tc.ctrlr.Finish()
			}
		})
	}
}

func TestOinment_Int64(t *testing.T) {
	is := assert.New(t)

	type args struct {
		unprefixed   string
		prefix       string
		defaultValue int64
	}

	type resStruct struct {
		res int64
	}

	type testCase struct {
		name  string
		inst  *oinment
		ctrlr *gomock.Controller
		args  args
		exp   resStruct
	}

	cases := []testCase{
		func() testCase {
			const name = "no prefix exists"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const (
				envName = "TEST_ENV_NAME"
				v       = int64(12345)
			)

			mockOs.EXPECT().LookupEnv(envName).Return(strconv.FormatInt(v, 10), true).AnyTimes()

			in := args{
				unprefixed:   envName,
				prefix:       "",
				defaultValue: 54321,
			}
			exp := resStruct{
				res: v,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
		func() testCase {
			const name = "no prefix does not exist"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const (
				envName = "TEST_ENV_NAME"
			)

			mockOs.EXPECT().LookupEnv(envName).Return("", false)

			in := args{
				unprefixed:   envName,
				prefix:       "",
				defaultValue: 54321,
			}
			exp := resStruct{
				res: in.defaultValue,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
		func() testCase {
			const name = "prefix exists"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const (
				envName = "TEST_ENV_NAME"
				prefix  = "ABCDE"
				v       = int64(12345)
			)

			mockOs.EXPECT().LookupEnv(prefix+"_"+envName).Return(strconv.FormatInt(v, 10), true).AnyTimes()

			in := args{
				unprefixed:   envName,
				prefix:       prefix,
				defaultValue: 54321,
			}
			exp := resStruct{
				res: v,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
		func() testCase {
			const name = "prefix does not exist"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const (
				envName = "TEST_ENV_NAME"
				prefix  = "ABCDE"
			)

			mockOs.EXPECT().LookupEnv(prefix+"_"+envName).Return("", false)

			in := args{
				unprefixed:   envName,
				prefix:       prefix,
				defaultValue: 54321,
			}
			exp := resStruct{
				res: in.defaultValue,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
	}

	t.Parallel()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res := resStruct{}
			res.res = tc.inst.Int64(tc.args.unprefixed, tc.args.defaultValue)
			is.Equal(tc.exp, res)

			if tc.ctrlr != nil {
				tc.ctrlr.Finish()
			}
		})
	}
}

func TestOinment_Uint64(t *testing.T) {
	is := assert.New(t)

	type args struct {
		unprefixed   string
		prefix       string
		defaultValue uint64
	}

	type resStruct struct {
		res uint64
	}

	type testCase struct {
		name  string
		inst  *oinment
		ctrlr *gomock.Controller
		args  args
		exp   resStruct
	}

	cases := []testCase{
		func() testCase {
			const name = "no prefix exists"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const (
				envName = "TEST_ENV_NAME"
				v       = uint64(12345)
			)

			mockOs.EXPECT().LookupEnv(envName).Return(strconv.FormatUint(v, 10), true).AnyTimes()

			in := args{
				unprefixed:   envName,
				prefix:       "",
				defaultValue: 54321,
			}
			exp := resStruct{
				res: v,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
		func() testCase {
			const name = "no prefix does not exist"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const (
				envName = "TEST_ENV_NAME"
			)

			mockOs.EXPECT().LookupEnv(envName).Return("", false)

			in := args{
				unprefixed:   envName,
				prefix:       "",
				defaultValue: 54321,
			}
			exp := resStruct{
				res: in.defaultValue,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
		func() testCase {
			const name = "prefix exists"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const (
				envName = "TEST_ENV_NAME"
				prefix  = "ABCDE"
				v       = uint64(12345)
			)

			mockOs.EXPECT().LookupEnv(prefix+"_"+envName).Return(strconv.FormatUint(v, 10), true).AnyTimes()

			in := args{
				unprefixed:   envName,
				prefix:       prefix,
				defaultValue: 54321,
			}
			exp := resStruct{
				res: v,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
		func() testCase {
			const name = "prefix does not exist"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const (
				envName = "TEST_ENV_NAME"
				prefix  = "ABCDE"
			)

			mockOs.EXPECT().LookupEnv(prefix+"_"+envName).Return("", false)

			in := args{
				unprefixed:   envName,
				prefix:       prefix,
				defaultValue: 54321,
			}
			exp := resStruct{
				res: in.defaultValue,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
	}

	t.Parallel()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res := resStruct{}
			res.res = tc.inst.Uint64(tc.args.unprefixed, tc.args.defaultValue)
			is.Equal(tc.exp, res)

			if tc.ctrlr != nil {
				tc.ctrlr.Finish()
			}
		})
	}
}

func TestOinment_Duration(t *testing.T) {
	is := assert.New(t)

	type args struct {
		unprefixed   string
		prefix       string
		defaultValue time.Duration
	}

	type resStruct struct {
		res time.Duration
	}

	type testCase struct {
		name  string
		inst  *oinment
		ctrlr *gomock.Controller
		args  args
		exp   resStruct
	}

	cases := []testCase{
		func() testCase {
			const name = "no prefix exists"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const (
				envName = "TEST_ENV_NAME"
				v       = "2h"
			)

			mockOs.EXPECT().LookupEnv(envName).Return(v, true).AnyTimes()

			in := args{
				unprefixed:   envName,
				prefix:       "",
				defaultValue: 54321,
			}
			exp := resStruct{
				res: 2 * time.Hour,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
		func() testCase {
			const name = "no prefix does not exist"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const (
				envName = "TEST_ENV_NAME"
			)

			mockOs.EXPECT().LookupEnv(envName).Return("", false)

			in := args{
				unprefixed:   envName,
				prefix:       "",
				defaultValue: 3 * time.Hour,
			}
			exp := resStruct{
				res: in.defaultValue,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
		func() testCase {
			const name = "prefix exists"

			ctrlr := gomock.NewController(t)
			mockOs := mocks.NewMockLookupInterface(ctrlr)

			const (
				envName = "TEST_ENV_NAME"
				prefix  = "ABCDE"
				v       = "6m"
			)

			mockOs.EXPECT().LookupEnv(prefix+"_"+envName).Return("6m", true).AnyTimes()

			in := args{
				unprefixed:   envName,
				prefix:       prefix,
				defaultValue: 3 * time.Hour,
			}
			exp := resStruct{
				res: 6 * time.Minute,
			}

			return testCase{
				name:  name,
				inst:  New(WithPrefix(in.prefix), WithLookup(mockOs)),
				args:  in,
				exp:   exp,
				ctrlr: ctrlr,
			}
		}(),
	}

	t.Parallel()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res := resStruct{}
			res.res = tc.inst.Duration(tc.args.unprefixed, tc.args.defaultValue)
			is.Equal(tc.exp, res)

			if tc.ctrlr != nil {
				tc.ctrlr.Finish()
			}
		})
	}
}
