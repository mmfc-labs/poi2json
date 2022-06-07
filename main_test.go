package main

import (
	"reflect"
	"testing"
)

func TestGetTowards(t *testing.T) {
	type args struct {
		poiss string
	}
	tests := []struct {
		name    string
		args    args
		want    []DirectionPoint
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "test", args: args{poiss: "22.577288,113.913726;22.578052,113.913598"}, want: []DirectionPoint{{Lat: 22.577288, Lon: 113.913726, Auth: User, Status: Open}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTowards(tt.args.poiss)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTowards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTowards() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTowardStringAndAuthStatus(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   Role
		want2   Status
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "test", args: args{s: "22.577288,113.913726"}, want: "22.577288,113.91372d", want1: User, want2: Open, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := GetTowardStringAndAuthStatus(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTowardStringAndAuthStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetTowardStringAndAuthStatus() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetTowardStringAndAuthStatus() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("GetTowardStringAndAuthStatus() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestParseLine(t *testing.T) {
	type args struct {
		l string
	}
	tests := []struct {
		name    string
		args    args
		want    Point
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "te", args: args{l: `22.571779,113.919897 广东省深圳市宝安区大宝路33号(广深高速入口闸道) - red 22.571816,113.920914`}, want: Point{}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseLine(tt.args.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
