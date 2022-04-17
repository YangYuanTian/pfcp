package pfcp

const (
	// PFCP Noderelated messages

	MsgHeartbeatRequest              MessageType = 1
	MsgHeartbeatResponse             MessageType = 2
	MsgPFDManagementRequest          MessageType = 3
	MsgPFDManagementResponse         MessageType = 4
	MsgAssociationSetupRequest       MessageType = 5
	MsgAssociationSetupResponse      MessageType = 6
	MsgAssociationUpdateRequest      MessageType = 7
	MsgAssociationUpdateResponse     MessageType = 8
	MsgAssociationReleaseRequest     MessageType = 9
	MsgAssociationReleaseResponse    MessageType = 10
	MsgVersionNotSupportedResponse   MessageType = 11
	MsgNodeReportRequest             MessageType = 12
	MsgNodeReportResponse            MessageType = 13
	MsgSessionSetDeletionRequest     MessageType = 14
	MsgSessionSetDeletionResponse    MessageType = 15
	MsgSessionSetModificationRequest MessageType = 16
	MsgSessionSetModifcationResponse MessageType = 17

	//For future use  18 to =49
	// PFCP Session related messages

	MsgSessionEstablishmentRequest  MessageType = 50
	MsgSessionEstablishmentResponse MessageType = 51
	MsgSessionModificationRequest   MessageType = 52
	MsgSessionModificationResponse  MessageType = 53
	MsgSessionDeletionRequest       MessageType = 54
	MsgSessionDeletionResponse      MessageType = 55
	MsgSessionReportRequest         MessageType = 56
	MsgSessionReportResponse        MessageType = 57

//Forfutureuse MessageType 58 to =99
//Othermessages
//Forfutureuse MessageType 100 to 2=55

)
