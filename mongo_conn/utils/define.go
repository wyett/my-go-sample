//
//Author: your name
//Date: 2021-10-25 18:01:59
//LastEditTime: 2021-11-02 15:38:35
//LastEditors: Please set LastEditors
//Description: In User Settings Edit
//FilePath: \my-go-sample\mongo_conn\utils\define.go
//
package utils

const (
	// mongo connect mode
	VarMongoConnectModePrimary            = "primary"
	VarMongoConnectModeSecondaryPreferred = "secondaryPreferred"
	VarMongoConnectModeSecondary          = "secondary"
	VarMongoConnectModeNearset            = "nearest"
	VarMongoConnectModeStandalone         = "standalone"
)

type Pair struct {
	First  interface{}
	Second interface{}
}
