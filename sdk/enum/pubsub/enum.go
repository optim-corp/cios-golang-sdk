package ciossdk_pusub_enum

type (
	PackerFormat  string
	MessagingMode string
)

const (
	PayloadOnly PackerFormat = "payload_only"
	Json        PackerFormat = "json"
	MsgPack     PackerFormat = "msgpack"

	Publish   MessagingMode = "publish"
	Subscribe MessagingMode = "subscribe"
	PubSub    MessagingMode = "pubsub"
)
