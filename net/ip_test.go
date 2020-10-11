package net

import (
	"bytes"
	gnet "net"
	"testing"
)

func TestParseIPv4Port(t *testing.T) {

	tests := []struct {
		name     string
		address  string
		wantIP   gnet.IP
		wantPort uint16
		wantErr  bool
	}{
		{
			name:     "Parse 127.0.0.1:1080",
			address:  "127.0.0.1:1080",
			wantIP:   []byte{127, 0, 0, 1},
			wantPort: 1080,
			wantErr:  false,
		},
		{
			name:     "Parse 127.0.0:1080",
			address:  "127.0.0:1080",
			wantIP:   []byte{},
			wantPort: 1080,
			wantErr:  true,
		},
		{
			name:     "Parse 127.0.0.1",
			address:  "127.0.0.1",
			wantIP:   []byte{},
			wantPort: 0,
			wantErr:  true,
		},
		{
			name:     "Parse 127.0.0.1:",
			address:  "127.0.0.1:",
			wantIP:   []byte{},
			wantPort: 0,
			wantErr:  true,
		},
		{
			name:     "Parse 9999.0.0.1:1080",
			address:  "9999.0.0.1:",
			wantIP:   []byte{},
			wantPort: 0,
			wantErr:  true,
		},
		{
			name:     "Parse 127.0.0.1:65535",
			address:  "127.0.0.1:65535",
			wantIP:   []byte{127, 0, 0, 1},
			wantPort: 65535,
			wantErr:  false,
		},
		{
			name:     "Parse 127.0.0.1:65536",
			address:  "127.0.0.1:65536",
			wantIP:   []byte{},
			wantPort: 0,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIP, gotPort, err := ParseIPv4Port(tt.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseIPv4Port() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if !bytes.Equal(gotIP, tt.wantIP) {
					t.Errorf("ParseIPv4Port() gotIp = %v, want %v", gotIP, tt.wantIP)
				}
				if gotPort != tt.wantPort {
					t.Errorf("ParseIPv4Port() gotPort = %v, want %v", gotPort, tt.wantPort)
				}
			}
		})
	}
}
