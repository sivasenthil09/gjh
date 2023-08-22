// this interface will provide the requried methods
package interfaces

import "rest-api/models"

type ITransaction interface{
	CreateTransaction(transaction *models.Transaction)(*models.DBResponse, error)
	

}