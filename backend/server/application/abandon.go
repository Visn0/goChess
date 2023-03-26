package application

type AbandonOutput struct {
	Action string `json:"action"`
}

type AbandonAction struct {
}

func NewAbandonAction() *AbandonAction {
	return &AbandonAction{}
}

func (uc *AbandonAction) Invoke() *AbandonOutput {
	output := &AbandonOutput{
		Action: "abandon",
	}

	return output
}
