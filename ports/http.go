package ports

import (
	"context"
	"log"
	"net/http"

	"github.com/Muchogoc/semezana/app"
	"github.com/Muchogoc/semezana/dto"
	"github.com/go-chi/render"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

const tracerName = "github.com/Muchogoc/semezana/ports"

type HttpServer struct {
	service      app.ChatService
	sessionStore *SessionStore
}

func NewHttpServer(service app.ChatService, sessionStore *SessionStore) HttpServer {
	return HttpServer{
		service:      service,
		sessionStore: sessionStore,
	}
}

func (h HttpServer) GetUserAccessToken(w http.ResponseWriter, r *http.Request) {
	creds := dto.NewToken{}
	if err := render.Decode(r, &creds); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	ctx := r.Context()
	ctx, span := otel.Tracer(tracerName).Start(ctx, "GetUserAccessToken()")
	defer span.End()

	token, err := h.service.GenerateAccessToken(ctx, creds)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	data := map[string]interface{}{}
	if err = mapstructure.Decode(token, &data); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	response := dto.APIResponseFormat{
		Data:    data,
		Message: "success",
		Status:  dto.Success,
	}

	render.Status(r, http.StatusOK)
	render.Respond(w, r, response)
}

func (h HttpServer) GetChannels(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx, span := otel.Tracer(tracerName).Start(ctx, "GetChannels()")
	defer span.End()

	channels, err := h.service.GetChannels(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	data := []map[string]interface{}{}
	if err = mapstructure.Decode(channels, &data); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
	}

	response := dto.APIResponseFormat{
		Data: map[string]interface{}{
			"channels": data,
		},
		Message: "success",
		Status:  dto.Success,
	}

	render.Status(r, http.StatusOK)
	render.Respond(w, r, response)
}

// Create a channel
func (h HttpServer) CreateChannel(w http.ResponseWriter, r *http.Request) {
	newChannel := dto.NewChannel{}
	if err := render.Decode(r, &newChannel); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	ctx := r.Context()
	ctx, span := otel.Tracer(tracerName).Start(ctx, "CreateChannel()")
	defer span.End()

	channel, err := h.service.CreateChannel(ctx, newChannel)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	data := map[string]interface{}{}
	if err = mapstructure.Decode(channel, &data); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
	}

	response := dto.APIResponseFormat{
		Data:    data,
		Message: "success",
		Status:  dto.Success,
	}

	render.Status(r, http.StatusCreated)
	render.Respond(w, r, response)
}

// Delete a channel
func (h HttpServer) DeleteChannelById(w http.ResponseWriter, r *http.Request, channelID string) {}

// Retrieve a channel
func (h HttpServer) GetChannelById(w http.ResponseWriter, r *http.Request, channelID string) {
	ctx := r.Context()
	ctx, span := otel.Tracer(tracerName).Start(ctx, "GetChannelById()")
	defer span.End()

	channel, err := h.service.GetChannelById(ctx, channelID)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	data := map[string]interface{}{}
	if err = mapstructure.Decode(channel, &data); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
	}

	response := dto.APIResponseFormat{
		Data:    data,
		Message: "success",
		Status:  dto.Success,
	}

	render.Status(r, http.StatusOK)
	render.Respond(w, r, response)
}

// Opens a WebSocket connection
func (h HttpServer) Websocket(w http.ResponseWriter, r *http.Request) {
	// Handles websocket requests from peers.
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// Allow connections from any Origin
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("ws: Invalid HTTP method", r.Method)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("ws: failed to Upgrade ", err)
		return
	}

	session, _ := h.sessionStore.NewWebsocketSession(conn, &h.service)

	ctx := context.Background()
	ctx, span := otel.Tracer(tracerName).Start(ctx, "Websocket")

	go session.Writer(ctx, span)
	go session.Reader(ctx, span)
	// go session.SubscriptionListener()
}

// Retrieve all users
func (h HttpServer) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx, span := otel.Tracer(tracerName).Start(ctx, "GetUsers()")
	defer span.End()

	users, err := h.service.GetUsers(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	data := []map[string]interface{}{}
	if err = mapstructure.Decode(users, &data); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	response := dto.APIResponseFormat{
		Data: map[string]interface{}{
			"users": data,
		},
		Message: "success",
		Status:  dto.Success,
	}

	render.Status(r, http.StatusOK)
	render.Respond(w, r, response)

}

// Create a user
func (h HttpServer) CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := dto.NewUser{}
	if err := render.Decode(r, &newUser); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	ctx := r.Context()
	ctx, span := otel.Tracer(tracerName).Start(ctx, "CreateUser()")
	defer span.End()

	user, err := h.service.CreateUser(ctx, newUser)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	data := map[string]interface{}{}
	if err = mapstructure.Decode(user, &data); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	response := dto.APIResponseFormat{
		Data:    data,
		Message: "success",
		Status:  dto.Success,
	}

	render.Status(r, http.StatusCreated)
	render.Respond(w, r, response)
}

// Delete a user
func (h HttpServer) DeleteUserById(w http.ResponseWriter, r *http.Request, userID string) {}

// Retrieve a user
func (h HttpServer) GetUserById(w http.ResponseWriter, r *http.Request, userID string) {
	ctx := r.Context()
	ctx, span := otel.Tracer(tracerName).Start(ctx, "GetUserById()")
	defer span.End()

	user, err := h.service.GetUserById(ctx, userID)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	data := map[string]interface{}{}
	if err = mapstructure.Decode(user, &data); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	response := dto.APIResponseFormat{
		Data:    data,
		Message: "success",
		Status:  dto.Success,
	}

	render.Status(r, http.StatusOK)
	render.Respond(w, r, response)
}

// Create a membership
func (h HttpServer) CreateMembership(w http.ResponseWriter, r *http.Request, channelID string) {
	newMembership := dto.NewMembership{}
	if err := render.Decode(r, &newMembership); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	ctx := r.Context()
	ctx, span := otel.Tracer(tracerName).Start(ctx, "CreateMembership()")
	defer span.End()

	membership, err := h.service.CreateMembership(ctx, newMembership)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	data := map[string]interface{}{}
	if err = mapstructure.Decode(membership, &data); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	response := dto.APIResponseFormat{
		Data:    data,
		Message: "success",
		Status:  dto.Success,
	}

	render.Status(r, http.StatusCreated)
	render.Respond(w, r, response)
}

// Remove a membership
func (h HttpServer) DeleteMembershipById(w http.ResponseWriter, r *http.Request, channelID string, membershipID string) {
}

// Retrieve a membership
func (h HttpServer) GetMembershipById(w http.ResponseWriter, r *http.Request, channelID string, membershipID string) {
}
