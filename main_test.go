package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ob-algdatii-ss19/leistungsnachweis-ateam/backend"
)

func Test_jsonInterfaceHandler(t *testing.T) {

	sut01 := backend.GuiRequestData{
		backend.Settings{Algorithm: backend.BASIC_GREEDY},
		backend.Intersection{
			Left: backend.Intersection_part{
				LeftLane:     true,
				StraightLane: true,
				RightLane:    true,
				Pedestrian:   backend.NORMAL,
			},
			Buttom: backend.Intersection_part{
				LeftLane:     true,
				StraightLane: true,
				RightLane:    true,
				Pedestrian:   backend.NORMAL,
			},
			Right: backend.Intersection_part{
				LeftLane:     true,
				StraightLane: true,
				RightLane:    true,
				Pedestrian:   backend.NORMAL,
			},
			Top: backend.Intersection_part{
				LeftLane:     true,
				StraightLane: true,
				RightLane:    true,
				Pedestrian:   backend.NORMAL,
			},
		},
	}

	data, _ := json.Marshal(sut01)

	type args struct {
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"Sending in correct JSON and receiving correct JSON response",
			args{
				httptest.NewRequest("GET", "/json", strings.NewReader(string(data)))},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			jsonInterfaceHandler(w, tt.args.req)

			got := w.Result()
			//gotBody, _ := ioutil.ReadAll(got.Body)

			decoder := json.NewDecoder(got.Body)
			var receivedData backend.JsonResponse
			err := decoder.Decode(&receivedData)

			if err != nil {
				t.Errorf("jsonInterfaceHandler() = nil")
			}

			if receivedData.TrafficLightPhases == nil {
				t.Errorf("jsonInterfaceHandler() = %v, but want object with TrafficLightPhases", receivedData)
			}

			if receivedData.ReceivedDataSuccessful == false {
				t.Errorf("jsonInterfaceHandler() = %v, but want object with ReceivedDataSuccessful == true", receivedData)
			}
		})
	}
}

func Test_viewHandler(t *testing.T) {
	type args struct {
		r *http.Request
	}
	type want struct {
		statusCode  int
		contentType string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			"Deliver index.html for default requests",
			args{httptest.NewRequest("GET", "/", nil)},
			want{200, "text/html; charset=utf-8"},
		},
		{
			"Deliver css file",
			args{httptest.NewRequest("GET", "/styles/styles.css", nil)},
			want{200, "text/css; charset=utf-8"},
		},
		{
			"Deliver specific html file",
			args{httptest.NewRequest("GET", "/intersection.html", nil)},
			want{200, "text/html; charset=utf-8"},
		},
		{
			"Deliver specific html file",
			args{httptest.NewRequest("GET", "/doesNotExist.html", nil)},
			want{404, ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			viewHandler(w, tt.args.r)
			got := w.Result()

			if w.Code != tt.want.statusCode {
				t.Errorf("viewHandler() = %v,  want %v", w.Code, tt.want.statusCode)
			}

			if string(got.Header.Get("Content-Type")) != tt.want.contentType {
				t.Errorf("viewHandler() = %v,  want %v", string(got.Header.Get("Content-Type")), tt.want.contentType)
			}
		})
	}
}
