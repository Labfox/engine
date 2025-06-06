// Copyright 2016 The G3N Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasm
// +build wasm

package audio

import (
	"github.com/Labfox/engine/core"
	"github.com/Labfox/engine/gls"
	"github.com/Labfox/engine/math32"
)

// Listener is an audio listener positioned in space.
type Listener struct {
	core.Node
}

// NewListener creates a Listener object.
func NewListener() *Listener {

	l := new(Listener)
	l.Node.Init(l)
	return l
}

// SetVelocity sets the velocity of the listener with x, y, z components.
func (l *Listener) SetVelocity(vx, vy, vz float32) {
	panic("unimplemented")
	// TODO
}

// SetVelocityVec sets the velocity of the listener with a vector.
func (l *Listener) SetVelocityVec(v *math32.Vector3) {
	panic("unimplemented")
	// TODO
}

// Velocity returns the velocity of the listener as x, y, z components.
func (l *Listener) Velocity() (float32, float32, float32) {
	panic("unimplemented")
	// TODO
}

// VelocityVec returns the velocity of the listener as a vector.
func (l *Listener) VelocityVec() math32.Vector3 {
	panic("unimplemented")
	// TODO
}

// SetGain sets the gain of the listener.
func (l *Listener) SetGain(gain float32) {
	panic("unimplemented")
	// TODO
}

// Gain returns the gain of the listener.
func (l *Listener) Gain() float32 {
	panic("unimplemented")
	// TODO
}

// Render is called by the renderer at each frame.
// Updates the position and orientation of the listener.
func (l *Listener) Render(gl *gls.GLS) {
	panic("unimplemented")
	// TODO
}
