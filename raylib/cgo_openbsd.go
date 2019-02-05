// +build openbsd

package rl

/*
#include "external/glfw/src/context.c"
#include "external/glfw/src/init.c"
#include "external/glfw/src/input.c"
#include "external/glfw/src/monitor.c"
#include "external/glfw/src/vulkan.c"
#include "external/glfw/src/null_joystick.c"
#include "external/glfw/src/window.c"

#ifdef _GLFW_X11
#include "external/glfw/src/x11_init.c"
#include "external/glfw/src/x11_monitor.c"
#include "external/glfw/src/x11_window.c"
#include "external/glfw/src/glx_context.c"
#endif

#include "external/glfw/src/posix_thread.c"
#include "external/glfw/src/posix_time.c"
#include "external/glfw/src/xkb_unicode.c"
#include "external/glfw/src/egl_context.c"
#include "external/glfw/src/osmesa_context.c"

#cgo openbsd LDFLAGS: -lGL -lm -pthread -lX11 -L/usr/X11R6/lib

#cgo openbsd CFLAGS: -D_GLFW_X11 -I/usr/X11R6/include -Iexternal/glfw/include -I/usr/local/include/inotify/ -DPLATFORM_DESKTOP

#cgo openbsd,opengl11 CFLAGS: -DGRAPHICS_API_OPENGL_11
#cgo openbsd,opengl21 CFLAGS: -DGRAPHICS_API_OPENGL_21
#cgo openbsd,!opengl11,!opengl21 CFLAGS: -DGRAPHICS_API_OPENGL_33
*/
import "C"
