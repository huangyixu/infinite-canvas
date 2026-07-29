package handler

import "testing"

func TestParseVideoTaskPayloadReadsConsoleVideoURL(t *testing.T) {
	tests := []struct {
		name    string
		payload string
	}{
		{name: "result url", payload: `{"status":"SUCCESS","result_url":"https://example.com/result.mp4"}`},
		{name: "nested content", payload: `{"status":"SUCCESS","data":{"content":{"video_url":"https://example.com/content.mp4"}}}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseVideoTaskPayload([]byte(tt.payload), "seedance-2.0")
			if got.Status != "completed" || got.VideoURL == "" {
				t.Fatalf("parsed task = %#v, want completed task with video URL", got)
			}
		})
	}
}
