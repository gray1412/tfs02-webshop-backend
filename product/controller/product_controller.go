package controller

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"project-backend/database"
	"regexp"
	"time"

	"project-backend/product/request"
	response "project-backend/product/response"
	resultResponse "project-backend/util/response"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

var db = database.ConnectDB()

const (
	DEFAULT_PAGE_SIZE  = 10
	DEFAULT_PAGE_INDEX = 1
	DEFAULT_SORT_TITLE = "id"
	DEFAULT_SORT_BY    = "asc"
)

func convertStringToAlias(s string) string {
	s = strings.ToLower(s)
	s = strings.Trim(s," !@#$%^&*()")
	s = strings.ReplaceAll(s, " ","-")
	re := regexp.MustCompile("[^\\w- ]+")
	s = re.ReplaceAllString(s, "")
	s += fmt.Sprint(time.Now().UnixNano() / int64(time.Millisecond))
	return s
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	requestCreateProduct := request.Product{}
	err := json.NewDecoder(r.Body).Decode(&requestCreateProduct)
	if err != nil {
		resultResponse.RespondWithJSON(w, 400, 0, " Hi Bad request", nil)
		return
	}
	requestCreateProduct.Alias = convertStringToAlias(requestCreateProduct.Name)
	tx := db.Begin()
	err = tx.Omit("Variants", "Image").Create(&requestCreateProduct).Error
	if err != nil {
		tx.Rollback()
		resultResponse.RespondWithJSON(w, 500, 0, err.Error(), nil)
		return
	}
	//productId := requestCreateProduct.Id
	//luu variant
	// map[position_option][map[postion_op_value][id_op_value]]
	options := make(map[int64]map[int64]int64)
	for _, o := range requestCreateProduct.Options {
		options[o.Position] = make(map[int64]int64)
		for _, ov := range o.OptionValues {
			options[o.Position][ov.Position] = ov.Id
		}
	}
	// set id product, ip option cho variant
	productId := requestCreateProduct.Id
	for i, _ := range requestCreateProduct.Variants {
		requestCreateProduct.Variants[i].ProductId = productId
		requestCreateProduct.Variants[i].Option1 = options[1][requestCreateProduct.Variants[i].Option1]
		requestCreateProduct.Variants[i].Option2 = options[2][requestCreateProduct.Variants[i].Option2]
		requestCreateProduct.Variants[i].Option3 = options[3][requestCreateProduct.Variants[i].Option3]
	}
	err = tx.Create(&requestCreateProduct.Variants).Error
	if err != nil {
		tx.Rollback()
		resultResponse.RespondWithJSON(w, 500, 0, err.Error(), nil)
		return
	}
	// get id variant// map[positionVariants][variantId]
	variantIds := make(map[int64]int64)
	for _, v := range requestCreateProduct.Variants {
		variantIds[v.Position] = v.Id
	}
	// luu image
		//1. set id variant, product cho image
	for i, _ := range requestCreateProduct.Image {
		requestCreateProduct.Image[i].ProductId = productId
		requestCreateProduct.Image[i].VariantId = variantIds[requestCreateProduct.Image[i].VariantId]
	}
	err = tx.Create(&requestCreateProduct.Image).Error
	if err != nil {
		tx.Rollback()
		resultResponse.RespondWithJSON(w, 500, 0, err.Error(), nil)
		return
	}
	tx.Commit()
	resultResponse.RespondWithJSON(w, 201, 1, "", requestCreateProduct)
}
func GetProductByAlias(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	alias := param["alias"]
	// productId, err := strconv.Atoi(param["id"])
	// if err != nil {
	// 	resultResponse.RespondWithJSON(w, 400, 0, "Invalid param", "")
	// 	return
	// }
	
	var result response.ResponseProductByID
	sql := "SELECT * FROM products WHERE active = 2 AND  alias = ?"
	resultQuery := db.Raw(sql, alias).Scan(&result)
	if resultQuery.RowsAffected < 1 {
		resultResponse.RespondWithJSON(w, 404, 1, "Not found", "")
		return
	}

	// get các options, option_value
	var options []response.Option
	
	db.Where("product_id =?",result.Id).Preload("OptionValues").Find(&options)
	result.SetOptions(&options)

	// get variants
	var variants []response.Variant
	sql = "SELECT * FROM `variants` WHERE product_id = ?"
	db.Raw(sql, result.Id).Scan(&variants)
	result.SetVariants(&variants)

	//get image
	var images []response.Image
	sql = "SELECT * FROM product_images WHERE product_id = ?"
	db.Raw(sql, result.Id).Scan(&images)
	result.SetImages(&images)

	resultResponse.RespondWithJSON(w, 200, 1, "", result)
}
func SearchProduct(w http.ResponseWriter, r *http.Request) {
	//  input/: param:
	// pageindex: default: 1
	// name: product(like) // defaullt: ""
	// fillter:
	//brand
	//category // default: all
	// sort: price | date  // default: id
	// order: asc | desc	// default: asc
	//---------------------------------------------

	// lấy param
	nameProduct := r.URL.Query().Get("name") // nameProduct theo tên product
	brand := r.URL.Query().Get("brand")
	category := r.URL.Query().Get("category")
	sortTitle := r.URL.Query().Get("sort")
	orderBy := r.URL.Query().Get("order")
	rawPageIndex := r.URL.Query().Get("page")
	rawLimit := r.URL.Query().Get("limit")

	// check param: số, ký tự đặc biệt
	nameProduct = strings.Replace(nameProduct, "%", "\\%", -1)
	// nameProduct = strings.Replace(nameProduct, "-", "\\-", -1)
	nameProduct = strings.Replace(nameProduct, "'", "\\'", -1)
	nameProduct = strings.Replace(nameProduct, "_", "\\_", -1)

	pageIndex, err := strconv.Atoi(rawPageIndex)
	if err != nil {
		pageIndex = DEFAULT_PAGE_INDEX
	}

	// query sql
	sql := "active = 2 AND products.name LIKE '%" + nameProduct + "%' "
	if brand != "" {
		sql += "AND products.brand_name = '" + brand + "' "
	}
	if category != "" {
		sql += "AND products.category_id = (SELECT id FROM categories WHERE categories.name = '" + category + "') "
	}
	switch sortTitle {
	case "price":
	case "date":
		sortTitle = "created_at"
	default:
		sortTitle = DEFAULT_SORT_TITLE
	}
	if orderBy != "asc" && orderBy != "desc" {
		orderBy = DEFAULT_SORT_BY
	}
	sql += "ORDER BY " + sortTitle + " " + orderBy + " "
	limit, err := strconv.Atoi(rawLimit)
	if err != nil {
		limit = DEFAULT_PAGE_SIZE
	}
	if limit < 1 {
		limit = DEFAULT_PAGE_SIZE
	}
	if pageIndex < 1 {
		pageIndex = DEFAULT_PAGE_INDEX
	}
	// check page total
	var totalElement int64
	db.Raw("SELECT COUNT(*) FROM products WHERE " + sql).Scan(&totalElement)
	if totalElement == 0 {
		resultResponse.RespondWithJSON(w, 404, 1, "Not found", nil)
		return
	}
	totalPage := int(math.Ceil(float64(totalElement) / float64(limit)))

	if pageIndex < 1 {
		pageIndex = DEFAULT_PAGE_INDEX
	}
	if pageIndex > totalPage {
		pageIndex = totalPage
	}
	// sql thu được
	//	SELECT id, name, image_url price,original_price,quantity FROM products WHERE active = 2
	//	products.name LIKE '% %' AND
	//	products.brand_name = '' AND
	//	products.category_id = (SELECT id FROM categories WHERE categories.name = 'dog')
	//	ORDER BY produc ASC
	//	LIMIT 0,10
	sql = "SELECT * FROM products WHERE " + sql + " LIMIT " + strconv.Itoa((pageIndex-1)*limit) + " , " + strconv.Itoa(limit)
	// Lấy data từ dât
	var products []response.Product
	db.Raw(sql).Scan(&products)

	res := response.ResponseSearch{
		TotalPage:    totalPage,
		TotalElement: int(totalElement),
		PageIndex:    pageIndex,
		PageSize:     limit,
		Products:     products,
	}
	resultResponse.RespondWithJSON(w, 200, 1, "", res)

}
