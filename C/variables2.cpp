#include <stdio.h>

//variable declarations
int global; //se puede usar en otros archivos
int a, b, c;

float f, g, h;

int main (){
    //variable definition
    int a, b, c;
    //variable initialization
    a = -10;
    b = -2147483647;
    c = a + b;

    g = 1000.999999;
    f = 1.111111;
    h = f + g;
    
    printf("%d + %d = %d", a, b, c);
    printf("\n %f + %f = %f", f, g, h);
    return 0;
}