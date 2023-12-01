package bot

const (
	FlowProductRecommendations = "product_recommendations"
	FlowReturnOrder            = "return_order"
	FlowReviewProduct          = "review_product"
)

type Service interface {
	ListenMessages() error
	SendMessage(userId interface{}, message string) error
	StartFlow(flow string, parameters map[string]interface{}) error
}

type Config struct {
	Type     string
	ApiToken string
}
