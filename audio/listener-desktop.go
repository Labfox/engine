// Copyright 2016 The G3N Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !wasm
// +build !wasm

package audio

import (
	"github.com/Labfox/engine/audio/al"
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

	al.Listener3f(al.Velocity, vx, vy, vz)
}

// SetVelocityVec sets the velocity of the listener with a vector.
func (l *Listener) SetVelocityVec(v *math32.Vector3) {

	al.Listener3f(al.Velocity, v.X, v.Y, v.Z)
}

// Velocity returns the velocity of the listener as x, y, z components.
func (l *Listener) Velocity() (float32, float32, float32) {

	return al.GetListener3f(al.Velocity)
}

// VelocityVec returns the velocity of the listener as a vector.
func (l *Listener) VelocityVec() math32.Vector3 {

	vx, vy, vz := al.GetListener3f(al.Velocity)
	return math32.Vector3{vx, vy, vz}
}

// SetGain sets the gain of the listener.
func (l *Listener) SetGain(gain float32) {

	al.Listenerf(al.Gain, gain)
}

// Gain returns the gain of the listener.
func (l *Listener) Gain() float32 {

	return al.GetListenerf(al.Gain)
}

// Render is called by the renderer at each frame.
// Updates the position and orientation of the listener.
func (l *Listener) Render(gl *gls.GLS) {

	// Sets the listener source world position
	var wpos math32.Vector3
	l.WorldPosition(&wpos)
	al.Listener3f(al.Position, wpos.X, wpos.Y, wpos.Z)

	// Get listener current world direction
	var vdir math32.Vector3
	l.WorldDirection(&vdir)

	// Assumes initial UP vector and recalculates current up vector
	vup := math32.Vector3{0, 1, 0}
	var vright math32.Vector3
	vright.CrossVectors(&vdir, &vup)
	vup.CrossVectors(&vright, &vdir)

	// Sets the listener orientation
	orientation := []float32{vdir.X, vdir.Y, vdir.Z, vup.X, vup.Y, vup.Z}
	al.Listenerfv(al.Orientation, orientation)
}
