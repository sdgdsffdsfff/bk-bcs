/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.,
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package template

import (
	"context"
	"errors"

	"github.com/spf13/viper"

	"bk-bscp/internal/database"
	"bk-bscp/internal/dbsharding"
	pbcommon "bk-bscp/internal/protocol/common"
	pb "bk-bscp/internal/protocol/datamanager"
	"bk-bscp/pkg/common"
)

// QueryAction query target config template object.
type QueryAction struct {
	ctx   context.Context
	viper *viper.Viper
	smgr  *dbsharding.ShardingManager

	req  *pb.QueryConfigTemplateReq
	resp *pb.QueryConfigTemplateResp

	sd *dbsharding.ShardingDB

	configTemplate database.ConfigTemplate
}

// NewQueryAction create new QueryAction
func NewQueryAction(ctx context.Context, viper *viper.Viper, smgr *dbsharding.ShardingManager,
	req *pb.QueryConfigTemplateReq, resp *pb.QueryConfigTemplateResp) *QueryAction {
	action := &QueryAction{ctx: ctx, viper: viper, smgr: smgr, req: req, resp: resp}

	action.resp.Seq = req.Seq
	action.resp.Code = pbcommon.ErrCode_E_OK
	action.resp.Message = "OK"

	return action
}

// Err setup error code message in response and return the error.
func (act *QueryAction) Err(errCode pbcommon.ErrCode, errMsg string) error {
	act.resp.Code = errCode
	act.resp.Message = errMsg
	return errors.New(errMsg)
}

// Input handles the input messages.
func (act *QueryAction) Input() error {
	if err := act.verify(); err != nil {
		return act.Err(pbcommon.ErrCode_E_DM_PARAMS_INVALID, err.Error())
	}
	return nil
}

// Output handles the output messages
func (act *QueryAction) Output() error {
	configTemplate := &pbcommon.ConfigTemplate{
		BizId:         act.configTemplate.BizID,
		TemplateId:    act.configTemplate.TemplateID,
		Name:          act.configTemplate.Name,
		CfgName:       act.configTemplate.CfgName,
		CfgFpath:      act.configTemplate.CfgFpath,
		User:          act.configTemplate.User,
		UserGroup:     act.configTemplate.UserGroup,
		FilePrivilege: act.configTemplate.FilePrivilege,
		FileFormat:    act.configTemplate.FileFormat,
		FileMode:      act.configTemplate.FileMode,
		EngineType:    act.configTemplate.EngineType,
		Memo:          act.configTemplate.Memo,
		Creator:       act.configTemplate.Creator,
		LastModifyBy:  act.configTemplate.LastModifyBy,
		State:         act.configTemplate.State,
		CreatedAt:     act.configTemplate.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:     act.configTemplate.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	act.resp.Data = configTemplate
	return nil
}

func (act *QueryAction) verify() error {
	var err error

	if err = common.ValidateString("biz_id", act.req.BizId,
		database.BSCPNOTEMPTY, database.BSCPIDLENLIMIT); err != nil {
		return err
	}
	if err = common.ValidateString("template_id", act.req.TemplateId,
		database.BSCPNOTEMPTY, database.BSCPIDLENLIMIT); err != nil {
		return err
	}
	return nil
}

func (act *QueryAction) queryConfigTemplate() (pbcommon.ErrCode, string) {
	err := act.sd.DB().
		Where(&database.ConfigTemplate{BizID: act.req.BizId, TemplateID: act.req.TemplateId}).
		Last(&act.configTemplate).Error

	if err == database.RECORDNOTFOUND {
		return pbcommon.ErrCode_E_DM_NOT_FOUND, "config template not found"
	}
	if err != nil {
		return pbcommon.ErrCode_E_DM_DB_EXEC_ERR, err.Error()
	}
	return pbcommon.ErrCode_E_OK, "OK"
}

// Do makes the workflows of this action base on input messages.
func (act *QueryAction) Do() error {
	// business sharding db.
	sd, err := act.smgr.ShardingDB(act.req.BizId)
	if err != nil {
		return act.Err(pbcommon.ErrCode_E_DM_ERR_DBSHARDING, err.Error())
	}
	act.sd = sd

	// query config template.
	if errCode, errMsg := act.queryConfigTemplate(); errCode != pbcommon.ErrCode_E_OK {
		return act.Err(errCode, errMsg)
	}
	return nil
}
