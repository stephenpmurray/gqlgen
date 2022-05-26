// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type APITokens struct {
	ID         string  `json:"id"`
	GoogleMaps *string `json:"googleMaps"`
	Mapbox     *string `json:"mapbox"`
}

type BBox struct {
	// Northwesternmost point of a bounding box
	Nw *LatLngInput `json:"nw"`
	// Southeasternmost point of a bounding box
	Se *LatLngInput `json:"se"`
}

type Block struct {
	ID              string          `json:"id"`
	DisplayName     *string         `json:"displayName"`
	Values          *string         `json:"values"`
	SelectedVersion *string         `json:"selectedVersion"`
	Chart           *BlockChart     `json:"chart"`
	Status          *Status         `json:"status"`
	Connections     []*Connection   `json:"connections"`
	Resources       []*OktoResource `json:"resources"`
}

type BlockChart struct {
	// Chart digest, used for client caching. Cannot fetch by ID!
	ID string `json:"id"`
	// Unique chart name in Helm repo. name field in Chart.yaml
	Name string `json:"name"`
	// version field in Chart.yaml
	Version           string   `json:"version"`
	AvailableVersions []string `json:"availableVersions"`
	// keywords field in Chart.yaml
	Categories []string `json:"categories"`
	// description field in Chart.yaml
	Description *string `json:"description"`
	// full contents of Chart.yaml
	ChartYaml *string `json:"chartYaml"`
	// annotations.displayName in Chart.yaml
	DisplayName *string `json:"displayName"`
	// values.yaml
	ValuesYaml *string `json:"valuesYaml"`
	// annotations.vendor field in Chart.yaml
	Vendor *string `json:"vendor"`
	// icon field in Chart.yaml
	LogoURL *string `json:"logoUrl"`
}

type BlockConnectionInput struct {
	// Field name
	Name string `json:"name"`
	// Connection ID
	ConnID string `json:"connId"`
}

type BlockInput struct {
	// ID of a previously-deployed block to update.
	// Providing this parameter implies an update to an existing block,
	// whereas inputs with no ID will be treated as new blocks.
	ID *string `json:"id"`
	// Display name of the block
	DisplayName string `json:"displayName"`
	// Chart name in Helm repo
	BlockChartName string `json:"blockChartName"`
	// Chart semantic version number
	BlockChartVersion string `json:"blockChartVersion"`
	// Override values in YAML format
	Values string `json:"values"`
}

type ChartKey struct {
	// Unique chart name in Helm repo. name field in Chart.yaml
	Name string `json:"name"`
	// version field in Chart.yaml
	Version string `json:"version"`
}

type Connection struct {
	ID          string  `json:"id"`
	Category    string  `json:"category"`
	Description *string `json:"description"`
	// Parent entity providing the connection
	Entity *ConnectionEntity `json:"entity"`
	Kind   string            `json:"kind"`
	Name   string            `json:"name"`
	Source string            `json:"source"`
}

type ConnectionEntity struct {
	ID   string      `json:"id"`
	Type *EntityType `json:"type"`
}

type ConnectionInput struct {
	Category string `json:"category"`
	Kind     string `json:"kind"`
	Name     string `json:"name"`
	Source   string `json:"source"`
}

type CreateUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Org   string `json:"org"`
	Role  Role   `json:"role"`
}

type Dependencies struct {
	Services []*Service `json:"services"`
}

type DeployTargetInput struct {
	Master string `json:"master"`
	Host   string `json:"host"`
}

