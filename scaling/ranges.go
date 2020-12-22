package scaling

const (
	// DefaultMinReplicas is the minimal amount of replicas for a service.
	DefaultMinReplicas = 1

	// DefaultMaxReplicas is the amount of replicas a service will auto-scale up to.
	DefaultMaxReplicas = 5

	// DefaultScalingFactor is the defining proportion for the scaling increments.
	DefaultScalingFactor = 5

	// MinScaleLabel label indicating min scale for a function
	MinScaleLabel = "com.openfx.scale.min"

	// MaxScaleLabel label indicating max scale for a function
	MaxScaleLabel = "com.openfx.scale.max"

	// ScalingFactorLabel label indicates the scaling factor for a function
	ScalingFactorLabel = "com.openfx.scale.factor"
)
