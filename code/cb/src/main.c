#include <stdio.h>

extern int f(int n);

int
main(int argc, char **argv)
{
	printf("f(5)=%d\n", f(5));
	return 0;
}
