// returns data points from Win32_PerfRawData_DirectoryServices_DirectoryServices
// <add link to documentation here> - Win32_PerfRawData_DirectoryServices_DirectoryServices class
package collector

import (
	"log"

	"github.com/StackExchange/wmi"
	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	Factories["ad"] = NewADCollector
}

// A ADCollector is a Prometheus collector for WMI Win32_PerfRawData_DirectoryServices_DirectoryServices metrics
type ADCollector struct {
}

// NewADCollector ...
func NewADCollector() (Collector, error) {
	const subsystem = "ad"
	return &ADCollector{
		ABANRPersec: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "abanr_persec"),
			"(ABANRPersec)",
			nil,
			nil,
		),
	}, nil
}

// Collect sends the metric values for each metric
// to the provided prometheus Metric channel.
func (c *ADCollector) Collect(ch chan<- prometheus.Metric) error {
	if desc, err := c.collect(ch); err != nil {
		log.Println("[ERROR] failed collecting ad metrics:", desc, err)
		return err
	}
	return nil
}

type Win32_PerfRawData_DirectoryServices_DirectoryServices struct {
	Name string

	ABANRPersec                                                      uint32
	ABBrowsesPersec                                                  uint32
	ABClientSessions                                                 uint32
	ABMatchesPersec                                                  uint32
	ABPropertyReadsPersec                                            uint32
	ABProxyLookupsPersec                                             uint32
	ABSearchesPersec                                                 uint32
	ApproximatehighestDNT                                            uint32
	ATQEstimatedQueueDelay                                           uint32
	ATQOutstandingQueuedRequests                                     uint32
	ATQRequestLatency                                                uint32
	ATQThreadsLDAP                                                   uint32
	ATQThreadsOther                                                  uint32
	ATQThreadsTotal                                                  uint32
	BasesearchesPersec                                               uint32
	DatabaseaddsPersec                                               uint32
	DatabasedeletesPersec                                            uint32
	DatabasemodifysPersec                                            uint32
	DatabaserecyclesPersec                                           uint32
	DigestBindsPersec                                                uint32
	DRAHighestUSNCommittedHighpart                                   uint32
	DRAHighestUSNCommittedLowpart                                    uint32
	DRAHighestUSNIssuedHighpart                                      uint32
	DRAHighestUSNIssuedLowpart                                       uint32
	DRAInboundBytesCompressedBetweenSitesAfterCompressionPersec      uint32
	DRAInboundBytesCompressedBetweenSitesAfterCompressionSinceBoot   uint32
	DRAInboundBytesCompressedBetweenSitesBeforeCompressionPersec     uint32
	DRAInboundBytesCompressedBetweenSitesBeforeCompressionSinceBoot  uint32
	DRAInboundBytesNotCompressedWithinSitePersec                     uint32
	DRAInboundBytesNotCompressedWithinSiteSinceBoot                  uint32
	DRAInboundBytesTotalPersec                                       uint32
	DRAInboundBytesTotalSinceBoot                                    uint32
	DRAInboundFullSyncObjectsRemaining                               uint32
	DRAInboundLinkValueUpdatesRemaininginPacket                      uint32
	DRAInboundObjectsAppliedPersec                                   uint32
	DRAInboundObjectsFilteredPersec                                  uint32
	DRAInboundObjectsPersec                                          uint32
	DRAInboundObjectUpdatesRemaininginPacket                         uint32
	DRAInboundPropertiesAppliedPersec                                uint32
	DRAInboundPropertiesFilteredPersec                               uint32
	DRAInboundPropertiesTotalPersec                                  uint32
	DRAInboundTotalUpdatesRemaininginPacket                          uint32
	DRAInboundValuesDNsonlyPersec                                    uint32
	DRAInboundValuesTotalPersec                                      uint32
	DRAOutboundBytesCompressedBetweenSitesAfterCompressionPersec     uint32
	DRAOutboundBytesCompressedBetweenSitesAfterCompressionSinceBoot  uint32
	DRAOutboundBytesCompressedBetweenSitesBeforeCompressionPersec    uint32
	DRAOutboundBytesCompressedBetweenSitesBeforeCompressionSinceBoot uint32
	DRAOutboundBytesNotCompressedWithinSitePersec                    uint32
	DRAOutboundBytesNotCompressedWithinSiteSinceBoot                 uint32
	DRAOutboundBytesTotalPersec                                      uint32
	DRAOutboundBytesTotalSinceBoot                                   uint32
	DRAOutboundObjectsFilteredPersec                                 uint32
	DRAOutboundObjectsPersec                                         uint32
	DRAOutboundPropertiesPersec                                      uint32
	DRAOutboundValuesDNsonlyPersec                                   uint32
	DRAOutboundValuesTotalPersec                                     uint32
	DRAPendingReplicationOperations                                  uint32
	DRAPendingReplicationSynchronizations                            uint32
	DRASyncFailuresonSchemaMismatch                                  uint32
	DRASyncRequestsMade                                              uint32
	DRASyncRequestsSuccessful                                        uint32
	DRAThreadsGettingNCChanges                                       uint32
	DRAThreadsGettingNCChangesHoldingSemaphore                       uint32
	DSClientBindsPersec                                              uint32
	DSClientNameTranslationsPersec                                   uint32
	DSDirectoryReadsPersec                                           uint32
	DSDirectorySearchesPersec                                        uint32
	DSDirectoryWritesPersec                                          uint32
	DSMonitorListSize                                                uint32
	DSNameCachehitrate                                               uint32
	DSNotifyQueueSize                                                uint32
	DSPercentReadsfromDRA                                            uint32
	DSPercentReadsfromKCC                                            uint32
	DSPercentReadsfromLSA                                            uint32
	DSPercentReadsfromNSPI                                           uint32
	DSPercentReadsfromNTDSAPI                                        uint32
	DSPercentReadsfromSAM                                            uint32
	DSPercentReadsOther                                              uint32
	DSPercentSearchesfromDRA                                         uint32
	DSPercentSearchesfromKCC                                         uint32
	DSPercentSearchesfromLDAP                                        uint32
	DSPercentSearchesfromLSA                                         uint32
	DSPercentSearchesfromNSPI                                        uint32
	DSPercentSearchesfromNTDSAPI                                     uint32
	DSPercentSearchesfromSAM                                         uint32
	DSPercentSearchesOther                                           uint32
	DSPercentWritesfromDRA                                           uint32
	DSPercentWritesfromKCC                                           uint32
	DSPercentWritesfromLDAP                                          uint32
	DSPercentWritesfromLSA                                           uint32
	DSPercentWritesfromNSPI                                          uint32
	DSPercentWritesfromNTDSAPI                                       uint32
	DSPercentWritesfromSAM                                           uint32
	DSPercentWritesOther                                             uint32
	DSSearchsuboperationsPersec                                      uint32
	DSSecurityDescriptorPropagationsEvents                           uint32
	DSSecurityDescriptorPropagatorAverageExclusionTime               uint32
	DSSecurityDescriptorPropagatorRuntimeQueue                       uint32
	DSSecurityDescriptorsuboperationsPersec                          uint32
	DSServerBindsPersec                                              uint32
	DSServerNameTranslationsPersec                                   uint32
	DSThreadsinUse                                                   uint32
	ExternalBindsPersec                                              uint32
	FastBindsPersec                                                  uint32
	LDAPActiveThreads                                                uint32
	LDAPBindTime                                                     uint32
	LDAPClientSessions                                               uint32
	LDAPClosedConnectionsPersec                                      uint32
	LDAPNewConnectionsPersec                                         uint32
	LDAPNewSSLConnectionsPersec                                      uint32
	LDAPSearchesPersec                                               uint32
	LDAPSuccessfulBindsPersec                                        uint32
	LDAPUDPoperationsPersec                                          uint32
	LDAPWritesPersec                                                 uint32
	LinkValuesCleanedPersec                                          uint32
	NegotiatedBindsPersec                                            uint32
	NTLMBindsPersec                                                  uint32
	OnelevelsearchesPersec                                           uint32
	PhantomsCleanedPersec                                            uint32
	PhantomsVisitedPersec                                            uint32
	SAMAccountGroupEvaluationLatency                                 uint32
	SAMDisplayInformationQueriesPersec                               uint32
	SAMDomainLocalGroupMembershipEvaluationsPersec                   uint32
	SAMEnumerationsPersec                                            uint32
	SAMGCEvaluationsPersec                                           uint32
	SAMGlobalGroupMembershipEvaluationsPersec                        uint32
	SAMMachineCreationAttemptsPersec                                 uint32
	SAMMembershipChangesPersec                                       uint32
	SAMNonTransitiveMembershipEvaluationsPersec                      uint32
	SAMPasswordChangesPersec                                         uint32
	SAMResourceGroupEvaluationLatency                                uint32
	SAMSuccessfulComputerCreationsPersecIncludesallrequests          uint32
	SAMSuccessfulUserCreationsPersec                                 uint32
	SAMTransitiveMembershipEvaluationsPersec                         uint32
	SAMUniversalGroupMembershipEvaluationsPersec                     uint32
	SAMUserCreationAttemptsPersec                                    uint32
	SimpleBindsPersec                                                uint32
	SubtreesearchesPersec                                            uint32
	TombstonesGarbageCollectedPersec                                 uint32
	TombstonesVisitedPersec                                          uint32
	Transitiveoperationsmillisecondsrun                              uint32
	TransitiveoperationsPersec                                       uint32
	TransitivesuboperationsPersec                                    uint32
}

