package repository

import (
	"fmt"
	"log"
	"portservices/config"
	"portservices/model"
	"time"
)

//CreateDataClient ...
func CreateDataClient(data *model.DataClient, userID uint64) (bool, error) {
	now := time.Now()
	sqlInsert := `
		INSERT INTO client (
			company_code, 
			sales_org, 
			dist_channel, 
			division_client, 
			sap_customer_number, 
			customer_name, 
			address, 
			telephone, 
			mobile_phone, 
			fax, 
			email, 
			status, 
			attachment, 
			created_at, 
			created_by, 
			user_id
		) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16);
	`
	// UPDATE client, users
	// 	SET client.created_by=users.email
	// 	WHERE client.user_id=$17;
	_, sqlError := config.DB.Exec(sqlInsert,
		data.CompanyCode,
		data.SalesOrg,
		data.DistChannel,
		data.DivisionClient,
		data.CustomerNumber,
		data.CustomerName,
		data.Address,
		data.Telephone,
		data.MobilePhone,
		data.Fax,
		data.Email,
		data.Status,
		data.Attachment,
		now,
		userID,
		userID)
	if sqlError != nil {
		log.Println("SQL error on data client =>", sqlError)
		return false, sqlError
	}
	fmt.Printf("Success to create data client %v\n", *data)
	return true, sqlError
}

// GetDataClient ...
func GetDataClient(userID uint64) ([]model.DataClient, error) {
	sqlQuery := `
		SELECT 
			_id,
			company_code, 
			sales_org, 
			dist_channel, 
			division_client, 
			sap_customer_number, 
			customer_name, 
			address, 
			telephone, 
			mobile_phone, 
			fax, 
			email, 
			status, 
			attachment,
			created_at,
			created_by,
			modified_at,
			modified_by,
			delete_at,
			delete_by
		FROM client where delete_at IS NULL;
	`
	rows, sqlError := config.DB.Query(sqlQuery)
	if sqlError != nil {
		log.Println("SQL error on data client =>", sqlError)
	}
	defer rows.Close()
	var res []model.DataClient
	for rows.Next() {
		var client = model.DataClient{}
		sqlError = rows.Scan(
			&client.ID,
			&client.CompanyCode,
			&client.SalesOrg,
			&client.DistChannel,
			&client.DivisionClient,
			&client.CustomerNumber,
			&client.CustomerName,
			&client.Address,
			&client.Telephone,
			&client.MobilePhone,
			&client.Fax,
			&client.Email,
			&client.Status,
			&client.Attachment,
			&client.CreatedAt,
			&client.CreatedBy,
			&client.ModifiedAt,
			&client.ModifiedBy,
			&client.DeleteAt,
			&client.DeleteBy,
		)
		if sqlError != nil {
			log.Println("SQL error on data client =>", sqlError)
		}
		res = append(res, client)
	}
	sqlError = rows.Err()
	if sqlError != nil {
		log.Println("SQL error on data client =>", sqlError)
	}
	return res, sqlError
}
