package api

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/okutsen/PasswordManager/internal/log"
)

// ContextSetter reads header, creates RequestContext and adds it to r.Context
func ContextSetter(ctx *APIContext, next http.HandlerFunc) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		corIDStr := r.Header.Get(CorrelationIDName)
		corID := parseRequestID(corIDStr, ctx.logger)
		ctx := context.WithValue(r.Context(), RequestContextName, &RequestContext{
			corID: corID,
			ps:    ps,
		})
		r = r.WithContext(ctx)
		next(rw, r)
	}
}

// parseRequestID
func parseRequestID(idStr string, logger log.Logger) uuid.UUID {
	id, err := uuid.Parse(idStr)
	if err != nil {
		logger.Warnf("Invalid corID <%s>: %s", idStr, err)
		newID := uuid.New()
		logger.Debugf("Setting new corID: %s", newID.String())
		return newID
	}
	return id
}
