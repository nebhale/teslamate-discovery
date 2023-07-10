package units

import (
	"reflect"
	"testing"

	"github.com/spf13/cobra"
)

func TestRangeType_Prefix(t *testing.T) {
	tests := []struct {
		name string
		r    RangeType
		want string
	}{
		{
			name: "estimated",
			r:    Estimated,
			want: "est",
		},
		{
			name: "ideal",
			r:    Ideal,
			want: "ideal",
		},
		{
			name: "rated",
			r:    Rated,
			want: "rated",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Prefix(); got != tt.want {
				t.Errorf("RangeType.Prefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRangeType_Set(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		args    args
		want    RangeType
		wantErr bool
	}{
		{
			name: "estimated",
			args: args{v: "estimated"},
			want: Estimated,
		},
		{
			name: "ideal",
			args: args{v: "ideal"},
			want: Ideal,
		},
		{
			name: "rated",
			args: args{v: "rated"},
			want: Rated,
		},
		{
			name:    "unknown",
			args:    args{v: "unknown"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r RangeType
			err := r.Set(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("RangeType.Set() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}

			if r != tt.want {
				t.Errorf("RangeType.Set() = %v, want %v", r, tt.want)
			}
		})
	}
}

func TestRangeType_String(t *testing.T) {
	tests := []struct {
		name string
		r    RangeType
		want string
	}{
		{
			name: "estimated",
			r:    Estimated,
			want: "estimated",
		},
		{
			name: "ideal",
			r:    Ideal,
			want: "ideal",
		},
		{
			name: "rated",
			r:    Rated,
			want: "rated",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.String(); got != tt.want {
				t.Errorf("RangeType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRangeType_Type(t *testing.T) {
	tests := []struct {
		name string
		r    RangeType
		want string
	}{
		{
			name: "estimated",
			r:    Estimated,
			want: "string",
		},
		{
			name: "ideal",
			r:    Ideal,
			want: "string",
		},
		{
			name: "rated",
			r:    Rated,
			want: "string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Type(); got != tt.want {
				t.Errorf("RangeType.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRangeTypeCompletion(t *testing.T) {
	type args struct {
		cmd        *cobra.Command
		args       []string
		toComplete string
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 cobra.ShellCompDirective
	}{
		{
			name:  "always",
			args:  args{},
			want:  []string{"estimated", "ideal", "rated"},
			want1: cobra.ShellCompDirectiveDefault,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := RangeTypeCompletion(tt.args.cmd, tt.args.args, tt.args.toComplete)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RangeTypeCompletion() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RangeTypeCompletion() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
