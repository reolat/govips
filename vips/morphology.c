#include "morphology.h"

int rank(VipsImage *in, VipsImage **out, int width, int height, int index) {
  return vips_rank(in, out, width, height, index, NULL);
}

int morph(VipsImage *in, VipsImage **out, VipsImage *mask, VipsOperationMorphology morph) {
  return vips_morph(in, out, mask, morph, NULL);
}
