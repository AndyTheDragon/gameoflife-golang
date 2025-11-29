package main

// Plane: out-of-bounds -> -1
func IndexPlane(x, y, rows, cols int) int {
    if x < 0 || x >= cols || y < 0 || y >= rows {
        return -1
    }
    return y*cols + x
}

// Torus: wrap both axes
func IndexTorus(x, y, rows, cols int) int {
    x = (x + cols) % cols
    y = (y + rows) % rows
    return y*cols + x
}

// Sphere: your current “sphere” behaviour
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

func NewGrid(rows, cols int) *GenericGrid {
    return NewGenericGrid(rows, cols, IndexPlane)
}

func NewTorusGrid(rows, cols int) *GenericGrid {
    return NewGenericGrid(rows, cols, IndexTorus)
}

func NewSphereGrid(rows, cols int) *GenericGrid {
    return NewGenericGrid(rows, cols, IndexSphere)
}