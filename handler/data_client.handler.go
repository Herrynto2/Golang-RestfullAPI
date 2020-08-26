package handler

import (
	"fmt"
	"portservices/model"
	"portservices/repository"
	"portservices/utils"

	"github.com/gin-gonic/gin"
)

//CreateDataClient ...
func CreateDataClient(c *gin.Context) {
	userID := uint64(c.MustGet("userID").(float64))
	var data model.DataClient
	c.ShouldBindJSON(&data)
	// c.PostForm(&data)
	rows, err := repository.CreateDataClient(&data, userID)
	if rows {
		utils.ResponseSuccess(c, true, "Create data client success", gin.H{
			"company_code":        data.CompanyCode,
			"sales_org":           data.SalesOrg,
			"dist_channel":        data.DistChannel,
			"division_client":     data.DivisionClient,
			"sap_customer_number": data.CustomerNumber,
			"customer_name":       data.CustomerName,
			"address":             data.Address,
			"telephone":           data.Telephone,
			"mobile_phone":        data.MobilePhone,
			"fax":                 data.Fax,
			"email":               data.Email,
			"status":              data.Status,
			"attachment":          data.Attachment,
		})
	} else {
		utils.ResponseBadRequest(c, false, "Customer name has been available", gin.H{
			"data":  rows,
			"error": err.Error(),
		})
	}
}

//GetDataClient ...
func GetDataClient(c *gin.Context) {
	userID := uint64(c.MustGet("userID").(float64))
	var dataClient []model.DataClient
	dataClient, err := repository.GetDataClient(userID)
	fmt.Println(dataClient, err)
	if err == nil {
		utils.ResponseSuccess(c, true, "Get all master data client", gin.H{
			"data":  dataClient,
			"error": "",
			"total": len(dataClient),
		})
	} else {
		utils.ResponseBadRequest(c, false, "Get all master data client", gin.H{
			"data":  dataClient,
			"error": err.Error(),
		})
	}

}
