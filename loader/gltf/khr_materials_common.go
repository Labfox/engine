package gltf

import (
	"github.com/Labfox/engine/material"
	"github.com/Labfox/engine/math32"
	"github.com/Labfox/engine/texture"
)

// loadMaterialCommon receives an interface value describing a KHR_materials_common extension,
// decodes it and returns a Material closest to the specified description.
// The specification of this extension is at:
// https://github.com/KhronosGroup/glTF/tree/master/extensions/1.0/Khronos/KHR_materials_common
func (g *GLTF) loadMaterialCommon(ext interface{}) (material.IMaterial, error) {

	// The extension must be an object
	m := ext.(map[string]interface{})

	// Double sided
	doubleSided := false
	val, ok := m["doubleSided"]
	if ok {
		doubleSided = val.(bool)
	}

	// Transparent
	transparent := false
	val, ok = m["transparent"]
	if ok {
		transparent = val.(bool)
	}

	// Defaul values
	ambient := []float32{0, 0, 0, 1}
	diffuse := []float32{0, 0, 0, 1}
	emission := []float32{0, 0, 0, 1}
	specular := []float32{0, 0, 0, 1}
	shininess := float32(0)
	transparency := float32(1)
	var texDiffuse *texture.Texture2D

	// Converts a slice of interface values which should be float64
	// to a slice of float32
	convIF32 := func(v interface{}) []float32 {

		si := v.([]interface{})
		res := make([]float32, 0)
		for i := 0; i < len(si); i++ {
			res = append(res, float32(si[i].(float64)))
		}
		return res
	}

	// Values
	values, ok := m["values"].(map[string]interface{})
	if ok {

		// Ambient light
		val, ok = values["ambient"]
		if ok {
			ambient = convIF32(val)
		}

		// Diffuse light
		val, ok = values["diffuse"]
		if ok {
			v := convIF32(val)
			// Checks for texture index
			if len(v) == 1 {
				var err error
				texDiffuse, err = g.LoadTexture(int(v[0]))
				if err != nil {
					return nil, err
				}
				diffuse = []float32{1, 1, 1, 1}
			}
		}

		// Emission light
		val, ok = values["emission"]
		if ok {
			emission = convIF32(val)
		}

		// Specular light
		val, ok = values["specular"]
		if ok {
			specular = convIF32(val)
		}

		// Shininess
		val, ok = values["shininess"]
		if ok {
			s := convIF32(val)
			shininess = s[0]
		}

		// Transparency
		val, ok = values["transparency"]
		if ok {
			s := convIF32(val)
			transparency = s[0]
		}
	}

	//log.Error("doubleSided:%v", doubleSided)
	//log.Error("technique:%v", technique)
	//log.Error("transparent:%v", transparent)
	//log.Error("values:%v", values)
	//log.Error("ambient:%v", ambient)
	//log.Error("diffuse:%v", diffuse)
	//log.Error("emission:%v", emission)
	//log.Error("specular:%v", specular)
	//log.Error("shininess:%v", shininess)
	//log.Error("transparency:%v", transparency)

	mat := material.NewStandard(&math32.Color{diffuse[0], diffuse[1], diffuse[2]})
	mat.SetAmbientColor(&math32.Color{ambient[0], ambient[1], ambient[2]})
	mat.SetEmissiveColor(&math32.Color{emission[0], emission[1], emission[2]})
	mat.SetSpecularColor(&math32.Color{specular[0], specular[1], specular[2]})
	mat.SetShininess(shininess)
	mat.SetOpacity(transparency)
	if texDiffuse != nil {
		mat.AddTexture(texDiffuse)
	}

	// Double Sided
	if doubleSided {
		mat.SetSide(material.SideDouble)
	} else {
		mat.SetSide(material.SideFront)
	}

	// Transparency
	if transparent {
		mat.SetDepthMask(true)
	} else {
		mat.SetDepthMask(false)
	}
	return mat, nil
}
