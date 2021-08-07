package main

import (
	"testing"
)

func TestCircle_Area(t *testing.T) {
	tests := []struct {
		name    string
		c       Circle
		want    float64
		wantErr bool
	}{
		{
			name: "usual",
			c: Circle{
				radius: 1,
			},
			want:    3.141592653589793,
			wantErr: false,
		},
		{
			name: "radius under or equal zero",
			c: Circle{
				radius: -1,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Area()
			if (err != nil) != tt.wantErr {
				t.Errorf("Circle.Area() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if got != tt.want {
				t.Errorf("Circle.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_Perimeter(t *testing.T) {
	tests := []struct {
		name    string
		c       Circle
		want    float64
		wantErr bool
	}{
		{
			name: "usual",
			c: Circle{
				radius: 1,
			},
			want:    6.283185307179586,
			wantErr: false,
		},
		{
			name: "radius under or equal zero",
			c: Circle{
				radius: 0,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Perimeter()
			if (err != nil) != tt.wantErr {
				t.Errorf("Circle.Perimeter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Circle.Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_Area(t *testing.T) {
	tests := []struct {
		name    string
		r       Rectangle
		want    float64
		wantErr bool
	}{
		{
			name: "usual",
			r: Rectangle{
				height: 2,
				width:  2,
			},
			want:    4,
			wantErr: false,
		},
		{
			name: "hight is under or equel zero ",
			r: Rectangle{
				height: 0,
				width:  2,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "width is under or equel zero ",
			r: Rectangle{
				height: 2,
				width:  0,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Area()
			if (err != nil) != tt.wantErr {
				t.Errorf("Rectangle.Area() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Rectangle.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_Perimeter(t *testing.T) {
	tests := []struct {
		name    string
		r       Rectangle
		want    float64
		wantErr bool
	}{
		{
			name: "usual",
			r: Rectangle{
				height: 2,
				width:  2,
			},
			want:    8,
			wantErr: false,
		},
		{
			name: "hight is under or equel zero ",
			r: Rectangle{
				height: -1,
				width:  2,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "width is under or equel zero ",
			r: Rectangle{
				height: 2,
				width:  -1,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Perimeter()
			if (err != nil) != tt.wantErr {
				t.Errorf("Rectangle.Perimeter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Rectangle.Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestCircle_String(t *testing.T) {
	tests := []struct {
		name string
		c    Circle
		want string
	}{
		{
			name: "usual",
			c:    Circle{2},
			want: "Cirlce : radius 2.000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("Circle.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestRectangle_String(t *testing.T) {
	tests := []struct {
		name string
		r    Rectangle
		want string
	}{
		{
			name: "regular",
			r:    Rectangle{2, 2},
			want: "Rectangle with height 2.000000 and width 2.000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.String(); got != tt.want {
				t.Errorf("Rectangle.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
