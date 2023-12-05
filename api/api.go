package api

import (
	"addons/addons"
	"net/http"
)

//	method: AddonS.IsBlackList, params: [aParty,bParty]

type AttrIsBlackList struct {
	AParty string
	BParty string
}
type AddonS int
type PingAttr int

// Ping AddonS for Gorilla
func (addon *AddonS) Ping(r *http.Request, params *PingAttr, reply *string) (err error) {
	*reply = "PONG.."
	return
}

// Ping for TraRPCServer
func (addon *AddonS) ZPing(params *PingAttr, reply *string) (err error) {
	*reply = "PONG.."
	return
}

// IsBlackList returns weather aParty or bParty is blacklisted
func (addon *AddonS) IsBlackList(r *http.Request, attr *AttrIsBlackList, reply *bool) (err error) {
	cOut, err := addons.IsBlackListed(attr.AParty, attr.BParty)
	if err != nil {
		return
	}
	*reply = cOut
	return
}
