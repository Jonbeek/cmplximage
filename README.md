cmplximage
==========

This is a package for making and manipulating images with complex functions in Go. Currently, there is only support for creating images from complex functions.

To draw an image:
-----------------

1. Create a ColorMap either explicitly, or by using one of the pre-defined `ComplexMap` to `ColorMap` functions found in `src/color.go`.
2. Define the size of the image using a `image.Rectangle`.
3. Define the domain of the function using a `ComplexRect`.
4. Draw the image by calling `Draw` with the `ColorMap` obtained in step 1, the image size obtained in step 2, and the function domain obtained in step 3.

And that's it.
