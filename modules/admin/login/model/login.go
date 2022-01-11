package model

import (
	"errors"
	"fmt"

	"github.com/jackjewl/squirrelBlog/db"
	"github.com/jackjewl/squirrelBlog/modules/admin/login/entities"
	"github.com/jackjewl/squirrelBlog/utils"
)

// 这里传入的admin的password是不需要md5加密的
//后续会用https  stl来进行安全的传输
func AdminLogin(admin entities.Admin) (success bool, err error) {

	var rawAdmin *entities.Admin
	err = db.DB.Model(&entities.Admin{}).Where("email = ", admin.Email).First(rawAdmin).Error

	if err != nil {
		fmt.Println("查找不到管理员信息")
		return false, err
	}

	if rawAdmin != nil && rawAdmin.Password == utils.MD5Encrypt(admin.Password) {
		return true, nil
	} else {
		return false, errors.New("密码错误")
	}
}
