package auth

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

// Check User Permission
func (permission *Auth) IsGranted(modulename string, permisison Action, tenantid int) (bool, error) {

	if (permission.RoleId != 1 || permission.RoleId !=2){ //if not an admin user

		// var modid string

		var module TblModule

		var modpermissions TblModulePermission

		if err := permission.DB.Debug().Model(TblModule{}).Where("module_name=? and parent_id !=0 and (tenant_id is NULL or tenant_id=?)", modulename, tenantid).First(&module).Error; err != nil {

			fmt.Println(err)
		}
		fmt.Println("module:", module)

		if err1 := permission.DB.Debug().Model(TblModulePermission{}).Where("display_name=? and (tenant_id is NULL or tenant_id=?)", modulename, tenantid).Find(&modpermissions).Error; err1 != nil {

			return false, err1
		}
		fmt.Println("modpermissions:", modpermissions)

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

			if err := permission.DB.Model(TblModulePermission{}).Where("(full_access_permission=1 or display_name='View' or display_name='Update' or  display_name='Create' or display_name='Delete') and (tenant_id is NULL or tenant_id=?)"+que+"", tenantid).Find(&modulepermission).Error; err != nil {

				return false, err
			}

		} else {

			if err := permission.DB.Table("tbl_module_permissions").Where("display_name=? and (tenant_id is NULL or tenant_id=?)"+que+"", permisison, tenantid).First(&modulepermission).Error; err != nil {

				return false, err
			}

		}

		if len(modulepermission) == 0 {

			return false, nil
		}

		for _, val := range modulepermission {

			var rolecheck TblRolePermission

			query := permission.DB.Model(TblRolePermission{}).Where("permission_id=? and role_id=? and (tenant_id is NULL or tenant_id=?)", val.Id, permission.RoleId, tenantid).First(&rolecheck)

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