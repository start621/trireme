package rpcWrapper

import (
	"time"

	"github.com/aporeto-inc/trireme/enforcer"
	"github.com/aporeto-inc/trireme/policy"
	"github.com/aporeto-inc/trireme/utils/tokens"
)

//Request exported
type Request struct {
	MethodIdentifier int
	Payload          interface{}
}

//exported
const (
	SUCCESS = 0
)

//Response exported
type Response struct {
	MethodIdentifier int
	Status           int
	Payload          interface{}
}

//InitRequestPayload exported
type InitRequestPayload struct {
	FqConfig   enforcer.FilterQueue
	MutualAuth bool
	Validity   time.Duration
	SecretType tokens.SecretsType
	ContextID  string
	CAPEM      []byte
	PublicPEM  []byte
	PrivatePEM []byte
}

//InitSupervisorPayload exported
type InitSupervisorPayload struct {
	NetworkQueues     string
	ApplicationQueues string
	TargetNetworks    []string
}

// EnforcePayload exported
type EnforcePayload struct {
	ContextID string
	PuPolicy  *policy.PUPolicy
}

//SuperviseRequestPayload exported
type SuperviseRequestPayload struct {
	ContextID string
	PuPolicy  *policy.PUPolicy
}

//UnEnforcePayload exported
type UnEnforcePayload struct {
	ContextID string
}

//InitResponsePayload exported
type InitResponsePayload struct {
	Status int
}

//EnforceResponsePayload exported
type EnforceResponsePayload struct {
	Status int
}

//SuperviseResponsePayload exported
type SuperviseResponsePayload struct {
	Status int
}

//UnEnforceResponsePayload exported
type UnEnforceResponsePayload struct {
	Status int
}

//RPCServer exported
type RPCServer interface {
	ProcessMessage(req *Request, resp *Response) error
}

//exported
const (
	InitEnforcer   = 0
	InitSupervisor = 1
	Enforce        = 2
	Supervise      = 3
)
