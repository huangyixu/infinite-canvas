package handler

import (
	"testing"

	"github.com/tigerowo/infinite-canvas/model"
)

func TestResolveAIProxyPathKeepsConsoleVideoContractForSeedance(t *testing.T) {
	channel := model.ModelChannel{Protocol: "openai", BaseURL: "http://console.example/v1"}
	if got := resolveAIProxyPath(channel, "seedance-2.0", "/video/generations"); got != "/video/generations" {
		t.Fatalf("create path = %q", got)
	}
	if got := resolveAIProxyPath(channel, "seedance-2.0", "/video/generations/task-1"); got != "/video/generations/task-1" {
		t.Fatalf("query path = %q", got)
	}
}

func TestResolveAIProxyPathAdaptsNativeVideoChannels(t *testing.T) {
	tests := []struct {
		name      string
		channel   model.ModelChannel
		modelName string
		path      string
		want      string
	}{
		{name: "ark create", channel: model.ModelChannel{BaseURL: "https://ark.cn-beijing.volces.com/api/v3"}, modelName: "seedance-2.0", path: "/video/generations", want: "/contents/generations/tasks"},
		{name: "ark query", channel: model.ModelChannel{BaseURL: "https://ark.cn-beijing.volces.com/api/v3"}, modelName: "seedance-2.0", path: "/video/generations/task-1", want: "/contents/generations/tasks/task-1"},
		{name: "kie create", channel: model.ModelChannel{Protocol: "kie"}, modelName: "kling-3-0-video", path: "/video/generations", want: "/jobs/createTask"},
		{name: "kie query", channel: model.ModelChannel{Protocol: "kie"}, modelName: "kling-3-0-video", path: "/video/generations/task-1", want: "/jobs/recordInfo?taskId=task-1"},
		{name: "apimart create", channel: model.ModelChannel{Protocol: "apimart"}, modelName: "kling-v2-6", path: "/video/generations", want: "/videos/generations"},
		{name: "apimart query", channel: model.ModelChannel{Protocol: "apimart"}, modelName: "kling-v2-6", path: "/video/generations/task-1", want: "/tasks/task-1?language=zh"},
		{name: "agnes create", channel: model.ModelChannel{}, modelName: "Agnes-Video-V2.0", path: "/video/generations", want: "/videos"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resolveAIProxyPath(tt.channel, tt.modelName, tt.path); got != tt.want {
				t.Fatalf("path = %q, want %q", got, tt.want)
			}
		})
	}
}
