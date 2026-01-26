package internal

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// TODO: better error messages

func NewGraphicsHandler(dir string) (*Graphics, error) {
	graphics := Graphics{
		images:   map[string]*rl.Image{},
		textures: map[string]rl.Texture2D{},
	}

	err := graphics.loadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to load directory %s", dir)
	}

	return &graphics, nil
}

type Graphics struct {
	images   map[string]*rl.Image
	textures map[string]rl.Texture2D
}

// TODO: check file ext

func (g *Graphics) loadImage(path string) error {
	// check that file exists
	_, err := os.Stat(path)
	if err != nil {
		return errors.New("file does not exist")
	}

	base := filepath.Base(path)

	// check that image already loaded
	_, exists := g.images[base]
	if exists {
		return errors.New("image already loaded")
	}

	g.images[base] = rl.LoadImage(path)
	g.textures[base] = rl.LoadTextureFromImage(g.images[base])

	return nil
}

func (g *Graphics) loadDir(dir string) error {
	return filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		return g.loadImage(path)
	})
}

func (g *Graphics) DrawTexture(image string, x int32, y int32) error {
	texture, ok := g.textures[image]
	if !ok {
		return errors.New("image not found")
	}

	rl.DrawTexture(texture, x, y, rl.White)

	return nil
}
