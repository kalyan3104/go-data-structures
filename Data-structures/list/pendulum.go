package main

import (
	"container/list"
	"fmt"
	"math"
)

// PendulumState represents the state of the pendulum at a specific time
type PendulumState struct {
	time           float64 // Time at this state
	angle          float64 // Angle (radians)
	angularVel     float64 // Angular velocity (radians/second)
	angularAccel   float64 // Angular acceleration (radians/second^2)
}

// Pendulum stores the properties of the pendulum and its history
type Pendulum struct {
	length   float64       // Length of the pendulum (meters)
	damping  float64       // Damping factor (optional)
	history  *list.List    // History of the pendulum's states (using container/list)
}

// NewPendulum creates a new Pendulum with initial values
func NewPendulum(length, initialAngle, damping float64) *Pendulum {
	return &Pendulum{
		length:   length,
		damping:  damping,
		history:  list.New(), // Initialize a new doubly linked list
	}
}

// Update calculates the new state of the pendulum and adds it to the history
func (p *Pendulum) Update(time, dt float64) {
	// Get the current state (last state in the history)
	var currentState PendulumState
	if p.history.Len() == 0 {
		// If no history, assume starting from rest
		currentState = PendulumState{time: 0, angle: math.Pi / 4, angularVel: 0.0, angularAccel: 0.0}
	} else {
		// Get the last element in the list
		lastElement := p.history.Back()
		currentState = lastElement.Value.(PendulumState)
	}

	// Calculate the angular acceleration
	currentState.angularAccel = -(9.81 / p.length) * math.Sin(currentState.angle)
	// Update the angular velocity
	currentState.angularVel += currentState.angularAccel * dt
	// Apply damping
	currentState.angularVel *= (1 - p.damping)
	// Update the angle
	currentState.angle += currentState.angularVel * dt
	// Set the time
	currentState.time = time

	// Append the new state to the history
	p.history.PushBack(currentState)
}

// GetHistory returns the history of the pendulum's states as a slice
func (p *Pendulum) GetHistory() []PendulumState {
	var states []PendulumState
	for e := p.history.Front(); e != nil; e = e.Next() {
		states = append(states, e.Value.(PendulumState))
	}
	return states
}

func main() {
	// Create a new pendulum with length 1.0 meter, initial angle 45 degrees, and a damping factor of 0.01
	pendulum := NewPendulum(1.0, math.Pi/4, 0.01)

	// Simulation parameters
	dt := 0.01      // Time step (seconds)
	totalTime := 10.0 // Total time to simulate (seconds)

	// Run the simulation
	for t := 0.0; t < totalTime; t += dt {
		pendulum.Update(t, dt)
	}

	// Retrieve the history and print it
	history := pendulum.GetHistory()
	for _, state := range history {
		fmt.Printf("Time: %.2f s, Angle: %.2f rad, Angular Velocity: %.2f rad/s, Angular Acceleration: %.2f rad/s^2\n",
			state.time, state.angle, state.angularVel, state.angularAccel)
	}
}
