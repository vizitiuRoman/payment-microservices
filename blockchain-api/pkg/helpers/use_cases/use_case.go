package use_cases

type UseCase[In, Out any] interface {
	Execute(input *In) (output *Out, err error)
}
