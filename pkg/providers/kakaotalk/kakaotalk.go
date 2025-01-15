package kakaotalk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/multierr"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/notify/pkg/utils"
	sliceutil "github.com/projectdiscovery/utils/slice"
)

type Provider struct {
	Kakaotalk []*Options `yaml:"kakaotalk,omitempty"`
	counter   int
}

type Options struct {
	ID                    string `yaml:"id,omitempty"`
	KakaotalkClientId     string `yaml:"kakaotalk_client_id,omitempty"`
	KakaotalkRedirectUri  string `yaml:"kakaotalk_redirect_uri,omitempty"`
	KakaotalkFormat       string `yaml:"kakaotalk_format,omitempty"`
}

func New(options []*Options, ids []string) (*Provider, error) {
	provider := &Provider{}

	for _, o := range options {
		if len(ids) == 0 || sliceutil.Contains(ids, o.ID) {
			provider.Kakaotalk = append(provider.Kakaotalk, o)
		}
	}

	provider.counter = 0
	return provider, nil
}