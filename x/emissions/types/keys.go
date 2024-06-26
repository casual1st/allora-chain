package types

import "cosmossdk.io/collections"

const (
	ModuleName                                 = "emissions"
	StoreKey                                   = "emissions"
	AlloraStakingAccountName                   = "allorastaking"
	AlloraRewardsAccountName                   = "allorarewards"
	AlloraPendingRewardForDelegatorAccountName = "allorapendingrewards"
)

var (
	ParamsKey                                   = collections.NewPrefix(0)
	TotalStakeKey                               = collections.NewPrefix(1)
	TopicStakeKey                               = collections.NewPrefix(2)
	RewardsKey                                  = collections.NewPrefix(3)
	NextTopicIdKey                              = collections.NewPrefix(4)
	TopicsKey                                   = collections.NewPrefix(5)
	TopicWorkersKey                             = collections.NewPrefix(6)
	TopicReputersKey                            = collections.NewPrefix(7)
	DelegatorStakeKey                           = collections.NewPrefix(8)
	DelegateStakePlacementKey                   = collections.NewPrefix(9)
	TargetStakeKey                              = collections.NewPrefix(10)
	InferencesKey                               = collections.NewPrefix(11)
	ForecastsKey                                = collections.NewPrefix(12)
	WorkerNodesKey                              = collections.NewPrefix(13)
	ReputerNodesKey                             = collections.NewPrefix(14)
	LatestInferencesTsKey                       = collections.NewPrefix(15)
	ActiveTopicsKey                             = collections.NewPrefix(16)
	AllInferencesKey                            = collections.NewPrefix(17)
	AllForecastsKey                             = collections.NewPrefix(18)
	AllLossBundlesKey                           = collections.NewPrefix(19)
	StakeRemovalKey                             = collections.NewPrefix(20)
	StakeByReputerAndTopicId                    = collections.NewPrefix(21)
	DelegateStakeRemovalKey                     = collections.NewPrefix(22)
	AllTopicStakeSumKey                         = collections.NewPrefix(23)
	AddressTopicsKey                            = collections.NewPrefix(24)
	WhitelistAdminsKey                          = collections.NewPrefix(24)
	ChurnableTopicsKey                          = collections.NewPrefix(25)
	RewardableTopicsKey                         = collections.NewPrefix(26)
	NetworkLossBundlesKey                       = collections.NewPrefix(27)
	NetworkRegretsKey                           = collections.NewPrefix(28)
	StakeByReputerAndTopicIdKey                 = collections.NewPrefix(29)
	ReputerScoresKey                            = collections.NewPrefix(30)
	InferenceScoresKey                          = collections.NewPrefix(31)
	ForecastScoresKey                           = collections.NewPrefix(32)
	ReputerListeningCoefficientKey              = collections.NewPrefix(33)
	InfererNetworkRegretsKey                    = collections.NewPrefix(34)
	ForecasterNetworkRegretsKey                 = collections.NewPrefix(35)
	OneInForecasterNetworkRegretsKey            = collections.NewPrefix(36)
	OneInForecasterSelfNetworkRegretsKey        = collections.NewPrefix(37)
	UnfulfilledWorkerNoncesKey                  = collections.NewPrefix(38)
	UnfulfilledReputerNoncesKey                 = collections.NewPrefix(39)
	FeeRevenueEpochKey                          = collections.NewPrefix(40)
	TopicFeeRevenueKey                          = collections.NewPrefix(41)
	PreviousTopicWeightKey                      = collections.NewPrefix(42)
	PreviousReputerRewardFractionKey            = collections.NewPrefix(43)
	PreviousInferenceRewardFractionKey          = collections.NewPrefix(44)
	PreviousForecastRewardFractionKey           = collections.NewPrefix(45)
	LatestInfererScoresByWorkerKey              = collections.NewPrefix(46)
	LatestForecasterScoresByWorkerKey           = collections.NewPrefix(47)
	LatestReputerScoresByReputerKey             = collections.NewPrefix(48)
	TopicRewardNonceKey                         = collections.NewPrefix(49)
	DelegateRewardPerShare                      = collections.NewPrefix(50)
	PreviousPercentageRewardToStakedReputersKey = collections.NewPrefix(51)
	StakeRemovalsByBlockKey                     = collections.NewPrefix(52)
	DelegateStakeRemovalsByBlockKey             = collections.NewPrefix(53)
	StakeRemovalsByActorKey                     = collections.NewPrefix(54)
	DelegateStakeRemovalsByActorKey             = collections.NewPrefix(55)
)