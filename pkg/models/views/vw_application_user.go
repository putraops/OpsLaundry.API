package views

import (
	"opslaundry/pkg/models"
	"strings"
)

type ApplicationUser struct {
	models.ApplicationUser
	OrganizationName string `json:"organization_name"`
	TenantName       string `json:"tenant_name"`
	Fullname         string `json:"fullname"`
	InitialName      string `json:"initial_name"`
	HasRole          bool   `json:"has_role"`

	RecordCreated   string `json:"record_created"`
	RecordUpdated   string `json:"record_updated"`
	RecordSubmitted string `json:"record_submitted"`
	RecordApproved  string `json:"record_approved"`
}

func (ApplicationUser) TableName() string {
	return "vw_application_user"
}

func (ApplicationUser) ViewModel() string {
	var sql strings.Builder
	sql.WriteString("SELECT")
	sql.WriteString("  r.id,")
	sql.WriteString("  r.is_active,")
	sql.WriteString("  r.is_locked,")
	sql.WriteString("  r.is_default,")
	sql.WriteString("  r.owner_id,")
	sql.WriteString("  r.created_at,")
	sql.WriteString("  r.created_by,")
	sql.WriteString("  r.updated_at,")
	sql.WriteString("  r.updated_by,")
	sql.WriteString("  r.approved_at,")
	sql.WriteString("  r.approved_by,")
	sql.WriteString("  r.submitted_at,")
	sql.WriteString("  r.submitted_by,")
	// sql.WriteString("  r.entity_id,")
	sql.WriteString("  r.organization_id,")
	sql.WriteString("  o1.name AS organization_name,")
	sql.WriteString("  r.tenant_id,")
	sql.WriteString("  t.name AS tenant_name,")
	sql.WriteString("  r.fb_uid,")
	sql.WriteString("  r.username,")
	sql.WriteString("  r.password,")
	sql.WriteString("  r.first_name,")
	sql.WriteString("  r.last_name,")
	sql.WriteString("  CASE WHEN r.last_name IS NULL OR r.last_name = '' THEN r.first_name ELSE CONCAT(r.first_name, ' ', r.last_name) END AS fullname,")
	sql.WriteString("  r.title,")
	sql.WriteString("  r.address,")
	sql.WriteString("  r.phone,")
	sql.WriteString("  r.email,")
	sql.WriteString("  r.total_point,")
	sql.WriteString("  r.is_email_verified,")
	sql.WriteString("  r.is_phone_verified,")
	sql.WriteString("  false AS has_role,")
	sql.WriteString("  r.is_system_admin,")
	sql.WriteString("  r.is_admin,")
	sql.WriteString("  r.user_type,")
	sql.WriteString("  r.gender,")
	sql.WriteString("  r.filepath,")
	sql.WriteString("  r.filepath_thumbnail,")
	sql.WriteString("  r.filename,")
	sql.WriteString("  r.extension,")
	sql.WriteString("  r.size,")
	sql.WriteString("  r.description,")
	sql.WriteString("  CONCAT(UPPER(LEFT(r.first_name, 1)), '', UPPER(LEFT(r.last_name, 1))) AS initial_name,")
	sql.WriteString("  CASE WHEN u1.last_name IS NULL OR u1.last_name = '' THEN u1.first_name ELSE concat(u1.first_name, ' ', u1.last_name) END AS record_created,")
	sql.WriteString("  CASE WHEN u2.last_name IS NULL OR u2.last_name = '' THEN u2.first_name ELSE concat(u2.first_name, ' ', u2.last_name) END AS record_updated,")
	sql.WriteString("  CASE WHEN u3.last_name IS NULL OR u3.last_name = '' THEN u3.first_name ELSE concat(u3.first_name, ' ', u3.last_name) END AS record_submitted,")
	sql.WriteString("  CASE WHEN u4.last_name IS NULL OR u4.last_name = '' THEN u4.first_name ELSE concat(u4.first_name, ' ', u4.last_name) END AS record_approved ")
	sql.WriteString("FROM application_user r ")
	sql.WriteString("LEFT JOIN organization o1 ON o1.id = r.organization_id ")
	sql.WriteString("LEFT JOIN tenant t ON t.id = r.tenant_id ")
	sql.WriteString("LEFT JOIN application_user u1 ON u1.id = r.created_by ")
	sql.WriteString("LEFT JOIN application_user u2 ON u2.id = r.updated_by ")
	sql.WriteString("LEFT JOIN application_user u3 ON u3.id = r.submitted_by ")
	sql.WriteString("LEFT JOIN application_user u4 ON u4.id = r.approved_by ")
	return sql.String()
}
func (ApplicationUser) Migration() map[string]string {
	var view = ApplicationUser{}
	var m = make(map[string]string)
	m["view_name"] = view.TableName()
	m["query"] = view.ViewModel()
	return m
}
