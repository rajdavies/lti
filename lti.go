// Package lti implements a general state-space representation of a linear system.
// A state-space representation can be expressed in a matrix form as
//
// x'(t) = A * x(t) + B * u(t)
// and
// y(t)  = C * x(t) + D * u(t)
//
// where x(t) represents the state and u(t) the control input vectors and the
// matrices are
// A: System matrix,
// B: Control matrix,
// C: Output matrix and
// D: Feedforward matrix
//
package lti

type LTI interface {
	Observable() (bool, error)
	Controllable() (bool, error)
}