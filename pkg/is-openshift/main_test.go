package is_openshift_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	is_openshift "github.com/iblancasa/gopenshift/pkg/is-openshift"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func TestIsOpenShift(t *testing.T) {
	for _, tt := range []struct {
		apiGroupList *metav1.APIGroupList
		expected     bool
	}{
		{
			&metav1.APIGroupList{},
			false,
		},
		{
			&metav1.APIGroupList{
				Groups: []metav1.APIGroup{
					{
						Name: "config.openshift.io",
					},
				},
			},
			true,
		},
	} {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			output, err := json.Marshal(tt.apiGroupList)
			require.NoError(t, err)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, err = w.Write(output)
			require.NoError(t, err)
		}))
		defer server.Close()

		// test
		routes, err := is_openshift.IsOpenShift(&rest.Config{Host: server.URL})

		// verify
		assert.NoError(t, err)
		assert.Equal(t, tt.expected, routes)
	}
}