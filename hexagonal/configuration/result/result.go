package result

type ResultType int

const (
	Success             ResultType = 200
	Created             ResultType = 201
	NoContent           ResultType = 204
	BadRequest          ResultType = 400
	Conflict            ResultType = 409
	UnprocessableEntity ResultType = 422
	Internal            ResultType = 500
)

type Metadata struct {
	Page  int `json:"page,omitempty"`
	Limit int `json:"limit,omitempty"`
}

type Body struct {
	Message  string    `json:"message,omitempty"`
	Data     any       `json:"data,omitempty"`
	Metadata *Metadata `json:"metadata,omitempty"`
}

type Result struct {
	StatusCode ResultType
	Body       Body
	Errors     error
}

func NewSuccessResult(data any, page int, limitPages int) Result {
	return Result{
		StatusCode: Success,
		Body: Body{
			Data: data,
			Metadata: &Metadata{
				Page:  page,
				Limit: limitPages,
			},
		},
	}
}

func NewCreatedResult() Result {
	return Result{
		StatusCode: Created,
		Body:       Body{},
	}
}

func NewNoContentResult() Result {
	return Result{
		StatusCode: NoContent,
		Body:       Body{},
	}
}

func NewBadRequestResult(message string) Result {
	return Result{
		StatusCode: BadRequest,
		Body: Body{
			Message: message,
		},
	}
}

func NewValidationResult(message string) Result {
	return Result{
		StatusCode: UnprocessableEntity,
		Body: Body{
			Message: message,
		},
	}
}

func NewConflictResult() Result {
	return Result{
		StatusCode: Conflict,
		Body:       Body{},
	}
}

func NewInternalResult(err error) Result {
	return Result{
		StatusCode: Internal,
		Body: Body{
			Message: "Ops! Ocorreu um erro inesperado.",
		},
		Errors: err,
	}
}