# Template for GLPK problem formulation

var A integer;
var B integer;

param Ax := {Ax};
param Ay := {Ay};
param Bx := {Bx};
param By := {By};
param Px := {Px};
param Py := {Py};

minimize obj: 3 * A + B;

s.t. c1: A * Ax + B * Bx = Px;
s.t. c2: A * Ay + B * By = Py;

solve;

display A, B;

end;