type Device struct {
	ID           string        `json:"id"`
	Tags         []string      `json:"tags"`
	Parents      []string      `json:"parents"`
	DisplayName  *string       `json:"displayName"`
	Img          *string       `json:"img"`
	IconType     *string       `json:"iconType"`
	Position     *LatLng       `json:"position"`
	Specs        *string       `json:"specs"`
	Status       *DeviceStatus `json:"status"`
	Progress     *Progress     `json:"progress"`
	Mac          *string       `json:"mac"`
	ProbeIP      *string       `json:"probeIp"`
	NztpWorkflow *string       `json:"nztpWorkflow"`
	IpmiAddress  *string       `json:"ipmiAddress"`
	IpmiUsername *string       `json:"ipmiUsername"`
	HostIP       *string       `json:"hostIp"`
	Links        []*DeviceLink `json:"links"`
	TelemetryURL *string       `json:"telemetryUrl"`
	Connections  []*Connection `json:"connections"`
	Site         *string       `json:"site"`
}

// Properties of a new device to be registered
type DeviceInput struct {
	// List of connections (e.g. video streams, APIs) available on the device
	Connections []*ConnectionInput `json:"connections"`
	// List of tags that describe the device
	Tags []string `json:"tags"`
	// List of IDs of parent devices
	Parents []string `json:"parents"`
	// Display name of device to be shown in the UI
	DisplayName *string `json:"displayName"`
	// Path to device image (currently unused)
	Img *string `json:"img"`
	// Map icon of device (currently unused)
	IconType *string `json:"iconType"`
	// Configuration file used to connect to Kubenetes Control Plane of a Cloudlet
	Kubeconfig *string `json:"kubeconfig"`
	// Geographical position of device
	Position *LatLngInput `json:"position"`
	// MAC address of device
	Mac *string `json:"mac"`
	// IP the probes should use to collect network metrics from the device
	ProbeIP *string `json:"probeIp"`
	// Near Zero-Touch Provisioning server address
	NztpServer *string `json:"nztpServer"`
	// Near Zero-Touch Provisioning workflow
	NztpWorkflow *string `json:"nztpWorkflow"`
	// Address of the IPMI management interface of the device
	IpmiAddress *string `json:"ipmiAddress"`
	// Username used to authenticate when issuing IPMI commands
	IpmiUsername *string `json:"ipmiUsername"`
	// Password used to authenticate when issuing IPMI commands
	IpmiPassword *string `json:"ipmiPassword"`
	// Site that device belongs to
	Site string `json:"site"`
}

type DeviceLink struct {
	ID      string    `json:"id"`
	Device  *Device   `json:"device"`
	Status  *string   `json:"status"`
	Metrics []*Metric `json:"metrics"`
}

type DevicesResult struct {
	Devices []*Device `json:"devices"`
	Hits    *int      `json:"hits"`
}

type EditConnectionInput struct {
	ID       string `json:"id"`
	Category string `json:"category"`
	Kind     string `json:"kind"`
	Name     string `json:"name"`
	Source   string `json:"source"`
}

type EditDeviceInput struct {
	// Display name of device to be shown in the UI
	DisplayName *string `json:"displayName"`
	// List of tags that describe the device
	Tags []string `json:"tags"`
	// Geographical position of device
	Position *LatLngInput `json:"position"`
	// Near Zero-Touch Provisioning workflow
	NztpWorkflow *string `json:"nztpWorkflow"`
	// Near Zero-Touch Provisioning flag to provision on next boot
	NztpProvision *bool `json:"nztpProvision"`
	// Address of the IPMI management interface of the device
	IpmiAddress *string `json:"ipmiAddress"`
	// Username used to authenticate when issuing IPMI commands
	IpmiUsername *string `json:"ipmiUsername"`
	// Password used to authenticate when issuing IPMI commands
	IpmiPassword       *string                `json:"ipmiPassword"`
	Connections        []*EditConnectionInput `json:"connections"`
	DeletedConnections []string               `json:"deletedConnections"`
	// Site that device belongs to
	Site *string `json:"site"`
}

type FieldParam struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Kpi struct {
	Name  *string `json:"name"`
	Type  *string `json:"type"`
	Value *string `json:"value"`
}

// pair of kind / category for search parameters
type KindCatPair struct {
	Category string `json:"category"`
	Kind     string `json:"kind"`
}

