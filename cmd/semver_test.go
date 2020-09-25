package cmd

import (
	"testing"
)

func TestParse(t *testing.T) {
	sv, err := Parse("v0.0.1")
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%v", sv)
}

func TestSemVer_Compare(t *testing.T) {
	type fields struct {
		Major int
		Minor int
		Patch int
	}
	type args struct {
		other SemVer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "Larger",
			fields: fields{
				Major: 0,
				Minor: 2,
				Patch: 1,
			},
			args: args{
				other: SemVer{
					Major: 0,
					Minor: 2,
					Patch: 0,
				},
			},
			want: 1,
		},
		{
			name: "Smaller",
			fields: fields{
				Major: 0,
				Minor: 2,
				Patch: 1,
			},
			args: args{
				other: SemVer{
					Major: 0,
					Minor: 2,
					Patch: 2,
				},
			},
			want: -1,
		},
		{
			name: "Equal",
			fields: fields{
				Major: 0,
				Minor: 2,
				Patch: 1,
			},
			args: args{
				other: SemVer{
					Major: 0,
					Minor: 2,
					Patch: 1,
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SemVer{
				Major: tt.fields.Major,
				Minor: tt.fields.Minor,
				Patch: tt.fields.Patch,
			}
			if got := s.Compare(tt.args.other); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
