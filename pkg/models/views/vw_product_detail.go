package views

import (
	"opslaundry/pkg/models"
	"strings"

	"gorm.io/gorm"
)

type ProductDetail struct {
	models.ProductDetail
	ProductName     string `json:"product_name"`
	ServiceTypeName string `json:"service_type_name"`
	UomName         string `json:"uom_name"`
	UomCode         string `json:"uom_code"`
	UomSymbol       string `json:"uom_symbol"`

	OrganizationName string `json:"organization_name"`
	RecordCreated    string `json:"record_created"`
	RecordUpdated    string `json:"record_updated"`
	RecordSubmitted  string `json:"record_submitted"`
	RecordApproved   string `json:"record_approved"`
}

func (ProductDetail) TableName() string {
	return "vw_product_detail"
}

func (ProductDetail) ViewQuery() *gorm.Statement {
	var DB *gorm.DB
	stmt := DB.Find(&ProductDetail{}).Statement
	// stmt.SQL.String() //=> SELECT * FROM `product_detail` WHERE `id` = $1 ORDER BY `id`
	// stmt.Vars
	return stmt
}

func (ProductDetail) ViewModel() string {
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
	sql.WriteString("  r.product_id,")
	sql.WriteString("  p.name AS product_name,")
	sql.WriteString("  r.item_name,")
	sql.WriteString("  r.service_type_id,")
	sql.WriteString("  st.name AS service_type_name,")
	sql.WriteString("  r.uom_id,")
	sql.WriteString("  uom1.uom_name,")
	sql.WriteString("  uom1.uom_code,")
	sql.WriteString("  uom1.uom_symbol,")
	sql.WriteString("  CASE WHEN u1.last_name IS NULL OR u1.last_name = '' THEN u1.first_name ELSE concat(u1.first_name, ' ', u1.last_name) END AS record_created,")
	sql.WriteString("  CASE WHEN u2.last_name IS NULL OR u2.last_name = '' THEN u2.first_name ELSE concat(u2.first_name, ' ', u2.last_name) END AS record_updated,")
	sql.WriteString("  CASE WHEN u3.last_name IS NULL OR u3.last_name = '' THEN u3.first_name ELSE concat(u3.first_name, ' ', u3.last_name) END AS record_submitted,")
	sql.WriteString("  CASE WHEN u4.last_name IS NULL OR u4.last_name = '' THEN u4.first_name ELSE concat(u4.first_name, ' ', u4.last_name) END AS record_approved ")
	sql.WriteString("FROM product_detail r ")
	sql.WriteString("LEFT JOIN organization o1 ON o1.id = r.organization_id ")
	sql.WriteString("LEFT JOIN product p ON p.id = r.product_id ")
	sql.WriteString("LEFT JOIN service_type st ON st.id = r.service_type_id ")
	sql.WriteString("LEFT JOIN uom uom1 ON uom1.id = r.uom_id ")
	sql.WriteString("LEFT JOIN application_user u1 ON u1.id = r.created_by ")
	sql.WriteString("LEFT JOIN application_user u2 ON u2.id = r.updated_by ")
	sql.WriteString("LEFT JOIN application_user u3 ON u3.id = r.submitted_by ")
	sql.WriteString("LEFT JOIN application_user u4 ON u4.id = r.approved_by ")
	return sql.String()
}
func (ProductDetail) Migration() map[string]string {
	var view = ProductDetail{}
	var m = make(map[string]string)
	m["view_name"] = view.TableName()
	m["query"] = view.ViewModel()
	return m
}
