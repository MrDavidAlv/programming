/*
Practica de Enums
*/

#include <stdio.h>

int i, j, k = 0;
int minutosEstacionados = 0;
int minimunMinutes = 10;

char y = 'y';
unsigned char dailyWorkedHours;

enum deck {
    club = 0,
    diamonds = 5,
    hearts = 10,
    spades = 15, 
} card;

int main (){
    card = spades;
    printf("Card Power %d", card);
    printf("\nSize of var %d", sizeof(card));
    return 0;
}