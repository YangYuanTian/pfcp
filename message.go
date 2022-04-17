package pfcp

type HeartbeatRequest struct {
	Header            MsgHeader
	RecoveryTimeStamp *RecoveryTimeStamp //M
	SourceIPAddress   *SourceIPAddress
}

func (h *HeartbeatRequest) Type() MessageType {
	return MsgHeartbeatRequest
}

type HeartbeatResponse struct {
	Header            MsgHeader
	RecoveryTimeStamp *RecoveryTimeStamp //M
}

func (h *HeartbeatResponse) Type() MessageType {
	return MsgHeartbeatResponse
}

type PFDManagementRequest struct {
	Header             MsgHeader
	ApplicationIDsPFDs []*ApplicationIDsPFDs
	NodeID             *NodeID
}

func (P *PFDManagementRequest) Type() MessageType {
	return MsgPFDManagementRequest
}

type PFDManagementResponse struct {
	Header MsgHeader
}

type AssociationSetupRequest struct {
	Header MsgHeader
}

type AssociationSetupResponse struct {
	Header MsgHeader
}

type AssociationUpdateRequest struct {
	Header MsgHeader
}

type AssociationUpdateResponse struct {
	Header MsgHeader
}
type AssociationReleaseRequest struct {
	Header MsgHeader
}
type AssociationReleaseResponse struct {
	Header MsgHeader
}
type VersionNotSupportedResponse struct {
	Header MsgHeader
}
type NodeReportRequest struct {
	Header MsgHeader
}
type NodeReportResponse struct {
	Header MsgHeader
}
type SessionSetDeletionRequest struct {
	Header MsgHeader
}
type SessionSetDeletionResponse struct {
	Header MsgHeader
}
type SessionSetModificationRequest struct {
	Header MsgHeader
}
type SessionSetModifcationResponse struct {
	Header MsgHeader
}
type SessionSetModificationReques struct {
	Header MsgHeader
}
type SessionSetModificationResponse struct {
	Header MsgHeader
}

type SessionEstablishmentRequest struct {
	Header MsgHeader
}
type SessionEstablishmentResponse struct {
	Header MsgHeader
}
type SessionModificationRequest struct {
	Header MsgHeader
}
type SessionModificationResponse struct {
	Header MsgHeader
}
type SessionDeletionRequest struct {
	Header MsgHeader
}
type SessionDeletionResponse struct {
	Header MsgHeader
}
type SessionReportRequest struct {
	Header MsgHeader
}
type SessionReportResponse struct {
	Header MsgHeader
}
