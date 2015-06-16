package comprise

import (
	"errors"
	"log"
	"math/rand"
	"strconv"
	"sync"
)

func NewServer() *Server {
	server := new(Server)
	clients := make(map[string]*Client)
	clientsMutex := new(sync.RWMutex)

	server.Clients = clients
	server.clientsMutex = clientsMutex

	return server
}

type Server struct {
	Clients      map[string]*Client
	clientsMutex *sync.RWMutex
}

func (server *Server) RegisterClient(name *string, reply *Client) error {
	reply.Name = *name

	id, err := server.nextClientID()
	if err != nil {
		return err
	}
	reply.ID = id

	server.addClient(reply)

	return nil
}

func (server *Server) nextClientID() (string, error) {
	var clientID string

	// TODO: Make it sleep between tries and have an upper limit when we wont
	// try anymore and throw an error instead.
	for {
		id, err := randomId()
		if err != nil {
			log.Println("can not get a random id: ", err)
			continue
		}

		server.clientsMutex.RLock()
		_, exists := server.Clients[id]
		server.clientsMutex.RUnlock()

		if !exists {
			clientID = id
			break
		}
	}

	return clientID, nil
}

func (server *Server) addClient(client *Client) error {
	server.clientsMutex.RLock()
	_, exists := server.Clients[client.ID]
	server.clientsMutex.RUnlock()

	if exists {
		return errors.New("the client is already registered with the server")
	}

	server.clientsMutex.Lock()
	server.Clients[client.ID] = client
	server.clientsMutex.Unlock()

	return nil
}

func randomId() (string, error) {
	number := rand.Uint32()
	id := formatNumberToID(number)

	return id, nil
}

func formatNumberToID(number uint32) string {
	var id string

	for i := 0; i < 10; i++ {
		// Get last piece of the number
		piece := number % 10

		// Convert piece to string and add a '-' after every third round
		id += strconv.FormatUint(uint64(piece), 10)
		if i%3 == 2 {
			id += "-"
		}

		// Cut of piece from the number
		number = number / 10
	}

	return id
}
