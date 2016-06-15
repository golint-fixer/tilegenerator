package entities

import (
	"fmt"

	"github.com/TerraFactory/wktparser"
	"github.com/TerraFactory/wktparser/geometry"
)

// MapObject represents a geometry on a map
type MapObject struct {
	ID        int
	StyleName string
<<<<<<< HEAD
	TypeID    int
=======
	Label     string
	Position  string
	Size      int
>>>>>>> 6bfe44e5dfdff73a5334cf15fa92da1455dcc76e
	Geometry  geometry.Geometry
}

// NewObject creates new MapObject with a parsed from WKT geometry
func NewObject(id int, typeId int, wkt string) (*MapObject, error) {
	geo, err := wktparser.Parse(wkt)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &MapObject{ID: id, TypeID: typeId, Geometry: geo}, nil
}
