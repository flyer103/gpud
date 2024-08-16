package fabricmanagerlog

import (
	"reflect"
	"testing"
	"time"
)

func TestExtractTimeFromLogLine(t *testing.T) {
	t.Parallel()

	type args struct {
		line []byte
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "expected log",
			args: args{
				line: []byte("[Jul 09 2024 18:14:07] [ERROR] [tid 12727] detected NVSwitch non-fatal error 12028 on fid 0 on NVSwitch pci bus id 00000000:86:00.0 physical id 3 port 61"),
			},
			want:    time.Date(2024, time.July, 9, 18, 14, 07, 0, time.UTC),
			wantErr: false,
		},
		{
			name: "unexpected log",
			args: args{
				line: []byte("[2024-07-09 18:14:07] [ERROR] [tid 12727] detected NVSwitch non-fatal error 12028 on fid 0 on NVSwitch pci bus id 00000000:86:00.0 physical id 3 port 61"),
			},
			want:    time.Time{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractTimeFromLogLine(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractTimeFromLogLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractTimeFromLogLine() got = %v, want %v", got, tt.want)
			}
		})
	}
}