package api

const (
	UnityAPIPrefix      = "/api"
	UnityAPITypesPrefix = UnityAPIPrefix + "/types"
)

type UnityAPITypesInstances string

const (
	UnityAPIBasicSystemInfoInstances UnityAPITypesInstances = UnityAPITypesPrefix + "/basicSystemInfo/instances?compact=true"
	UnityAPISystemCapacityInstances  UnityAPITypesInstances = UnityAPITypesPrefix + "/systemCapacity/instances?compact=true"
	UnityAPISystemInstances          UnityAPITypesInstances = UnityAPITypesPrefix + "/system/instances?compact=true"
	UnityAPILunInstances             UnityAPITypesInstances = UnityAPITypesPrefix + "/lun/instances?compact=true"
	UnityAPIPoolInstances            UnityAPITypesInstances = UnityAPITypesPrefix + "/pool/instances?compact=true"
	UnityAPIStorageResourceInstances UnityAPITypesInstances = UnityAPITypesPrefix + "/storageResource/instances?compact=true"
)

/*
	Request

Endpoint = Unity Console URL (ex: https://127.0.0.1:443)
Params = (ex: Key->"fields", Value->"id,name")
Logined = if you want login, set true
Auth = when you want login, that will use this auth(encoded base64)
*/
type Request struct {
	Endpoint string
	Params   map[string]string
	Logined  bool
	Auth     string
}
