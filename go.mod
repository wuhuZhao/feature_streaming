module github.com/wuhuZhao/feature_streaming

go 1.20

require (
	github.com/bytedance/gopkg v0.0.0-20230324090325-a00d8057bef9
	github.com/segmentio/kafka-go v0.4.39
	github.com/wuhuZhao/feature_extractor v0.0.0-20230319054137-74bc86bac8e0
)

require (
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
)

replace github.com/wuhuZhao/feature_extractor v0.0.0-20230319054137-74bc86bac8e0 => ../feature_extractor
