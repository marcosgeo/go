package link

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func Test_dfs(t *testing.T) {
	htmlStr := strings.NewReader("<p>This is an HTML paragraph.</a>")
	htmlNode, _ := html.Parse(htmlStr)
	type args struct {
		n       *html.Node
		padding string
		want    error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Should return no errors",
			args: args{
				n:       htmlNode,
				padding: "  ",
				want:    nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := dfs(tt.args.n, tt.args.padding)
			if err != tt.args.want {
				t.Errorf("Fail. want: %v, got%v", tt.args.want, err)
			}
		})
	}
}
