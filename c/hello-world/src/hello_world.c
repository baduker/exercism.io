#include "hello_world.h"

const char *hello(void)
{
    static char hw[] = "Hello, World!";
    return hw;
}
