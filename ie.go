package pfcp

import "net"

type CreatePDR struct {
	Header IEHeader
}
type PDI struct {
	Header IEHeader
}
type CreateFAR struct {
	Header IEHeader
}
type ForwardingParameters struct {
	Header IEHeader
}
type DuplicatingParameters struct {
	Header IEHeader
}
type CreateURR struct {
	Header IEHeader
}
type CreateQER struct {
	Header IEHeader
}
type CreatedPDR struct {
	Header IEHeader
}
type UpdatePDR struct {
	Header IEHeader
}
type UpdateFAR struct {
	Header IEHeader
}
type UpdateForwardingParameters struct {
	Header IEHeader
}
type UpdateBARPFCPSessionReportResponse struct {
	Header IEHeader
}
type UpdateURR struct {
	Header IEHeader
}
type UpdateQER struct {
	Header IEHeader
}
type RemovePDR struct {
	Header IEHeader
}
type RemoveFAR struct {
	Header IEHeader
}
type RemoveURR struct {
	Header IEHeader
}
type RemoveQER struct {
	Header IEHeader
}
type Cause struct {
	Header IEHeader
}
type SourceInterface struct {
	Header IEHeader
}
type FTEID struct {
	Header IEHeader
}
type NetworkInstance struct {
	Header IEHeader
}
type SDFFilter struct {
	Header IEHeader
}
type ApplicationID struct {
	Header IEHeader
}
type GateStatus struct {
	Header IEHeader
}
type MBR struct {
	Header IEHeader
}
type GBR struct {
	Header IEHeader
}
type QERCorrelationID struct {
	Header IEHeader
}
type Precedence struct {
	Header IEHeader
}
type TransportLevelMarking struct {
	Header IEHeader
}
type VolumeThreshold struct {
	Header IEHeader
}
type TimeThreshold struct {
	Header IEHeader
}
type MonitoringTime struct {
	Header IEHeader
}
type SubsequentVolumeThreshold struct {
	Header IEHeader
}
type SubsequentTimeThreshold struct {
	Header IEHeader
}
type InactivityDetectionTime struct {
	Header IEHeader
}
type ReportingTriggers struct {
	Header IEHeader
}
type RedirectInformation struct {
	Header IEHeader
}
type ReportType struct {
	Header IEHeader
}
type OffendingIE struct {
	Header IEHeader
}
type ForwardingPolicy struct {
	Header IEHeader
}
type DestinationInterface struct {
	Header IEHeader
}
type UPFunctionFeatures struct {
	Header IEHeader
}
type ApplyAction struct {
	Header IEHeader
}
type DownlinkDataServiceInformation struct {
	Header IEHeader
}
type DownlinkDataNotificationDelay struct {
	Header IEHeader
}
type DLBufferingDuration struct {
	Header IEHeader
}
type DLBufferingSuggestedPacketCount struct {
	Header IEHeader
}
type PFCPSMReqFlags struct {
	Header IEHeader
}
type PFCPSRRspFlags struct {
	Header IEHeader
}
type LoadControlInformation struct {
	Header IEHeader
}
type SequenceNumber struct {
	Header IEHeader
}
type Metric struct {
	Header IEHeader
}
type OverloadControlInformation struct {
	Header IEHeader
}
type Timer struct {
	Header IEHeader
}
type PDRID struct {
	Header IEHeader
}
type FSEID struct {
	Header IEHeader
}
type PFDContext struct {
	Header IEHeader
}
type NodeID struct {
	Header IEHeader
}
type PFDContents struct {
	Header IEHeader
}
type MeasurementMethod struct {
	Header IEHeader
}
type UsageReportTrigger struct {
	Header IEHeader
}
type MeasurementPeriod struct {
	Header IEHeader
}
type FQCSID struct {
	Header IEHeader
}
type VolumeMeasurement struct {
	Header IEHeader
}
type DurationMeasurement struct {
	Header IEHeader
}
type ApplicationDetectionInformation struct {
	Header IEHeader
}
type TimeOfFirstPacket struct {
	Header IEHeader
}
type TimeOfLastPacket struct {
	Header IEHeader
}
type QuotaHoldingTime struct {
	Header IEHeader
}
type DroppedDLTrafficThreshold struct {
	Header IEHeader
}
type VolumeQuota struct {
	Header IEHeader
}
type TimeQuota struct {
	Header IEHeader
}
type StartTime struct {
	Header IEHeader
}
type EndTime struct {
	Header IEHeader
}
type QueryURR struct {
	Header IEHeader
}
type UsageReportSessionModificationResponse struct {
	Header IEHeader
}
type UsageReportSessionDeletionResponse struct {
	Header IEHeader
}
type UsageReportSessionReportRequest struct {
	Header IEHeader
}
type URRID struct {
	Header IEHeader
}
type LinkedURRID struct {
	Header IEHeader
}
type DownLinkDataReport struct {
	Header IEHeader
}
type OuterHeaderCreation struct {
	Header IEHeader
}
type CreateBAR struct {
	Header IEHeader
}
type UpdateBARSessionModificationRequest struct {
	Header IEHeader
}
type RemoveBAR struct {
	Header IEHeader
}
type BARID struct {
	Header IEHeader
}
type CPFunctionFeatures struct {
	Header IEHeader
}
type UsageInformation struct {
	Header IEHeader
}
type ApplicationInstanceID struct {
	Header IEHeader
}
type FlowInformation struct {
	Header IEHeader
}
type UEIPAddress struct {
	Header IEHeader
}

