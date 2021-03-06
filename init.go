package pfcp

func init() {
	fct.store(IeCreatePDR, new(CreatePDR))
	fct.store(IePDI, new(PDI))
	fct.store(IeCreateFAR, new(CreateFAR))
	fct.store(IeForwardingParameters, new(ForwardingParameters))
	fct.store(IeDuplicatingParameters, new(DuplicatingParameters))
	fct.store(IeCreateURR, new(CreateURR))
	fct.store(IeCreateQER, new(CreateQER))
	fct.store(IeCreatedPDR, new(CreatedPDR))
	fct.store(IeUpdatePDR, new(UpdatePDR))
	fct.store(IeUpdateFAR, new(UpdateFAR))
	fct.store(IeUpdateForwardingParameters, new(UpdateForwardingParameters))
	fct.store(IeUpdateBARPFCPSessionReportResponse, new(UpdateBARPFCPSessionReportResponse))
	fct.store(IeUpdateURR, new(UpdateURR))
	fct.store(IeUpdateQER, new(UpdateQER))
	fct.store(IeRemovePDR, new(RemovePDR))
	fct.store(IeRemoveFAR, new(RemoveFAR))
	fct.store(IeRemoveURR, new(RemoveURR))
	fct.store(IeRemoveQER, new(RemoveQER))
	fct.store(IeCause, new(Cause))
	fct.store(IeSourceInterface, new(SourceInterface))
	fct.store(IeFTEID, new(FTEID))
	fct.store(IeNetworkInstance, new(NetworkInstance))
	fct.store(IeSDFFilter, new(SDFFilter))
	fct.store(IeApplicationID, new(ApplicationID))
	fct.store(IeGateStatus, new(GateStatus))
	fct.store(IeMBR, new(MBR))
	fct.store(IeGBR, new(GBR))
	fct.store(IeQERCorrelationID, new(QERCorrelationID))
	fct.store(IePrecedence, new(Precedence))
	fct.store(IeTransportLevelMarking, new(TransportLevelMarking))
	fct.store(IeVolumeThreshold, new(VolumeThreshold))
	fct.store(IeTimeThreshold, new(TimeThreshold))
	fct.store(IeMonitoringTime, new(MonitoringTime))
	fct.store(IeSubsequentVolumeThreshold, new(SubsequentVolumeThreshold))
	fct.store(IeSubsequentTimeThreshold, new(SubsequentTimeThreshold))
	fct.store(IeInactivityDetectionTime, new(InactivityDetectionTime))
	fct.store(IeReportingTriggers, new(ReportingTriggers))
	fct.store(IeRedirectInformation, new(RedirectInformation))
	fct.store(IeReportType, new(ReportType))
	fct.store(IeOffendingIE, new(OffendingIE))
	fct.store(IeForwardingPolicy, new(ForwardingPolicy))
	fct.store(IeDestinationInterface, new(DestinationInterface))
	fct.store(IeUPFunctionFeatures, new(UPFunctionFeatures))
	fct.store(IeApplyAction, new(ApplyAction))
	fct.store(IeDownlinkDataServiceInformation, new(DownlinkDataServiceInformation))
	fct.store(IeDownlinkDataNotificationDelay, new(DownlinkDataNotificationDelay))
	fct.store(IeDLBufferingDuration, new(DLBufferingDuration))
	fct.store(IeDLBufferingSuggestedPacketCount, new(DLBufferingSuggestedPacketCount))
	fct.store(IePFCPSMReqFlags, new(PFCPSMReqFlags))
	fct.store(IePFCPSRRspFlags, new(PFCPSRRspFlags))
	fct.store(IeLoadControlInformation, new(LoadControlInformation))
	fct.store(IeSequenceNumber, new(SequenceNumber))
	fct.store(IeMetric, new(Metric))
	fct.store(IeOverloadControlInformation, new(OverloadControlInformation))
	fct.store(IeTimer, new(Timer))
	fct.store(IePDRID, new(PDRID))
	fct.store(IeFSEID, new(FSEID))
	fct.store(IeApplicationIDsPFDs, new(ApplicationIDsPFDs))
	fct.store(IePFDContext, new(PFDContext))
	fct.store(IeNodeID, new(NodeID))
	fct.store(IePFDContents, new(PFDContents))
	fct.store(IeMeasurementMethod, new(MeasurementMethod))
	fct.store(IeUsageReportTrigger, new(UsageReportTrigger))
	fct.store(IeMeasurementPeriod, new(MeasurementPeriod))
	fct.store(IeFQCSID, new(FQCSID))
	fct.store(IeVolumeMeasurement, new(VolumeMeasurement))
	fct.store(IeDurationMeasurement, new(DurationMeasurement))
	fct.store(IeApplicationDetectionInformation, new(ApplicationDetectionInformation))
	fct.store(IeTimeOfFirstPacket, new(TimeOfFirstPacket))
	fct.store(IeTimeOfLastPacket, new(TimeOfLastPacket))
	fct.store(IeQuotaHoldingTime, new(QuotaHoldingTime))
	fct.store(IeDroppedDLTrafficThreshold, new(DroppedDLTrafficThreshold))
	fct.store(IeVolumeQuota, new(VolumeQuota))
	fct.store(IeTimeQuota, new(TimeQuota))
	fct.store(IeStartTime, new(StartTime))
	fct.store(IeEndTime, new(EndTime))
	fct.store(IeQueryURR, new(QueryURR))
	fct.store(IeUsageReportSessionModificationResponse, new(UsageReportSessionModificationResponse))
	fct.store(IeUsageReportSessionDeletionResponse, new(UsageReportSessionDeletionResponse))
	fct.store(IeUsageReportSessionReportRequest, new(UsageReportSessionReportRequest))
	fct.store(IeURRID, new(URRID))
	fct.store(IeLinkedURRID, new(LinkedURRID))
	fct.store(IeDownLinkDataReport, new(DownLinkDataReport))
	fct.store(IeOuterHeaderCreation, new(OuterHeaderCreation))
	fct.store(IeCreateBAR, new(CreateBAR))
	fct.store(IeUpdateBARSessionModificationRequest, new(UpdateBARSessionModificationRequest))
	fct.store(IeRemoveBAR, new(RemoveBAR))
	fct.store(IeBARID, new(BARID))
	fct.store(IeCPFunctionFeatures, new(CPFunctionFeatures))
	fct.store(IeUsageInformation, new(UsageInformation))
	fct.store(IeApplicationInstanceID, new(ApplicationInstanceID))
	fct.store(IeFlowInformation, new(FlowInformation))
	fct.store(IeUEIPAddress, new(UEIPAddress))
	fct.store(IePacketRate, new(PacketRate))
	fct.store(IeOuterHeaderRemoval, new(OuterHeaderRemoval))
	fct.store(IeRecoveryTimeStamp, new(RecoveryTimeStamp))
	fct.store(IeDLFlowLevelMarking, new(DLFlowLevelMarking))
	fct.store(IeHeaderEnrichment, new(HeaderEnrichment))
	fct.store(IeErrorIndicationReport, new(ErrorIndicationReport))
	fct.store(IeMeasurementInformation, new(MeasurementInformation))
	fct.store(IeNodeReportType, new(NodeReportType))
	fct.store(IeUserPlanePathFailureReport, new(UserPlanePathFailureReport))
	fct.store(IeRemoteGTPUPeer, new(RemoteGTPUPeer))
	fct.store(IeURSEQN, new(URSEQN))
	fct.store(IeUpdateDuplicatingParameters, new(UpdateDuplicatingParameters))
	fct.store(IeActivatePredefinedRules, new(ActivatePredefinedRules))
	fct.store(IeDeactivatePredefinedRules, new(DeactivatePredefinedRules))
	fct.store(IeFARID, new(FARID))
	fct.store(IeQERID, new(QERID))
	fct.store(IeOCIFlags, new(OCIFlags))
	fct.store(IePFCPAssociationReleaseRequest, new(PFCPAssociationReleaseRequest))
	fct.store(IeGracefulReleasePeriod, new(GracefulReleasePeriod))
	fct.store(IePDNType, new(PDNType))
	fct.store(IeFailedRuleID, new(FailedRuleID))
	fct.store(IeTimeQuotaMechanism, new(TimeQuotaMechanism))
	fct.store(IeUserPlaneInactivityTimer, new(UserPlaneInactivityTimer))
	fct.store(IeAggregatedURRs, new(AggregatedURRs))
	fct.store(IeMultiplier, new(Multiplier))
	fct.store(IeAggregatedURRID, new(AggregatedURRID))
	fct.store(IeSubsequentVolumeQuota, new(SubsequentVolumeQuota))
	fct.store(IeSubsequentTimeQuota, new(SubsequentTimeQuota))
	fct.store(IeRQI, new(RQI))
	fct.store(IeQFI, new(QFI))
	fct.store(IeQueryURRReference, new(QueryURRReference))
	fct.store(IeAdditionalUsageReportsInformation, new(AdditionalUsageReportsInformation))
	fct.store(IeCreateTrafficEndpoint, new(CreateTrafficEndpoint))
	fct.store(IeCreatedTrafficEndpoint, new(CreatedTrafficEndpoint))
	fct.store(IeUpdateTrafficEndpoint, new(UpdateTrafficEndpoint))
	fct.store(IeRemoveTrafficEndpoint, new(RemoveTrafficEndpoint))
	fct.store(IeTrafficEndpointID, new(TrafficEndpointID))
	fct.store(IeEthernetPacketFilter, new(EthernetPacketFilter))
	fct.store(IeMACAddress, new(MACAddress))
	fct.store(IeCTAG, new(CTAG))
	fct.store(IeSTAG, new(STAG))
	fct.store(IeEtherType, new(EtherType))
	fct.store(IeProxying, new(Proxying))
	fct.store(IeEthernetFilterID, new(EthernetFilterID))
	fct.store(IeEthernetFilterProperties, new(EthernetFilterProperties))
	fct.store(IeSuggestedBufferingPacketsCount, new(SuggestedBufferingPacketsCount))
	fct.store(IeUserID, new(UserID))
	fct.store(IeEthernetPDUSessionInformation, new(EthernetPDUSessionInformation))
	fct.store(IeEthernetTrafficInformation, new(EthernetTrafficInformation))
	fct.store(IeMACAddressesDetected, new(MACAddressesDetected))
	fct.store(IeMACAddressesRemoved, new(MACAddressesRemoved))
	fct.store(IeEthernetInactivityTimer, new(EthernetInactivityTimer))
	fct.store(IeAdditionalMonitoringTime, new(AdditionalMonitoringTime))
	fct.store(IeEventQuota, new(EventQuota))
	fct.store(IeEventThreshold, new(EventThreshold))
	fct.store(IeSubsequentEventQuota, new(SubsequentEventQuota))
	fct.store(IeSubsequentEventThreshold, new(SubsequentEventThreshold))
	fct.store(IeTraceInformation, new(TraceInformation))
	fct.store(IeFramedRoute, new(FramedRoute))
	fct.store(IeFramedRouting, new(FramedRouting))
	fct.store(IeFramedIPv6Route, new(FramedIPv6Route))
	fct.store(IeTimeStamp, new(TimeStamp))
	fct.store(IeAveragingWindow, new(AveragingWindow))
	fct.store(IePagingPolicyIndicator, new(PagingPolicyIndicator))
	fct.store(IeAPNDNN, new(APNDNN))
	fct.store(IeThirdGPPInterfaceType, new(ThirdGPPInterfaceType))
	fct.store(IePFCPSRReqFlags, new(PFCPSRReqFlags))
	fct.store(IePFCPAUReqFlags, new(PFCPAUReqFlags))
	fct.store(IeActivationTime, new(ActivationTime))
	fct.store(IeDeactivationTime, new(DeactivationTime))
	fct.store(IeCreateMAR, new(CreateMAR))
	fct.store(IeThirdGPPAccessForwardingActionInformation, new(ThirdGPPAccessForwardingActionInformation))
	fct.store(IeNon3GPPAccessForwardingActionInformation, new(Non3GPPAccessForwardingActionInformation))
	fct.store(IeRemoveMAR, new(RemoveMAR))
	fct.store(IeUpdateMAR, new(UpdateMAR))
	fct.store(IeMARID, new(MARID))
	fct.store(IeSteeringFunctionality, new(SteeringFunctionality))
	fct.store(IeSteeringMode, new(SteeringMode))
	fct.store(IeWeight, new(Weight))
	fct.store(IePriority, new(Priority))
	fct.store(IeUpdate3GPPAccessForwardingActionInformation, new(Update3GPPAccessForwardingActionInformation))
	fct.store(IeUpdateNon3GPPAccessForwardingActionInformation, new(UpdateNon3GPPAccessForwardingActionInformation))
	fct.store(IeUEIPAddressPoolIdentity, new(UEIPAddressPoolIdentity))
	fct.store(IeAlternativeSMFIPAddress, new(AlternativeSMFIPAddress))
	fct.store(IePacketReplicationAndDetectionCarryOnInformation, new(PacketReplicationAndDetectionCarryOnInformation))
	fct.store(IeSMFSetID, new(SMFSetID))
	fct.store(IeQuotaValidityTime, new(QuotaValidityTime))
	fct.store(IeNumberOfReports, new(NumberOfReports))
	fct.store(IePFCPSessionRetentionInformationWithinPFCPAssociationSetupRequest, new(PFCPSessionRetentionInformationWithinPFCPAssociationSetupRequest))
	fct.store(IePFCPASRspFlags, new(PFCPASRspFlags))
	fct.store(IeCPPFCPEntityIPAddress, new(CPPFCPEntityIPAddress))
	fct.store(IePFCPSEReqFlags, new(PFCPSEReqFlags))
	fct.store(IeUserPlanePathRecoveryReport, new(UserPlanePathRecoveryReport))
	fct.store(IeIPMulticastAddressingInfoWithinPFCPSessionEstablishmentRequest, new(IPMulticastAddressingInfoWithinPFCPSessionEstablishmentRequest))
	fct.store(IeJoinIPMulticastInformationIEWithinUsageReport, new(JoinIPMulticastInformationIEWithinUsageReport))
	fct.store(IeLeaveIPMulticastInformationIEWithinUsageReport, new(LeaveIPMulticastInformationIEWithinUsageReport))
	fct.store(IeIPMulticastAddress, new(IPMulticastAddress))
	fct.store(IeSourceIPAddress, new(SourceIPAddress))
	fct.store(IePacketRateStatus, new(PacketRateStatus))
	fct.store(IeCreateBridgeInfoForTSC, new(CreateBridgeInfoForTSC))
	fct.store(IeCreatedBridgeInfoForTSC, new(CreatedBridgeInfoForTSC))
	fct.store(IeDSTTPortNumber, new(DSTTPortNumber))
	fct.store(IeNWTTPortNumber, new(NWTTPortNumber))
	fct.store(IeFiveGSUserPlaneNode, new(FiveGSUserPlaneNode))
	fct.store(IeTSCManagementInformationIEWithinPFCPSessionModificationRequest, new(TSCManagementInformationIEWithinPFCPSessionModificationRequest))
	fct.store(IeTSCManagementInformationIEWithinPFCPSessionModificationResponse, new(TSCManagementInformationIEWithinPFCPSessionModificationResponse))
	fct.store(IeTSCManagementInformationIEWithinPFCPSessionReportRequest, new(TSCManagementInformationIEWithinPFCPSessionReportRequest))
	fct.store(IePortManagementInformationContainer, new(PortManagementInformationContainer))
	fct.store(IeClockDriftControlInformation, new(ClockDriftControlInformation))
	fct.store(IeRequestedClockDriftInformation, new(RequestedClockDriftInformation))
	fct.store(IeClockDriftReport, new(ClockDriftReport))
	fct.store(IeTimeDomainNumber, new(TimeDomainNumber))
	fct.store(IeTimeOffsetThreshold, new(TimeOffsetThreshold))
	fct.store(IeCumulativeRateRatioThreshold, new(CumulativeRateRatioThreshold))
	fct.store(IeTimeOffsetMeasurement, new(TimeOffsetMeasurement))
	fct.store(IeCumulativeRateRatioMeasurement, new(CumulativeRateRatioMeasurement))
	fct.store(IeRemoveSRR, new(RemoveSRR))
	fct.store(IeCreateSRR, new(CreateSRR))
	fct.store(IeUpdateSRR, new(UpdateSRR))
	fct.store(IeSessionReport, new(SessionReport))
	fct.store(IeSRRID, new(SRRID))
	fct.store(IeAccessAvailabilityControlInformation, new(AccessAvailabilityControlInformation))
	fct.store(IeRequestedAccessAvailabilityInformation, new(RequestedAccessAvailabilityInformation))
	fct.store(IeAccessAvailabilityReport, new(AccessAvailabilityReport))
	fct.store(IeAccessAvailabilityInformation, new(AccessAvailabilityInformation))
	fct.store(IeProvideATSSSControlInformation, new(ProvideATSSSControlInformation))
	fct.store(IeATSSSControlParameters, new(ATSSSControlParameters))
	fct.store(IeMPTCPControlInformation, new(MPTCPControlInformation))
	fct.store(IeATSSSLLControlInformation, new(ATSSSLLControlInformation))
	fct.store(IePMFControlInformation, new(PMFControlInformation))
	fct.store(IeMPTCPParameters, new(MPTCPParameters))
	fct.store(IeATSSSLLParameters, new(ATSSSLLParameters))
	fct.store(IePMFParameters, new(PMFParameters))
	fct.store(IeMPTCPAddressInformation, new(MPTCPAddressInformation))
	fct.store(IeUELinkSpecificIPAddress, new(UELinkSpecificIPAddress))
	fct.store(IePMFAddressInformation, new(PMFAddressInformation))
	fct.store(IeATSSSLLInformation, new(ATSSSLLInformation))
	fct.store(IeDataNetworkAccessIdentifier, new(DataNetworkAccessIdentifier))
	fct.store(IeUEIPAddressPoolInformation, new(UEIPAddressPoolInformation))
	fct.store(IeAveragePacketDelay, new(AveragePacketDelay))
	fct.store(IeMinimumPacketDelay, new(MinimumPacketDelay))
	fct.store(IeMaximumPacketDelay, new(MaximumPacketDelay))
	fct.store(IeQoSReportTrigger, new(QoSReportTrigger))
	fct.store(IeGTPUPathQoSControlInformation, new(GTPUPathQoSControlInformation))
	fct.store(IeGTPUPathQoSReportPFCPNodeReportRequest, new(GTPUPathQoSReportPFCPNodeReportRequest))
	fct.store(IeQoSInformationInGTPUPathQoSReport, new(QoSInformationInGTPUPathQoSReport))
	fct.store(IeGTPUPathInterfaceType, new(GTPUPathInterfaceType))
	fct.store(IeQoSMonitoringPerQoSFlowControlInformation, new(QoSMonitoringPerQoSFlowControlInformation))
	fct.store(IeRequestedQoSMonitoring, new(RequestedQoSMonitoring))
	fct.store(IeReportingFrequency, new(ReportingFrequency))
	fct.store(IePacketDelayThresholds, new(PacketDelayThresholds))
	fct.store(IeMinimumWaitTime, new(MinimumWaitTime))
	fct.store(IeQoSMonitoringReport, new(QoSMonitoringReport))
	fct.store(IeQoSMonitoringMeasurement, new(QoSMonitoringMeasurement))
	fct.store(IeMTEDTControlInformation, new(MTEDTControlInformation))
	fct.store(IeDLDataPacketsSize, new(DLDataPacketsSize))
	fct.store(IeQERControlIndications, new(QERControlIndications))
	fct.store(IePacketRateStatusReport, new(PacketRateStatusReport))
	fct.store(IeNFInstanceID, new(NFInstanceID))
	fct.store(IeEthernetContextInformation, new(EthernetContextInformation))
	fct.store(IeRedundantTransmissionParameters, new(RedundantTransmissionParameters))
	fct.store(IeUpdatedPDR, new(UpdatedPDR))
	fct.store(IeSNSSAI, new(SNSSAI))
	fct.store(IeIPVersion, new(IPVersion))
	fct.store(IePFCPASReqFlags, new(PFCPASReqFlags))
	fct.store(IeDataStatus, new(DataStatus))
	fct.store(IeProvideRDSConfigurationInformation, new(ProvideRDSConfigurationInformation))
	fct.store(IeRDSConfigurationInformation, new(RDSConfigurationInformation))
	fct.store(IeQueryPacketRateStatusIEWithinPFCPSessionModificationRequest, new(QueryPacketRateStatusIEWithinPFCPSessionModificationRequest))
	fct.store(IePacketRateStatusReportIEWithinPFCPSessionModificationResponse, new(PacketRateStatusReportIEWithinPFCPSessionModificationResponse))
	fct.store(IeMPTCPApplicableIndication, new(MPTCPApplicableIndication))
	fct.store(IeBridgeManagementInformationContainer, new(BridgeManagementInformationContainer))
	fct.store(IeUEIPAddressUsageInformation, new(UEIPAddressUsageInformation))
	fct.store(IeNumberOfUEIPAddresses, new(NumberOfUEIPAddresses))
	fct.store(IeValidityTimer, new(ValidityTimer))
	fct.store(IeRedundantTransmissionForwardingParameters, new(RedundantTransmissionForwardingParameters))
	fct.store(IeTransportDelayReporting, new(TransportDelayReporting))
	fct.store(IePartialFailureInformation, new(PartialFailureInformation))
	fct.store(IeSpare, new(Spare))
	fct.store(IeOffendingIEInformation, new(OffendingIEInformation))
	fct.store(IeRATType, new(RATType))
	fct.store(IeL2TPTunnelInformation, new(L2TPTunnelInformation))
	fct.store(IeL2TPSessionInformation, new(L2TPSessionInformation))
	fct.store(IeL2TPUserAuthenticationIE, new(L2TPUserAuthenticationIE))
	fct.store(IeCreatedL2TPSession, new(CreatedL2TPSession))
	fct.store(IeLNSAddress, new(LNSAddress))
	fct.store(IeTunnelPreference, new(TunnelPreference))
	fct.store(IeCallingNumber, new(CallingNumber))
	fct.store(IeCalledNumber, new(CalledNumber))
	fct.store(IeL2TPSessionIndications, new(L2TPSessionIndications))
	fct.store(IeDNSServerAddress, new(DNSServerAddress))
	fct.store(IeNBNSServerAddress, new(NBNSServerAddress))
	fct.store(IeMaximumReceiveUnit, new(MaximumReceiveUnit))
	fct.store(IeThresholds, new(Thresholds))
	fct.store(IeSteeringModeIndicator, new(SteeringModeIndicator))
	fct.store(IePFCPSessionChangeInfo, new(PFCPSessionChangeInfo))
	fct.store(IeGroupID, new(GroupID))
	fct.store(IeCPIPAddress, new(CPIPAddress))
	fct.store(IeIPAddressAndPortNumberReplacement, new(IPAddressAndPortNumberReplacement))
	fct.store(IeDNSQueryFilter, new(DNSQueryFilter))
	fct.store(IeDirectReportingInformation, new(DirectReportingInformation))
	fct.store(IeEventNotificationURI, new(EventNotificationURI))
	fct.store(IeNotificationCorrelationID, new(NotificationCorrelationID))
	fct.store(IeReportingFlags, new(ReportingFlags))
	fct.store(IePredefinedRulesName, new(PredefinedRulesName))
	fct.store(IeMBSSessionN4mbControlInformation, new(MBSSessionN4mbControlInformation))
	fct.store(IeMBSMulticastParameters, new(MBSMulticastParameters))
	fct.store(IeAddMBSUniCastParameters, new(AddMBSUniCastParameters))
	fct.store(IeMBSSessionN4mbInformation, new(MBSSessionN4mbInformation))
	fct.store(IeRemoveMBSUniCastParameters, new(RemoveMBSUniCastParameters))
	fct.store(IeMBSSessionIdentifier, new(MBSSessionIdentifier))
	fct.store(IeMulticastTransportInformation, new(MulticastTransportInformation))
	fct.store(IeMBSN4mbReqFlags, new(MBSN4mbReqFlags))
	fct.store(IeLocalIngressTunnel, new(LocalIngressTunnel))
	fct.store(IeMBSUniCastParametersID, new(MBSUniCastParametersID))
	fct.store(IeMBSSessionN4ControlInformation, new(MBSSessionN4ControlInformation))
	fct.store(IeMBSSessionN4Information, new(MBSSessionN4Information))
	fct.store(IeMBSN4RespFlags, new(MBSN4RespFlags))
	fct.store(IeTunnelPassword, new(TunnelPassword))
	fct.store(IeAreaSessionID, new(AreaSessionID))
	fct.store(IePeerUPRestartReport, new(PeerUPRestartReport))
	fct.store(IeDSCPtoPPIControlInformation, new(DSCPtoPPIControlInformation))
	fct.store(IeDSCPtoPPIMappingInformation, new(DSCPtoPPIMappingInformation))
	fct.store(IePFCPSDRspFlags, new(PFCPSDRspFlags))
	fct.store(MsgHeartbeatRequest, new(HeartbeatRequest))
	fct.store(MsgHeartbeatResponse, new(HeartbeatResponse))
	fct.store(MsgPFDManagementRequest, new(PFDManagementRequest))
	fct.store(MsgPFDManagementResponse, new(PFDManagementResponse))
	fct.store(MsgAssociationSetupRequest, new(AssociationSetupRequest))
	fct.store(MsgAssociationSetupResponse, new(AssociationSetupResponse))
	fct.store(MsgAssociationUpdateRequest, new(AssociationUpdateRequest))
	fct.store(MsgAssociationUpdateResponse, new(AssociationUpdateResponse))
	fct.store(MsgAssociationReleaseRequest, new(AssociationReleaseRequest))
	fct.store(MsgAssociationReleaseResponse, new(AssociationReleaseResponse))
	fct.store(MsgVersionNotSupportedResponse, new(VersionNotSupportedResponse))
	fct.store(MsgNodeReportRequest, new(NodeReportRequest))
	fct.store(MsgNodeReportResponse, new(NodeReportResponse))
	fct.store(MsgSessionSetDeletionRequest, new(SessionSetDeletionRequest))
	fct.store(MsgSessionSetDeletionResponse, new(SessionSetDeletionResponse))
	fct.store(MsgSessionSetModificationRequest, new(SessionSetModificationRequest))
	fct.store(MsgSessionSetModifcationResponse, new(SessionSetModifcationResponse))
	fct.store(MsgSessionEstablishmentRequest, new(SessionEstablishmentRequest))
	fct.store(MsgSessionEstablishmentResponse, new(SessionEstablishmentResponse))
	fct.store(MsgSessionModificationRequest, new(SessionModificationRequest))
	fct.store(MsgSessionModificationResponse, new(SessionModificationResponse))
	fct.store(MsgSessionDeletionRequest, new(SessionDeletionRequest))
	fct.store(MsgSessionDeletionResponse, new(SessionDeletionResponse))
	fct.store(MsgSessionReportRequest, new(SessionReportRequest))
	fct.store(MsgSessionReportResponse, new(SessionReportResponse))
}
