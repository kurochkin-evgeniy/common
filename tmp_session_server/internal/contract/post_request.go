package tmp_session_contract

type CreationRequest struct {
	AuxKey      string `json:"aux_key"`
	SessionData string `json:"session_data"`
}

type CreationResponse struct {
	UserCode string `json:"user_code"`
}
