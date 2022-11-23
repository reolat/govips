package vips

// #include "morphology.h"
import "C"

type OperationMorphology int

const (
	OperationMorphologyErode  = C.VIPS_OPERATION_MORPHOLOGY_ERODE
	OperationMorphologyDilate = C.VIPS_OPERATION_MORPHOLOGY_DILATE
	OperationMorphologyLast   = C.VIPS_OPERATION_MORPHOLOGY_LAST
)

// https://libvips.github.io/libvips/API/current/libvips-morphology.html#vips-rank
func vipsRank(in *C.VipsImage, width int, height int, index int) (*C.VipsImage, error) {
	incOpCounter("rank")
	var out *C.VipsImage

	err := C.rank(in, &out, C.int(width), C.int(height), C.int(index))
	if int(err) != 0 {
		return nil, handleImageError(out)
	}

	return out, nil
}

func vipsMorph(in *C.VipsImage, mask *C.VipsImage, morph OperationMorphology) (*C.VipsImage, error) {
	incOpCounter("morph")
	var out *C.VipsImage

	err := C.morph(in, &out, mask, C.VipsOperationMorphology(morph))
	if int(err) != 0 {
		return nil, handleImageError(out)
	}

	return out, nil
}
