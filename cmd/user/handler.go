package main

import (
	"byteDemo/cmd/user/dao"
	user "byteDemo/kitex_gen/user"
	"context"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreatUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreatUser(ctx context.Context, req *user.Req) (resp *user.Resp, err error) {
	// TODO: Your code here...
	err = dao.CreatUser(req)
	if err != nil {
		return nil, err
	}
	resp = new(user.Resp)
	resp.Code = 200
	resp.Msg = "ok"

	return resp, nil
}