func (c *ADCollector) collect(ch chan<- prometheus.Metric) (*prometheus.Desc, error) {
	var dst []Win32_PerfRawData_DirectoryServices_DirectoryServices
	q := wmi.CreateQuery(&dst, "")
	if err := wmi.Query(q, &dst); err != nil {
		return nil, err
	}

	ch <- prometheus.MustNewConstMetric(
		c.ABANRPersec,
		prometheus.GaugeValue,
		float64(dst[0].ABANRPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.ABBrowsesPersec,
		prometheus.GaugeValue,
		float64(dst[0].ABBrowsesPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.ABClientSessions,
		prometheus.GaugeValue,
		float64(dst[0].ABClientSessions),
	)

	ch <- prometheus.MustNewConstMetric(
		c.ABMatchesPersec,
		prometheus.GaugeValue,
		float64(dst[0].ABMatchesPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.ABPropertyReadsPersec,
		prometheus.GaugeValue,
		float64(dst[0].ABPropertyReadsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.ABProxyLookupsPersec,
		prometheus.GaugeValue,
		float64(dst[0].ABProxyLookupsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.ABSearchesPersec,
		prometheus.GaugeValue,
		float64(dst[0].ABSearchesPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.ApproximatehighestDNT,
		prometheus.GaugeValue,
		float64(dst[0].ApproximatehighestDNT),
	)

	ch <- prometheus.MustNewConstMetric(
		c.ATQEstimatedQueueDelay,
		prometheus.GaugeValue,
		float64(dst[0].ATQEstimatedQueueDelay),
	)

	ch <- prometheus.MustNewConstMetric(
		c.ATQOutstandingQueuedRequests,
		prometheus.GaugeValue,
		float64(dst[0].ATQOutstandingQueuedRequests),
	)

	ch <- prometheus.MustNewConstMetric(
		c.ATQRequestLatency,
		prometheus.GaugeValue,
		float64(dst[0].ATQRequestLatency),
	)

	ch <- prometheus.MustNewConstMetric(
		c.ATQThreadsLDAP,
		prometheus.GaugeValue,
		float64(dst[0].ATQThreadsLDAP),
	)

	ch <- prometheus.MustNewConstMetric(
		c.ATQThreadsOther,
		prometheus.GaugeValue,
		float64(dst[0].ATQThreadsOther),
	)

	ch <- prometheus.MustNewConstMetric(
		c.ATQThreadsTotal,
		prometheus.GaugeValue,
		float64(dst[0].ATQThreadsTotal),
	)

	ch <- prometheus.MustNewConstMetric(
		c.BasesearchesPersec,
		prometheus.GaugeValue,
		float64(dst[0].BasesearchesPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DatabaseaddsPersec,
		prometheus.GaugeValue,
		float64(dst[0].DatabaseaddsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DatabasedeletesPersec,
		prometheus.GaugeValue,
		float64(dst[0].DatabasedeletesPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DatabasemodifysPersec,
		prometheus.GaugeValue,
		float64(dst[0].DatabasemodifysPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DatabaserecyclesPersec,
		prometheus.GaugeValue,
		float64(dst[0].DatabaserecyclesPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DigestBindsPersec,
		prometheus.GaugeValue,
		float64(dst[0].DigestBindsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAHighestUSNCommittedHighpart,
		prometheus.GaugeValue,
		float64(dst[0].DRAHighestUSNCommittedHighpart),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAHighestUSNCommittedLowpart,
		prometheus.GaugeValue,
		float64(dst[0].DRAHighestUSNCommittedLowpart),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAHighestUSNIssuedHighpart,
		prometheus.GaugeValue,
		float64(dst[0].DRAHighestUSNIssuedHighpart),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAHighestUSNIssuedLowpart,
		prometheus.GaugeValue,
		float64(dst[0].DRAHighestUSNIssuedLowpart),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAInboundBytesCompressedBetweenSitesAfterCompressionPersec,
		prometheus.GaugeValue,
		float64(dst[0].DRAInboundBytesCompressedBetweenSitesAfterCompressionPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAInboundBytesCompressedBetweenSitesAfterCompressionSinceBoot,
		prometheus.GaugeValue,
		float64(dst[0].DRAInboundBytesCompressedBetweenSitesAfterCompressionSinceBoot),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAInboundBytesCompressedBetweenSitesBeforeCompressionPersec,
		prometheus.GaugeValue,
		float64(dst[0].DRAInboundBytesCompressedBetweenSitesBeforeCompressionPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAInboundBytesCompressedBetweenSitesBeforeCompressionSinceBoot,
		prometheus.GaugeValue,
		float64(dst[0].DRAInboundBytesCompressedBetweenSitesBeforeCompressionSinceBoot),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAInboundBytesNotCompressedWithinSitePersec,
		prometheus.GaugeValue,
		float64(dst[0].DRAInboundBytesNotCompressedWithinSitePersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAInboundBytesNotCompressedWithinSiteSinceBoot,
		prometheus.GaugeValue,
		float64(dst[0].DRAInboundBytesNotCompressedWithinSiteSinceBoot),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAInboundBytesTotalPersec,
		prometheus.GaugeValue,
		float64(dst[0].DRAInboundBytesTotalPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAInboundBytesTotalSinceBoot,
		prometheus.GaugeValue,
		float64(dst[0].DRAInboundBytesTotalSinceBoot),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAInboundFullSyncObjectsRemaining,
		prometheus.GaugeValue,
		float64(dst[0].DRAInboundFullSyncObjectsRemaining),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAInboundLinkValueUpdatesRemaininginPacket,
		prometheus.GaugeValue,
		float64(dst[0].DRAInboundLinkValueUpdatesRemaininginPacket),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAInboundObjectsAppliedPersec,
		prometheus.GaugeValue,
		float64(dst[0].DRAInboundObjectsAppliedPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAInboundObjectsFilteredPersec,
		prometheus.GaugeValue,
		float64(dst[0].DRAInboundObjectsFilteredPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAInboundObjectsPersec,
		prometheus.GaugeValue,
		float64(dst[0].DRAInboundObjectsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAInboundObjectUpdatesRemaininginPacket,
		prometheus.GaugeValue,
		float64(dst[0].DRAInboundObjectUpdatesRemaininginPacket),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAInboundPropertiesAppliedPersec,
		prometheus.GaugeValue,
		float64(dst[0].DRAInboundPropertiesAppliedPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAInboundPropertiesFilteredPersec,
		prometheus.GaugeValue,
		float64(dst[0].DRAInboundPropertiesFilteredPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAInboundPropertiesTotalPersec,
		prometheus.GaugeValue,
		float64(dst[0].DRAInboundPropertiesTotalPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAInboundTotalUpdatesRemaininginPacket,
		prometheus.GaugeValue,
		float64(dst[0].DRAInboundTotalUpdatesRemaininginPacket),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAInboundValuesDNsonlyPersec,
		prometheus.GaugeValue,
		float64(dst[0].DRAInboundValuesDNsonlyPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAInboundValuesTotalPersec,
		prometheus.GaugeValue,
		float64(dst[0].DRAInboundValuesTotalPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAOutboundBytesCompressedBetweenSitesAfterCompressionPersec,
		prometheus.GaugeValue,
		float64(dst[0].DRAOutboundBytesCompressedBetweenSitesAfterCompressionPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAOutboundBytesCompressedBetweenSitesAfterCompressionSinceBoot,
		prometheus.GaugeValue,
		float64(dst[0].DRAOutboundBytesCompressedBetweenSitesAfterCompressionSinceBoot),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAOutboundBytesCompressedBetweenSitesBeforeCompressionPersec,
		prometheus.GaugeValue,
		float64(dst[0].DRAOutboundBytesCompressedBetweenSitesBeforeCompressionPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAOutboundBytesCompressedBetweenSitesBeforeCompressionSinceBoot,
		prometheus.GaugeValue,
		float64(dst[0].DRAOutboundBytesCompressedBetweenSitesBeforeCompressionSinceBoot),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAOutboundBytesNotCompressedWithinSitePersec,
		prometheus.GaugeValue,
		float64(dst[0].DRAOutboundBytesNotCompressedWithinSitePersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAOutboundBytesNotCompressedWithinSiteSinceBoot,
		prometheus.GaugeValue,
		float64(dst[0].DRAOutboundBytesNotCompressedWithinSiteSinceBoot),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAOutboundBytesTotalPersec,
		prometheus.GaugeValue,
		float64(dst[0].DRAOutboundBytesTotalPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAOutboundBytesTotalSinceBoot,
		prometheus.GaugeValue,
		float64(dst[0].DRAOutboundBytesTotalSinceBoot),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAOutboundObjectsFilteredPersec,
		prometheus.GaugeValue,
		float64(dst[0].DRAOutboundObjectsFilteredPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAOutboundObjectsPersec,
		prometheus.GaugeValue,
		float64(dst[0].DRAOutboundObjectsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAOutboundPropertiesPersec,
		prometheus.GaugeValue,
		float64(dst[0].DRAOutboundPropertiesPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAOutboundValuesDNsonlyPersec,
		prometheus.GaugeValue,
		float64(dst[0].DRAOutboundValuesDNsonlyPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAOutboundValuesTotalPersec,
		prometheus.GaugeValue,
		float64(dst[0].DRAOutboundValuesTotalPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAPendingReplicationOperations,
		prometheus.GaugeValue,
		float64(dst[0].DRAPendingReplicationOperations),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAPendingReplicationSynchronizations,
		prometheus.GaugeValue,
		float64(dst[0].DRAPendingReplicationSynchronizations),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRASyncFailuresonSchemaMismatch,
		prometheus.GaugeValue,
		float64(dst[0].DRASyncFailuresonSchemaMismatch),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRASyncRequestsMade,
		prometheus.GaugeValue,
		float64(dst[0].DRASyncRequestsMade),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRASyncRequestsSuccessful,
		prometheus.GaugeValue,
		float64(dst[0].DRASyncRequestsSuccessful),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAThreadsGettingNCChanges,
		prometheus.GaugeValue,
		float64(dst[0].DRAThreadsGettingNCChanges),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DRAThreadsGettingNCChangesHoldingSemaphore,
		prometheus.GaugeValue,
		float64(dst[0].DRAThreadsGettingNCChangesHoldingSemaphore),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSClientBindsPersec,
		prometheus.GaugeValue,
		float64(dst[0].DSClientBindsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSClientNameTranslationsPersec,
		prometheus.GaugeValue,
		float64(dst[0].DSClientNameTranslationsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSDirectoryReadsPersec,
		prometheus.GaugeValue,
		float64(dst[0].DSDirectoryReadsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSDirectorySearchesPersec,
		prometheus.GaugeValue,
		float64(dst[0].DSDirectorySearchesPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSDirectoryWritesPersec,
		prometheus.GaugeValue,
		float64(dst[0].DSDirectoryWritesPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSMonitorListSize,
		prometheus.GaugeValue,
		float64(dst[0].DSMonitorListSize),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSNameCachehitrate,
		prometheus.GaugeValue,
		float64(dst[0].DSNameCachehitrate),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSNotifyQueueSize,
		prometheus.GaugeValue,
		float64(dst[0].DSNotifyQueueSize),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSPercentReadsfromDRA,
		prometheus.GaugeValue,
		float64(dst[0].DSPercentReadsfromDRA),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSPercentReadsfromKCC,
		prometheus.GaugeValue,
		float64(dst[0].DSPercentReadsfromKCC),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSPercentReadsfromLSA,
		prometheus.GaugeValue,
		float64(dst[0].DSPercentReadsfromLSA),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSPercentReadsfromNSPI,
		prometheus.GaugeValue,
		float64(dst[0].DSPercentReadsfromNSPI),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSPercentReadsfromNTDSAPI,
		prometheus.GaugeValue,
		float64(dst[0].DSPercentReadsfromNTDSAPI),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSPercentReadsfromSAM,
		prometheus.GaugeValue,
		float64(dst[0].DSPercentReadsfromSAM),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSPercentReadsOther,
		prometheus.GaugeValue,
		float64(dst[0].DSPercentReadsOther),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSPercentSearchesfromDRA,
		prometheus.GaugeValue,
		float64(dst[0].DSPercentSearchesfromDRA),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSPercentSearchesfromKCC,
		prometheus.GaugeValue,
		float64(dst[0].DSPercentSearchesfromKCC),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSPercentSearchesfromLDAP,
		prometheus.GaugeValue,
		float64(dst[0].DSPercentSearchesfromLDAP),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSPercentSearchesfromLSA,
		prometheus.GaugeValue,
		float64(dst[0].DSPercentSearchesfromLSA),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSPercentSearchesfromNSPI,
		prometheus.GaugeValue,
		float64(dst[0].DSPercentSearchesfromNSPI),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSPercentSearchesfromNTDSAPI,
		prometheus.GaugeValue,
		float64(dst[0].DSPercentSearchesfromNTDSAPI),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSPercentSearchesfromSAM,
		prometheus.GaugeValue,
		float64(dst[0].DSPercentSearchesfromSAM),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSPercentSearchesOther,
		prometheus.GaugeValue,
		float64(dst[0].DSPercentSearchesOther),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSPercentWritesfromDRA,
		prometheus.GaugeValue,
		float64(dst[0].DSPercentWritesfromDRA),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSPercentWritesfromKCC,
		prometheus.GaugeValue,
		float64(dst[0].DSPercentWritesfromKCC),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSPercentWritesfromLDAP,
		prometheus.GaugeValue,
		float64(dst[0].DSPercentWritesfromLDAP),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSPercentWritesfromLSA,
		prometheus.GaugeValue,
		float64(dst[0].DSPercentWritesfromLSA),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSPercentWritesfromNSPI,
		prometheus.GaugeValue,
		float64(dst[0].DSPercentWritesfromNSPI),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSPercentWritesfromNTDSAPI,
		prometheus.GaugeValue,
		float64(dst[0].DSPercentWritesfromNTDSAPI),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSPercentWritesfromSAM,
		prometheus.GaugeValue,
		float64(dst[0].DSPercentWritesfromSAM),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSPercentWritesOther,
		prometheus.GaugeValue,
		float64(dst[0].DSPercentWritesOther),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSSearchsuboperationsPersec,
		prometheus.GaugeValue,
		float64(dst[0].DSSearchsuboperationsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSSecurityDescriptorPropagationsEvents,
		prometheus.GaugeValue,
		float64(dst[0].DSSecurityDescriptorPropagationsEvents),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSSecurityDescriptorPropagatorAverageExclusionTime,
		prometheus.GaugeValue,
		float64(dst[0].DSSecurityDescriptorPropagatorAverageExclusionTime),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSSecurityDescriptorPropagatorRuntimeQueue,
		prometheus.GaugeValue,
		float64(dst[0].DSSecurityDescriptorPropagatorRuntimeQueue),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSSecurityDescriptorsuboperationsPersec,
		prometheus.GaugeValue,
		float64(dst[0].DSSecurityDescriptorsuboperationsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSServerBindsPersec,
		prometheus.GaugeValue,
		float64(dst[0].DSServerBindsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSServerNameTranslationsPersec,
		prometheus.GaugeValue,
		float64(dst[0].DSServerNameTranslationsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DSThreadsinUse,
		prometheus.GaugeValue,
		float64(dst[0].DSThreadsinUse),
	)

	ch <- prometheus.MustNewConstMetric(
		c.ExternalBindsPersec,
		prometheus.GaugeValue,
		float64(dst[0].ExternalBindsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.FastBindsPersec,
		prometheus.GaugeValue,
		float64(dst[0].FastBindsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.LDAPActiveThreads,
		prometheus.GaugeValue,
		float64(dst[0].LDAPActiveThreads),
	)

	ch <- prometheus.MustNewConstMetric(
		c.LDAPBindTime,
		prometheus.GaugeValue,
		float64(dst[0].LDAPBindTime),
	)

	ch <- prometheus.MustNewConstMetric(
		c.LDAPClientSessions,
		prometheus.GaugeValue,
		float64(dst[0].LDAPClientSessions),
	)

	ch <- prometheus.MustNewConstMetric(
		c.LDAPClosedConnectionsPersec,
		prometheus.GaugeValue,
		float64(dst[0].LDAPClosedConnectionsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.LDAPNewConnectionsPersec,
		prometheus.GaugeValue,
		float64(dst[0].LDAPNewConnectionsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.LDAPNewSSLConnectionsPersec,
		prometheus.GaugeValue,
		float64(dst[0].LDAPNewSSLConnectionsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.LDAPSearchesPersec,
		prometheus.GaugeValue,
		float64(dst[0].LDAPSearchesPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.LDAPSuccessfulBindsPersec,
		prometheus.GaugeValue,
		float64(dst[0].LDAPSuccessfulBindsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.LDAPUDPoperationsPersec,
		prometheus.GaugeValue,
		float64(dst[0].LDAPUDPoperationsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.LDAPWritesPersec,
		prometheus.GaugeValue,
		float64(dst[0].LDAPWritesPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.LinkValuesCleanedPersec,
		prometheus.GaugeValue,
		float64(dst[0].LinkValuesCleanedPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.NegotiatedBindsPersec,
		prometheus.GaugeValue,
		float64(dst[0].NegotiatedBindsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.NTLMBindsPersec,
		prometheus.GaugeValue,
		float64(dst[0].NTLMBindsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.OnelevelsearchesPersec,
		prometheus.GaugeValue,
		float64(dst[0].OnelevelsearchesPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.PhantomsCleanedPersec,
		prometheus.GaugeValue,
		float64(dst[0].PhantomsCleanedPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.PhantomsVisitedPersec,
		prometheus.GaugeValue,
		float64(dst[0].PhantomsVisitedPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.SAMAccountGroupEvaluationLatency,
		prometheus.GaugeValue,
		float64(dst[0].SAMAccountGroupEvaluationLatency),
	)

	ch <- prometheus.MustNewConstMetric(
		c.SAMDisplayInformationQueriesPersec,
		prometheus.GaugeValue,
		float64(dst[0].SAMDisplayInformationQueriesPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.SAMDomainLocalGroupMembershipEvaluationsPersec,
		prometheus.GaugeValue,
		float64(dst[0].SAMDomainLocalGroupMembershipEvaluationsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.SAMEnumerationsPersec,
		prometheus.GaugeValue,
		float64(dst[0].SAMEnumerationsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.SAMGCEvaluationsPersec,
		prometheus.GaugeValue,
		float64(dst[0].SAMGCEvaluationsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.SAMGlobalGroupMembershipEvaluationsPersec,
		prometheus.GaugeValue,
		float64(dst[0].SAMGlobalGroupMembershipEvaluationsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.SAMMachineCreationAttemptsPersec,
		prometheus.GaugeValue,
		float64(dst[0].SAMMachineCreationAttemptsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.SAMMembershipChangesPersec,
		prometheus.GaugeValue,
		float64(dst[0].SAMMembershipChangesPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.SAMNonTransitiveMembershipEvaluationsPersec,
		prometheus.GaugeValue,
		float64(dst[0].SAMNonTransitiveMembershipEvaluationsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.SAMPasswordChangesPersec,
		prometheus.GaugeValue,
		float64(dst[0].SAMPasswordChangesPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.SAMResourceGroupEvaluationLatency,
		prometheus.GaugeValue,
		float64(dst[0].SAMResourceGroupEvaluationLatency),
	)

	ch <- prometheus.MustNewConstMetric(
		c.SAMSuccessfulComputerCreationsPersecIncludesallrequests,
		prometheus.GaugeValue,
		float64(dst[0].SAMSuccessfulComputerCreationsPersecIncludesallrequests),
	)

	ch <- prometheus.MustNewConstMetric(
		c.SAMSuccessfulUserCreationsPersec,
		prometheus.GaugeValue,
		float64(dst[0].SAMSuccessfulUserCreationsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.SAMTransitiveMembershipEvaluationsPersec,
		prometheus.GaugeValue,
		float64(dst[0].SAMTransitiveMembershipEvaluationsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.SAMUniversalGroupMembershipEvaluationsPersec,
		prometheus.GaugeValue,
		float64(dst[0].SAMUniversalGroupMembershipEvaluationsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.SAMUserCreationAttemptsPersec,
		prometheus.GaugeValue,
		float64(dst[0].SAMUserCreationAttemptsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.SimpleBindsPersec,
		prometheus.GaugeValue,
		float64(dst[0].SimpleBindsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.SubtreesearchesPersec,
		prometheus.GaugeValue,
		float64(dst[0].SubtreesearchesPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.TombstonesGarbageCollectedPersec,
		prometheus.GaugeValue,
		float64(dst[0].TombstonesGarbageCollectedPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.TombstonesVisitedPersec,
		prometheus.GaugeValue,
		float64(dst[0].TombstonesVisitedPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.Transitiveoperationsmillisecondsrun,
		prometheus.GaugeValue,
		float64(dst[0].Transitiveoperationsmillisecondsrun),
	)

	ch <- prometheus.MustNewConstMetric(
		c.TransitiveoperationsPersec,
		prometheus.GaugeValue,
		float64(dst[0].TransitiveoperationsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.TransitivesuboperationsPersec,
		prometheus.GaugeValue,
		float64(dst[0].TransitivesuboperationsPersec),
	)

	return nil, nil
}
