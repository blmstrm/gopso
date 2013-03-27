package main

const (
    NROFVARIABLES = 5
    )

  type Particle struct{
    velocity, direction float64
      solution  [NROFVARIABLES]float64
  }

func main() {

particles := make([][]Particle,2)
	     for i := range particles {
	       particles[i] =  make([]Particle,2)
	     }
}
