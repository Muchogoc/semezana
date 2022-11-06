package ports

import (
	"log"
	"net/http"

	"github.com/Muchogoc/semezana/app"
	"github.com/Muchogoc/semezana/dto"
	"github.com/go-chi/render"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

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

	token, err := h.service.GenerateAccessToken(r.Context(), creds)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	data := map[string]interface{}{}
	mapstructure.Decode(token, &data)

	response := dto.APIResponseFormat{
		Data:    data,
		Message: "success",
		Status:  dto.Success,
	}

	render.Status(r, http.StatusOK)
	render.Respond(w, r, response)
}

func (h HttpServer) GetChannels(w http.ResponseWriter, r *http.Request) {
	channels, err := h.service.GetChannels(r.Context())
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	data := []map[string]interface{}{}
	mapstructure.Decode(channels, &data)

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

	channel, err := h.service.CreateChannel(r.Context(), newChannel)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	data := map[string]interface{}{}
	mapstructure.Decode(channel, &data)

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
	channel, err := h.service.GetChannelById(r.Context(), channelID)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	data := map[string]interface{}{}
	mapstructure.Decode(channel, &data)

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

	go session.Writer()
	go session.Reader()
	// go session.SubscriptionListener()
}

// Retrieve all users
func (h HttpServer) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetUsers(r.Context())
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	data := []map[string]interface{}{}
	mapstructure.Decode(users, &data)

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

	user, err := h.service.CreateUser(r.Context(), newUser)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	data := map[string]interface{}{}
	mapstructure.Decode(user, &data)

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
	user, err := h.service.GetUserById(r.Context(), userID)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	data := map[string]interface{}{}
	mapstructure.Decode(user, &data)

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

	membership, err := h.service.CreateMembership(r.Context(), newMembership)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	data := map[string]interface{}{}
	mapstructure.Decode(membership, &data)

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
