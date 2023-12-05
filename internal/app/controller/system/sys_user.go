package system

import (
	"github.com/championlong/go-quick-start/internal/app/model/common/request"
	response2 "github.com/championlong/go-quick-start/internal/app/model/common/response"
	system2 "github.com/championlong/go-quick-start/internal/app/model/system"
	request2 "github.com/championlong/go-quick-start/internal/app/model/system/request"
	"github.com/championlong/go-quick-start/internal/app/model/system/response"
	global2 "github.com/championlong/go-quick-start/internal/pkg/global"
	utils2 "github.com/championlong/go-quick-start/internal/pkg/utils"
	"github.com/mojocn/base64Captcha"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

var store = base64Captcha.DefaultMemStore

type BaseApi struct{}

// Captcha
// @Tags Base
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=systemRes.SysCaptchaResponse,msg=string} "生成验证码,返回包括随机数id,base64,验证码长度"
// @Router /base/captcha [post]
func (b *BaseApi) Captcha(c *gin.Context) {
	// 判断验证码是否开启
	openCaptcha := global2.GVA_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global2.GVA_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	key := c.ClientIP()
	v, ok := global2.BlackCache.Get(key)
	if !ok {
		global2.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc bool
	if openCaptcha == 0 || openCaptcha < interfaceToInt(v) {
		oc = true
	}

	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global2.GVA_CONFIG.Captcha.ImgHeight, global2.GVA_CONFIG.Captcha.ImgWidth, global2.GVA_CONFIG.Captcha.KeyLong, 0.7, 80)
	// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		global2.GVA_LOG.Error("验证码获取失败!", zap.Error(err))
		response2.FailWithMessage("验证码获取失败", c)
		return
	}
	response2.OkWithDetailed(response.SysCaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: global2.GVA_CONFIG.Captcha.KeyLong,
		OpenCaptcha:   oc,
	}, "验证码获取成功", c)
}

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body systemReq.Login true "用户名, 密码, 验证码"
// @Success 200 {object} response.Response{data=systemRes.LoginResponse,msg=string} "返回包括用户信息,token,过期时间"
// @Router /base/login [post]
func (b *BaseApi) Login(c *gin.Context) {
	var l request2.Login
	_ = c.ShouldBindJSON(&l)
	if err := utils2.Verify(l, utils2.LoginVerify); err != nil {
		response2.FailWithMessage(err.Error(), c)
		return
	}
	if store.Verify(l.CaptchaId, l.Captcha, true) {
		u := &system2.SysUser{Username: l.Username, Password: l.Password}
		if err, user := userService.Login(u); err != nil {
			global2.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			response2.FailWithMessage("用户名不存在或者密码错误", c)
		} else {
			b.tokenNext(c, *user)
		}
	} else {
		response2.FailWithMessage("验证码错误", c)
	}
}

// 登录以后签发jwt
func (b *BaseApi) tokenNext(c *gin.Context, user system2.SysUser) {
	j := &utils2.JWT{SigningKey: []byte(global2.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(request2.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global2.GVA_LOG.Error("获取token失败!", zap.Error(err))
		response2.FailWithMessage("获取token失败", c)
		return
	}
	if !global2.GVA_CONFIG.System.UseMultipoint {
		response2.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}

	if err, jwtStr := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			global2.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
			response2.FailWithMessage("设置登录状态失败", c)
			return
		}
		response2.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global2.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
		response2.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT system2.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response2.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			response2.FailWithMessage("设置登录状态失败", c)
			return
		}
		response2.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}

// @Tags SysUser
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body systemReq.Register true "用户名, 昵称, 密码, 角色ID"
// @Success 200 {object} response.Response{data=systemRes.SysUserResponse,msg=string} "用户注册账号,返回包括用户信息"
// @Router /user/admin_register [post]
func (b *BaseApi) Register(c *gin.Context) {
	var r request2.Register
	_ = c.ShouldBindJSON(&r)
	if err := utils2.Verify(r, utils2.RegisterVerify); err != nil {
		response2.FailWithMessage(err.Error(), c)
		return
	}
	var authorities []system2.SysAuthority
	for _, v := range r.AuthorityIds {
		authorities = append(authorities, system2.SysAuthority{
			AuthorityId: v,
		})
	}
	user := &system2.SysUser{Username: r.Username, NickName: r.NickName, Password: r.Password, HeaderImg: r.HeaderImg, AuthorityId: r.AuthorityId}
	err, userReturn := userService.Register(*user)
	if err != nil {
		global2.GVA_LOG.Error("注册失败!", zap.Error(err))
		response2.FailWithDetailed(response.SysUserResponse{User: userReturn}, "注册失败", c)
	} else {
		response2.OkWithDetailed(response.SysUserResponse{User: userReturn}, "注册成功", c)
	}
}

