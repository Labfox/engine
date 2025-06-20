// Copyright 2016 The G3N Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package helper

import (
	"github.com/Labfox/engine/core"
	"github.com/Labfox/engine/geometry"
	"github.com/Labfox/engine/gls"
	"github.com/Labfox/engine/graphic"
	"github.com/Labfox/engine/material"
	"github.com/Labfox/engine/math32"
)

// Normals is the visual representation of the normals of a target object.
type Normals struct {
	graphic.Lines
	size           float32
	targetNode     *core.Node
	targetGeometry *geometry.Geometry
}

// NewNormals creates a normals helper for the specified IGraphic, with the specified size, color, and lineWidth.
func NewNormals(ig graphic.IGraphic, size float32, color *math32.Color, lineWidth float32) *Normals {

	// Creates new Normals helper
	nh := new(Normals)
	nh.size = size

	// Save the object to show the normals
	nh.targetNode = ig.GetNode()

	// Get the geometry of the target object
	nh.targetGeometry = ig.GetGeometry()

	// Get the number of target vertex positions
	vertices := nh.targetGeometry.VBO(gls.VertexPosition)
	n := vertices.Buffer().Size() * 2

	// Creates this helper geometry
	geom := geometry.NewGeometry()
	positions := math32.NewArrayF32(n, n)
	geom.AddVBO(gls.NewVBO(positions).AddAttrib(gls.VertexPosition))

	// Creates this helper material
	mat := material.NewStandard(color)
	mat.SetLineWidth(lineWidth)

	// Initialize graphic
	nh.Lines.Init(geom, mat)

	nh.Update()
	return nh
}

// Update should be called in the render loop to
// update the normals based on the target object.
func (nh *Normals) Update() {

	var v1 math32.Vector3
	var v2 math32.Vector3
	var normalMatrix math32.Matrix3

	// Updates the target object matrix and get its normal matrix
	matrixWorld := nh.targetNode.MatrixWorld()
	normalMatrix.GetNormalMatrix(&matrixWorld)

	// Get the target positions and normals buffers
	tPosVBO := nh.targetGeometry.VBO(gls.VertexPosition)
	tPositions := tPosVBO.Buffer()
	tNormVBO := nh.targetGeometry.VBO(gls.VertexNormal)
	tNormals := tNormVBO.Buffer()

	// Get this object positions buffer
	geom := nh.GetGeometry()
	posVBO := geom.VBO(gls.VertexPosition)
	positions := posVBO.Buffer()

	// For each target object vertex position:
	for pos := 0; pos < tPositions.Size(); pos += 3 {
		// Get the target vertex position and apply the current world matrix transform
		// to get the base for this normal line segment.
		tPositions.GetVector3(pos, &v1)
		v1.ApplyMatrix4(&matrixWorld)

		// Calculates the end position of the normal line segment
		tNormals.GetVector3(pos, &v2)
		v2.ApplyMatrix3(&normalMatrix).Normalize().MultiplyScalar(nh.size).Add(&v1)

		// Sets the line segment representing the normal of the current target position
		// at this helper VBO
		positions.SetVector3(2*pos, &v1)
		positions.SetVector3(2*pos+3, &v2)
	}
	posVBO.Update()
}
