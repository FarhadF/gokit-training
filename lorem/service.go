package lorem

import (
	"context"
	gl "github.com/drhodes/golorem"
	"errors"
)

type Service interface {
	Lorem(ctx context.Context, requestType string, min int, max int)(string, error)
}

//implementation with empty struct (stateless)
type loremService struct {}

//constructor - we can later add initialization if needed
func NewService() Service {
	return loremService{}
}

//implementation
func (loremService) Lorem (ctx context.Context, requestType string, min int, max int) (string, error) {
	switch requestType {
	case "word":
		return gl.Word(min, max), nil
	case "sentence":
		return gl.Sentence(min, max), nil
	case "paragraph":
		return gl.Paragraph(min, max), nil
	default:
		return "", errors.New("request type should be word, sentence or paragraph")
	}
}