type PacketRate struct {
	Header IEHeader
}

type OuterHeaderRemoval struct {
	Header IEHeader
}

type RecoveryTimeStamp struct {
	Header IEHeader
	Value  uint32
}

func (r *RecoveryTimeStamp) encode(w Writer) error {
	w.MWrite(r.Value)
	return w.Error(r)
}

func (r *RecoveryTimeStamp) decode(R Reader) error {
	r.Value = R.ReadUint32()
	return R.Error(r)
}

func (r *RecoveryTimeStamp) length() uint16 {
	return 4
}

func (r *RecoveryTimeStamp) Type() IEType {
	return IeRecoveryTimeStamp
}

type DLFlowLevelMarking struct {
	Header IEHeader
}
type HeaderEnrichment struct {
	Header IEHeader
}
type ErrorIndicationReport struct {
	Header IEHeader
}
type MeasurementInformation struct {
	Header IEHeader
}
type NodeReportType struct {
	Header IEHeader
}
type UserPlanePathFailureReport struct {
	Header IEHeader
}
type RemoteGTPUPeer struct {
	Header IEHeader
}
type URSEQN struct {
	Header IEHeader
}
type UpdateDuplicatingParameters struct {
	Header IEHeader
}
type ActivatePredefinedRules struct {
	Header IEHeader
}
type DeactivatePredefinedRules struct {
	Header IEHeader
}
type FARID struct {
	Header IEHeader
}
type QERID struct {
	Header IEHeader
}
type OCIFlags struct {
	Header IEHeader
}
type PFCPAssociationReleaseRequest struct {
	Header IEHeader
}
type GracefulReleasePeriod struct {
	Header IEHeader
}
type PDNType struct {
	Header IEHeader
}
type FailedRuleID struct {
	Header IEHeader
}
type TimeQuotaMechanism struct {
	Header IEHeader
}

