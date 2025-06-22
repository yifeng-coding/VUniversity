package util

import "testing"

func TestSensitiveFilter_Filter(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{
			name: "无敏感词",
			text: "这是一个普通的文本，没有敏感词。",
			want: "这是一个普通的文本，没有敏感词。",
		},
		{
			name: "包含单个敏感词",
			text: "这是一个包含敏感词的文本，敏感词是海洛因。",
			want: "这是一个包含敏感词的文本，敏感词是***。",
		},
		{
			name: "包含多个敏感词",
			text: "这是一个包含多个敏感词的文本，敏感词是海洛因、法轮功和日本鬼子。",
			want: "这是一个包含多个敏感词的文本，敏感词是***、***和***。",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSensitiveFilter().Filter(tt.text); got != tt.want {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
