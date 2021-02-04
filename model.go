package cweed

import (
	"encoding/json"
	"math/rand"
)

// UploadResult contains upload result after put file to SeaweedFS
// Raw response: {"name":"go1.8.3.linux-amd64.tar.gz","size":82565628,"error":""}
type UploadResult struct {
	Name  string `json:"name,omitempty"`
	Size  int64  `json:"size,omitempty"`
	Error string `json:"error,omitempty"`
}

// AssignResult contains assign result.
// Raw response: {"fid":"1,0a1653fd0f","url":"localhost:8899","publicUrl":"localhost:8899","count":1,"error":""}
type AssignResult struct {
	FileID    string `json:"fid,omitempty"`
	URL       string `json:"url,omitempty"`
	PublicURL string `json:"publicUrl,omitempty"`
	Count     uint64 `json:"count,omitempty"`
	Error     string `json:"error,omitempty"`
}

// SubmitResult result of submit operation.
type SubmitResult struct {
	FileName string `json:"fileName,omitempty"`
	FileURL  string `json:"fileUrl,omitempty"`
	FileID   string `json:"fid,omitempty"`
	Size     int64  `json:"size,omitempty"`
	Error    string `json:"error,omitempty"`
}

// ClusterStatus result of getting status of cluster
type ClusterStatus struct {
	IsLeader bool
	Leader   string
	Peers    []string
}

// SystemStatus result of getting status of system
type SystemStatus struct {
	Topology Topology
	Version  string
	Error    string
}

// Topology result of topology stats request
type Topology struct {
	DataCenters []*DataCenter
	Free        int
	Max         int
	Layouts     []*Layout
}

// DataCenter stats of a datacenter
type DataCenter struct {
	Free  int
	Max   int
	Racks []*Rack
}

// Rack stats of racks
type Rack struct {
	DataNodes []*DataNode
	Free      int
	Max       int
}

// DataNode stats of data node
type DataNode struct {
	Free      int
	Max       int
	PublicURL string `json:"PublicUrl"`
	URL       string `json:"Url"`
	Volumes   int
}

// Layout of replication/collection stats. According to https://github.com/chrislusf/seaweedfs/wiki/Master-Server-API
type Layout struct {
	Replication string
	Writables   []uint64
}

// VolumeLocation location of volume responsed from master API. According to https://github.com/chrislusf/seaweedfs/wiki/Master-Server-API
type VolumeLocation struct {
	URL       string `json:"url,omitempty"`
	PublicURL string `json:"publicUrl,omitempty"`
}

// VolumeLocations returned VolumeLocations (volumes)
type VolumeLocations []*VolumeLocation

// Head get first location in list
func (c VolumeLocations) Head() *VolumeLocation {
	if len(c) == 0 {
		return nil
	}

	return c[0]
}

// RandomPickForRead random pick a location for further read request
func (c VolumeLocations) RandomPickForRead() *VolumeLocation {
	if len(c) == 0 {
		return nil
	}

	return c[rand.Intn(len(c))]
}

// LookupResult the result of looking up volume. According to https://github.com/chrislusf/seaweedfs/wiki/Master-Server-API
type LookupResult struct {
	VolumeLocations VolumeLocations `json:"locations,omitempty"`
	Error           string          `json:"error,omitempty"`
}


// ChunkInfo chunk information. According to https://github.com/chrislusf/seaweedfs/wiki/Large-File-Handling.
type ChunkInfo struct {
	Fid    string `json:"fid"`
	Offset int64  `json:"offset"`
	Size   int64  `json:"size"`
}

// ChunkManifest chunk manifest. According to https://github.com/chrislusf/seaweedfs/wiki/Large-File-Handling.
type ChunkManifest struct {
	Name   string       `json:"name,omitempty"`
	Mime   string       `json:"mime,omitempty"`
	Size   int64        `json:"size,omitempty"`
	Chunks []*ChunkInfo `json:"chunks,omitempty"`
}

// Marshal marshal whole chunk manifest
func (c *ChunkManifest) Marshal() ([]byte, error) {
	return json.Marshal(c)
}
