package github

import "testing"


func TestParseUrl(t *testing.T){
	client := NewClient(nil, "https://github.com", "test")

	tests := []struct{
		name string
		rawUrl string
		wantOwner string
		wantName string
		wantErr bool
	}{
		{
			name: "valid",
			rawUrl: "https://github.com/gin-gonic/gin",
			wantOwner: "gin-gonic",
			wantName: "gin",
		},
		{
			name: ".git suffix",
			rawUrl:  "https://github.com/gin-gonic/gin.git",
			wantOwner: "gin-gonic",
			wantName: "gin",
		},
		{
			name: "with slash",
			rawUrl:  "https://github.com/gin-gonic/gin/",
			wantOwner: "gin-gonic",
			wantName: "gin",
		},
		{
			name: "not github",
			rawUrl: "https://gitlab.com/owner/repo",
			wantErr: true,
		},
		{
			name: "missing",
			rawUrl: "https://github.com/gin-gonic",
			wantErr: true,
		},
		{
			name: "invalid url",
			rawUrl: "fklsp-d//:fpsdk",
			wantErr: true,
		},
	}

	for _, v := range tests{
		t.Run(v.name, func(t *testing.T){
			res, err := client.ParseURL(v.rawUrl)
			if v.wantErr{
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return
			}
			if err != nil{
				t.Errorf("unexpected errror: %v", err)
			}
			if res.Owner != v.wantOwner{
				t.Errorf("owner got  %v , want %v", res.Owner, v.wantOwner, )
			}
			if res.Name != v.wantName{
				t.Errorf("name got %v, want %v", res.Name, v.wantName)
			}
		})
	}
}