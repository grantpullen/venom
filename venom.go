package venom

import (
	"container/heap"
	"encoding/json"
	"strings"
)

// Default the default set of available config levels
const (
	DefaultLevel ConfigLevel = iota
	FileLevel
	EnvironmentLevel
	FlagLevel
	OverrideLevel
)

const defaultDelim = "."

// Delim is the delimiter used for separating nested key spaces. The default is
// to separate on "." characters.
var Delim = defaultDelim

// ConfigLevel is a type alias used to identify various configuration levels
type ConfigLevel int

// Resolve resolves
func (c ConfigLevel) Resolve(key string, d ConfigMap) (val interface{}, ok bool) {
	return
}

// ConfigMap defines the inner map type which holds actual config data. These
// are nested under a ConfigLevel which determines their priority
type ConfigMap map[string]interface{}

func (c ConfigMap) merge(d ConfigMap) {
	for key, val := range d {
		c[key] = val
	}
}

// ConfigLevelMap is a mapping of config levels to the maps which contain
// various configuration values at those levels
type ConfigLevelMap map[ConfigLevel]ConfigMap

// Venom is the configuration registry responsible for storing and managing
// arbitrary configuration keys and values.
type Venom struct {
	// config is the global config space. Values are stored at a specified
	// ConfigLevel for prioritized retrieval
	config ConfigLevelMap

	// usedLevels is a sorted slice of all ConfigLevels currently stored in the
	// config map
	usedLevels *ConfigLevelHeap

	// envKeys is the list of registered keys to pull from the environment
	envKeys []string

	// resolvers is the definitive list of any customer ConfigLevel resolvers
	// provided to this Venom instance
	resolvers map[ConfigLevel]Resolver
}

// New returns a newly initialized Venom instance.
//
// The internal config map is created empty, only allocating space for a given
// config level once a value is set to that level.
func New() *Venom {
	return &Venom{
		config:     make(ConfigLevelMap),
		usedLevels: NewConfigLevelHeap(),
		envKeys:    make([]string, 0, 0),
		resolvers:  make(map[ConfigLevel]Resolver),
	}
}

// RegisterResolver registers a custom config resolver for the specified
// ConfigLevel.
//
// Additionally, if the provided level is not already in the current collection
// of active config levels, it will be added automatically
func (v *Venom) RegisterResolver(level ConfigLevel, r Resolver) {
	v.resolvers[level] = r
	heap.Push(v.usedLevels, level)
}

// SetLevel is a generic key/value setter method. It sets the provided k/v at
// the specified level inside the map, conditionally creating a new ConfigMap if
// one didn't previously exist.
func (v *Venom) SetLevel(level ConfigLevel, key string, value interface{}) {
	v.setIfNotExists(level, key, value)
}

// SetDefault sets the provided key and value into the DefaultLevel of the
// config collection.
func (v *Venom) SetDefault(key string, value interface{}) {
	v.SetLevel(DefaultLevel, key, value)
}

// SetOverride sets the provided key and value into the OverrideLevel of the
// config collection.
func (v *Venom) SetOverride(key string, value interface{}) {
	v.SetLevel(OverrideLevel, key, value)
}

// Get performs a fetch on a given key from the inner config collection.
func (v *Venom) Get(key string) interface{} {
	val, _ := v.find(strings.Split(key, Delim))
	return val
}

// Find searches for the given key, returning the discovered value and a
// boolean indicating whether or not the key was found
func (v *Venom) Find(key string) (interface{}, bool) {
	return v.find(strings.Split(key, Delim))
}

// setIfNotExists inserts the key and value into the config map, allocating the
// ConfigMap for that level if one was not already allocated.
func (v *Venom) setIfNotExists(l ConfigLevel, key string, value interface{}) {
	if _, ok := v.config[l]; !ok {
		v.config[l] = make(ConfigMap)
		heap.Push(v.usedLevels, l)
	}
	setNested(v.config[l], strings.Split(key, Delim), value)
}

// mergeIfNotExists merges the provided config map into the ConfigLevel l,
// allocating space for ConfigLevel l if the level hasn't already been
// allocated
func (v *Venom) mergeIfNotExists(l ConfigLevel, data ConfigMap) {
	if _, ok := v.config[l]; !ok {
		v.config[l] = make(ConfigMap)
		heap.Push(v.usedLevels, l)
	}
	v.config[l].merge(data)
}

// setNested inserts the provided value into the nested keyspace as defined by
// the delim separated keys
func setNested(config ConfigMap, keys []string, value interface{}) {
	for _, k := range keys {
		// if we're at the end of our slice of keys, set the value and return
		if len(keys) == 1 {
			config[k] = value
			return
		}
		// otherwise, create a new ConfigMap at the current node and continue
		config[k] = make(ConfigMap)
		setNested(config[k].(ConfigMap), keys[1:], value)
	}
	return
}

// find iterates over each ConfigLevel, in order, and returns the first value
// that matches or nil
func (v *Venom) find(keys []string) (val interface{}, ok bool) {
	for _, level := range *v.usedLevels {
		resolver, resolverExists := v.resolvers[level]
		if !resolverExists {
			resolver = DefaultResolver
		}

		if val, ok = resolver(keys, v.config[level]); ok {
			return
		}
	}
	return nil, false
}

// Clear removes all data from the ConfigLevelMap and resets the heap of config
// levels
func (v *Venom) Clear() {
	v.config = make(ConfigLevelMap)
	v.usedLevels = NewConfigLevelHeap()
}

// Debug returns the current venom ConfigLevelMap as a pretty-printed JSON string
func (v *Venom) Debug() string {
	b, _ := json.MarshalIndent(v.config, "", "  ")
	return string(b)
}
