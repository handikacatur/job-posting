package model_error

import (
	customError "github.com/handikacatur/jobs-api/tools/custom_error"
)

var errorInternalServer = customError.New("something wrong in system. please wait for developer support")
