grammar Calc;

// Operators
ADD: 'nyya';      // 2y
SUB: 'nyyya';     // 3y
MUL: 'nyyyya';    // 4y
DIV: 'nyyyyya';   // 5y

// Numbers
NUMBER: 'ny' ('a')* | 'nyu';

// Ignore whitespace
WS: [ \r\n\t]+ -> skip;

// Parser rules
start : expression EOF;

expression
   : expression op=(MUL|DIV) expression # MulDiv
   | expression op=(ADD|SUB) expression # AddSub
   | NUMBER                             # Number
   ;