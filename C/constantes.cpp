#include <stdio.h>

const int COSTHOTDOG = 100u;
const int HOTDOG = 30l;
const float TOTALBILL =  100.05;
const char NEWLINE = '\n';

#define PAPERCOST 3
#define PIZZACOST 1.5

int main (){
    float costoPizas = 1.3;
    float numberOfSlices = 3;
    costoPizas = PIZZACOST * numberOfSlices;
    printf("Number pizzas %f", numberOfSlices);
    printf("%cPrice %2f", NEWLINE, costoPizas);
    printf("%CTotal bill %3f", NEWLINE, costoPizas);
    return 0;
}