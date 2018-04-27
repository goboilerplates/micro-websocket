package api

import (
	"github.com/gin-gonic/gin"
	"github.com/goboilerplates/core"
	"github.com/gonitor/gonitor-websocket/env"
	"github.com/gorilla/websocket"
)

// GetSamplesAPI is the interface for GetSamplesAPI.
type GetSamplesAPI interface {
	HanldeWebSocket(context *gin.Context)
	HandleMessage(conn *websocket.Conn, messageType int, messageBytes []byte)
}

// GetSamplesAPIImpl is the implementation of GetSamplesAPI interface.
type GetSamplesAPIImpl struct {
	Interactor core.GetSamplesInteractor
}

// HanldeWebSocket handles websocket connection.
func (api GetSamplesAPIImpl) HanldeWebSocket(context *gin.Context) {
	wsupgrader := api.GetUgrader()
	conn, err := wsupgrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		return
	}
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			break
		}
		go api.HandleMessage(conn, messageType, message)
	}
}

// GetUgrader .
func (api GetSamplesAPIImpl) GetUgrader() websocket.Upgrader {
	return websocket.Upgrader{
		EnableCompression: env.EnableCompression,
		ReadBufferSize:    env.ReadBufferSize,
		WriteBufferSize:   env.WriteBufferSize,
	}
}

// HandleMessage handles message from client.
func (api GetSamplesAPIImpl) HandleMessage(conn *websocket.Conn, messageType int, messageBytes []byte) {
	var response []byte
	message, messageErr := ConvertBytesToMap(messageBytes)
	if messageErr != nil {
		conn.WriteMessage(messageType, []byte(messageErr.Error()))
		return
	}

	switch message["name"] {
	// CPU
	case "All":
		samples, err := api.Interactor.All()
		response = ConvertInterfaceToJSONBytes(samples, err)
	case "AllByName":
		keyword := message["keyword"].(string)
		samples, err := api.Interactor.AllByName(keyword)
		response = ConvertInterfaceToJSONBytes(samples, err)
	}
	conn.WriteMessage(messageType, response)
}