type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type LatLngInput struct {
	// Degrees latitude
	Lat float64 `json:"lat"`
	// Degrees longitude
	Lng float64 `json:"lng"`
}

type MapSettings struct {
	Center *LatLng `json:"center"`
	Zoom   *int    `json:"zoom"`
}

type Metric struct {
	Dimension string     `json:"dimension"`
	Value     *float64   `json:"value"`
	Timestamp *time.Time `json:"timestamp"`
}

type OktoResource struct {
	ID        string  `json:"id"`
	Name      *string `json:"name"`
	Kind      *string `json:"kind"`
	Status    *Status `json:"status"`
	Namespace *string `json:"namespace"`
	Manifest  *string `json:"manifest"`
}

type Progress struct {
	Step int `json:"step"`
	Goal int `json:"goal"`
}

type ProvisionChart struct {
	// Unique chart name in Helm repo. name field in Chart.yaml
	Name string `json:"name"`
	// version field in Chart.yaml
	Version         string  `json:"version"`
	BaremetalValues *string `json:"baremetalValues"`
	SoftwareValues  *string `json:"softwareValues"`
}

type ProvisionInput struct {
	// Chart name in Helm repo
	ChartName string `json:"chartName"`
	// Chart semantic version number
	ChartVersion string `json:"chartVersion"`
	// Site selector for Provisioner block deployment
	ProvisionerSite    string `json:"provisionerSite"`
	ProvisionInventory string `json:"provisionInventory"`
	// Display name of device to be shown in the UI
	DisplayName string `json:"displayName"`
	// Geographical position of device
	Position *LatLngInput `json:"position"`
	// Site that device belongs to
	Site string `json:"site"`
}

type Server struct {
	DiagramNode string  `json:"diagramNode"`
	Device      *Device `json:"device"`
	Pinned      bool    `json:"pinned"`
}

type ServerInput struct {
	DiagramNode string `json:"diagramNode"`
	Device      string `json:"device"`
	Pinned      *bool  `json:"pinned"`
}

type Service struct {
	ID           string        `json:"id"`
	DiagramNode  *string       `json:"diagramNode"`
	Name         string        `json:"name"`
	Using        *Dependencies `json:"using"`
	Connections  []*Connection `json:"connections"`
	PlatformKPIs []*Kpi        `json:"platformKPIs"`
	Kpis         []*Kpi        `json:"kpis"`
	Status       *Status       `json:"status"`
	Fields       []*FieldParam `json:"fields"`
	Tenant       *string       `json:"tenant"`
	Owner        *string       `json:"owner"`
	CreatedAt    *time.Time    `json:"createdAt"`
}

type ServiceChain struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Blocks       []*Block      `json:"blocks"`
	Using        *Dependencies `json:"using"`
	Connections  []*Connection `json:"connections"`
	PlatformKPIs []*Kpi        `json:"platformKPIs"`
	Kpis         []*Kpi        `json:"kpis"`
	Servers      []*Server     `json:"servers"`
	Status       *Status       `json:"status"`
	Tenant       *string       `json:"tenant"`
	Owner        *string       `json:"owner"`
	CreatedAt    *time.Time    `json:"createdAt"`
}

type ServiceParam struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Settings struct {
	Map *MapSettings `json:"map"`
}

type Site struct {
	ID          string         `json:"id"`
	DisplayName *string        `json:"displayName"`
	Description *string        `json:"description"`
	Ancestors   []*Site        `json:"ancestors"`
	Sites       []*Site        `json:"sites"`
	Devices     *DevicesResult `json:"devices"`
}

type SiteInput struct {
	DisplayName *string `json:"displayName"`
	Description *string `json:"description"`
	// Top-down ordered list of ancestor site IDs. If empty, site is created at the top level
	Ancestors []string `json:"ancestors"`
}

