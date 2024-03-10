package rest

type MessageCode struct {
	Code    string
	Message string
}

var (
	SUCCESS                                      = MessageCode{"200", "general.success"}
	MESSAGE_CODE_FAIL                            = MessageCode{"500", "general.fail"}
	BAD_REQUEST                                  = MessageCode{"400", "general.badRequest"}
	NOT_FOUND                                    = MessageCode{"404", "general.notFound"}
	UNAUTHORIZED                                 = MessageCode{"401", "general.unauthorized"}
	ENTITY_NOT_FOUND                             = MessageCode{"404", "general.entityNotFound"}
	ENTITY_ALREADY_EXISTS                        = MessageCode{"409", "general.entityAlreadyExists"}
	FIELD_NOT_FOUND                              = MessageCode{"1001", "general.fieldNotFound"}
	BAD_ALGORITHM                                = MessageCode{"1002", "general.badAlgorithm"}
	NEW_PASSWORD_CANNOT_SAME_LAST_THREE_PASSWORD = MessageCode{"1003", "general.newPasswordCannotSameLastThreePassword"}
	PRODUCT_QUANTITY_ABSENT                      = MessageCode{"1004", "general.productQuantityInsufficient"}
	DEFAULT_CANNOT_DELETE                        = MessageCode{"1005", "general.defaultCannotDelete"}
	FIELD_UNIQUE                                 = MessageCode{"1006", "general.uniqueField"}
	NEED_VERIFICATION                            = MessageCode{"1007", "general.needVerification"}
	OPTION_SET_REQUIRED                          = MessageCode{"1008", "general.optionSetRequired"}
	WRONG_CREDS                                  = MessageCode{"1009", "general.wrongCreds"}
)
