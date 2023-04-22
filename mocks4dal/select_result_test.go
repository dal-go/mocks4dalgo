package mocks4dal

import (
	"github.com/dal-go/dalgo/dal"
	"testing"
)

func TestNewSingleRecordReader(t *testing.T) {
	NewSingleRecordReader(nil, "", nil)
}

func TestNewSelectResult(t *testing.T) {
	type args struct {
		getReader func(into func() dal.Record) dal.Reader
		err       error
	}
	tests := []struct {
		name string
		args args
		want SelectResult
	}{
		{"empty", args{func(into func() dal.Record) dal.Reader {
			return nil
		}, nil}, SelectResult{nil, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = NewSelectResult(tt.args.getReader, tt.args.err)
		})
	}
}