type Tag struct {
	Key string `json:"key"`
	Inc bool   `json:"inc"`
}

type TagInput struct {
	Key string `json:"key"`
	Inc bool   `json:"inc"`
}

type User struct {
	ID    string  `json:"id"`
	Name  *string `json:"name"`
	Email *string `json:"email"`
	Org   *Site   `json:"org"`
	Role  *Role   `json:"role"`
}

type DeviceStatus string

const (
	DeviceStatusUnknown   DeviceStatus = "UNKNOWN"
	DeviceStatusReady     DeviceStatus = "READY"
	DeviceStatusDeploying DeviceStatus = "DEPLOYING"
)

var AllDeviceStatus = []DeviceStatus{
	DeviceStatusUnknown,
	DeviceStatusReady,
	DeviceStatusDeploying,
}

func (e DeviceStatus) IsValid() bool {
	switch e {
	case DeviceStatusUnknown, DeviceStatusReady, DeviceStatusDeploying:
		return true
	}
	return false
}

func (e DeviceStatus) String() string {
	return string(e)
}

func (e *DeviceStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DeviceStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DeviceStatus", str)
	}
	return nil
}

func (e DeviceStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type EntityType string

const (
	EntityTypeUnknown EntityType = "UNKNOWN"
	EntityTypeBlock   EntityType = "BLOCK"
	EntityTypeDevice  EntityType = "DEVICE"
)

var AllEntityType = []EntityType{
	EntityTypeUnknown,
	EntityTypeBlock,
	EntityTypeDevice,
}

func (e EntityType) IsValid() bool {
	switch e {
	case EntityTypeUnknown, EntityTypeBlock, EntityTypeDevice:
		return true
	}
	return false
}

func (e EntityType) String() string {
	return string(e)
}

func (e *EntityType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EntityType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EntityType", str)
	}
	return nil
}

func (e EntityType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type IpmiCommand string

const (
	IpmiCommandRestart  IpmiCommand = "RESTART"
	IpmiCommandReset    IpmiCommand = "RESET"
	IpmiCommandPowerOn  IpmiCommand = "POWER_ON"
	IpmiCommandPowerOff IpmiCommand = "POWER_OFF"
)

var AllIpmiCommand = []IpmiCommand{
	IpmiCommandRestart,
	IpmiCommandReset,
	IpmiCommandPowerOn,
	IpmiCommandPowerOff,
}

func (e IpmiCommand) IsValid() bool {
	switch e {
	case IpmiCommandRestart, IpmiCommandReset, IpmiCommandPowerOn, IpmiCommandPowerOff:
		return true
	}
	return false
}

func (e IpmiCommand) String() string {
	return string(e)
}

func (e *IpmiCommand) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = IpmiCommand(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid IpmiCommand", str)
	}
	return nil
}

func (e IpmiCommand) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Role string

const (
	RoleAdmin    Role = "ADMIN"
	RoleOperator Role = "OPERATOR"
	RoleViewer   Role = "VIEWER"
)

var AllRole = []Role{
	RoleAdmin,
	RoleOperator,
	RoleViewer,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleOperator, RoleViewer:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Status string

const (
	StatusReady       Status = "READY"
	StatusDeploying   Status = "DEPLOYING"
	StatusUndeploying Status = "UNDEPLOYING"
	StatusFailure     Status = "FAILURE"
	StatusUnknown     Status = "UNKNOWN"
	StatusPending     Status = "PENDING"
)

var AllStatus = []Status{
	StatusReady,
	StatusDeploying,
	StatusUndeploying,
	StatusFailure,
	StatusUnknown,
	StatusPending,
}

func (e Status) IsValid() bool {
	switch e {
	case StatusReady, StatusDeploying, StatusUndeploying, StatusFailure, StatusUnknown, StatusPending:
		return true
	}
	return false
}

func (e Status) String() string {
	return string(e)
}

func (e *Status) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Status(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

func (e Status) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}