package typeoffns

import "testing"

func TestIsString(t *testing.T) {
	tests := []struct {
		name string
		str  interface{}
		want bool
	}{
		{
			name: "Valid string",
			str:  "Hello, World!",
			want: true,
		},
		{
			name: "Invalid string",
			str:  42,
			want: false,
		},
		{
			name: "Empty string",
			str:  "",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsString(tt.str); got != tt.want {
				t.Errorf("IsString(%v) = %v, want %v", tt.str, got, tt.want)
			}
		})
	}
}
func TestIsUint32(t *testing.T) {
	tests := []struct {
		name string
		num  interface{}
		want bool
	}{
		{
			name: "Valid uint32",
			num:  uint32(42),
			want: true,
		},
		{
			name: "Invalid uint32",
			num:  42,
			want: false,
		},
		{
			name: "Valid uint64",
			num:  uint64(42),
			want: false,
		},
		{
			name: "Valid int",
			num:  int(42),
			want: false,
		},
		{
			name: "Valid uint",
			num:  uint(42),
			want: false,
		},
		{
			name: "Valid bool",
			num:  true,
			want: false,
		},
		{
			name: "Valid map",
			num:  map[string]interface{}{},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUint32(tt.num); got != tt.want {
				t.Errorf("IsUint32(%v) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}
func TestIsUint64(t *testing.T) {
	tests := []struct {
		name string
		num  interface{}
		want bool
	}{
		{
			name: "Valid uint64",
			num:  uint64(42),
			want: true,
		},
		{
			name: "Invalid uint64",
			num:  42,
			want: false,
		},
		{
			name: "Valid uint32",
			num:  uint32(42),
			want: false,
		},
		{
			name: "Valid int",
			num:  int(42),
			want: false,
		},
		{
			name: "Valid uint",
			num:  uint(42),
			want: false,
		},
		{
			name: "Valid bool",
			num:  true,
			want: false,
		},
		{
			name: "Valid map",
			num:  map[string]interface{}{},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUint64(tt.num); got != tt.want {
				t.Errorf("IsUint64(%v) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}
func TestIsUint(t *testing.T) {
	tests := []struct {
		name string
		num  interface{}
		want bool
	}{
		{
			name: "Valid uint",
			num:  uint(42),
			want: true,
		},
		{
			name: "Invalid uint",
			num:  42,
			want: false,
		},
		{
			name: "Valid uint32",
			num:  uint32(42),
			want: false,
		},
		{
			name: "Valid uint64",
			num:  uint64(42),
			want: false,
		},
		{
			name: "Valid int",
			num:  int(42),
			want: false,
		},
		{
			name: "Valid bool",
			num:  true,
			want: false,
		},
		{
			name: "Valid map",
			num:  map[string]interface{}{},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUint(tt.num); got != tt.want {
				t.Errorf("IsUint(%v) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}
func TestIsBool(t *testing.T) {
	tests := []struct {
		name string
		b    interface{}
		want bool
	}{
		{
			name: "Valid bool",
			b:    true,
			want: true,
		},
		{
			name: "Invalid bool",
			b:    42,
			want: false,
		},
		{
			name: "Invalid bool",
			b:    "true",
			want: false,
		},
		{
			name: "Invalid bool",
			b:    nil,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBool(tt.b); got != tt.want {
				t.Errorf("IsBool(%v) = %v, want %v", tt.b, got, tt.want)
			}
		})
	}
}
func TestIsMap(t *testing.T) {
	tests := []struct {
		name string
		m    interface{}
		want bool
	}{
		{
			name: "Valid map",
			m:    map[string]interface{}{},
			want: true,
		},
		{
			name: "Invalid map",
			m:    42,
			want: false,
		},
		{
			name: "Invalid map",
			m:    "map",
			want: false,
		},
		{
			name: "Invalid map",
			m:    nil,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, got := IsMap(tt.m); got != tt.want {
				t.Errorf("IsMap(%v) = %v, want %v", tt.m, got, tt.want)
			}
		})
	}
}
