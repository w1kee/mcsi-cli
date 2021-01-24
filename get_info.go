package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Status struct {
	Online   bool       `json:"online,omitempty"`
	Ip       string     `json:"ip,omitempty"`
	Port     int        `json:"port,omitempty"`
	Debug    DebugInfo  `json:"debug,omitempty"`
	Motd     MOTD       `json:"motd,omitempty"`
	Players  Players    `json:"players,omitempty"`
	Version  string     `json:"version,omitempty"`
	Protocol int        `json:"protocol,omitempty"`
	Hostname string     `json:"hostname,omitempty"`
	Software string     `json:"software,omitempty"`
	Map      string     `json:"map,omitempty"`
	Plugins  PluginInfo `json:"plugins,omitempty"`
	Mods     ModInfo    `json:"mods,omitempty"`
	Info     InfoInfo   `json:"info,omitempty"`
}

type DebugInfo struct {
	Ping          bool `json:"ping,omitempty"`
	Query         bool `json:"query,omitempty"`
	Srv           bool `json:"srv,omitempty"`
	QueryMismatch bool `json:"querymismatch,omitempty"`
	IpInSrv       bool `json:"ipinsrv,omitempty"`
	CnameInSrv    bool `json:"cnameinsrv,omitempty"`
	AnimatedMOTD  bool `json:"animatedmotd,omitempty"`
	CacheTime     int  `json:"cachetime,omitempty"`
}

type MOTD struct {
	Raw   []string `json:"raw,omitempty"`
	Clean []string `json:"clean,omitempty"`
	Html  []string `json:"html,omitempty"`
}

type Players struct {
	Online int               `json:"online,omitempty"`
	Max    int               `json:"max,omitempty"`
	List   []string          `json:"list,omitempty"`
	Uuid   map[string]string `json:"uuid,omitempty"`
}

type PluginInfo struct {
	Names []string `json:"names,omitempty"`
	Raw   []string `json:"raw,omitempty"`
}

type ModInfo struct {
	Names []string `json:"names,omitempty"`
	Raw   []string `json:"raw,omitempty"`
}

type InfoInfo struct {
	Names []string `json:"names,omitempty"`
	Raw   []string `json:"raw,omitempty"`
	Html  []string `json:"html,omitempty"`
}

func getInfo(url string) (*Status, error) {
	result, err := http.Get(fmt.Sprintf("https://api.mcsrvstat.us/2/%s", url))

	if err != nil {
		return nil, err
	}

	var status Status
	if err := json.NewDecoder(result.Body).Decode(&status); err != nil {
		return nil, nil
	}

	return &status, nil
}
