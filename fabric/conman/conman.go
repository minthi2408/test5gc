package conman

import (
	"etrib5gc/fabric/common"
	"etrib5gc/fabric/httpdp"
)

// abstraction of protocol-specific service clients
type RemoteConnection interface {
	Send(common.Request) (common.Response, error)
	Addr() common.AgentAddr
}

type ConnectionManager interface {
	common.InternalService
	// create a new connection or resuse an existing one
	// add security layer if needs
	// NOTE: it is just a connection preparation, no interaction is done now
	Connect(common.AgentAddr) (RemoteConnection, error)

	// drop an unresponsive connection
	Drop(common.AgentAddr)
}

type connectionManager struct {
	connections map[string]RemoteConnection
}

func NewConnectionManager() *connectionManager {
	ret := &connectionManager{}
	return ret
}

// create a connection to a remote agent
func (cm *connectionManager) Connect(addr common.AgentAddr) (RemoteConnection, error) {
	//TODO: this is very simplified implementation. Netvision should implement
	//added features such as connection reuse, connection sercurity
	id := addr.String()
	if conn, ok := cm.connections[id]; ok { //connection existed
		return conn, nil
	}
	conn := httpdp.NewHttpRemoteAgent(addr)
	cm.connections[id] = conn
	return conn, nil
}

func (cm *connectionManager) Drop(addr common.AgentAddr) {
	delete(cm.connections, addr.String())
}

func (cm *connectionManager) Start() error {
	return nil
}

func (cm *connectionManager) Terminate() {
}
