#include <stdio.h>
#include "foo.h"

int main() {
  printf("Hello from C. The meaning of life is %d", meaningOfLife());
}

int meaningOfLife() {
  return 42;
}
