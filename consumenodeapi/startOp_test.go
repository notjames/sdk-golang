package consumenodeapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/sdk-golang/model"
	"github.com/opspec-io/sdk-golang/util/vhttp"
	"io/ioutil"
	netHttp "net/http"
)

var _ = Describe("StartOp", func() {

	It("should call httpClient.Do() with expected args", func() {

		/* arrange */
		providedStartOpReq := model.StartOpReq{
			Args:   map[string]*model.Data{},
			PkgRef: "dummyPkgRef",
		}

		expectedReqBytes, _ := json.Marshal(providedStartOpReq)
		expectedResult := "dummyOpId"

		expectedHttpReq, _ := netHttp.NewRequest(
			"POST",
			fmt.Sprintf("http://%v/ops/starts", "localhost:42224"),
			bytes.NewBuffer(expectedReqBytes),
		)

		fakeHTTPClient := new(vhttp.Fake)
		fakeHTTPClient.DoReturns(&netHttp.Response{Body: ioutil.NopCloser(bytes.NewReader([]byte(expectedResult)))}, nil)

		objectUnderTest := consumeNodeApi{
			httpClient: fakeHTTPClient,
		}

		/* act */
		actualResult, _ := objectUnderTest.StartOp(providedStartOpReq)

		/* assert */
		Expect(expectedHttpReq).To(Equal(fakeHTTPClient.DoArgsForCall(0)))
		Expect(expectedResult).To(Equal(actualResult))

	})
})