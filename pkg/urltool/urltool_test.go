package urltool

import "testing"

func TestGetBasePath(t *testing.T) {
	type args struct {
		targetUrl string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases. 编写测试用例
		{name: "基本示例", args: args{targetUrl: "https://www.liwenzhou.com/posts/Go/unit-test-1/"}, want: "unit-test-1", wantErr: false},
		{name: "相对路径url示例", args: args{targetUrl: "/xxx/1233/"}, want: "", wantErr: true},
		{name: "空字符串", args: args{targetUrl: ""}, want: "", wantErr: true},
		{name: "带query的url", args: args{targetUrl: "https://www.liwenzhou.com/posts/Go/unit-test-1/?a=1&b=2"}, want: "unit-test-1", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBasePath(tt.args.targetUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBasePath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetBasePath() got = %v, want %v", got, tt.want)
			}
		})
	}
}