// @Tags SysUser
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body systemReq.ChangePasswordStruct true "用户名, 原密码, 新密码"
// @Success 200 {object} response.Response{msg=string} "用户修改密码"
// @Router /user/changePassword [post]
func (b *BaseApi) ChangePassword(c *gin.Context) {
	var user request2.ChangePasswordStruct
	_ = c.ShouldBindJSON(&user)
	if err := utils2.Verify(user, utils2.ChangePasswordVerify); err != nil {
		response2.FailWithMessage(err.Error(), c)
		return
	}
	u := &system2.SysUser{Username: user.Username, Password: user.Password}
	if err, _ := userService.ChangePassword(u, user.NewPassword); err != nil {
		global2.GVA_LOG.Error("修改失败!", zap.Error(err))
		response2.FailWithMessage("修改失败，原密码与当前账户不符", c)
	} else {
		response2.OkWithMessage("修改成功", c)
	}
}

// @Tags SysUser
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取用户列表,返回包括列表,总数,页码,每页数量"
// @Router /user/getUserList [post]
func (b *BaseApi) GetUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils2.Verify(pageInfo, utils2.PageInfoVerify); err != nil {
		response2.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := userService.GetUserInfoList(pageInfo); err != nil {
		global2.GVA_LOG.Error("获取失败!", zap.Error(err))
		response2.FailWithMessage("获取失败", c)
	} else {
		response2.OkWithDetailed(response2.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// @Tags SysUser
// @Summary 更改用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.SetUserAuth true "用户UUID, 角色ID"
// @Success 200 {object} response.Response{msg=string} "设置用户权限"
// @Router /user/setUserAuthority [post]
func (b *BaseApi) SetUserAuthority(c *gin.Context) {
	var sua request2.SetUserAuth
	_ = c.ShouldBindJSON(&sua)
	if UserVerifyErr := utils2.Verify(sua, utils2.SetUserAuthorityVerify); UserVerifyErr != nil {
		response2.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	userID := utils2.GetUserID(c)
	uuid := utils2.GetUserUuid(c)
	if err := userService.SetUserAuthority(userID, uuid, sua.AuthorityId); err != nil {
		global2.GVA_LOG.Error("修改失败!", zap.Error(err))
		response2.FailWithMessage(err.Error(), c)
	} else {
		claims := utils2.GetUserInfo(c)
		j := &utils2.JWT{SigningKey: []byte(global2.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
		claims.AuthorityId = sua.AuthorityId
		if token, err := j.CreateToken(*claims); err != nil {
			global2.GVA_LOG.Error("修改失败!", zap.Error(err))
			response2.FailWithMessage(err.Error(), c)
		} else {
			c.Header("new-token", token)
			c.Header("new-expires-at", strconv.FormatInt(claims.ExpiresAt, 10))
			response2.OkWithMessage("修改成功", c)
		}

	}
}

// @Tags SysUser
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.SetUserAuthorities true "用户UUID, 角色ID"
// @Success 200 {object} response.Response{msg=string} "设置用户权限"
// @Router /user/setUserAuthorities [post]
func (b *BaseApi) SetUserAuthorities(c *gin.Context) {
	var sua request2.SetUserAuthorities
	_ = c.ShouldBindJSON(&sua)
	if err := userService.SetUserAuthorities(sua.ID, sua.AuthorityIds); err != nil {
		global2.GVA_LOG.Error("修改失败!", zap.Error(err))
		response2.FailWithMessage("修改失败", c)
	} else {
		response2.OkWithMessage("修改成功", c)
	}
}

// @Tags SysUser
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "用户ID"
// @Success 200 {object} response.Response{msg=string} "删除用户"
// @Router /user/deleteUser [delete]
func (b *BaseApi) DeleteUser(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)
	if err := utils2.Verify(reqId, utils2.IdVerify); err != nil {
		response2.FailWithMessage(err.Error(), c)
		return
	}
	jwtId := utils2.GetUserID(c)
	if jwtId == uint(reqId.ID) {
		response2.FailWithMessage("删除失败, 自杀失败", c)
		return
	}
	if err := userService.DeleteUser(reqId.ID); err != nil {
		global2.GVA_LOG.Error("删除失败!", zap.Error(err))
		response2.FailWithMessage("删除失败", c)
	} else {
		response2.OkWithMessage("删除成功", c)
	}
}

// @Tags SysUser
// @Summary 设置用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysUser true "ID, 用户名, 昵称, 头像链接"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "设置用户信息"
// @Router /user/setUserInfo [put]
func (b *BaseApi) SetUserInfo(c *gin.Context) {
	var user request2.ChangeUserInfo
	_ = c.ShouldBindJSON(&user)
	if err := utils2.Verify(user, utils2.IdVerify); err != nil {
		response2.FailWithMessage(err.Error(), c)
		return
	}

	if len(user.AuthorityIds) != 0 {
		err := userService.SetUserAuthorities(user.ID, user.AuthorityIds)
		if err != nil {
			global2.GVA_LOG.Error("设置失败!", zap.Error(err))
			response2.FailWithMessage("设置失败", c)
		}
	}

	if err := userService.SetUserInfo(system2.SysUser{
		GVA_MODEL: global2.GVA_MODEL{
			ID: user.ID,
		},
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		Phone:     user.Phone,
		Email:     user.Email,
	}); err != nil {
		global2.GVA_LOG.Error("设置失败!", zap.Error(err))
		response2.FailWithMessage("设置失败", c)
	} else {
		response2.OkWithMessage("设置成功", c)
	}
}

// @Tags SysUser
// @Summary 设置用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysUser true "ID, 用户名, 昵称, 头像链接"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "设置用户信息"
// @Router /user/SetSelfInfo [put]
func (b *BaseApi) SetSelfInfo(c *gin.Context) {
	var user request2.ChangeUserInfo
	_ = c.ShouldBindJSON(&user)
	user.ID = utils2.GetUserID(c)
	if err := userService.SetUserInfo(system2.SysUser{
		GVA_MODEL: global2.GVA_MODEL{
			ID: user.ID,
		},
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		Phone:     user.Phone,
		Email:     user.Email,
	}); err != nil {
		global2.GVA_LOG.Error("设置失败!", zap.Error(err))
		response2.FailWithMessage("设置失败", c)
	} else {
		response2.OkWithMessage("设置成功", c)
	}
}

// @Tags SysUser
// @Summary 获取用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取用户信息"
// @Router /user/getUserInfo [get]
func (b *BaseApi) GetUserInfo(c *gin.Context) {
	uuid := utils2.GetUserUuid(c)
	if err, ReqUser := userService.GetUserInfo(uuid); err != nil {
		global2.GVA_LOG.Error("获取失败!", zap.Error(err))
		response2.FailWithMessage("获取失败", c)
	} else {
		response2.OkWithDetailed(gin.H{"userInfo": ReqUser}, "获取成功", c)
	}
}

// @Tags SysUser
// @Summary 重置用户密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body system.SysUser true "ID"
// @Success 200 {object} response.Response{msg=string} "重置用户密码"
// @Router /user/resetPassword [post]
func (b *BaseApi) ResetPassword(c *gin.Context) {
	var user system2.SysUser
	_ = c.ShouldBindJSON(&user)
	if err := userService.ResetPassword(user.ID); err != nil {
		global2.GVA_LOG.Error("重置失败!", zap.Error(err))
		response2.FailWithMessage("重置失败"+err.Error(), c)
	} else {
		response2.OkWithMessage("重置成功", c)
	}
}

// 类型转换
func interfaceToInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		i = v
	default:
		i = 0
	}
	return
}