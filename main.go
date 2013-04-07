package main

import (
    "math/rand"
    "time"
    )

const (
    NROFVARIABLES = 7
    NROFPARTICLES = 100
    )

type Particle struct{
  velocity float64
    currentSolution [NROFVARIABLES]float64
    bestSolution [NROFVARIABLES]float64
    /*TODO Velocity*/
}

func main() {
/*TODO Uncomment me
  var globalSolution  [NROFVARIABLES]float64
 */

    rand.Seed(time.Now().UnixNano())
    var particles = make([][]Particle,NROFPARTICLES)
    for i := range particles {
      particles[i] =  make([]Particle,NROFPARTICLES)
	for j := range particles[i] {
	  for k := range particles[i][j].currentSolution{
	    /*TODO Should be ~U(b_lo,bo_up)*/
	    particles[i][j].currentSolution[k] = rand.Float64()
	  }
	  particles[i][j].bestSolution = particles[i][j].currentSolution
	    /*TODO Compare bestSolution to globalSolution, if better overwrite
	      Uncomment him. 
	     */
	}
    }

}

/*
   For each particle:
   Random variables for each variable in particle.

   Init each particle with vector values picked
   randomly from ~U(b_lo,b_up) 
   (~U(a,b) = 1/b-a for x in [a,b] else 0) where
   b_lo is lower bound of search space and b_up is higher bound.
   Set particles best position to initial position.
   If applicable set that position as swarms best position.
   Set particle initial velocity with:
   ~U(-|b_up-b_lo|,|b_up-b_lo|)
   Until satisified:
   For each dimension in every particle:
   Generate random numbers r_p,r_g from ~U(0,1)
   Update particles velocity according to:
   velocity = butt_var*v_i,d + var_p*r_p(p_i,d-x_i,d)+var_g*r_p
   v_i,d current velocity[d].
   p_i,d current best known position[d]. 
   x_i,d current position[d]. 
   butt_var, var_p and var_g is decided by developer.
   New position equals old position + velocity
   If this new position is better than this particle bestSolution,
   replace the two.
   If this new position is better than the global particle, replace the two.
 */

