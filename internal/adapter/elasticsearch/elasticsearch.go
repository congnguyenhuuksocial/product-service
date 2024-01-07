package elasticsearch

import (
	"crypto/tls"
	"github.com/opensearch-project/opensearch-go"
	"go.uber.org/zap"
	"net/http"
	"product-service/pkg/config"
)

func NewSearchClient(conf *config.Config, log *zap.Logger) *opensearch.Client {
	search, err := opensearch.NewClient(opensearch.Config{
		Addresses: []string{conf.Search.Url},
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Username:      conf.Search.Username,
		Password:      conf.Search.Password,
		MaxRetries:    conf.Search.MaxRetries,
		RetryOnStatus: []int{502, 503, 504, 429},
	})

	if err != nil {
		log.Fatal("Failed to create opensearch client", zap.Error(err))
	}
	return search
}
