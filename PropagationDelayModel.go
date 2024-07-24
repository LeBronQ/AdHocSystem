package AdHocSystem

const (
	c = 299800000
)

func ConstantSpeedPropagationDelayModel(node1 Node, node2 Node) float64 {
	dis := Distance(node1, node2)
	delay := dis / c
	return delay
}
