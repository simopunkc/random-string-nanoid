package service

import (
	"random-string-generator-service/internal/app/nanoid/config"
	"random-string-generator-service/internal/app/nanoid/repository"
	"random-string-generator-service/internal/app/nanoid/stub"
	"reflect"
	"testing"
)

func TestNanoidService_RunGenerateRandomString(t *testing.T) {
	type fields struct {
		nanoidRepo NanoidRepository
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "case length random string is not match",
			fields: fields{
				nanoidRepo: &NanoidRepositoryMock{
					GenerateUniqueIDFunc: func() string {
						return "ap"
					},
				},
			},
			want: false,
		},
		{
			name: "case length random string is match",
			fields: fields{
				nanoidRepo: &NanoidRepositoryMock{
					GenerateUniqueIDFunc: func() string {
						return stub.NanoidResult
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ns := NanoidService{
				nanoidRepo: tt.fields.nanoidRepo,
			}
			if _, ok := ns.RunGenerateRandomString(); !reflect.DeepEqual(ok, tt.want) {
				t.Errorf("NanoidService.RunGenerateRandomString() = %v, want %v", ok, tt.want)
			}
		})
	}
}

func BenchmarkNanoidService_RunGenerateRandomString(b *testing.B) {
	limit := config.LimitCharacter
	prefix := config.Prefix
	repoNano := repository.NewCmdNanoidRepository(limit, prefix)
	for i := 0; i < b.N; i++ {
		ns := NanoidService{
			nanoidRepo: repoNano,
		}
		ns.RunGenerateRandomString()
	}
}
