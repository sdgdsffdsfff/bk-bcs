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
	"fmt"
	"math"

	"github.com/spf13/viper"

	"bk-bscp/internal/database"
	pbcommon "bk-bscp/internal/protocol/common"
	pbdatamanager "bk-bscp/internal/protocol/datamanager"
	pb "bk-bscp/internal/protocol/templateserver"
	"bk-bscp/pkg/common"
	"bk-bscp/pkg/kit"
	"bk-bscp/pkg/logger"
)

// ListAction list target config template objects.
type ListAction struct {
	kit        kit.Kit
	viper      *viper.Viper
	dataMgrCli pbdatamanager.DataManagerClient

	req  *pb.QueryConfigTemplateListReq
	resp *pb.QueryConfigTemplateListResp

	totalCount uint32
	templates  []*pbcommon.ConfigTemplate
}

// NewListAction creates new ListAction.
func NewListAction(kit kit.Kit, viper *viper.Viper, dataMgrCli pbdatamanager.DataManagerClient,
	req *pb.QueryConfigTemplateListReq, resp *pb.QueryConfigTemplateListResp) *ListAction {
	action := &ListAction{kit: kit, viper: viper, dataMgrCli: dataMgrCli, req: req, resp: resp}

	action.resp.Result = true
	action.resp.Code = pbcommon.ErrCode_E_OK
	action.resp.Message = "OK"

	return action
}

// Err setup error code message in response and return the error.
func (act *ListAction) Err(errCode pbcommon.ErrCode, errMsg string) error {
	if errCode != pbcommon.ErrCode_E_OK {
		act.resp.Result = false
	}
	act.resp.Code = errCode
	act.resp.Message = errMsg
	return errors.New(errMsg)
}

// Input handles the input messages.
func (act *ListAction) Input() error {
	if err := act.verify(); err != nil {
		return act.Err(pbcommon.ErrCode_E_TPL_PARAMS_INVALID, err.Error())
	}
	return nil
}

// Output handles the output messages.
func (act *ListAction) Output() error {
	act.resp.Data = &pb.QueryConfigTemplateListResp_RespData{TotalCount: act.totalCount, Info: act.templates}
	return nil
}

func (act *ListAction) verify() error {
	var err error

	if err = common.ValidateString("biz_id", act.req.BizId,
		database.BSCPNOTEMPTY, database.BSCPIDLENLIMIT); err != nil {
		return err
	}

	if act.req.Page == nil {
		return errors.New("invalid input data, page is required")
	}
	if err = common.ValidateInt32("page.start", act.req.Page.Start, 0, math.MaxInt32); err != nil {
		return err
	}
	if err = common.ValidateInt32("page.limit", act.req.Page.Limit, 1, database.BSCPQUERYLIMIT); err != nil {
		return err
	}
	return nil
}

func (act *ListAction) queryConfigTemplateList() (pbcommon.ErrCode, string) {
	r := &pbdatamanager.QueryConfigTemplateListReq{
		Seq:   act.kit.Rid,
		BizId: act.req.BizId,
		Page:  act.req.Page,
	}

	ctx, cancel := context.WithTimeout(act.kit.Ctx, act.viper.GetDuration("datamanager.callTimeout"))
	defer cancel()

	logger.V(4).Infof("QueryConfigTemplateList[%s]| request to DataManager, %+v", r.Seq, r)

	resp, err := act.dataMgrCli.QueryConfigTemplateList(ctx, r)
	if err != nil {
		return pbcommon.ErrCode_E_TPL_SYSTEM_UNKNOWN,
			fmt.Sprintf("request to DataManager QueryConfigTemplateList, %+v", err)
	}
	if resp.Code != pbcommon.ErrCode_E_OK {
		return resp.Code, resp.Message
	}
	act.totalCount = resp.Data.TotalCount
	act.templates = resp.Data.Info

	return pbcommon.ErrCode_E_OK, ""
}

// Do makes the workflows of this action base on input messages.
func (act *ListAction) Do() error {
	// query config template list.
	if errCode, errMsg := act.queryConfigTemplateList(); errCode != pbcommon.ErrCode_E_OK {
		return act.Err(errCode, errMsg)
	}
	return nil
}
