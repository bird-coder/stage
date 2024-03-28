package schema

import (
	"time"
)

type AdminUserRoles struct {
	RoleID    int64     `gorm:"primary_key" json:"role_id"`
	ModelType string    `gorm:"primary_key" json:"model_type"`
	ModelID   int64     `gorm:"primary_key" json:"model_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FailedJobs struct {
	ID         int64     `gorm:"primary_key" json:"-"`
	UUID       string    `gorm:"unique" json:"uuid"`
	Connection string    `json:"connection"`
	Queue      string    `json:"queue"`
	Payload    string    `json:"payload"`
	Exception  string    `json:"exception"`
	FailedAt   time.Time `json:"failed_at"`
}

type AdminUserPermissions struct {
	PermissionID int64  `gorm:"primary_key" json:"permission_id"`
	ModelType    string `gorm:"primary_key" json:"model_type"`
	ModelID      int64  `gorm:"primary_key" json:"model_id"`
}

type AdminOperationLog struct {
	ID          int64     `gorm:"primary_key" json:"-"`
	UserID      int64     `gorm:"index" json:"user_id"`
	AssignUsers string    `json:"assign_users"`
	Path        string    `json:"path"`
	Method      string    `json:"method"`
	IP          string    `json:"ip"`
	Input       string    `json:"input"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Code        int       `json:"code"`
}

type AdminRoles struct {
	ID        int64     `gorm:"primary_key" json:"-"`
	Name      string    `gorm:"unique_index:admin_roles_name_guard_name_unique" json:"name"`
	GuardName string    `gorm:"unique_index:admin_roles_name_guard_name_unique" json:"guard_name"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AdminUser struct {
	ID            int       `gorm:"primary_key" json:"-"`
	Username      string    `gorm:"unique" json:"username"`
	Password      string    `json:"password"`
	Name          string    `json:"name"`
	Avatar        string    `json:"avatar"`
	RememberToken string    `json:"remember_token"`
	Phone         string    `json:"phone"`
	Email         string    `json:"email"`
	Status        bool      `json:"status"` // 用户状态 1可用 2离职 3其它
	HomeMenu      int       `json:"home_menu"`
	Attempt       bool      `json:"attempt"` // 登录失败次数
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Migrations struct {
	ID        int    `gorm:"primary_key" json:"-"`
	Migration string `json:"migration"`
	Batch     int    `json:"batch"`
}

type OauthClients struct {
	ID           int       `gorm:"primary_key" json:"-"`
	Name         string    `json:"name"` // 系统名
	ClientID     string    `gorm:"index" json:"client_id"`
	ClientSecret string    `json:"client_secret"`
	RedirectURI  string    `json:"redirect_uri"` // 重定向路径
	Status       bool      `json:"status"`       // 系统状态 0未启用1使用中
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type PersonalAccessTokens struct {
	ID            int64     `gorm:"primary_key" json:"-"`
	TokenableType string    `gorm:"index:personal_access_tokens_tokenable_type_tokenable_id_index" json:"tokenable_type"`
	TokenableID   int64     `gorm:"index:personal_access_tokens_tokenable_type_tokenable_id_index" json:"tokenable_id"`
	Name          string    `json:"name"`
	Token         string    `gorm:"unique" json:"token"`
	Abilities     string    `json:"abilities"`
	LastUsedAt    time.Time `json:"last_used_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type AdminMenu struct {
	ID        int       `gorm:"primary_key" json:"-"`
	ParentID  int       `gorm:"index" json:"parent_id"`
	Order     int       `json:"order"`
	Title     string    `json:"title"`
	Icon      string    `json:"icon"`
	URI       string    `gorm:"index" json:"uri"`
	Show      int8      `json:"show"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AdminPermissionMenu struct {
	PermissionID int       `gorm:"primary_key" json:"permission_id"`
	MenuID       int       `gorm:"primary_key" json:"menu_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type AdminRoleButton struct {
	RoleID    int       `gorm:"primary_key" json:"role_id"`
	ButtonID  int       `gorm:"primary_key" json:"button_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AdminRoleMenu struct {
	RoleID    int       `gorm:"primary_key" json:"role_id"`
	MenuID    int       `gorm:"primary_key" json:"menu_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AdminRolePermissions struct {
	PermissionID int64     `gorm:"primary_key" json:"permission_id"`
	RoleID       int64     `gorm:"primary_key" json:"role_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type AdminRoute struct {
	ID         int       `gorm:"primary_key" json:"-"`
	Name       string    `json:"name"`
	Slug       string    `gorm:"unique" json:"slug"`
	HTTPMethod string    `gorm:"index:idx_http_path_http_method" json:"http_method"`
	HTTPPath   string    `gorm:"index:idx_http_path_http_method" json:"http_path"`
	Order      int       `json:"order"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type AdminUserAssign struct {
	UserID       int64     `gorm:"primary_key" json:"user_id"`
	AssignUserID int64     `gorm:"primary_key" json:"assign_user_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	AssignStart  time.Time `gorm:"index:idx_assign_start_assign_end" json:"assign_start"`
	AssignEnd    time.Time `gorm:"index:idx_assign_start_assign_end" json:"assign_end"`
}

type OauthAuthCodes struct {
	ID        int       `gorm:"primary_key" json:"-"`
	Code      string    `gorm:"unique" json:"code"` // 授权码
	ClientID  string    `json:"client_id"`
	UserID    int       `json:"user_id"` // 用户id
	IP        string    `json:"ip"`      // 用户ip
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AdminMenuButton struct {
	ID        int       `gorm:"primary_key" json:"-"`
	MenuID    int       `gorm:"index" json:"menu_id"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	Icon      string    `json:"icon"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 限制登录用户记录表
type ThrottlesLogin struct {
	ID         int       `gorm:"primary_key" json:"-"`
	Username   string    `gorm:"unique_index:unique_username_ip" json:"username"` // 用户名
	IP         string    `gorm:"unique_index:unique_username_ip" json:"ip"`       // ip地址
	LoginTime  time.Time `json:"login_time"`                                      // 登录时间
	ExpireTime int       `json:"expire_time"`                                     // 过期时间(s)
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	IsLock     bool      `json:"is_lock"` // 是否封禁（1:限制，0:解禁）
}

type Users struct {
	ID              int64     `gorm:"primary_key" json:"-"`
	Name            string    `json:"name"`
	Email           string    `gorm:"unique" json:"email"`
	EmailVerifiedAt time.Time `json:"email_verified_at"`
	Password        string    `json:"password"`
	RememberToken   string    `json:"remember_token"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type PasswordResets struct {
	Email     string    `gorm:"index" json:"email"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}

type AdminPermissions struct {
	ID        int64     `gorm:"primary_key" json:"-"`
	Name      string    `gorm:"unique_index:admin_permissions_name_guard_name_unique" json:"name"`
	GuardName string    `gorm:"unique_index:admin_permissions_name_guard_name_unique" json:"guard_name"`
	Slug      string    `json:"slug"`
	Order     int       `json:"order"`
	ParentID  int       `gorm:"index" json:"parent_id"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AdminPermissionRoute struct {
	PermissionID int       `gorm:"primary_key" json:"permission_id"`
	RouteID      int       `gorm:"primary_key" json:"route_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
