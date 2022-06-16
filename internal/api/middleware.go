package api

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/okutsen/PasswordManager/internal/log"
)

// ContextSetter reads header, creates RequestContext and adds it to r.Context
func ContextSetter(logger log.Logger, next http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		corIDStr := r.Header.Get(CorrelationIDHPN)
		corID, err := uuid.Parse(corIDStr)
		if err != nil {
			logger.Warnf("Invalid corID <%s>: %s", corIDStr, err)
			corID = uuid.New()
			logger.Debugf("Setting new corID: %s", corID.String())
		}
		ctx := context.WithValue(r.Context(), RequestContextName, &RequestContext{
			corID: corID,
			ps:    ps,
		})
		r = r.WithContext(ctx)
		next(w, r)
	}
}
