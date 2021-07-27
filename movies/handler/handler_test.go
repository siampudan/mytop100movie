package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mocks "github.com/siampudan/mytop100movie/movies/mock"
	. "github.com/smartystreets/goconvey/convey"
)

func PrepareTestContext(t *testing.T, r *gin.Engine, wantResp interface{}, method, path string, payload io.Reader) (*http.Response, []byte, []byte) {
	ts := httptest.NewServer(r)
	defer ts.Close()

	pathReq := ts.URL + path

	client := &http.Client{}
	req, err := http.NewRequest(method, pathReq, payload)
	if err != nil {
		fmt.Println(err)
		return nil, nil, nil
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "mock-test-user-agent")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, nil, nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	want, err := json.Marshal(wantResp)
	if err != nil {
		t.Fatal(err)
	}

	return res, body, want
}

func TestGetMovieDetails(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	r := gin.New()
	mockMovieUsecase := mocks.NewMockMovieUseCase(mockCtrl)
	MoviesRoute(mockMovieUsecase, r.Group("v1"))

	result := map[string]interface{}{
		"id":       1,
		"name":     "Avengger",
		"category": "Actions",
	}
	wantResp := map[string]interface{}{
		"data": result,
	}

	Convey("When select movie detail", t, func() {
		mockMovieUsecase.EXPECT().GetMovieDetail(gomock.Any()).Return(result, nil).Times(1)
		res, body, want := PrepareTestContext(t, r, wantResp, "GET", "/v1/movies/1", nil)
		Convey("It should return success and data", func() {
			So(res.StatusCode, ShouldEqual, http.StatusOK)
			So(string(body), ShouldResemble, string(want))
		})
	})

}
