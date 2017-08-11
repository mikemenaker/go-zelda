[![Go Report Card](https://goreportcard.com/badge/github.com/faiface/pixel)](https://goreportcard.com/report/github.com/faiface/pixel)
# go-zelda

Zelda clone using faiface/pixel

![Alt text](https://ibin.co/3WXbFPOMGM10.gif "Go Zelda")

## Dependencies


```
go get github.com/faiface/pixel
```

[PixelGL](https://godoc.org/github.com/faiface/pixel/pixelgl) backend uses OpenGL to render
graphics. Because of that, OpenGL development libraries are needed for compilation. The dependencies
are same as for [GLFW](https://github.com/go-gl/glfw).

- On macOS, you need Xcode or Command Line Tools for Xcode (`xcode-select --install`) for required
  headers and libraries.
- On Ubuntu/Debian-like Linux distributions, you need `libgl1-mesa-dev` and `xorg-dev` packages.
- On CentOS/Fedora-like Linux distributions, you need `libX11-devel libXcursor-devel libXrandr-devel
  libXinerama-devel mesa-libGL-devel libXi-devel` packages.
- See [here](http://www.glfw.org/docs/latest/compile.html#compile_deps) for full details.


