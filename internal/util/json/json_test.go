package json

import (
	"github.com/boywei/go-zero-check/internal/model"
	"github.com/boywei/go-zero-check/internal/test"
	"os"
	"reflect"
	"testing"
)

func TestConvertUppaal2Json(t *testing.T) {
	bytes, err := os.ReadFile("../../test/model1.json")
	if err != nil {
		panic(err)
	}
	model1JSON := string(bytes)

	type args struct {
		uppaalModel model.Uppaal
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				uppaalModel: *test.Model1,
			},
			want:    model1JSON,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertUppaal2Json(tt.args.uppaalModel)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertUppaal2Json() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertUppaal2Json() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertJson2Uppaal(t *testing.T) {
	bytes, err := os.ReadFile("../../test/model1.json")
	if err != nil {
		panic(err)
	}
	model1JSON := string(bytes)

	type args struct {
		jsonInput string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Uppaal
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				jsonInput: model1JSON,
			},
			want:    test.Model1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertJson2Uppaal(tt.args.jsonInput)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertJson2Uppaal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertJson2Uppaal() got = %v, want %v", got, tt.want)
			}
		})
	}
}
