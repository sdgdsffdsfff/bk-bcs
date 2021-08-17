/*
Tencent is pleased to support the open source community by making Blueking Container Service available.
Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
Licensed under the MIT License (the "License"); you may not use this file except
in compliance with the License. You may obtain a copy of the License at
http://opensource.org/licenses/MIT
Unless required by applicable law or agreed to in writing, software distributed under
the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
either express or implied. See the License for the specific language governing permissions and
limitations under the License.
*/

package shardingdb

import (
	"context"
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"

	"bk-bscp/internal/database"
	"bk-bscp/internal/dbsharding"
	pbcommon "bk-bscp/internal/protocol/common"
	pb "bk-bscp/internal/protocol/datamanager"
	"bk-bscp/pkg/common"
)

// CreateAction is shardingdb create action object.
type CreateAction struct {
	ctx   context.Context
	viper *viper.Viper
	smgr  *dbsharding.ShardingManager

	req  *pb.CreateShardingDBReq
	resp *pb.CreateShardingDBResp
}

// NewCreateAction creates new CreateAction.
func NewCreateAction(ctx context.Context, viper *viper.Viper, smgr *dbsharding.ShardingManager,
	req *pb.CreateShardingDBReq, resp *pb.CreateShardingDBResp) *CreateAction {
	action := &CreateAction{ctx: ctx, viper: viper, smgr: smgr, req: req, resp: resp}

	action.resp.Seq = req.Seq
	action.resp.Code = pbcommon.ErrCode_E_OK
	action.resp.Message = "OK"

	return action
}

// Err setup error code message in response and return the error.
func (act *CreateAction) Err(errCode pbcommon.ErrCode, errMsg string) error {
	act.resp.Code = errCode
	act.resp.Message = errMsg
	return errors.New(errMsg)
}

// Input handles the input messages.
func (act *CreateAction) Input() error {
	if err := act.verify(); err != nil {
		return act.Err(pbcommon.ErrCode_E_DM_PARAMS_INVALID, err.Error())
	}
	return nil
}

// Output handles the output messages.
func (act *CreateAction) Output() error {
	// do nothing.
	return nil
}

func (act *CreateAction) verify() error {
	var err error

	if err = common.ValidateString("db_id", act.req.DbId,
		database.BSCPNOTEMPTY, database.BSCPLONGSTRLENLIMIT); err != nil {
		return err
	}
	if err = common.ValidateString("host", act.req.Host,
		database.BSCPNOTEMPTY, database.BSCPLONGSTRLENLIMIT); err != nil {
		return err
	}
	if act.req.Port == 0 {
		return errors.New("invalid input data, port is required")
	}
	if err = common.ValidateString("user", act.req.User,
		database.BSCPNOTEMPTY, database.BSCPNORMALSTRLENLIMIT); err != nil {
		return err
	}
	if err = common.ValidateString("password", act.req.Password,
		database.BSCPEMPTY, database.BSCPNORMALSTRLENLIMIT); err != nil {
		return err
	}
	if err = common.ValidateString("memo", act.req.Memo,
		database.BSCPEMPTY, database.BSCPLONGSTRLENLIMIT); err != nil {
		return err
	}
	return nil
}

func (act *CreateAction) createShardingDB() (pbcommon.ErrCode, string) {
	db := &pbcommon.ShardingDB{
		DbId:     act.req.DbId,
		Host:     act.req.Host,
		Port:     act.req.Port,
		User:     act.req.User,
		Password: act.req.Password,
		Memo:     act.req.Memo,
	}

	err := act.smgr.CreateShardingDB(db)
	if err != nil {
		e, ok := err.(*mysql.MySQLError)
		if !ok {
			return pbcommon.ErrCode_E_DM_DB_EXEC_ERR, err.Error()
		}
		if e.Number == 1062 {
			return pbcommon.ErrCode_E_DM_ALREADY_EXISTS, "the sharding database of target dbid already exist."
		}
		return pbcommon.ErrCode_E_DM_DB_EXEC_ERR, err.Error()
	}
	return pbcommon.ErrCode_E_OK, ""
}

// Do makes the workflows of this action base on input messages.
func (act *CreateAction) Do() error {
	// create shardingdb.
	if errCode, errMsg := act.createShardingDB(); errCode != pbcommon.ErrCode_E_OK {
		return act.Err(errCode, errMsg)
	}
	return nil
}
