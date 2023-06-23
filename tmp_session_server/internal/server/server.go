package tmp_session

import (
	"math/rand"
	"net/http"
	"strconv"
	"tmp_session_contract"

	"github.com/gin-gonic/gin"
)

type SessionServer struct {
	httpEngine *gin.Engine
	sessionMap map[string]string
}

func (server *SessionServer) Run() error {
	server.sessionMap = make(map[string]string)
	server.httpEngine = gin.Default()

	server.httpEngine.POST("/api/sessions/:id", server.createSessionHandler)
	server.httpEngine.GET("/api/sessions/:id", server.getSessionHandler)
	server.httpEngine.POST("/api/sessions/:id/retrive", server.retriveSessionHandler)

	return server.httpEngine.Run(":5008")
}

func makeComplexKey(id, auxKey, userCode string) string {
	return id + "::" + auxKey + "::" + userCode
}

func (server *SessionServer) createSessionHandler(c *gin.Context) {
	id := c.Params.ByName("id")

	var request tmp_session_contract.CreationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	var response tmp_session_contract.CreationResponse
	response.UserCode = strconv.Itoa(rand.Intn(10000))

	key := makeComplexKey(id, request.AuxKey, response.UserCode)
	server.sessionMap[key] = request.SessionData

	c.JSON(http.StatusOK, response)
}

func (server *SessionServer) getSessionHandler(c *gin.Context) {
	id := c.Params.ByName("id")

	auxKey, auxKeyExist := c.GetQuery("aux_key")
	if !auxKeyExist {
		c.Status(http.StatusBadRequest)
		return
	}

	userCode, userCodeExist := c.GetQuery("user_code")
	if !userCodeExist {
		c.Status(http.StatusBadRequest)
		return
	}

	key := makeComplexKey(id, auxKey, userCode)

	value, has := server.sessionMap[key]
	if !has {
		c.Status(http.StatusNotFound)
		return
	}

	var response tmp_session_contract.GetResponse
	response.SessionData = value

	c.JSON(http.StatusOK, response)
}

func (server *SessionServer) retriveSessionHandler(c *gin.Context) {
	id := c.Params.ByName("id")

	auxKey, auxKeyExist := c.GetQuery("aux_key")
	if !auxKeyExist {
		c.Status(http.StatusBadRequest)
		return
	}

	userCode, userCodeExist := c.GetQuery("user_code")
	if !userCodeExist {
		c.Status(http.StatusBadRequest)
		return
	}

	key := makeComplexKey(id, auxKey, userCode)

	value, has := server.sessionMap[key]
	if !has {
		c.Status(http.StatusNotFound)
		return
	}

	var response tmp_session_contract.GetResponse
	response.SessionData = value

	delete(server.sessionMap, key)

	c.JSON(http.StatusOK, response)
}

func (server *SessionServer) Shutdown() error {
	return nil
}
