package day12


// ========== DEFINITION ==================================

type Ferry struct {
  position  Coord
  waypoint  Coord

  commands  []Command
}


// ========== RECEIVERS ===================================

func (f Ferry) Move () int {
  position := f.position

  for _, cmd := range f.commands {
    switch cmd.dir {
    case "E", "N", "S", "W":
      position = position.MoveCardinally(cmd)
    case "F":
      position = position.MoveForward(cmd)
    case "L":
      position = position.TurnLeft(cmd)
    case "R":
      position = position.TurnRight(cmd)
    }
    f.position = position
  }

  return f.DistanceTraveled()
}

func (f Ferry) Navigate () int {
  position := f.position
  waypoint := f.waypoint

  for _, cmd := range f.commands {
    switch cmd.dir {
    case "E", "N", "S", "W":
      waypoint = waypoint.MoveCardinally(cmd)
    case "F":
      position = position.MoveRelatively(cmd, waypoint)
    case "L":
      waypoint = waypoint.RotateLeft(cmd)
    case "R":
      waypoint = waypoint.RotateRight(cmd)
    }
    f.position = position
    f.waypoint = waypoint
  }

  return f.DistanceTraveled()
}


// ---------- UTILITIES -----------------------------------

func (f Ferry) DistanceTraveled () int {
  return f.position.DistanceFromOrigin()
}
