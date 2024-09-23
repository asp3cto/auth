package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/asp3cto/auth/internal/api/user"
	"github.com/asp3cto/auth/internal/model"
	"github.com/asp3cto/auth/internal/service"
	serviceMocks "github.com/asp3cto/auth/internal/service/mocks"
	desc "github.com/asp3cto/auth/pkg/user_v1"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id    = gofakeit.Int64()
		name  = gofakeit.Name()
		email = gofakeit.Email()

		serviceErr = fmt.Errorf("serviec error")

		req = &desc.CreateRequest{
			Info: &desc.UserInfo{
				Name:  name,
				Email: email,
			},
		}

		info = &model.UserInfo{
			Name:  name,
			Email: email,
		}

		res = &desc.CreateResponse{
			Id: id,
		}
	)

	// defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateResponse
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.CreateMock.Expect(ctx, info).Return(id, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.CreateMock.Expect(ctx, info).Return(0, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				userServiceMock := tt.userServiceMock(mc)
				api := user.NewImplementation(userServiceMock)

				resHandler, err := api.Create(tt.args.ctx, tt.args.req)
				require.Equal(t, tt.err, err)
				require.Equal(t, tt.want, resHandler)
			},
		)
	}

}
