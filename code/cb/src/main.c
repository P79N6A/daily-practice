#include <stdio.h>

extern int f(int n);
extern int g(int n);

int
main(int argc, char **argv)
{
	printf("f(5)=%d\n", f(5));
	printf("g(2)=%d\n", g(2));
	return 0;
}
