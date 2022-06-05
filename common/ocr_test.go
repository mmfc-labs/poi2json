package common

import (
	"testing"
)

func TestGetBase64FromImage(t *testing.T) {
	type args struct {
		imagePath string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{imagePath: `D:\src\bookbook\minimap\Screenshot_2022-06-05-09-55-29-689_com.tencent.mm.jpg`}, want: "", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBase64FromImage(tt.args.imagePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBase64FromImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetBase64FromImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getOcr(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{name: "test1", want: "红色"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getOcr(); got != tt.want {
				t.Errorf("getOcr() = %v, want %v", got, tt.want)
			}
		})
	}
}
