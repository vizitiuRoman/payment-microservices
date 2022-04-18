package use_cases

type UseCase[In, Out any] interface {
	Execute(input *In) (output *Out, err error)
}

type UseCaseW struct {
}

// create use case with input and output
func NewUseCase(input *struct{}, output *struct{}) UseCase[struct{}, struct{}] {
	return &UseCaseW{}
}

func (u *UseCaseW) Execute(input *struct{}) (output *struct{}, err error) {
	return &struct{}{}, nil
}
