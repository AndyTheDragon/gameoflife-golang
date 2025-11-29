package main

type GridIface interface {
    Size() (int, int)
    Get(x, y int) uint8
    Set(x, y int, value uint8)
    SumOfNeighbors(x, y int) uint8
    Clear()
    Randomize(probability float32)
    CopyFrom(other GridIface)
    CopyTo(indexFunc indexFunc) GridIface
    CreateSpaceship(name string, x, y int)
}