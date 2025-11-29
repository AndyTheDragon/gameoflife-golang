package main

// Plane: out-of-bounds -> -1
func IndexPlane(x, y, rows, cols int) int {
    if x < 0 || x >= cols || y < 0 || y >= rows {
        return -1
    }
    return y*cols + x
}

func NewGrid(rows, cols int) *GenericGrid {
	return NewGenericGrid(rows, cols, IndexPlane)
}

// Torus: wrap both axes
func IndexTorus(x, y, rows, cols int) int {
    x = (x + cols) % cols
    y = (y + rows) % rows
    return y*cols + x
}

func NewTorusGrid(rows, cols int) *GenericGrid {
	return NewGenericGrid(rows, cols, IndexTorus)
}

// "Sphere": wrap horizontally, mirror vertically
func IndexSphere(x, y, rows, cols int) int {
    if y < 0 {
        y = 0
        x = x + (cols / 2)
    }
    if y > rows-1 {
        y = rows - 1
        x = x + (cols / 2)
    }
    x = (x + cols) % cols
    return y*cols + x
}

func NewSphereGrid(rows, cols int) *GenericGrid {
    return NewGenericGrid(rows, cols, IndexSphere)
}

// Cylinder: wrap horizontally, clamp vertically
func IndexCylinder(x, y, rows, cols int) int {
    if y < 0 || y >= rows {
        return -1
    }
    x = (x + cols) % cols
    return y*cols + x
}

func NewCylinderGrid(rows, cols int) *GenericGrid {
    return NewGenericGrid(rows, cols, IndexCylinder)
}

// Klein bottle: horizontal normal wrap, vertical wrap with horizontal flip
func IndexKlein(x, y, rows, cols int) int {
    // vertical wrapping first
    if y < 0 {
        y = rows - 1
        // flip horizontally when crossing top/bottom
        x = cols - 1 - x
    } else if y >= rows {
        y = 0
        x = cols - 1 - x
    }

    // horizontal normal wrap
    x = (x + cols) % cols

    // if after flip & wrap x goes out of [0,cols), treat as empty
    if x < 0 || x >= cols {
        return -1
    }
    return y*cols + x
}

func NewKleinGrid(rows, cols int) *GenericGrid {
    return NewGenericGrid(rows, cols, IndexKlein)
}

// Möbius strip along X: X wraps with flip, Y has hard edges
func IndexMoebiusX(x, y, rows, cols int) int {
    if y < 0 || y >= rows {
        return -1
    }

    // when going left of 0, appear at right with flipped y
    if x < 0 {
        x = cols - 1
        y = rows - 1 - y
    } else if x >= cols {
        x = 0
        y = rows - 1 - y
    }

    if x < 0 || x >= cols {
        return -1
    }
    return y*cols + x
}

func NewMoebiusXGrid(rows, cols int) *GenericGrid {
    return NewGenericGrid(rows, cols, IndexMoebiusX)
}

// Reflecting edges: bounce coordinates back into the grid
func IndexReflect(x, y, rows, cols int) int {
    // reflect x
    for x < 0 || x >= cols {
        if x < 0 {
            x = -x - 1
        } else if x >= cols {
            x = 2*cols - x - 1
        }
    }
    // reflect y
    for y < 0 || y >= rows {
        if y < 0 {
            y = -y - 1
        } else if y >= rows {
            y = 2*rows - y - 1
        }
    }
    return y*cols + x
}

func NewReflectGrid(rows, cols int) *GenericGrid {
    return NewGenericGrid(rows, cols, IndexReflect)
}
