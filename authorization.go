package auth

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

// Check User Permission
func (permission *Auth) IsGranted(modulename string, permisison Action) (bool, error) {

	if permission.RoleId != 1 { //if not an admin user

		// var modid string

		var module TblModule

		var modpermissions TblModulePermission

		if err := permission.DB.Model(TblModule{}).Where("module_name=? and parent_id !=0", modulename).First(&module).Error; err != nil {

			fmt.Println(err)
		}

		if err1 := permission.DB.Model(TblModulePermission{}).Where("display_name=?", modulename).Find(&modpermissions).Error; err1 != nil {

			return false, err1
		}

		var que string

		if module.Id != 0 {

			// modid = module.Id

			que = `and module_id =` + strconv.Itoa(module.Id)

		} else {

			// modid = modpermissions.Id

			que = `and id =` + strconv.Itoa(modpermissions.Id)
		}

		var modulepermission []TblModulePermission

		if permisison == "CRUD" {

			if err := permission.DB.Model(TblModulePermission{}).Where("(full_access_permission=1 or display_name='View' or display_name='Update' or  display_name='Create' or display_name='Delete')" + que + "").Find(&modulepermission).Error; err != nil {

				return false, err
			}

		} else {

			if err := permission.DB.Table("tbl_module_permissions").Where("display_name=? "+que+"", permisison).First(&modulepermission).Error; err != nil {

				return false, err
			}

		}

		if len(modulepermission) == 0 {

			return false, nil
		}

		for _, val := range modulepermission {

			var rolecheck TblRolePermission

			query := permission.DB.Model(TblRolePermission{}).Where("permission_id=? and role_id=?", val.Id, permission.RoleId).First(&rolecheck)

			if query.Error == gorm.ErrRecordNotFound {

				permission.PermissionFlg = true

				return false, ErrorUnauthorized

			}

			if err := query.Error; err != nil {

				return false, err
			}

		}

		permission.PermissionFlg = true

	}

	permission.PermissionFlg = true

	return true, nil

}
