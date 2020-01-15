package world

import (
	"github.com/dragonfly-tech/dragonfly/dragonfly/world/chunk"
	"github.com/dragonfly-tech/dragonfly/dragonfly/world/gamemode"
	"io"
)

// Provider represents a value that may provide world data to a World value. It usually does the reading and
// writing of the world data so that the World may use it.
type Provider interface {
	io.Closer
	// WorldName returns the name of the world that the provider provides for. When setting the provider of a
	// World, the World will replace its current name with this one.
	WorldName() string
	// SetWorldName sets the name of the world to a new name.
	SetWorldName(name string)
	// WorldSpawn returns the spawn position of the world. Although players may spawn at different positions,
	// every new player spawns at this position.
	WorldSpawn() BlockPos
	// SetWorldSpawn sets the spawn of a world to a new position.
	SetWorldSpawn(pos BlockPos)
	// LoadChunk attempts to load a chunk from the chunk position passed. If successful, a non-nil chunk is
	// returned and exists is true and err nil. If no chunk was saved at the chunk position passed, the chunk
	// returned is nil, and so is the error. If the chunk did exist, but if the data was invalid, nil is
	// returned for the chunk and true, with a non-nil error.
	// If exists ends up false, the chunk at the position is instead newly generated by the world.
	LoadChunk(position ChunkPos) (c *chunk.Chunk, exists bool, err error)
	// SaveChunk saves a chunk at a specific position in the provider. If writing was not successful, an error
	// is returned.
	SaveChunk(position ChunkPos, c *chunk.Chunk) error
	// LoadEntities loads all entities stored at a particular chunk position. If the entities cannot be read,
	// LoadEntities returns a non-nil error.
	LoadEntities(position ChunkPos) ([]Entity, error)
	// SaveEntities saves a list of entities in a chunk position. If writing is not successful, an error is
	// returned.
	SaveEntities(position ChunkPos, entities []Entity) error
	// LoadTime loads the time of the world.
	LoadTime() int64
	// SaveTime saves the time of the world.
	SaveTime(time int64)
	// SaveTimeCycle saves the state of the time cycle: Either stopped or started. If true is passed, the time
	// is running. If false, the time is stopped.
	SaveTimeCycle(running bool)
	// LoadTimeCycle loads the state of the time cycle: If time is running, true is returned. If the time
	// cycle is stopped, false is returned.
	LoadTimeCycle() bool
	// DefaultGameMode loads the default game mode of the world.
	DefaultGameMode() gamemode.GameMode
	// SetDefaultGameMode sets the default game mode of the world.
	SetDefaultGameMode(mode gamemode.GameMode)
}

// NoIOProvider implements a Provider while not performing any disk I/O. It generates values on the run and
// dynamically, instead of reading and writing data, and returns otherwise empty values.
type NoIOProvider struct{}

// DefaultGameMode ...
func (p NoIOProvider) DefaultGameMode() gamemode.GameMode { return gamemode.Adventure{} }

// SetDefaultGameMode ...
func (p NoIOProvider) SetDefaultGameMode(mode gamemode.GameMode) {}

// SetWorldSpawn ...
func (p NoIOProvider) SetWorldSpawn(pos BlockPos) {}

// SaveTimeCycle ...
func (p NoIOProvider) SaveTimeCycle(running bool) {}

// LoadTimeCycle ...
func (p NoIOProvider) LoadTimeCycle() bool {
	return true
}

// LoadTime ...
func (p NoIOProvider) LoadTime() int64 {
	return 0
}

// SaveTime ...
func (p NoIOProvider) SaveTime(time int64) {}

// LoadEntities ...
func (p NoIOProvider) LoadEntities(position ChunkPos) ([]Entity, error) {
	return nil, nil
}

// SaveEntities ...
func (p NoIOProvider) SaveEntities(position ChunkPos, entities []Entity) error {
	return nil
}

// SaveChunk ...
func (p NoIOProvider) SaveChunk(position ChunkPos, c *chunk.Chunk) error {
	return nil
}

// LoadChunk ...
func (p NoIOProvider) LoadChunk(position ChunkPos) (*chunk.Chunk, bool, error) {
	return nil, false, nil
}

// WorldName ...
func (p NoIOProvider) WorldName() string {
	return ""
}

// SetWorldName ...
func (p NoIOProvider) SetWorldName(name string) {}

// WorldSpawn ...
func (p NoIOProvider) WorldSpawn() BlockPos {
	return BlockPos{0, 30, 0}
}

// Close ...
func (p NoIOProvider) Close() error {
	return nil
}