type UserPlaneInactivityTimer struct {
	Header IEHeader
}
type AggregatedURRs struct {
	Header IEHeader
}
type Multiplier struct {
	Header IEHeader
}
type AggregatedURRID struct {
	Header IEHeader
}
type SubsequentVolumeQuota struct {
	Header IEHeader
}
type SubsequentTimeQuota struct {
	Header IEHeader
}
type RQI struct {
	Header IEHeader
}
type QFI struct {
	Header IEHeader
}
type QueryURRReference struct {
	Header IEHeader
}
type AdditionalUsageReportsInformation struct {
	Header IEHeader
}
type CreateTrafficEndpoint struct {
	Header IEHeader
}
type CreatedTrafficEndpoint struct {
	Header IEHeader
}
type UpdateTrafficEndpoint struct {
	Header IEHeader
}
type RemoveTrafficEndpoint struct {
	Header IEHeader
}
type TrafficEndpointID struct {
	Header IEHeader
}
type EthernetPacketFilter struct {
	Header IEHeader
}
type MACAddress struct {
	Header IEHeader
}
type CTAG struct {
	Header IEHeader
}
type STAG struct {
	Header IEHeader
}
type EtherType struct {
	Header IEHeader
}
type Proxying struct {
	Header IEHeader
}
type EthernetFilterID struct {
	Header IEHeader
}
type EthernetFilterProperties struct {
	Header IEHeader
}
type SuggestedBufferingPacketsCount struct {
	Header IEHeader
}
type UserID struct {
	Header IEHeader
}
type EthernetPDUSessionInformation struct {
	Header IEHeader
}
type EthernetTrafficInformation struct {
	Header IEHeader
}
type MACAddressesDetected struct {
	Header IEHeader
}
type MACAddressesRemoved struct {
	Header IEHeader
}
type EthernetInactivityTimer struct {
	Header IEHeader
}
type AdditionalMonitoringTime struct {
	Header IEHeader
}
type EventQuota struct {
	Header IEHeader
}
type EventThreshold struct {
	Header IEHeader
}
type SubsequentEventQuota struct {
	Header IEHeader
}
type SubsequentEventThreshold struct {
	Header IEHeader
}
type TraceInformation struct {
	Header IEHeader
}
type FramedRoute struct {
	Header IEHeader
}
type FramedRouting struct {
	Header IEHeader
}
type FramedIPv6Route struct {
	Header IEHeader
}
type TimeStamp struct {
	Header IEHeader
}
type AveragingWindow struct {
	Header IEHeader
}
type PagingPolicyIndicator struct {
	Header IEHeader
}
type APNDNN struct {
	Header IEHeader
}
type ThirdGPPInterfaceType struct {
	Header IEHeader
}
type PFCPSRReqFlags struct {
	Header IEHeader
}
type PFCPAUReqFlags struct {
	Header IEHeader
}
type ActivationTime struct {
	Header IEHeader
}
type DeactivationTime struct {
	Header IEHeader
}
type CreateMAR struct {
	Header IEHeader
}
type ThirdGPPAccessForwardingActionInformation struct {
	Header IEHeader
}
type Non3GPPAccessForwardingActionInformation struct {
	Header IEHeader
}
type RemoveMAR struct {
	Header IEHeader
}
type UpdateMAR struct {
	Header IEHeader
}
type MARID struct {
	Header IEHeader
}
type SteeringFunctionality struct {
	Header IEHeader
}
type SteeringMode struct {
	Header IEHeader
}
type Weight struct {
	Header IEHeader
}
type Priority struct {
	Header IEHeader
}
type Update3GPPAccessForwardingActionInformation struct {
	Header IEHeader
}
type UpdateNon3GPPAccessForwardingActionInformation struct {
	Header IEHeader
}
type UEIPAddressPoolIdentity struct {
	Header IEHeader
}
type AlternativeSMFIPAddress struct {
	Header IEHeader
}
type PacketReplicationAndDetectionCarryOnInformation struct {
	Header IEHeader
}
type SMFSetID struct {
	Header IEHeader
}
type QuotaValidityTime struct {
	Header IEHeader
}
type NumberOfReports struct {
	Header IEHeader
}
type PFCPSessionRetentionInformationWithinPFCPAssociationSetupRequest struct {
	Header IEHeader
}
type PFCPASRspFlags struct {
	Header IEHeader
}
type CPPFCPEntityIPAddress struct {
	Header IEHeader
}
type PFCPSEReqFlags struct {
	Header IEHeader
}
type UserPlanePathRecoveryReport struct {
	Header IEHeader
}
type IPMulticastAddressingInfoWithinPFCPSessionEstablishmentRequest struct {
	Header IEHeader
}
type JoinIPMulticastInformationIEWithinUsageReport struct {
	Header IEHeader
}
type LeaveIPMulticastInformationIEWithinUsageReport struct {
	Header IEHeader
}
type IPMulticastAddress struct {
	Header IEHeader
}

type SourceIPAddress struct {
	Header      IEHeader
	MPL, V4, V6 bool
	IPv4Address net.IP
	IPv6Address net.IP
	MaskLen     uint8
}

func (s *SourceIPAddress) Type() IEType {
	return IeSourceIPAddress
}

func (s *SourceIPAddress) encode(w Writer) error {
	w.MWrite(Bool2Uint8(s.MPL)<<2 | Bool2Uint8(s.V4)<<1 | Bool2Uint8(s.V6))
	w.TestWrite(s.V4, s.IPv4Address)
	w.TestWrite(s.V6, s.IPv6Address)
	w.TestWrite(s.MPL, s.MaskLen)
	return w.Error(s)
}

func (s *SourceIPAddress) decode(r Reader) error {
	_, _, _, _, _, s.MPL, s.V4, s.V6 = r.ReadMBool()
	r.TestRead(s.V4, 4, &s.IPv4Address)
	r.TestRead(s.V6, 16, &s.IPv6Address)
	r.TestRead(s.MPL, 0, &s.MaskLen)
	return r.Error(s)
}

