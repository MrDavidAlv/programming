/*
Practica de Enums
*/

#include <stdio.h>

//variable declarations
extern int global; //se puede usar en otros archivos
int a, b, c;

float f, g, h;

int main (){
    //variable initialization
    a = 10;
    b = 2147483647;

    c = a + b;
    printf("%d + %d = %d", a, b, c);
    return 0;
}