package emailutils

import (
	"context"
	"time"

	"github.com/haroldcampbell/go_utils/utils"
	"golang.org/x/time/rate"
)

type Email struct {
	FromName         string
	FromAddress      string
	RecipientName    string
	RecipientAddress string
	Subject          string
	ContentPlainText string
	ContentHtml      string
}

func (e Email) GetFromName() string {
	if len(e.FromName) == 0 {
		return e.FromAddress
	}

	return e.FromName
}

func (e Email) GetRecipientName() string {
	if len(e.RecipientName) == 0 {
		return e.RecipientAddress
	}

	return e.RecipientName
}

type EmailHandlerWithAccessToken[T any] func(e Email, accessToken string) (*T, error)

type LimitedQueue[T any] struct {
	EmailHandler EmailHandlerWithAccessToken[T]
	Ratelimiter  *rate.Limiter
}

func NewLimitedQueue[T any](r time.Duration, bucketSize int, hanlder EmailHandlerWithAccessToken[T]) *LimitedQueue[T] {
	// bucketSize request every r seconds
	limiter := rate.NewLimiter(rate.Every(r*time.Second), bucketSize)

	q := &LimitedQueue[T]{
		EmailHandler: hanlder,
		Ratelimiter:  limiter,
	}
	return q
}

func (q *LimitedQueue[T]) EnqueueRequestWithAccessToken(e Email, accessToken string) (*T, error) {
	ctx := context.Background()
	err := q.Ratelimiter.Wait(ctx) // This is a blocking call. Honors the rate limit
	if err != nil {
		utils.Error("EnqueueRequestWithAccessToken", "Rate limiter failing. Email: %+v Err: %v", e, err)
		return nil, err
	}

	return q.EmailHandler(e, accessToken)
}