func (s *SourceIPAddress) length() uint16 {
	return 1 + uint16(Bool2Uint8(s.V6))*16 + uint16(Bool2Uint8(s.V4))*4 + uint16(Bool2Uint8(s.MPL)*1)
}

type PacketRateStatus struct {
	Header IEHeader
}
type CreateBridgeInfoForTSC struct {
	Header IEHeader
}
type CreatedBridgeInfoForTSC struct {
	Header IEHeader
}
type DSTTPortNumber struct {
	Header IEHeader
}
type NWTTPortNumber struct {
	Header IEHeader
}
type FiveGSUserPlaneNode struct {
	Header IEHeader
}
type TSCManagementInformationIEWithinPFCPSessionModificationRequest struct {
	Header IEHeader
}
type TSCManagementInformationIEWithinPFCPSessionModificationResponse struct {
	Header IEHeader
}
type TSCManagementInformationIEWithinPFCPSessionReportRequest struct {
	Header IEHeader
}
type PortManagementInformationContainer struct {
	Header IEHeader
}
type ClockDriftControlInformation struct {
	Header IEHeader
}
type RequestedClockDriftInformation struct {
	Header IEHeader
}
type ClockDriftReport struct {
	Header IEHeader
}
type TimeDomainNumber struct {
	Header IEHeader
}
type TimeOffsetThreshold struct {
	Header IEHeader
}
type CumulativeRateRatioThreshold struct {
	Header IEHeader
}
type TimeOffsetMeasurement struct {
	Header IEHeader
}
type CumulativeRateRatioMeasurement struct {
	Header IEHeader
}
type RemoveSRR struct {
	Header IEHeader
}
type CreateSRR struct {
	Header IEHeader
}
type UpdateSRR struct {
	Header IEHeader
}
type SessionReport struct {
	Header IEHeader
}
type SRRID struct {
	Header IEHeader
}
type AccessAvailabilityControlInformation struct {
	Header IEHeader
}
type RequestedAccessAvailabilityInformation struct {
	Header IEHeader
}
type AccessAvailabilityReport struct {
	Header IEHeader
}
type AccessAvailabilityInformation struct {
	Header IEHeader
}
type ProvideATSSSControlInformation struct {
	Header IEHeader
}
type ATSSSControlParameters struct {
	Header IEHeader
}
type MPTCPControlInformation struct {
	Header IEHeader
}
type ATSSSLLControlInformation struct {
	Header IEHeader
}
type PMFControlInformation struct {
	Header IEHeader
}
type MPTCPParameters struct {
	Header IEHeader
}
type ATSSSLLParameters struct {
	Header IEHeader
}
type PMFParameters struct {
	Header IEHeader
}
type MPTCPAddressInformation struct {
	Header IEHeader
}
type UELinkSpecificIPAddress struct {
	Header IEHeader
}
type PMFAddressInformation struct {
	Header IEHeader
}
type ATSSSLLInformation struct {
	Header IEHeader
}
type DataNetworkAccessIdentifier struct {
	Header IEHeader
}
type UEIPAddressPoolInformation struct {
	Header IEHeader
}
type AveragePacketDelay struct {
	Header IEHeader
}
type MinimumPacketDelay struct {
	Header IEHeader
}
type MaximumPacketDelay struct {
	Header IEHeader
}
type QoSReportTrigger struct {
	Header IEHeader
}
type GTPUPathQoSControlInformation struct {
	Header IEHeader
}
type GTPUPathQoSReportPFCPNodeReportRequest struct {
	Header IEHeader
}
type QoSInformationInGTPUPathQoSReport struct {
	Header IEHeader
}
type GTPUPathInterfaceType struct {
	Header IEHeader
}
type QoSMonitoringPerQoSFlowControlInformation struct {
	Header IEHeader
}
type RequestedQoSMonitoring struct {
	Header IEHeader
}
type ReportingFrequency struct {
	Header IEHeader
}
type PacketDelayThresholds struct {
	Header IEHeader
}
type MinimumWaitTime struct {
	Header IEHeader
}
type QoSMonitoringReport struct {
	Header IEHeader
}
type QoSMonitoringMeasurement struct {
	Header IEHeader
}
type MTEDTControlInformation struct {
	Header IEHeader
}
type DLDataPacketsSize struct {
	Header IEHeader
}
type QERControlIndications struct {
	Header IEHeader
}
type PacketRateStatusReport struct {
	Header IEHeader
}
type NFInstanceID struct {
	Header IEHeader
}
type EthernetContextInformation struct {
	Header IEHeader
}
type RedundantTransmissionParameters struct {
	Header IEHeader
}
type UpdatedPDR struct {
	Header IEHeader
}
type SNSSAI struct {
	Header IEHeader
}
type IPVersion struct {
	Header IEHeader
}
type PFCPASReqFlags struct {
	Header IEHeader
}
type DataStatus struct {
	Header IEHeader
}
type ProvideRDSConfigurationInformation struct {
	Header IEHeader
}
type RDSConfigurationInformation struct {
	Header IEHeader
}
type QueryPacketRateStatusIEWithinPFCPSessionModificationRequest struct {
	Header IEHeader
}
type PacketRateStatusReportIEWithinPFCPSessionModificationResponse struct {
	Header IEHeader
}
type MPTCPApplicableIndication struct {
	Header IEHeader
}
type BridgeManagementInformationContainer struct {
	Header IEHeader
}
type UEIPAddressUsageInformation struct {
	Header IEHeader
}
type NumberOfUEIPAddresses struct {
	Header IEHeader
}
type ValidityTimer struct {
	Header IEHeader
}
type RedundantTransmissionForwardingParameters struct {
	Header IEHeader
}
type TransportDelayReporting struct {
	Header IEHeader
}
type PartialFailureInformation struct {
	Header IEHeader
}
type Spare struct {
	Header IEHeader
}
type OffendingIEInformation struct {
	Header IEHeader
}
type RATType struct {
	Header IEHeader
}
type L2TPTunnelInformation struct {
	Header IEHeader
}
type L2TPSessionInformation struct {
	Header IEHeader
}
type L2TPUserAuthenticationIE struct {
	Header IEHeader
}
type CreatedL2TPSession struct {
	Header IEHeader
}
type LNSAddress struct {
	Header IEHeader
}
type TunnelPreference struct {
	Header IEHeader
}
type CallingNumber struct {
	Header IEHeader
}
type CalledNumber struct {
	Header IEHeader
}
type L2TPSessionIndications struct {
	Header IEHeader
}
type DNSServerAddress struct {
	Header IEHeader
}
type NBNSServerAddress struct {
	Header IEHeader
}
type MaximumReceiveUnit struct {
	Header IEHeader
}
type Thresholds struct {
	Header IEHeader
}
type SteeringModeIndicator struct {
	Header IEHeader
}
type PFCPSessionChangeInfo struct {
	Header IEHeader
}
type GroupID struct {
	Header IEHeader
}
type CPIPAddress struct {
	Header IEHeader
}
type IPAddressAndPortNumberReplacement struct {
	Header IEHeader
}
type DNSQueryFilter struct {
	Header IEHeader
}
type DirectReportingInformation struct {
	Header IEHeader
}
type EventNotificationURI struct {
	Header IEHeader
}
type NotificationCorrelationID struct {
	Header IEHeader
}
type ReportingFlags struct {
	Header IEHeader
}
type PredefinedRulesName struct {
	Header IEHeader
}
type MBSSessionN4mbControlInformation struct {
	Header IEHeader
}
type MBSMulticastParameters struct {
	Header IEHeader
}
type AddMBSUniCastParameters struct {
	Header IEHeader
}
type MBSSessionN4mbInformation struct {
	Header IEHeader
}
type RemoveMBSUniCastParameters struct {
	Header IEHeader
}
type MBSSessionIdentifier struct {
	Header IEHeader
}
type MulticastTransportInformation struct {
	Header IEHeader
}
type MBSN4mbReqFlags struct {
	Header IEHeader
}
type LocalIngressTunnel struct {
	Header IEHeader
}
type MBSUniCastParametersID struct {
	Header IEHeader
}
type MBSSessionN4ControlInformation struct {
	Header IEHeader
}
type MBSSessionN4Information struct {
	Header IEHeader
}
type MBSN4RespFlags struct {
	Header IEHeader
}
type TunnelPassword struct {
	Header IEHeader
}
type AreaSessionID struct {
	Header IEHeader
}
type PeerUPRestartReport struct {
	Header IEHeader
}
type DSCPtoPPIControlInformation struct {
	Header IEHeader
}
type DSCPtoPPIMappingInformation struct {
	Header IEHeader
}
type PFCPSDRspFlags struct {
	Header IEHeader
}
