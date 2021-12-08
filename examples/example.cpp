#include <stdio.h>
#include <stdlib.h>
#include "fun.h"

// MY OLD C++ CODE :D

int main()
{
	printf("Input the highest degree of the first polynomial: ");
	int P1 = DueToUserEntersAcceptableDegree();
	int* ArrayCoefP1 = (int*)malloc((P1 + 1) * sizeof(int));
	int Check1 = -1;
	printf("Input coefficients starting from the highest monomial. Coefficients can be 0 or 1.\n");
	for (int i = P1; i >= 0; i--)
	{
		printf(">>>Coefficient of monomial of degree %d: ", i);
		ArrayCoefP1[i] = DueToUserEntersAcceptableNumber();
		if ((ArrayCoefP1[i] == 1) && (i >= Check1))
		{
			Check1 = i;
		}
	}
	printf("Input the highest degree of the second polynomial: ");
	int P2 = DueToUserEntersAcceptableDegree();
	int* ArrayCoefP2 = (int*)malloc((P2 + 1) * sizeof(int));
	int Check2 = -1;
	printf("Input coefficients starting from the highest monomial. Coefficients can be 0 or 1.\n");
	for (int i = P2; i >= 0; i--)
	{
		printf(">>>Coefficient of monomial of degree %d: ", i);
		ArrayCoefP2[i] = DueToUserEntersAcceptableNumber();
		if ((ArrayCoefP2[i] == 1) && (i >= Check2))
		{
			Check2 = i;
		}
	}
	printf("---------------------------------------------\n");
	DividePolyByPoly(ArrayCoefP1, ArrayCoefP2, Check1, Check2);
	printf("\n---------------------------------------------");

	free(ArrayCoefP1);
	free(ArrayCoefP2);
	return 0;
}