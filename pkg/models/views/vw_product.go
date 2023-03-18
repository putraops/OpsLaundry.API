package views

import (
	"opslaundry/pkg/models"
	"strings"
)

type Product struct {
	models.Product
	OrganizationName           string `json:"organization_name"`
	ProductCategoryName        string `json:"product_category_name"`
	ProductCategoryDescription string `json:"product_category_description"`
	RecordCreated              string `json:"record_created"`
	RecordUpdated              string `json:"record_updated"`
	RecordSubmitted            string `json:"record_submitted"`
	RecordApproved             string `json:"record_approved"`
}

func (Product) TableName() string {
	return "vw_product"
}

func (Product) ViewModel() string {
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
	sql.WriteString("  r.entity_id,")
	sql.WriteString("  r.organization_id,")
	sql.WriteString("  o1.name AS organization_name,")
	sql.WriteString("  r.name,")
	sql.WriteString("  r.description,")
	sql.WriteString("  r.product_category_id,")
	sql.WriteString("  pc.name AS product_category_name,")
	sql.WriteString("  pc.description AS product_category_description,")
	sql.WriteString("  CASE WHEN u1.last_name IS NULL OR u1.last_name = '' THEN u1.first_name ELSE concat(u1.first_name, ' ', u1.last_name) END AS record_created,")
	sql.WriteString("  CASE WHEN u2.last_name IS NULL OR u2.last_name = '' THEN u2.first_name ELSE concat(u2.first_name, ' ', u2.last_name) END AS record_updated,")
	sql.WriteString("  CASE WHEN u3.last_name IS NULL OR u3.last_name = '' THEN u3.first_name ELSE concat(u3.first_name, ' ', u3.last_name) END AS record_submitted,")
	sql.WriteString("  CASE WHEN u4.last_name IS NULL OR u4.last_name = '' THEN u4.first_name ELSE concat(u4.first_name, ' ', u4.last_name) END AS record_approved ")
	sql.WriteString("FROM product r ")
	sql.WriteString("LEFT JOIN organization o1 ON o1.id = r.organization_id ")
	sql.WriteString("LEFT JOIN product_category pc ON pc.id = r.product_category_id ")
	sql.WriteString("LEFT JOIN application_user u1 ON u1.id = r.created_by ")
	sql.WriteString("LEFT JOIN application_user u2 ON u2.id = r.updated_by ")
	sql.WriteString("LEFT JOIN application_user u3 ON u3.id = r.submitted_by ")
	sql.WriteString("LEFT JOIN application_user u4 ON u4.id = r.approved_by ")
	return sql.String()
}
func (Product) Migration() map[string]string {
	var view = Product{}
	var m = make(map[string]string)
	m["view_name"] = view.TableName()
	m["query"] = view.ViewModel()
	return m
}
