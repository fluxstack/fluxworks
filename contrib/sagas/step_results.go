package sagas

type stepResults struct {
	updatedSagaData    SagaData
	commands           []Command
	updatedStepContext stepContext
	local              bool
	failure            error
}
